package yahoo

import (
	"context"
	"testing"

	"github.com/Nimsaja/PortfolioPerformance/portfolio"
)

var svc = Default()
var c = context.TODO()

func TestGetQuote(t *testing.T) {
	s := portfolio.Stock{Name: "Google", Symbol: "ABEC.DE"}

	q, err := svc.GetQuote(c, s)

	if err != nil {
		t.Errorf("No error expected %v", err)
	}

	if q.Close == 0 {
		t.Errorf("Expected a close value different from zero, got %v", q.Close)
	}
}
func TestPrivateGetQuote(t *testing.T) {
	r, _ := getQuote(svc.client(c), "ABEC.DE")

	if r.Name != "Alphabet Inc." {
		t.Errorf("Name of stock should be %v. Got %v.", "Alphabet Inc.", r.Name)
	}
	if r.Cur != "EUR" {
		t.Errorf("Currency of stock should be %v. Got %v.", "EUR", r.Cur)
	}
}

func TestGetAllQuotes(t *testing.T) {
	google := portfolio.Stock{Name: "Google", Symbol: "ABEC.DE"}
	amazon := portfolio.Stock{Name: "Amazon", Symbol: "AMZ.DE"}

	ql := svc.GetAllQuotes(c, []portfolio.Stock{google, amazon})

	if len(ql) != 2 {
		t.Errorf("Expected a length of %v, got %v", 2, len(ql))
	}

	if (ql[0].Price * ql[1].Price) == 0 {
		t.Errorf("Expected both prices to be greater than zero. Got %v and %v.", ql[0].Price, ql[1].Price)
	}
	if (ql[0].Close * ql[1].Close) == 0 {
		t.Errorf("Expected both close values to be greater than zero. Got %v and %v.", ql[0].Close, ql[1].Close)
	}
}

func TestConvertJSON2ResultIfResultArrayIsEmpty(t *testing.T) {
	b := []byte(`{"quoteResponse":{"result":[],"error":null}}`)

	_, err := convertJSON2Result(b)
	if err == nil {
		t.Errorf("No error expected")
	}
}

func TestConvertJSON2ResultIfStockIsFound(t *testing.T) {
	b := []byte(`{"quoteResponse":{"result":
		[{"currency":"USD",
		"regularMarketPrice":1199.89,
		"regularMarketPreviousClose":1180.49,
		"longName":"Alphabet Inc.",
		"shortName":"Alphabet Inc.",
		"symbol":"GOOG"}],
		"error":null}}`)

	r, err := convertJSON2Result(b)
	if err != nil {
		t.Errorf("No error expected")
	}

	if r.Price != 1199.89 {
		t.Errorf("Expected %v, got %v", 1199.89, r.Price)
	}
}
