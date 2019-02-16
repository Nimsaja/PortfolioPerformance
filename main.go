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

	//Save Values to Database
	d := store.New(false, jasmin.Name+"_DB")
	d.Save(context.TODO(), jasmin.GetTodaySum(qs), jasmin.BuySum(), jasmin.RegularMarketTime(qs))

	//Load Values
	a, err := f.Load(context.TODO())
	if err != nil {
		fmt.Println("Error ", err)
	}
	for _, d := range a {
		fmt.Println(d.TimeHuman, ": ", d.Value, ", ", d.Diff)
	}

	plot.Create(a)

	//show List a Stock prices
	fmt.Println("\n****************************")
	for _, p := range jasmin.PortFolio {
		buy := p.Count * p.Buy

		q, _ := urlServ.GetQuote(context.TODO(), p.Stock)

		price := p.Count * q.Price
		tab := "\t\t"
		if len(p.Name) > 6 {
			tab = "\t"
		}

		fmt.Printf("%s:%s Preis: %6.2f, \tWert: %.2f, \tDiff: %.2f \n", p.Name, tab, q.Price, price, price-buy)
	}

	fmt.Println("\n****************************")
	fmt.Println("Quote/Diff/Time: ", jasmin.GetTodaySum(qs), jasmin.GetTodaySum(qs)-jasmin.BuySum(),
		time.Unix(jasmin.RegularMarketTime(qs), 0))
	fmt.Println("")

	// //play with google datastore
	// ctx, client := store.NewDB()
	// store.SaveQuote(ctx, client, jasmin.GetTodaySum(qs), jasmin.GetTodaySum(qs)-jasmin.BuySum(), time.Unix(jasmin.RegularMarketTime(qs), 0))

}
