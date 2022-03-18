package db

import (
	"testing"

	"github.com/ggrrrr/trade-executor-app/binance/models"
)

func TestMain(t *testing.T) {
	t.Setenv("SQL_DRIVER", "sqlite3")
	t.Setenv("SQL_URL", "sqlite3.sqlite")

	r1 := models.WsBookData{Id: 123123}

	err := Config()
	if err != nil {
		t.Fatalf("%v", err)
	}
	err = InsertBookTrade(&r1)
	if err != nil {
		t.Errorf("%v", err)
	}

	found, err := SelectBookData(r1.Id)
	if err != nil {
		t.Errorf("%v", err)
	}
	if found.Id != r1.Id {
		t.Errorf("id %v %v", r1.Id, found.Id)
	}
}
