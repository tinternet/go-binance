package binance

import "time"

// RateLimit struct
type RateLimit struct {
	RateLimitType RateLimiterType   `json:"rateLimitType"`
	Interval      RateLimitInterval `json:"interval"`
	Limit         int               `json:"limit"`
}

// SymbolFilter struct
type SymbolFilter struct {
	FilterType  SymbolFilterType `json:"filterType"`
	MinNotional float64          `json:"minNotional,string,omitempty"`
	MinPrice    float64          `json:"minPrice,string,omitempty"`
	MaxPrice    float64          `json:"maxPrice,string,omitempty"`
	TickSize    float64          `json:"tickSize,string,omitempty"`
	MinQty      float64          `json:"minQty,string,omitempty"`
	MaxQty      float64          `json:"maxQty,string,omitempty"`
	StepSize    float64          `json:"stepSize,string,omitempty"`
}

// Symbol struct
type Symbol struct {
	Name               string         `json:"symbol"`
	Status             string         `json:"status"`
	BaseAsset          string         `json:"baseAsset"`
	BaseAssetPrecision int            `json:"baseAssetPrecision"`
	QuoteAsset         string         `json:"quoteAsset"`
	QuotePrecision     int            `json:"quotePrecision"`
	OrderTypes         []OrderType    `json:"orderTypes"`
	IcebergAllowed     bool           `json:"icebergAllowed"`
	Filters            []SymbolFilter `json:"filters"`
}

// ExchangeInfo represents current general cryptocurrency trade information
type ExchangeInfo struct {
	Timezone        string      `json:"timezone"`
	ServerTime      uint64      `json:"serverTime"`
	RateLimits      []RateLimit `json:"rateLimits"`
	ExchangeFilters interface{} `json:"exchangeFilters"` // Ignored
	Symbols         []Symbol    `json:"symbols"`
}

// Order struct
type Order struct {
	Price    float64
	Quantity float64
}

// OrderBook represents all orders for given Symbol
type OrderBook struct {
	Symbol       string
	LastUpdateID uint64
	Bids         []Order
	Asks         []Order
}

