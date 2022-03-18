package models

import (
	"encoding/json"
	"fmt"

	"github.com/shopspring/decimal"
)

var (
	DEC_Z = decimal.NewFromInt(0)
)

type WsResponse struct {
	Result string `bson:"result"`
	Id     int    `bson:"id"`
	WsBookData
}

type WsBookData struct {
	Id       uint64          `json:"u" db:"u"`
	Symbol   string          `json:"s" db:"symbol"`
	Price    decimal.Decimal `json:"b" db:"price"`
	Qty      decimal.Decimal `json:"B" db:"qty"`
	AskPrice decimal.Decimal `json:"a" db:"ask_price"`
	AskQty   decimal.Decimal `json:"A" db:"ask_qty"`
}

func Parse(payload []byte) (*WsResponse, error) {
	var out WsResponse
	err := json.Unmarshal(payload, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func IsBookData(response *WsResponse) (*WsBookData, error) {
	if response.Symbol == "" {
		return nil, fmt.Errorf("symbol is nil")
	}
	if response.Price.Cmp(DEC_Z) == -1 {
		return nil, fmt.Errorf("price < 0")
	}
	if response.Qty.Cmp(DEC_Z) == -1 {
		return nil, fmt.Errorf("qty < 0")
	}
	return &response.WsBookData, nil
}
