package yahoo

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Nimsaja/PortfolioPerformance/portfolio"
	"google.golang.org/appengine/urlfetch"
)

// createClientFunc create a http.Client (in the cloud urlfetch.Client)
type createClientFunc func(c context.Context) *http.Client

// URLService ...
type URLService struct {
	client createClientFunc
}

//New factory to create URLService
func New(inCloud bool) URLService {
	var s URLService
	if inCloud {
		s = URLService{client: func(c context.Context) *http.Client {
			return urlfetch.Client(c)
		}}
	} else {
		s = URLService{client: func(_ context.Context) *http.Client {
			return http.DefaultClient
		}}
	}
	return s
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
