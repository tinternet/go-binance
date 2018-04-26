package binance

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

const (
	addrPing                 = "https://api.binance.com/api/v1/ping"
	addrServerTime           = "https://api.binance.com/api/v1/time"
	addrExchangeInfo         = "https://api.binance.com/api/v1/exchangeInfo"
	addrOrderBook            = "https://api.binance.com/api/v1/depth"
	addrRecentTradesList     = "https://api.binance.com/api/v1/trades"
	addrOldTradeLookup       = "https://api.binance.com/api/v1/historicalTrades"
	addrAggregatedTrades     = "https://api.binance.com/api/v1/aggTrades"
	addrKlineCandlestickData = "https://api.binance.com/api/v1/klines"
	addrExchangeData24H      = "https://api.binance.com/api/v1/ticker/24hr"
	addrSymbolPriceTicker    = "https://api.binance.com/api/v3/ticker/price"
	addrBookTicker           = "https://api.binance.com/api/v3/ticker/bookTicker"
)

type params map[string]string

func encodeQuery(u *url.URL, p params) {
	q := u.Query()
	for k, v := range p {
		if v != "" {
			q.Add(k, v)
		}
	}
	u.RawQuery = q.Encode()
}

func detectError(res *http.Response) error {
	if res.StatusCode == 200 {
		return nil
	}
	var reply struct {
		Message string `json:"msg"`
	}
	if err := json.NewDecoder(res.Body).Decode(&reply); err != nil {
		return err
	}
	return errors.New(reply.Message)
}

func fetch(url string, p params, reply interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	if p != nil {
		encodeQuery(req.URL, p)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if err := detectError(res); err != nil {
		return err
	}
	return json.NewDecoder(res.Body).Decode(&reply)
}
