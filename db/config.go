package db

import (
	"fmt"
	"strings"

	"database/sql"

	"github.com/sirupsen/logrus"

	"github.com/ggrrrr/trade-executor-app/binance/models"
	"github.com/ggrrrr/trade-executor-app/utils"
)

var (
	conn *sql.DB
)

func Config() error {
	driver := strings.TrimSpace(utils.GetString("sql", "driver"))
	url := strings.TrimSpace(utils.GetString("sql", "url"))

	switch driver {
	case "sqlite3":
		return connectSqlite3(url)
	default:
		return fmt.Errorf("unkown sql driver %v", driver)
	}
}

func Close() {
	if conn != nil {
		conn.Close()
	}
}

func InsertBookTrade(trade *models.WsBookData) error {
	if conn == nil {
		return fmt.Errorf("db.conn is nil")
	}
	sql := `INSERT INTO bookTrade(u, symbol, price, qty, ask_price, ask_qty) VALUES (?, ?, ?, ?, ?, ?)`
	statement, err := conn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = statement.Exec(trade.Id, trade.Symbol,
		trade.Price.String(),
		trade.Qty.String(),
		trade.AskPrice.String(),
		trade.AskQty.String(),
	)
	if err != nil {
		return err
	}
	return nil
}

func SelectBookData(u uint64) (*models.WsBookData, error) {
	sql := `select  u, price, qty from bookTrade where u = ?`
	row, err := conn.Query(sql, u)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		// asd, _ := row.Columns()
		var id uint64
		var price string
		var qty string
		err := row.Scan(&id, &price, &qty)
		if err != nil {
			return nil, err
		}
		out := models.WsBookData{
			Id:    id,
			Price: *utils.StoDec(price),
			Qty:   *utils.StoDec(qty),
		}
		return &out, nil
	}
	return nil, nil
}

func createTable() error {
	createStudentTableSQL := `CREATE TABLE bookTrade (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"u" integer NOT NULL,
		"symbol" TEXT,
		"price" DECIMAL,
		"qty" DECIMAL,
		"ask_price" DECIMAL,
		"ask_qty" DECIMAL
	  );`

	logrus.Infof("Create...")
	statement, err := conn.Prepare(createStudentTableSQL) // Prepare SQL Statement
	if err != nil {
		return err
	}
	_, err = statement.Exec() // Execute SQL Statements
	return err
}
