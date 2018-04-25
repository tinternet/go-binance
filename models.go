package binance

// RateLimit struct
type RateLimit struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	Limit         int    `json:"limit"`
}

// SymbolFilter struct
type SymbolFilter struct {
	FilterType  string  `json:"filterType"`
	MinNotional float64 `json:"minNotional,string,omitempty"`
	MinPrice    float64 `json:"minPrice,string,omitempty"`
	MaxPrice    float64 `json:"maxPrice,string,omitempty"`
	TickSize    float64 `json:"tickSize,string,omitempty"`
	MinQty      float64 `json:"minQty,string,omitempty"`
	MaxQty      float64 `json:"maxQty,string,omitempty"`
	StepSize    float64 `json:"stepSize,string,omitempty"`
}

// Symbol struct
type Symbol struct {
	Name               string         `json:"symbol"`
	Status             string         `json:"status"`
	BaseAsset          string         `json:"baseAsset"`
	BaseAssetPrecision int            `json:"baseAssetPrecision"`
	QuoteAsset         string         `json:"quoteAsset"`
	QuotePrecision     int            `json:"quotePrecision"`
	OrderTypes         []string       `json:"orderTypes"`
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

// AggregatedTrade struct
type AggregatedTrade struct {
	ID               uint64  `json:"a"`
	Price            float64 `json:"p"`
	Quantity         float64 `json:"q"`
	FirstTradeID     uint64  `json:"f"`
	LastTradeID      uint64  `json:"l"`
	Timestamp        uint64  `json:"T"`
	IsBuyerMaker     bool    `json:"m"`
	IsBestPriceMatch bool    `json:"M"`
}

// TODO

// TradeEvent represents trade change event
type TradeEvent struct {
	EventType     string      `json:"e"`
	EventTime     uint64      `json:"E"`
	Symbol        string      `json:"s"`
	TradeID       uint64      `json:"t"`
	Price         string      `json:"p"`
	Quantity      string      `json:"q"`
	BuyerOrderID  uint64      `json:"b"`
	SellerOrderID uint64      `json:"a"`
	TradeTime     uint64      `json:"T"`
	IsMarketMaker bool        `json:"m"`
	Ignore        interface{} `json:"M"`
}

// KindleCandlestick represents updates to the current klines/candlestick
type KindleCandlestick struct {
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
		Ignire                   string `json:"B"`
	} `json:"k"`
}

// TickerEvent represents ticker stream update
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

// Ticker struct
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

// PartialBook represents top <levels> bids and asks, pushed every second. Valid <levels> are 5, 10, or 20.
type PartialBook struct {
	LastUpdateID uint64          `json:"lastUpdateId"`
	Bids         [][]interface{} `json:"bids"`
	Asks         [][]interface{} `json:"asks"`
}

// DiffDepth represents order book price and quantity depth updates used to locally manage an order book
type DiffDepth struct {
	EventType     string          `json:"e"`
	EventTime     uint64          `json:"E"`
	Symbol        string          `json:"s"`
	FirstUpdateID uint64          `json:"U"`
	FinalUpdateID uint64          `json:"u"`
	Bids          [][]interface{} `json:"b"`
	Asks          [][]interface{} `json:"a"`
}
