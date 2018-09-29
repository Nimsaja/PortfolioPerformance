package main

import (
	"fmt"

	"github.com/Nimsaja/Portfolio/store"
	"github.com/Nimsaja/Portfolio/yahoo"

	"github.com/Nimsaja/Portfolio/portfolio"
)

var (
	google    = portfolio.Stock{Name: "Google", Symbol: "ABEC.DE"}
	amazon    = portfolio.Stock{Name: "Amazon", Symbol: "AMZ.DE"}
	netflix   = portfolio.Stock{Name: "Netflix", Symbol: "NFC.DE"}
	siemens   = portfolio.Stock{Name: "Siemens", Symbol: "SIE.DE"}
	xING      = portfolio.Stock{Name: "XING", Symbol: "O1BC.F"}
	biotech   = portfolio.Stock{Name: "Biotech", Symbol: "DWWD.SG"}
	auto      = portfolio.Stock{Name: "Auto&Robotic", Symbol: "2B76.F"}
	tecDax    = portfolio.Stock{Name: "TecDax", Symbol: "EXS2.F"}
	oekoworld = portfolio.Stock{Name: "Oekoworld", Symbol: "OE7A.SG"}
)

var stockValue = []portfolio.StockValue{
	portfolio.StockValue{Stock: google, Count: 0.211, Buy: 1069.138},
	portfolio.StockValue{Stock: amazon, Count: 0.056, Buy: 1776.515},
	portfolio.StockValue{Stock: netflix, Count: 2, Buy: 224.25},
	portfolio.StockValue{Stock: siemens, Count: 5, Buy: 106.02},
	portfolio.StockValue{Stock: xING, Count: 2, Buy: 328.765},
	portfolio.StockValue{Stock: biotech, Count: 3, Buy: 195.693},
	portfolio.StockValue{Stock: auto, Count: 33, Buy: 7.051},
	portfolio.StockValue{Stock: tecDax, Count: 10, Buy: 25.01},
	portfolio.StockValue{Stock: oekoworld, Count: 0.523, Buy: 191.051},
}

func main() {
	jasmin := portfolio.Owner{Name: "Jasmin", PortFolio: stockValue}

	qs := yahoo.GetAllQuotes(jasmin.Stocks())
	fmt.Println("Quotes Today/Yesterday: ", jasmin.GetTodaySum(qs), jasmin.GetYesterdaySum(qs))
	fmt.Println("Diffs Today/Yesterday: ", jasmin.GetTodaySum(qs)-jasmin.BuySum(), jasmin.GetYesterdaySum(qs)-jasmin.BuySum())

	f := store.NewFile(jasmin.Name)
	f.Save(jasmin.GetYesterdaySum(qs))
}
