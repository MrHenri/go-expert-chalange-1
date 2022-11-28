package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/MrHenri/go-expert-chalange-1/model"
	_ "github.com/mattn/go-sqlite3"
)

func InsertDolarQuotation(quote *model.Usdbrl) error {
	db, err := sql.Open("sqlite3", "database/quotes.db")
	if err != nil {
		return err
	}
	defer db.Close()
	sqlStmt := `CREATE TABLE IF NOT EXISTS quotes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		code TEXT,
		codein TEXT,
		name TEXT,
		high TEXT,
		low TEXT,
		varBid TEXT,
		pctChange TEXT,
		bid TEXT,
		ask TEXT,
		timestamp TIMESTAMP,
		create_date TIMESTAMP
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("INSERT INTO quotes (code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, create_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	_, err = stmt.ExecContext(ctx, quote.Code, quote.Codein, quote.Name, quote.High, quote.Low, quote.VarBid, quote.PctChange, quote.Bid, quote.Ask, quote.Timestamp, quote.CreateDate)
	return err
}
