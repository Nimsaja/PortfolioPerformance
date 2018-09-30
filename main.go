package main

import (
	"fmt"
	"time"

	"github.com/Nimsaja/PortfolioPerformance/data"
	"github.com/Nimsaja/PortfolioPerformance/store"
	"github.com/Nimsaja/PortfolioPerformance/yahoo"
)

func main() {
	jasmin := data.Jasmin()

	start := time.Now()
	qs := yahoo.GetAllQuotes(jasmin.Stocks())
	fmt.Println("Elapsed time: ", time.Now().Sub(start))

	fmt.Println("Quotes Today/Yesterday: ", jasmin.GetTodaySum(qs), jasmin.GetYesterdaySum(qs))
	fmt.Println("Diffs Today/Yesterday: ", jasmin.GetTodaySum(qs)-jasmin.BuySum(), jasmin.GetYesterdaySum(qs)-jasmin.BuySum())

	//Save Values
	f := store.NewFile(jasmin.Name)
	f.Save(jasmin.GetYesterdaySum(qs))

	//Load Values
	a, err := f.Load()
	if err != nil {
		fmt.Println("Error ", err)
	}
	for _, d := range a {
		fmt.Println(d.TimeHuman, ": ", d.Value)
	}
}
