package yahoo

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Nimsaja/PortfolioPerformance/portfolio"
)

// CreateClientFunc create a http.Client (in the cloud urlfetch.Client)
type CreateClientFunc func(c context.Context) *http.Client

// defaultClientFunc is the default, if not run in the cloud
var defaultClientFunc = func(c context.Context) *http.Client {
	return http.DefaultClient
}

// URLService ...
type URLService struct {
	client CreateClientFunc
}

// New instace of YahooService
func New(client CreateClientFunc) URLService {
	return URLService{client: client}
}

// Default instance of URLService
func Default() URLService {
	return URLService{client: defaultClientFunc}
}

var url = "https://query1.finance.yahoo.com/v7/finance/quote?lang=en-US&region=US&corsDomain=finance.yahoo.com&symbols="

//GetQuote gets single quote
func (svc URLService) GetQuote(c context.Context, s portfolio.Stock) (portfolio.Quote, error) {
	r, err := getQuote(svc.client(c), s.Symbol)

	return portfolio.Quote{Stock: s, Close: r.Close, Price: r.Price}, err
}

//GetAllQuotes ...
func (svc URLService) GetAllQuotes(c context.Context, sl []portfolio.Stock) []portfolio.Quote {
	ql := portfolio.New(len(sl))

	for _, s := range sl {
		go func(s portfolio.Stock) {
			defer ql.Done()

			q, err := svc.GetQuote(c, s)
			if err != nil {
				log.Printf("could not get quote for %v, %v", s.Name, err)
			}
			ql.Add(q)
		}(s)
	}

	//here we wait for all the go routines to be done
	return ql.Wait()
}

func getQuote(client *http.Client, s string) (Result, error) {
	var result Result

	u := fmt.Sprintf(url+"%v", s)

	resp, err := client.Get(u)
	if err != nil {
		return result, fmt.Errorf("Error during http.Get(%v): %v", u, err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, fmt.Errorf("Error reading body: %v", err)
	}

	r, err := convertJSON2Result(body)
	if err != nil {
		return result, fmt.Errorf("Error during conversion from json to quote result. Stock: %v. Error: %v", s, err)
	}
	return r, nil
}

func convertJSON2Result(b []byte) (result Result, err error) {
	out := QuoteResponse{}
	err = json.Unmarshal(b, &out)
	if err != nil {
		log.Println(err.Error())
		return result, fmt.Errorf("Error during json unmarshalling")
	}

	if len(out.QR.Res) == 0 {
		return result, fmt.Errorf("Can not find quotes")
	}

	return out.QR.Res[0], nil
}
