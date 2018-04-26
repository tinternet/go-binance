package binance

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// Ping tests connection to the Binance server
func Ping() error {
	return fetch(addrPing, nil, nil)
}

// GetServerTime gets Binance server time
func GetServerTime() (t *time.Time, err error) {
	var reply struct{ ServerTime int64 }
	if err := fetch(addrServerTime, nil, &reply); err != nil {
		return nil, err
	}
	tval := time.Unix(reply.ServerTime, 0)
	return &tval, nil
}

// GetExchangeInfo gets trade information for all symbols
func GetExchangeInfo() (info *ExchangeInfo, err error) {
	err = fetch(addrExchangeInfo, nil, &info)
	return
}

// GetOrderBook gets orders for given symbol.
// Weight is adjusted based on the limit where
// [Limit 5, 10, 20, 50, 100] = [Weight 1];
// [Limit 500] = [Weight 5];
// [Limit 1000] = [Weight 10]
func GetOrderBook(symbol, limit string) (*OrderBook, error) {
	var reply struct {
		LastUpdateID uint64          `json:"lastUpdateId"`
		Bids         [][]interface{} `json:"bids"`
		Asks         [][]interface{} `json:"asks"`
	}

	p := params{"symbol": symbol, "limit": limit}

	if err := fetch(addrOrderBook, p, &reply); err != nil {
		return nil, err
	}

	fmt.Println(reply)

	parseOrder := func(v [][]interface{}) ([]Order, error) {
		orders := make([]Order, len(v))
		for i, bid := range v {
			var err error
			var order Order
			if order.Price, err = strconv.ParseFloat(bid[0].(string), 64); err != nil {
				return nil, err
			}
			if order.Quantity, err = strconv.ParseFloat(bid[1].(string), 64); err != nil {
				return nil, err
			}
			orders[i] = order
		}
		return orders, nil
	}

	var bids []Order
	var asks []Order
	var err error

	if bids, err = parseOrder(reply.Bids); err != nil {
		return nil, err
	}
	if asks, err = parseOrder(reply.Asks); err != nil {
		return nil, err
	}

	book := new(OrderBook)
	book.LastUpdateID = reply.LastUpdateID
	book.Bids = bids
	book.Asks = asks
	return book, nil
}

// GetRecentTrades gets up to 500 trades
func GetRecentTrades(symbol, limit string) (list []Trade, err error) {
	p := params{"symbol": symbol, "limit": limit}
	err = fetch(addrRecentTradesList, p, &list)
	return
}

// GetOldTrades gets up to 500 older trades
func GetOldTrades(symbol, limit, fromID string) (list []Trade, err error) {
	p := params{"symbol": symbol, "limit": limit, "fromId": fromID}
	err = fetch(addrOldTradeLookup, p, &list)
	return
}

// GetAggregatedTrades get compressed, aggregate trades.
// Trades that fill at the time, from the same order,
// with the same price will have the quantity aggregated.
func GetAggregatedTrades(symbol, limit, fromID, startTime, endTime string) (list []AggregatedTrade, err error) {
	p := params{"symbol": symbol, "limit": limit, "fromId": fromID, "startTime": startTime, "endTime": endTime}
	err = fetch(addrOldTradeLookup, p, &list)
	return
}

// GetKlines gets lline/candlestick bars for a symbol.
func GetKlines(symbol, interval, limit, startTime, endTime string) ([]Kline, error) {
	var reply [][]interface{}
	p := params{"symbol": symbol, "interval": interval, "limit": limit, "startTime": startTime, "endTime": endTime}
	err := fetch(addrKlineCandlestickData, p, &reply)

	if err != nil {
		return nil, err
	}

	list := make([]Kline, len(reply))
	for i, v := range reply {
		var kline Kline
		var err error

		t, ok := v[0].(float64)
		if !ok {
			return nil, errors.New("Invalid open time")
		}
		kline.OpenTime = time.Unix(int64(t), 0)

		if kline.Open, err = strconv.ParseFloat(v[1].(string), 64); err != nil {
			return nil, err
		}
		if kline.High, err = strconv.ParseFloat(v[2].(string), 64); err != nil {
			return nil, err
		}
		if kline.Low, err = strconv.ParseFloat(v[3].(string), 64); err != nil {
			return nil, err
		}
		if kline.Close, err = strconv.ParseFloat(v[4].(string), 64); err != nil {
			return nil, err
		}
		if kline.Volume, err = strconv.ParseFloat(v[5].(string), 64); err != nil {
			return nil, err
		}

		t, ok = v[6].(float64)
		if !ok {
			return nil, errors.New("Invalid close time")
		}
		kline.CloseTime = time.Unix(int64(t), 0)

		if kline.QuoteAssetVolume, err = strconv.ParseFloat(v[7].(string), 64); err != nil {
			return nil, err
		}

		c, ok := v[8].(float64)
		if !ok {
			return nil, errors.New("Invalid trades count")
		}
		kline.TradesCount = int(c)

		if kline.TakerBuyBaseAssetVol, err = strconv.ParseFloat(v[9].(string), 64); err != nil {
			return nil, err
		}
		if kline.TakerBuyQuoteAssetVol, err = strconv.ParseFloat(v[10].(string), 64); err != nil {
			return nil, err
		}
		list[i] = kline
	}

	return list, err
}

// GetTicker returns 24hr statistics for symbol
func GetTicker(symbol string) (t *Ticker, err error) {
	if symbol == "" {
		return nil, errors.New("Empty symbol")
	}
	p := params{"symbol": symbol}
	err = fetch(addrExchangeData24H, p, &t)
	return
}

// GetTickers gets tickers for all symbols
func GetTickers() (list []Ticker, err error) {
	err = fetch(addrExchangeData24H, nil, &list)
	return
}

// GetPrice gets latest price for a symbol
func GetPrice(symbol string) (price *Price, err error) {
	p := params{"symbol": symbol}
	err = fetch(addrSymbolPriceTicker, p, &price)
	return
}

// GetPrices gets latest price for all symbols
func GetPrices() (list []Price, err error) {
	err = fetch(addrSymbolPriceTicker, nil, &list)
	return
}

// GetOrderBookTicker gets best price/qty on the order book for a symbol
func GetOrderBookTicker(symbol string) (t *OrderBookTicker, err error) {
	p := params{"symbol": symbol}
	err = fetch(addrBookTicker, p, &t)
	return
}

// GetOrderBookTickers gets best price/qty on the order book for all symbols
func GetOrderBookTickers() (list []OrderBookTicker, err error) {
	err = fetch(addrBookTicker, nil, &list)
	return
}
