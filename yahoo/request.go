package yahoo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Nimsaja/Portfolio/portfolio"
)

var url = "https://query1.finance.yahoo.com/v7/finance/quote?lang=en-US&region=US&corsDomain=finance.yahoo.com&symbols="

//GetQuote ...
func GetQuote(s portfolio.Stock) (portfolio.Quote, error) {
	r, err := getQuote(s.Symbol)

	return portfolio.Quote{Stock: s, Close: r.Close, Price: r.Price}, err
}

//GetAllQuotes ...
func GetAllQuotes(sl []portfolio.Stock) []portfolio.Quote {
	ql := portfolio.New(sl)

	for _, s := range sl {
		go func(s portfolio.Stock) {
			defer ql.Done()

			q, err := GetQuote(s)
			if err != nil {
				log.Printf("could not get quote for %v, %v", s.Name, err)
			}
			ql.Add(q)
		}(s)
	}

	//here we wait for all the go routines to be done
	return ql.Wait()
}

func getQuote(s string) (Result, error) {
	var result Result

	u := fmt.Sprintf(url+"%v", s)

	resp, err := http.Get(u)
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
