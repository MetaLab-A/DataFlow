package models

import "database/sql"

type Stock struct {
	ID           string `db:"ID"`
	Name         string `db:"Name"`
	GroupID      string `db:"GroupID"`
	Cost         string	`db:"Cost"`
	Price        string	`db:"Price"`
	StockQty     string	`db:"StockQty"`
	StockValue   string	`db:"StockValue"`
	LastBuyDate  sql.NullString	`db:"LastBuyDate"`
	LastSellDate sql.NullString	`db:"LastsellDate"`
	EditDate     string	`db:"EditDate"`
}