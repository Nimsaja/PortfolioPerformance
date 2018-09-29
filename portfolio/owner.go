package portfolio

//StockValue structure
type StockValue struct {
	Stock
	Count float32
	Buy   float32
}

//Owner ...
type Owner struct {
	Name      string
	PortFolio []StockValue
}

// Stocks ...
func (o Owner) Stocks() []Stock {
	s := make([]Stock, 0)
	for _, sv := range o.PortFolio {
		s = append(s, sv.Stock)
	}
	return s
}

//BuySum gets the sum of spended money
func (o Owner) BuySum() float32 {
	var sum float32
	for _, p := range o.PortFolio {
		sum += p.Buy * p.Count
	}
	return sum
}

//GetYesterdaySum calculates the sum for yesterday
func (o Owner) GetYesterdaySum(qs Quotes) (sum float32) {
	for _, sv := range o.PortFolio {
		sum += sv.Count * qs.FindQuote(sv.Stock).Close
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
