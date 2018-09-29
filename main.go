package main

import (
	"fmt"

	"github.com/Nimsaja/Portfolio/portfolio"
)

func main() {
	google := portfolio.Stock{Name: "Google", Symbol: "ABEC.DE"}
	amazon := portfolio.Stock{Name: "Amazon", Symbol: "AMZ.DE"}

	jasmin := portfolio.Owner{Name: "Jasmin"}
	jasmin.PortFolio = []portfolio.StockValue{
		portfolio.StockValue{Stock: google, Count: 2, Buy: 2000},
		portfolio.StockValue{Stock: amazon, Count: 7, Buy: 5000},
	}

	sum := jasmin.SumBuy()
	fmt.Println("Jasmin: ", sum)

	mario := portfolio.Owner{Name: "Mario"}
	fmt.Println("Mario ", mario.SumBuy())
}
