package portfolio

import (
	"context"
	"fmt"

	"google.golang.org/appengine/datastore"
)

//StockValue structure
type StockValue struct {
	Stock
	Count float32
	Buy   float32
}

//StockValueStorage structure
type StockValueStorage struct {
	Owner  string
	Name   string
	Symbol string
	Count  float32
	Buy    float32
}

//Owner ...
type Owner struct {
	Name      string
	PortFolio []StockValue
}

// Stocks ...
func (o *Owner) Stocks(c context.Context) (s []Stock, err error) {
	err = o.loadPortfolioFromDB(c)
	if err != nil {
		return s, err
	}

	for _, sv := range o.PortFolio {
		s = append(s, sv.Stock)
	}
	return s, nil
}

func (o *Owner) loadPortfolioFromDB(c context.Context) error {
	// get all stocks for this owner
	query := datastore.NewQuery("StockValueStorage").Filter("Owner =", o.Name)

	data := []StockValueStorage{}
	_, err := query.GetAll(c, &data)
	if err != nil {
		return fmt.Errorf("Error during GetAll %v", err)
	}

	sv := make([]StockValue, len(data))
	for i, d := range data {
		sv[i] = StockValue{
			Stock: Stock{Name: d.Name, Symbol: d.Symbol},
			Count: d.Count,
			Buy:   d.Buy,
		}
	}

	o.PortFolio = sv

	return nil
}

//BuySum gets the sum of spended money
func (o Owner) BuySum() float32 {
	var sum float32
	for _, p := range o.PortFolio {
		sum += p.Buy * p.Count
	}
	return sum
}

//GetTodaySum calculates the sum for today
func (o Owner) GetTodaySum(qs Quotes) (sum float32) {
	for _, sv := range o.PortFolio {
		sum += sv.Count * qs.FindQuote(sv.Stock).Price
	}
	return sum
}

//RegularMarketTime gets the latest
func (o Owner) RegularMarketTime(qs Quotes) (t int64) {
	for _, q := range qs {
		if q.Time > t {
			t = q.Time
		}
	}
	return t
}
