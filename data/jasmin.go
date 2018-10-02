package data

import "github.com/Nimsaja/PortfolioPerformance/portfolio"

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
	portfolio.StockValue{Stock: google, Count: 0.329, Buy: 1065.028},
	portfolio.StockValue{Stock: amazon, Count: 0.113, Buy: 1777.077},
	portfolio.StockValue{Stock: netflix, Count: 2, Buy: 224.25},
	portfolio.StockValue{Stock: siemens, Count: 5, Buy: 106.02},
	portfolio.StockValue{Stock: xING, Count: 2, Buy: 328.765},
	portfolio.StockValue{Stock: biotech, Count: 3.707, Buy: 198.831},
	portfolio.StockValue{Stock: auto, Count: 33, Buy: 7.051},
	portfolio.StockValue{Stock: tecDax, Count: 10, Buy: 25.01},
	portfolio.StockValue{Stock: oekoworld, Count: 1.055, Buy: 189.631},
}

//Jasmin portfolio
func Jasmin() portfolio.Owner {
	return portfolio.Owner{Name: "Jasmin", PortFolio: stockValue}
}
