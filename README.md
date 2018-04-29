# Go Binance API
go-binance is Binance public API implementation in golang

# Installation
```
go get github.com/bloc4ain/go-binance
```

# Examples
## REST API

### Test connectivity to the Rest API.
```golang
err := binance.Ping()
```

### Test connectivity to the Rest API and get the current server time.
```golang
time, err := binance.GetServerTime()
```

### Get current exchange trading rules and symbol information
```golang
info, err := binance.GetExchangeInfo()
```

### Get order book
```golang
book, err := binance.GetOrderBook("TRXBTC", "1000")
```

### Get recent trade list
```golang
trades, err := binance.GetRecentTrades("TRXBTC", "500")
```
### Old trade lookup
```golang
trades, err := binance.GetRecentTrades("TRXBTC", "500")
```

### Get compressed, aggregate trades
```golang
trades, err := binance.GetAggregateTrades("TRXBTC", "100", "", "", "")
```

### Query order status
Orders are assigned with order ID when issued and can later be queried using it
```golang
trades, err := binance.GetKlines("TRXBTC", "100", "", "", "")
```

## Streams examples
### Aggregate Trade Streams
```golang
// The Aggregate Trade Streams push trade information that is aggregated for a single taker order.
stream, err := binance.OpenAggregateTradeStream("TRXBTC")
if err != nil {
	fmt.Printf("Stream open error: %s\n", err)
	return
}
defer stream.Close()
for {
	update, err := stream.Read()
	if err != nil {
		fmt.Printf("Stream read error: %s\n", err)
		return
	}
	fmt.Printf("%+v\n", update)
}
```

### Trade Streams
```golang
// The Trade Streams push raw trade information; each trade has a unique buyer and seller.
stream, err := binance.OpenTradeStream("TRXBTC")
if err != nil {
	fmt.Printf("Stream open error: %s\n", err)
	return
}
defer stream.Close()
for {
	update, err := stream.Read()
	if err != nil {
		fmt.Printf("Stream read error: %s\n", err)
		return
	}
	fmt.Printf("%+v\n", update)
}
```

### Kline/Candlestick Streams
```golang
// The Kline/Candlestick Stream push updates to the current klines/candlestick every second.
stream, err := binance.OpenChartStream("TRXBTC", binance.ChartIntervalOneMin)
if err != nil {
	fmt.Printf("Stream open error: %s\n", err)
	return
}
defer stream.Close()
for {
	update, err := stream.Read()
	if err != nil {
		fmt.Printf("Stream read error: %s\n", err)
		return
	}
	fmt.Printf("%+v\n", update)
}
```

### Individual Symbol Ticker Streams
```golang
// 24hr Ticker statistics for a single symbol pushed every second
stream, err := binance.OpenTickerStream("TRXBTC")
if err != nil {
	fmt.Printf("Stream open error: %s\n", err)
	return
}
defer stream.Close()
for {
	update, err := stream.Read()
	if err != nil {
		fmt.Printf("Stream read error: %s\n", err)
		return
	}
	fmt.Printf("%+v\n", update)
}
```

### All Market Tickers Stream
```golang
// 24hr Ticker statistics for all symbols in an array pushed every second
stream, err := binance.OpenTickersStream()
if err != nil {
	fmt.Printf("Stream open error: %s\n", err)
	return
}
defer stream.Close()
for {
	update, err := stream.Read()
	if err != nil {
		fmt.Printf("Stream read error: %s\n", err)
		return
	}
	fmt.Printf("%+v\n", update)
}
```

### Partial Book Depth Streams
```golang
// Top <levels> bids and asks, pushed every second. Valid <levels> are 5, 10, or 20.
stream, err := binance.OpenPartialBookStream("TRXBTC", "20")
if err != nil {
	fmt.Printf("Stream open error: %s\n", err)
	return
}
defer stream.Close()
for {
	update, err := stream.Read()
	if err != nil {
		fmt.Printf("Stream read error: %s\n", err)
		return
	}
	fmt.Printf("%+v\n", update)
}
```

### Diff. Depth Stream
```golang
// Order book price and quantity depth updates used to locally manage an order book pushed every second.
stream, err := binance.OpenDiffDepthStream("TRXBTC")
if err != nil {
	fmt.Printf("Stream open error: %s\n", err)
	return
}
defer stream.Close()
for {
	update, err := stream.Read()
	if err != nil {
		fmt.Printf("Stream read error: %s\n", err)
		return
	}
	fmt.Printf("%+v\n", update)
}
```


## How to manage a local order book correctly
1. Open a stream using binance.OpenDiffDepthStream
2. Buffer the events you receive from the stream
3. Get a depth snapshot using binance.GetOrderBook
4. Drop any event where `FinalUpdateID` is <= `LastUpdateID` in the snapshot
5. The first processed should have `FirstUpdateID` <= `LastUpdateID`+1 **AND** `FinalUpdateID` >= `LastUpdateID`+1
6. While listening to the stream, each new event's `FirstUpdateID` should be equal to the previous event's `FinalUpdateID`+1
7. The data in each event is the **absolute** quantity for a price level
8. If the quantity is 0, **remove** the price level
9. Receiving an event that removes a price level that is not in your local order book can happen and is normal.

# License
This project is licensed under the [MIT License](http://opensource.org/licenses/MIT). See the [LICENSE](LICENSE) file for more info.

