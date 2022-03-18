package models

import (
	"fmt"
)

const (
	BOOK_TICKER = "bookTicker"
)

type E_METHODS string

var (
	METHOD_SUBSCRIBE          E_METHODS = "SUBSCRIBE"
	METHOD_UNSUBSCRIBE                  = "UNSUBSCRIBE"
	METHOD_LIST_SUBSCRIPTIONS           = "LIST_SUBSCRIPTIONS"
)

type WsRequest struct {
	Method string   `json:"method"`
	Params []string `json:"params"`
	Id     int32    `json:"id"`
}

func NewBookSubscriptionRequest(symbols []string) *WsRequest {

	params := []string{}

	for _, v := range symbols {
		params = append(params, fmt.Sprintf("%s@%s", v, BOOK_TICKER))
	}

	out := WsRequest{
		Method: string(METHOD_SUBSCRIBE),
		Params: params,
	}
	return &out
}
