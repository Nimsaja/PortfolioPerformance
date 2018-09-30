package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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

	url := "/forcecall"
	router.HandleFunc(url, getHistData).Methods("GET")

	return corsAndOptionHandler(router)
}

var jasmin portfolio.Owner

func main() {
	jasmin = data.Jasmin()

	http.Handle("/", handler())

	fmt.Println("*******Open http://localhost:8080/forcecall*******")
	fmt.Println()
	appengine.Main()
}

func getHistData(w http.ResponseWriter, r *http.Request) {
	qs := yahoo.GetAllQuotes(jasmin.Stocks())

	//Save Values
	f := store.NewFile(jasmin.Name)
	f.Save(jasmin.GetYesterdaySum(qs))

	//Load Values
	a, err := f.Load()
	if err != nil {
		fmt.Println("Error ", err)
	}

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
