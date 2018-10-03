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

	url := "/portfolio/forcecall"
	router.HandleFunc(url, loadHistData).Methods("GET")
	url = "/portfolio/table"
	router.HandleFunc(url, getTableData).Methods("GET")

	return corsAndOptionHandler(router)
}

var jasmin portfolio.Owner
var inCloud bool
var storage store.StorageService
var urlService yahoo.URLService

func main() {
	inCloud, _ = strconv.ParseBool(os.Getenv("RUN_IN_CLOUD"))
	jasmin = data.Jasmin()
	storage = store.New(inCloud, jasmin.Name)
	urlService = yahoo.New(inCloud)

	http.Handle("/", handler())

	if !inCloud {
		fmt.Println("*******Open http://localhost:8080/portfolio/table*******")
		fmt.Println()
	}
	appengine.Main()
}

func loadHistData(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	qs := urlService.GetAllQuotes(c, jasmin.Stocks())

	//Save Values
	err := storage.Save(c, jasmin.GetYesterdaySum(qs), jasmin.BuySum())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		writeOutAsJSON(w, err.Error())
	} else {
		s := fmt.Sprintf("Successfully wrote new data to storage. Sum from Yesterday: %v", jasmin.GetYesterdaySum(qs))
		writeOutAsJSON(w, s)
	}
}

func getTableData(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	//need to get all quotes for the current value
	qs := urlService.GetAllQuotes(c, jasmin.Stocks())

	//Load Historical Data from File
	a, err := storage.Load(c)
	if err != nil {
		fmt.Println("Error ", err)
	}

	//add current value
	d := store.Data{TimeHuman: time.Now(), Value: jasmin.GetTodaySum(qs), Diff: jasmin.GetTodaySum(qs) - jasmin.BuySum()}
	a = append(a, d)

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
