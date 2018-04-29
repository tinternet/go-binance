package binance

import "strconv"

type rawOrderBook struct {
	LastUpdateID uint64          `json:"lastUpdateId"`
	Bids         [][]interface{} `json:"bids"`
	Asks         [][]interface{} `json:"asks"`
}

func parseOrders(v [][]interface{}) ([]Order, error) {
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

func parseOrderBook(raw *rawOrderBook) (*OrderBook, error) {
	var bids []Order
	var asks []Order
	var err error
	if bids, err = parseOrders(raw.Bids); err != nil {
		return nil, err
	}
	if asks, err = parseOrders(raw.Asks); err != nil {
		return nil, err
	}
	book := new(OrderBook)
	book.LastUpdateID = raw.LastUpdateID
	book.Bids = bids
	book.Asks = asks
	return book, nil
}
