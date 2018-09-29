package portfolio

import (
	"testing"
)

var (
	google = Stock{Name: "Google", Symbol: "ABEC.DE"}
	amazon = Stock{Name: "Amazon", Symbol: "AMZ.DE"}
)

func owner() Owner {
	o := Owner{Name: "Test"}
	o.PortFolio = []StockValue{
		StockValue{Stock: google, Count: 0.4, Buy: 10.5},
		StockValue{Stock: amazon, Count: 7, Buy: 5000},
	}
	return o
}

func quotes() Quotes {
	return []Quote{
		Quote{Stock: google, Close: 100, Price: 110},
		Quote{Stock: amazon, Close: 1.5, Price: 1.4},
	}
}

func TestGetBuySum(t *testing.T) {
	s := owner().BuySum()

	if s != 35004.2 {
		t.Errorf("Expected a sum of %v, got %v", 35004.2, s)
	}
}
func TestGetYesterdaySum(t *testing.T) {
	s := owner().GetYesterdaySum(quotes())

	if s != 50.5 {
		t.Errorf("Expected a sum of %v, got %v", 50.5, s)
	}
}
func TestGetTodaySum(t *testing.T) {
	s := owner().GetTodaySum(quotes())

	if s != 53.8 {
		t.Errorf("Expected a sum of %v, got %v", 53.8, s)
	}
}
