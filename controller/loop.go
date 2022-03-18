package controller

import "github.com/ggrrrr/trade-executor-app/binance/models"

func Loop(order *MarketOrder, pd chan *models.WsBookData) {
	config(order)
	for {
		bookData := <-pd
		ok := process(bookData)
		if ok {
			return
		}
	}
}
