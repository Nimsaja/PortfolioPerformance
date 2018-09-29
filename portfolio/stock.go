package portfolio

//Stock structure
type Stock struct {
	Name   string
	Symbol string
}

//Quote structure
type Quote struct {
	Stock
	Close float32
	Price float32
}
