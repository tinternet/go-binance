package binance

import "github.com/gorilla/websocket"

type stream struct {
	socket *websocket.Conn
}

// Close function closes underlying websocket connection
func (s stream) Close() error {
	return s.socket.Close()
}

// AggregateTradeStream struct
type AggregateTradeStream struct {
	stream
}

func (s AggregateTradeStream) Read() (event *AggregateTradeEvent, err error) {
	err = s.socket.ReadJSON(&event)
	return
}

// TradeStream struct
type TradeStream struct {
	stream
}

func (s TradeStream) Read() (event *TradeEvent, err error) {
	err = s.socket.ReadJSON(&event)
	return
}

// ChartStream struct
type ChartStream struct {
	stream
}

func (s ChartStream) Read() (event ChartEvent, err error) {
	err = s.socket.ReadJSON(&event)
	return
}

// TickerStream struct
type TickerStream struct {
	stream
}

func (s TickerStream) Read() (event TickerEvent, err error) {
	err = s.socket.ReadJSON(&event)
	return
}

// TickersStream struct
type TickersStream struct {
	stream
}

func (s TickersStream) Read() (events []TickerEvent, err error) {
	err = s.socket.ReadJSON(&events)
	return
}

// PartialBookStream struct
type PartialBookStream struct {
	stream
}

func (s PartialBookStream) Read() (event *OrderBook, err error) {
	r := new(rawOrderBook)
	if err := s.socket.ReadJSON(r); err != nil {
		return nil, err
	}
	return parseOrderBook(r)
}

// DiffDepthStream struct
type DiffDepthStream struct {
	stream
}

func (s DiffDepthStream) Read() (event *DiffDepth, err error) {
	var rawBook struct {
		EventType     string          `json:"e"`
		EventTime     uint64          `json:"E"`
		Symbol        string          `json:"s"`
		FirstUpdateID uint64          `json:"U"`
		FinalUpdateID uint64          `json:"u"`
		Bids          [][]interface{} `json:"b"`
		Asks          [][]interface{} `json:"a"`
	}
	if err = s.socket.ReadJSON(&rawBook); err != nil {
		return
	}
	var bids []Order
	var asks []Order
	if bids, err = parseOrders(rawBook.Bids); err != nil {
		return nil, err
	}
	if asks, err = parseOrders(rawBook.Asks); err != nil {
		return nil, err
	}
	return &DiffDepth{
		EventType:     rawBook.EventType,
		EventTime:     rawBook.EventTime,
		Symbol:        rawBook.Symbol,
		FirstUpdateID: rawBook.FirstUpdateID,
		FinalUpdateID: rawBook.FinalUpdateID,
		Bids:          bids,
		Asks:          asks,
	}, nil
}
