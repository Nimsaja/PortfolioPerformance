package portfolio

import "sync"

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

//Quotes ...
type Quotes []Quote

//FindQuote finds the quote for this stock
func (qs Quotes) FindQuote(s Stock) Quote {
	for _, q := range qs {
		if s == q.Stock {
			return q
		}
	}
	return Quote{}
}

// QuoteList ...
type QuoteList struct {
	ql    chan Quote
	cap   int
	ready int
	//mutual exclusion - so that only on go routine can have access to QuoteList
	//block which should be mutual exclusive must be surrounded by lock and unlock
	sync.Mutex
}

// Yesterday value of stock for yesterday
// func (q Quote) Yesterday(count float32) float32 {
// 	return q.Close * count
// }

// Today value of stock for today
// func (q Quote) Today(count float32) float32 {
// 	return q.Price * count
// }

// New stock value list (channel with stock values)
// cap: nb of stocks
func New(sl []Stock) *QuoteList {
	return &QuoteList{
		cap: len(sl),
		ql:  make(chan Quote, len(sl)),
	}
}

// Add a stock value to channel
func (ql *QuoteList) Add(q Quote) {
	ql.ql <- q
}

// Done counter for get stock values are done
func (ql *QuoteList) Done() {
	ql.Lock()
	defer ql.Unlock()
	ql.ready++
	if ql.ready == ql.cap {
		close(ql.ql)
	}
}

// Wait read all stock values from channel
// and create sum today and yesterday
func (ql *QuoteList) Wait() []Quote {
	res := make([]Quote, 0)
	for q := range ql.ql {
		res = append(res, q)
	}
	return res
}
