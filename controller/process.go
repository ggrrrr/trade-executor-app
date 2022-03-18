package controller

import (
	"strings"

	"github.com/ggrrrr/trade-executor-app/binance/models"
	"github.com/ggrrrr/trade-executor-app/db"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type MarketOrder struct {
	Symbol       string          `json:"symbol"`
	Price        decimal.Decimal `json:"price"`
	Qty          decimal.Decimal `json:"qty"`
	transactions int
}

var (
	order *MarketOrder
)

func config(o *MarketOrder) {
	order = o
	order.transactions = 0
	logrus.Infof("s: %v p: %v: q: %v", order.Symbol, order.Price, order.Qty)
}

func add(pd *models.WsBookData) {
	calcQty := calc(pd)
	order.Qty = order.Qty.Sub(calcQty)
	pd.Qty = calcQty
	logrus.Infof("transation: %+v", pd)
	go db.InsertBookTrade(pd)
}

// Calc transaction qty
func calc(pd *models.WsBookData) decimal.Decimal {
	if order.Qty.Cmp(pd.Qty) == -1 {
		return order.Qty
	}
	return pd.Qty
}

// Process single trade return true if order is complete
func process(bookData *models.WsBookData) bool {
	logrus.Debugf("order.qty: %v bookData: %+v", order.Qty, bookData)
	if order.Symbol != strings.ToLower(bookData.Symbol) {
		return false
	}
	if bookData.Price.Cmp(order.Price) != 0 {
		return false
	}
	add(bookData)
	logrus.Debugf("order.qty: %v bookData: %+v", order.Qty, bookData)
	if order.Qty.Cmp(decimal.Zero) == 0 || order.Qty.Cmp(decimal.Zero) == -1 {
		return true
	}
	return false
}
