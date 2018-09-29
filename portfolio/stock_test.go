package portfolio

import "testing"

var (
	a = Stock{Name: "AAA"}
	b = Stock{Name: "BBB"}
	c = Stock{Name: "CCC"}
	x = Stock{Name: "XXX"}
)

var qs Quotes = []Quote{
	Quote{Stock: a},
	Quote{Stock: b},
	Quote{Stock: c},
}

func TestFindQuoteSuccess(t *testing.T) {
	qb := qs.FindQuote(b)

	if qb.Name != "BBB" {
		t.Errorf("Expected name %v, got %v", "AAA", qb.Name)
	}
}
func TestFindQuoteFail(t *testing.T) {
	qx := qs.FindQuote(x)

	if qx.Name != "" {
		t.Errorf("Expected name %v, got %v", "", qx.Name)
	}
}

func TestNew(t *testing.T) {
	ql := New(3)

	if ql.cap != 3 {
		t.Errorf("Expected %v, got %v", 3, ql.cap)
	}
}

func TestAdd(t *testing.T) {
	ql := New(3)

	ql.Add(qs[0])
	if len(ql.ql) != 1 {
		t.Errorf("Expected length of %v, got %v", 1, len(ql.ql))
	}

	ql.Add(qs[1])
	if len(ql.ql) != 2 {
		t.Errorf("Expected length of %v, got %v", 2, len(ql.ql))
	}

	chanQ := <-ql.ql
	if chanQ.Stock.Name != "AAA" {
		t.Errorf("Expected Stock with name %v, got %v", "AAA", chanQ.Stock.Name)
	}

	chanQ = <-ql.ql
	if chanQ.Stock.Name != "BBB" {
		t.Errorf("Expected Stock with name %v, got %v", "BBB", chanQ.Stock.Name)
	}
}

func TestDoneAndWait(t *testing.T) {
	ql := New(2)

	go func() {
		ql.Add(qs[0])
		ql.Done()
	}()

	ql.Add(qs[1])
	ql.Done()

	quotes := ql.Wait()

	if len(quotes) != 2 {
		t.Errorf("Expected length of %v, got %v", 2, len(quotes))
	}
}
