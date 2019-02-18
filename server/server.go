package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Nimsaja/PortfolioPerformance/data"
	"github.com/Nimsaja/PortfolioPerformance/portfolio"
	"github.com/Nimsaja/PortfolioPerformance/store"
	"github.com/Nimsaja/PortfolioPerformance/yahoo"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
)

// handle CORS and the OPTION method
func corsAndOptionHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		} else {
			h.ServeHTTP(w, r)
		}
	}
}

// create all used Handler
func handler() http.Handler {
	router := mux.NewRouter()

	router.Handle("/", http.RedirectHandler("/portfolio/table", http.StatusFound))

	url := "/portfolio/hist"
	router.HandleFunc(url, loadHistData).Methods("GET")
	url = "/portfolio/table"
	router.HandleFunc(url, getTableData).Methods("GET")

	return corsAndOptionHandler(router)
}

// App ...
type App struct {
	storage2BK store.StorageService
	storage2DB store.StorageService // TODO replace storage2BK with this later
	urlService yahoo.URLService
}

func inCloud() bool {
	b, _ := strconv.ParseBool(os.Getenv("RUN_IN_CLOUD"))
	return b
}

// New creates new App with services
func New() App {
	inCloud := inCloud()
	return App{store.New(inCloud, jasmin.Name), store.New(inCloud, jasmin.Name+"_DB"), yahoo.New(inCloud)}
}

var jasmin portfolio.Owner
var app App

func main() {
	jasmin = data.Jasmin()
	app = New()

	http.Handle("/", handler())

	if !inCloud() {
		fmt.Println("*******Open http://localhost:8080/portfolio/table*******")
		fmt.Println()
	}
	appengine.Main()
}

// this is called by the cron job to save the daily datas
func loadHistData(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	qs := app.urlService.GetAllQuotes(c, jasmin.Stocks())

	//Save Values
	err := app.storage2BK.Save(c, jasmin.GetTodaySum(qs), jasmin.BuySum(), jasmin.RegularMarketTime(qs))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		writeOutAsJSON(w, err.Error())
	} else {
		s := fmt.Sprintf("Successfully wrote new data to bucket. Sum from today: %v", jasmin.GetTodaySum(qs))
		writeOutAsJSON(w, s)
	}

	//Save Values to Database
	err = app.storage2DB.Save(c, jasmin.GetTodaySum(qs), jasmin.BuySum(), jasmin.RegularMarketTime(qs))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		writeOutAsJSON(w, err.Error())
	} else {
		s := fmt.Sprintf("Successfully wrote new data to database. Sum from today: %v. Log time %v.", jasmin.GetTodaySum(qs), time.Now())
		writeOutAsJSON(w, s)
	}
}

func getTableData(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	//need to get all quotes for the current value
	qs := app.urlService.GetAllQuotes(c, jasmin.Stocks())

	//Load Historical Data from File
	a, err := app.storage2BK.Load(c)
	if err != nil {
		fmt.Println("Error ", err)
	}

	//add current value to array (not to file - this is done a few times a day by a cron job)
	t := jasmin.RegularMarketTime(qs)
	d := store.Data{Value: jasmin.GetTodaySum(qs), Diff: jasmin.GetTodaySum(qs) - jasmin.BuySum()}
	a = store.AppendToList(a, d, t)

	writeOutAsJSON(w, a)
}

func writeOutAsJSON(w http.ResponseWriter, v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s\n", string(b))
}
