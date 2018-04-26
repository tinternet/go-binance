package binance

// SymbolStatus represents symbol trading status
type SymbolStatus string

// Symbol trading statuses
const (
	PreTradingSymbol   = SymbolStatus("PRE_TRADING")
	TradingSymbol      = SymbolStatus("TRADING")
	PostTradingSymbol  = SymbolStatus("POST_TRADING")
	EndOfDaySymbol     = SymbolStatus("END_OF_DAY")
	HaltSymbol         = SymbolStatus("HALT")
	AuctionMatchSymbol = SymbolStatus("AUCTION_MATCH")
	BreakSymbol        = SymbolStatus("BREAK")
)

// SymbolType string
type SymbolType string

// Spot symbol type
const Spot = SymbolStatus("SPOT")

// OrderStatus string
type OrderStatus string

// Order statuses
const (
	NewOrder             = OrderStatus("NEW")
	PartiallyFilledOrder = OrderStatus("PARTIALLY_FILLED")
	FilledOrder          = OrderStatus("FILLED")
	CanceledOrder        = OrderStatus("CANCELED")
	PendingCancelOrder   = OrderStatus("PENDING_CANCEL")
	RejectedOrder        = OrderStatus("REJECTED")
	ExpiredOrder         = OrderStatus("EXPIRED")
)

// OrderType string
type OrderType string

// Order types
const (
	LimitOrder           = OrderType("LIMIT")
	MarketOrder          = OrderType("MARKET")
	StopLossOrder        = OrderType("STOP_LOSS")
	StopLossLimitOrder   = OrderType("STOP_LOSS_LIMIT")
	TakeProfitOrder      = OrderType("TAKE_PROFIT")
	TakeProfitLimitOrder = OrderType("TAKE_PROFIT_LIMIT")
	LimitMakerOrder      = OrderType("LIMIT_MAKER")
)

// OrderSide string
type OrderSide string

// Order sides
const (
	BuyOrder  = OrderSide("BUY")
	SellOrder = OrderSide("SELL")
)

// ChartInterval string
type ChartInterval string

// Chart intervals
const (
	OneMinInterval     = ChartInterval("1m")
	ThreeMinInterval   = ChartInterval("3m")
	FiveMinInterval    = ChartInterval("5m")
	FifteenMinInterval = ChartInterval("15m")
	ThirtyMinInterval  = ChartInterval("30m")
	OneHourInterval    = ChartInterval("1h")
	TwoHourInterval    = ChartInterval("2h")
	FourHourInterval   = ChartInterval("4h")
	SixHourInterval    = ChartInterval("6h")
	EightHourInterval  = ChartInterval("8h")
	TwelveHourInterval = ChartInterval("12h")
	OneDayInterval     = ChartInterval("1d")
	ThreeDayInterval   = ChartInterval("3d")
	OneWeekInterval    = ChartInterval("1w")
	OneMonthInterval   = ChartInterval("1M")
)

// RateLimiterType string
type RateLimiterType string

// Rate limiter types
const (
	RequestsRLType = RateLimiterType("REQUESTS")
	OrdersRLType   = RateLimiterType("ORDERS")
)

// RateLimitInterval string``
type RateLimitInterval string

// Rate limit intervals
const (
	SecondRLInterval = RateLimitInterval("SECOND")
	MinuteRLInterval = RateLimitInterval("MINUTE")
	DayRLInterval    = RateLimitInterval("DAY")
)
