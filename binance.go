package binance

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Ping tests connection to the Rest API
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
	r := new(rawOrderBook)
	p := params{"symbol": symbol, "limit": limit}
	if err := fetch(addrOrderBook, p, r); err != nil {
		return nil, err
	}
	return parseOrderBook(r)
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

// GetAggregateTrades get compressed, aggregate trades.
// Trades that fill at the time, from the same order,
// with the same price will have the quantity aggregated.
func GetAggregateTrades(symbol, limit, fromID, startTime, endTime string) (list []AggregateTrade, err error) {
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

// OpenAggregateTradeStream opens websocket with trade information that is aggregated for a single taker order.
func OpenAggregateTradeStream(symbol string) (*AggregateTradeStream, error) {
	ws, err := connectWebsocket(strings.ToLower(symbol) + "@aggTrade")
	if err != nil {
		return nil, err
	}
	return &AggregateTradeStream{stream{ws}}, nil
}

// OpenTradeStream opens websocket with raw trade information; each trade has a unique buyer and seller.
func OpenTradeStream(symbol string) (*TradeStream, error) {
	ws, err := connectWebsocket(strings.ToLower(symbol) + "@trade")
	if err != nil {
		return nil, err
	}
	return &TradeStream{stream{ws}}, nil
}

// OpenChartStream pushes trade information that is aggregated for a single taker order.
func OpenChartStream(symbol string, interval ChartInterval) (*ChartStream, error) {
	ws, err := connectWebsocket(fmt.Sprintf("%s@kline_%s", strings.ToLower(symbol), interval))
	if err != nil {
		return nil, err
	}
	return &ChartStream{stream{ws}}, nil
}

// OpenTickerStream pushes trade information that is aggregated for a single taker order.
func OpenTickerStream(symbol string) (*TickerStream, error) {
	ws, err := connectWebsocket(strings.ToLower(symbol) + "@ticker")
	if err != nil {
		return nil, err
	}
	return &TickerStream{stream{ws}}, nil
}

// OpenTickersStream pushes trade information that is aggregated for a single taker order.
func OpenTickersStream() (*TickersStream, error) {
	ws, err := connectWebsocket("!ticker@arr")
	if err != nil {
		return nil, err
	}
	return &TickersStream{stream{ws}}, nil
}

// OpenPartialBookStream pushes trade information that is aggregated for a single taker order.
func OpenPartialBookStream(symbol, level string) (*PartialBookStream, error) {
	ws, err := connectWebsocket(fmt.Sprintf("%s@depth%s", strings.ToLower(symbol), level))
	if err != nil {
		return nil, err
	}
	return &PartialBookStream{stream{ws}}, nil
}

// OpenDiffDepthStream pushes trade information that is aggregated for a single taker order.
func OpenDiffDepthStream(symbol string) (*DiffDepthStream, error) {
	ws, err := connectWebsocket(strings.ToLower(symbol) + "@depth")
	if err != nil {
		return nil, err
	}
	return &DiffDepthStream{stream{ws}}, nil
}
