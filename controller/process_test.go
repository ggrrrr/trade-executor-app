package controller

import (
	"testing"

	"github.com/ggrrrr/trade-executor-app/binance/models"
	"github.com/shopspring/decimal"
)

func TestLoop(t *testing.T) {
	order1 := MarketOrder{Symbol: "asd", Price: decimal.NewFromInt(5), Qty: decimal.NewFromInt(3)}
	config(&order1)

	process(&models.WsBookData{Symbol: "asd", Price: decimal.NewFromInt(5), Qty: decimal.NewFromInt(1)})
	t.Logf("order q: %v", order.Qty)
	if order.Qty.Cmp(decimal.NewFromInt(2)) != 0 {
		t.Errorf("order q is not 2: %v", order.Qty)
	}
	process(&models.WsBookData{Symbol: "ASD", Price: decimal.NewFromInt(10), Qty: decimal.NewFromInt(1)})
	t.Logf("order q: %v", order.Qty)
	if order.Qty.Cmp(decimal.NewFromInt(2)) != 0 {
		t.Errorf("order q is not 2: %v", order.Qty)
	}
	process(&models.WsBookData{Symbol: "ASDa", Price: decimal.NewFromInt(10), Qty: decimal.NewFromInt(1)})
	t.Logf("order q: %v", order.Qty)
	if order.Qty.Cmp(decimal.NewFromInt(2)) != 0 {
		t.Errorf("order q is not 2: %v", order.Qty)
	}
	process(&models.WsBookData{Symbol: "ASD", Price: decimal.NewFromInt(5), Qty: decimal.NewFromInt(10)})
	t.Logf("order q: %v", order.Qty)
	if order.Qty.Cmp(decimal.NewFromInt(0)) != 0 {
		t.Errorf("order q is not 0: %v", order.Qty)
	}

}

func TestCalc(t *testing.T) {
	book1 := models.WsBookData{Symbol: "ASD", Price: decimal.NewFromInt(5), Qty: decimal.RequireFromString("1")}
	book2 := models.WsBookData{Symbol: "ASD", Price: decimal.NewFromInt(5), Qty: decimal.RequireFromString("10")}
	book3 := models.WsBookData{Symbol: "ASD", Price: decimal.NewFromInt(5), Qty: decimal.RequireFromString("3")}

	order = &MarketOrder{Symbol: "ASD", Price: decimal.NewFromInt(5), Qty: decimal.NewFromInt(3)}

	calc1 := calc(&book1)
	if calc1.Cmp(decimal.NewFromInt(1)) != 0 {
		t.Errorf("calc not 1 %v", calc1)
	}
	calc2 := calc(&book2)
	if calc2.Cmp(decimal.NewFromInt(3)) != 0 {
		t.Errorf("calc not 3 %v", calc2)
	}
	calc3 := calc(&book3)
	if calc3.Cmp(decimal.NewFromInt(3)) != 0 {
		t.Errorf("calc not 3 %v", calc3)
	}

}
