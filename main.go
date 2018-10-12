package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Nimsaja/PortfolioPerformance/data"
	"github.com/Nimsaja/PortfolioPerformance/plot"
	"github.com/Nimsaja/PortfolioPerformance/store"
	"github.com/Nimsaja/PortfolioPerformance/yahoo"
)

func main() {
	jasmin := data.Jasmin()

	start := time.Now()
	urlServ := yahoo.New(false)
	qs := urlServ.GetAllQuotes(context.TODO(), jasmin.Stocks())
	fmt.Println("Elapsed time: ", time.Now().Sub(start))

	fmt.Println("Quote/Diff/Time: ", jasmin.GetTodaySum(qs), jasmin.GetTodaySum(qs)-jasmin.BuySum(),
		time.Unix(jasmin.RegularMarketTime(qs), 0))
	fmt.Println("")

	//Save Values
	f := store.New(false, jasmin.Name)
	f.Save(context.TODO(), jasmin.GetTodaySum(qs), jasmin.BuySum(), jasmin.RegularMarketTime(qs))

	//Load Values
	a, err := f.Load(context.TODO())
	if err != nil {
		fmt.Println("Error ", err)
	}
	for _, d := range a {
		fmt.Println(d.TimeHuman, ": ", d.Value, ", ", d.Diff)
	}

	plot.Create(a, jasmin.GetTodaySum(qs), jasmin.BuySum())
}
