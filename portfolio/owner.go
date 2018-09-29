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

//SumBuy gets the sum of spended money
func (o Owner) SumBuy() float32 {
	var sum float32
	for _, p := range o.PortFolio {
		sum += p.Buy * p.Count
	}
	return sum
}
