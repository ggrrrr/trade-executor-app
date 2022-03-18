package main

import (
	"fmt"
	"os"

	"github.com/ggrrrr/trade-executor-app/binance/models"
	"github.com/ggrrrr/trade-executor-app/binance/streams"
	"github.com/ggrrrr/trade-executor-app/controller"
	"github.com/ggrrrr/trade-executor-app/db"
	"github.com/ggrrrr/trade-executor-app/utils"
	"github.com/sirupsen/logrus"
)

var (
	err       error
	osSignals = make(chan os.Signal, 1)
)

func main() {
	fmt.Printf("main")

	err = db.Config()
	if err != nil {
		logrus.Panicf("db: %v", err)
	}

	order := controller.MarketOrder{
		Symbol: utils.GetString("order", "symbol"),
		Qty:    utils.GetDec("order", "qty", 0),
		Price:  utils.GetDec("order", "price", 0),
	}

	if order.Symbol == "" {
		logrus.Panicf("order symbol not set")
	}
	if order.Qty.IntPart() == 0 {
		logrus.Panicf("order qty not set")
	}
	if order.Price.IntPart() == 0 {
		logrus.Panicf("order price not set")
	}

	shutdown := make(chan struct{})
	bookData := make(chan *models.WsBookData)
	err = streams.Config(shutdown, &order)
	if err != nil {
		panic(err)
	}

	// go controller.Loop(bookData)

	// signal.Notify(osSignals, os.Interrupt)
	go streams.Start(bookData)

	s := models.NewBookSubscriptionRequest([]string{"btcusdt"})
	streams.Request(s)

	// logrus.Printf("os.signal: %v", <-osSignals)
	controller.Loop(&order, bookData)
	shutdown <- struct{}{}
	logrus.Printf("end.")

}
