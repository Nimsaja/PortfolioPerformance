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
	Price float32
	Time  int64
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

// New stock value list (channel with stock values)
// cap: nb of stocks
func New(n int) *QuoteList {
	return &QuoteList{
		cap: n,
		ql:  make(chan Quote, n),
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

	if ql.cap == 0 {
		return res
	}

	for q := range ql.ql {
		res = append(res, q)
	}
	return res
}
