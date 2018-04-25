package binance

import (
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
			var order Order
			var price, qty float64
			var err error
			if price, err = strconv.ParseFloat(bid[0].(string), 64); err != nil {
				return nil, err
			}
			if qty, err = strconv.ParseFloat(bid[1].(string), 64); err != nil {
				return nil, err
			}
			order.Price = price
			order.Quantity = qty
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
