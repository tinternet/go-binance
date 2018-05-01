package binance

// SymbolFilterType defines trading rules on a symbol.
type SymbolFilterType string

const (
	// SymbolFilterTypePrice defines the price rules for a symbol.
	SymbolFilterTypePrice = SymbolFilterType("PRICE_FILTER")

	// SymbolFilterTypeLotSize filter defines the quantity (aka "lots" in auction terms) rules for a symbol.
	SymbolFilterTypeLotSize = SymbolFilterType("LOT_SIZE")

	// SymbolFilterMinNotional filter defines the minimum notional value allowed for an order on a symbol.
	// An order's notional value is the price * quantity.
	SymbolFilterMinNotional = SymbolFilterType("MIN_NOTIONAL")

	// SymbolFilterMaxNumOrders filter defines the maximum number of orders an account is allowed to have open on a symbol.
	// Note that both "algo" orders and normal orders are counted for this filter.
	SymbolFilterMaxNumOrders = SymbolFilterType("MAX_NUM_ORDERS")

	// SymbolFilterMaxAlgoOrders filter defines the maximum number of "algo" orders an account is allowed to have open on a symbol.
	// "Algo" orders are STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, and TAKE_PROFIT_LIMIT orders.
	SymbolFilterMaxAlgoOrders = SymbolFilterType("MAX_ALGO_ORDERS")
)

// ExchangeFilterType defines trading rules on exchange.
type ExchangeFilterType string

const (
	// ExchangeFTMaxNumOrders filter defines the maximum number of orders an account is allowed to have open on the exchange.
	// Note that both "algo" orders and normal orders are counted for this filter.
	ExchangeFTMaxNumOrders = ExchangeFilterType("EXCHANGE_MAX_NUM_ORDERS")

	// ExchangeFTMaxAlgoOrders filter defines the maximum number of "algo" orders an account is allowed to have open on the exchange.
	// "Algo" orders are STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, and TAKE_PROFIT_LIMIT orders.
	ExchangeFTMaxAlgoOrders = ExchangeFilterType("EXCHANGE_MAX_ALGO_ORDERS")
)

// SymbolStatus represents symbol trading status
type SymbolStatus string

const (
	// SymbolStatusPreTrading represents symbol that will traded in the future.
	SymbolStatusPreTrading = SymbolStatus("PRE_TRADING")

	// SymbolStatusTrading represents symbol that can be currently traded.
	SymbolStatusTrading = SymbolStatus("TRADING")

	// SymbolStatusPostTrading represents POST_TRADING status
	SymbolStatusPostTrading = SymbolStatus("POST_TRADING")

	// SymbolStatusEndOfDay represents END_OF_DAY status
	SymbolStatusEndOfDay = SymbolStatus("END_OF_DAY")

	// SymbolStatusHalt represents HALT status
	SymbolStatusHalt = SymbolStatus("HALT")

	// SymbolStatusAuctionMatch represents AUCTION_MATCH status
	SymbolStatusAuctionMatch = SymbolStatus("AUCTION_MATCH")

	// SymbolStatusBreak represents BREAK status
	SymbolStatusBreak = SymbolStatus("BREAK")
)

// SymbolType string
type SymbolType string

// Spot symbol type
const Spot = SymbolStatus("SPOT")

// OrderStatus string
type OrderStatus string

// Order statuses
const (
	// OrderStatusNew represents NEW status
	OrderStatusNew = OrderStatus("NEW")

	// OrderStatusPartiallyFilled represents PARTIALLY_FILLED status
	OrderStatusPartiallyFilled = OrderStatus("PARTIALLY_FILLED")

	// OrderStatusFilled represents FILLED status
	OrderStatusFilled = OrderStatus("FILLED")

	// OrderStatusCanceled represents CANCELED
	OrderStatusCanceled = OrderStatus("CANCELED")

	// OrderStatusPendingCancel represents PENDING_CANCEL status
	OrderStatusPendingCancel = OrderStatus("PENDING_CANCEL")

	// OrderStatusRejected represents REJECTED status
	OrderStatusRejected = OrderStatus("REJECTED")

	// OrderStatusExpired represents EXPIRED status
	OrderStatusExpired = OrderStatus("EXPIRED")
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
	ChartIntervalOneMin     = ChartInterval("1m")
	ChartIntervalThreeMin   = ChartInterval("3m")
	ChartIntervalFiveMin    = ChartInterval("5m")
	ChartIntervalFifteenMin = ChartInterval("15m")
	ChartIntervalThirtyMin  = ChartInterval("30m")
	ChartIntervalOneHour    = ChartInterval("1h")
	ChartIntervalTwoHour    = ChartInterval("2h")
	ChartIntervalFourHour   = ChartInterval("4h")
	ChartIntervalSixHour    = ChartInterval("6h")
	ChartIntervalEightHour  = ChartInterval("8h")
	ChartIntervalTwelveHour = ChartInterval("12h")
	ChartIntervalOneDay     = ChartInterval("1d")
	ChartIntervalThreeDay   = ChartInterval("3d")
	ChartIntervalOneWeek    = ChartInterval("1w")
	ChartIntervalOneMonth   = ChartInterval("1M")
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
