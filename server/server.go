package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func main() {
	jasmin = data.Jasmin()

	http.Handle("/", handler())

	fmt.Println("*******Open http://localhost:8080/portfolio/table*******")
	fmt.Println()
	appengine.Main()
}

func loadHistData(w http.ResponseWriter, r *http.Request) {
	qs := yahoo.GetAllQuotes(jasmin.Stocks())

	//Save Values
	f := store.NewFile(jasmin.Name)
	f.Save(jasmin.GetYesterdaySum(qs))
}

func getTableData(w http.ResponseWriter, r *http.Request) {
	//need to get all quotes for the current value
	qs := yahoo.GetAllQuotes(jasmin.Stocks())

	//Load Historical Data from File
	f := store.NewFile(jasmin.Name)
	a, err := f.Load()
	if err != nil {
		fmt.Println("Error ", err)
	}

	//add current value
	d := store.Data{TimeHuman: time.Now(), Value: jasmin.GetTodaySum(qs)}
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
