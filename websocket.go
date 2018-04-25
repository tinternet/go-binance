package binance

import (
	"net/url"

	"github.com/gorilla/websocket"
)

const binanceWsAddr = "stream.binance.com:9443"
const binanceWsScheme = "wss"

// Stream interface
type Stream interface {
	Read() ([]byte, error)
	Close() error
}

type stream struct {
	socket *websocket.Conn
}

// NewStream connects to binance websocket server
func NewStream(path string) (Stream, error) {
	socketURL := url.URL{Scheme: binanceWsScheme, Host: binanceWsAddr, Path: path}
	socket, _, err := websocket.DefaultDialer.Dial(socketURL.String(), nil)
	if err != nil {
		return nil, err
	}
	s := new(stream)
	s.socket = socket
	return s, nil
}

// Read function reads next websocket message
func (s stream) Read() ([]byte, error) {
	_, message, err := s.socket.ReadMessage()
	return message, err
}

// Close function closes underlying websocket connection
func (s stream) Close() error {
	return s.socket.Close()
}