// Trade represents raw trade information
type Trade struct {
	ID           uint64 `json:"id"`
	Price        string `json:"price"`
	Quantity     string `json:"qty"`
	Time         uint64 `json:"time"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
	IsBestMatch  bool   `json:"isBestMatch"`
}

// TradeEvent struct
type TradeEvent struct {
	EventType     string  `json:"e"`
	EventTime     uint64  `json:"E"`
	Symbol        string  `json:"s"`
	TradeID       uint64  `json:"t"`
	Price         float64 `json:"p,string"`
	Quantity      float64 `json:"q,string"`
	BuyerOrderID  uint64  `json:"b"`
	SellerOrderID uint64  `json:"a"`
	TradeTime     int64   `json:"T"`
	IsBuyerMaker  bool    `json:"m"`
}

// AggregateTrade struct
type AggregateTrade struct {
	ID               uint64  `json:"a"`
	Price            float64 `json:"p"`
	Quantity         float64 `json:"q"`
	FirstTradeID     uint64  `json:"f"`
	LastTradeID      uint64  `json:"l"`
	Timestamp        uint64  `json:"T"`
	IsBuyerMaker     bool    `json:"m"`
	IsBestPriceMatch bool    `json:"M"`
}

// AggregateTradeEvent struct
type AggregateTradeEvent struct {
	EventType     string      `json:"e"`
	EventTime     uint64      `json:"E"`
	Symbol        string      `json:"s"`
	TradeID       uint64      `json:"t"`
	Price         float64     `json:"p,string"`
	Quantity      float64     `json:"q,string"`
	BuyerOrderID  uint64      `json:"b"`
	SellerOrderID uint64      `json:"a"`
	TradeTime     uint64      `json:"T"`
	IsMarketMaker bool        `json:"m"`
	Ignore        interface{} `json:"M"`
}

// Kline struct
type Kline struct {
	OpenTime              time.Time
	Open                  float64
	High                  float64
	Low                   float64
	Close                 float64
	Volume                float64
	CloseTime             time.Time
	QuoteAssetVolume      float64
	TradesCount           int
	TakerBuyBaseAssetVol  float64
	TakerBuyQuoteAssetVol float64
}

// ChartEvent represents updates to the current klines/candlestick
type ChartEvent struct {
	EventType string `json:"e"`
	EventTime uint64 `json:"E"`
	Symbol    string `json:"s"`
	Kline     struct {
		KlineStart               uint64 `json:"t"`
		KlineClose               uint64 `json:"T"`
		Symbol                   string `json:"s"`
		Interval                 string `json:"i"`
		FirstTradeID             uint64 `json:"f"`
		LastTradeID              uint64 `json:"L"`
		OpenPrice                string `json:"o"`
		ClosePrice               string `json:"c"`
		HighPrice                string `json:"h"`
		LowPrice                 string `json:"l"`
		BaseAssetVolume          string `json:"v"`
		NumberOfTrades           int    `json:"n"`
		IsClosed                 bool   `json:"x"`
		QuoteAssetVolume         string `json:"q"`
		TakerBuyBaseAssetVolume  string `json:"V"`
		TakerBuyQuoteAssetVolume string `json:"Q"`
	} `json:"k"`
}

// Ticker represents 24 hour price change statistics
type Ticker struct {
	Symbol             string  `json:"symbol"`
	PriceChange        float64 `json:"priceChange,string"`
	PriceChangePercent float64 `json:"priceChangePercent,string"`
	WeightedAvgPrice   float64 `json:"weightedAvgPrice,string"`
	PrevClosePrice     float64 `json:"prevClosePrice,string"`
	LastPrice          float64 `json:"lastPrice,string"`
	LastQty            float64 `json:"lastQty,string"`
	BidPrice           float64 `json:"bidPrice,string"`
	AskPrice           float64 `json:"askPrice,string"`
	OpenPrice          float64 `json:"openPrice,string"`
	HighPrice          float64 `json:"highPrice,string"`
	LowPrice           float64 `json:"lowPrice,string"`
	Volume             float64 `json:"volume,string"`
	QuoteVolume        float64 `json:"quoteVolume,string"`
	OpenTime           int64   `json:"openTime"`
	CloseTime          int64   `json:"closeTime"`
	FirstTradeID       int     `json:"firstId"`
	LastTradeID        int     `json:"lastId"`
	TradeCount         int     `json:"count"`
}

// TickerEvent represents ticker change event
type TickerEvent struct {
	EventType                string `json:"e"`
	EventTime                uint64 `json:"E"`
	Symbol                   string `json:"s"`
	PriceChange              string `json:"p"`
	PriceChangePercent       string `json:"P"`
	WeightedAvgPrice         string `json:"w"`
	PrevDayClosePrice        string `json:"x"`
	CurrDayClosePrice        string `json:"c"`
	CLoseTradeQuantity       string `json:"Q"`
	BestBidPrice             string `json:"b"`
	BidQuantity              string `json:"B"`
	BestAskPrice             string `json:"a"`
	BestAskQuantity          string `json:"A"`
	OpenPrice                string `json:"o"`
	ClosePrice               string `json:"h"`
	LowPrice                 string `json:"l"`
	TotalTradedBaseAssetVol  string `json:"v"`
	TotalTradedQuoteAssetVol string `json:"q"`
	StatOpenTime             uint64 `json:"O"`
	StatCloseTime            uint64 `json:"C"`
	FirstTradeID             uint64 `json:"F"`
	LastTradeID              uint64 `json:"L"`
	TotalTrades              int    `json:"n"`
}

// Price struct
type Price struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price,string"`
}

// OrderBookTicker represents best price/qty for a symbol
type OrderBookTicker struct {
	Symbol   string  `json:"symbol"`
	BidPrice float64 `json:"bidPrice,string"`
	BidQty   float64 `json:"bidQty,string"`
	AskPrice float64 `json:"askPrice,string"`
	AskQty   float64 `json:"askQty,string"`
}

// DiffDepth represents order book price and quantity depth updates used to locally manage an order book
type DiffDepth struct {
	EventType     string
	EventTime     uint64
	Symbol        string
	FirstUpdateID uint64
	FinalUpdateID uint64
	Bids          []Order
	Asks          []Order
}
