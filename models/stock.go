package models

import "database/sql"

type Stock struct {
	RowOrder       string         `db:"RowOrder"`
	ID             string         `db:"ID"`
	Barcode        string         `db:"Barcode"`
	Status         string         `db:"Status"`
	IsShow         string         `db:"IsShow"`
	Name           string         `db:"Name"`
	Name2          string         `db:"Name2"`
	ShortName      sql.NullString `db:"ShortName"`
	StockID        string         `db:"StockID"`
	GroupID        string         `db:"GroupID"`
	UnitID         string         `db:"UnitID"`
	PurchaseUnitID sql.NullString `db:"PurchaseUnitID"`
	ApID           string         `db:"ApID"`
	Cost           string         `db:"Cost"`
	Price          string         `db:"Price"`
	StockQty       string         `db:"StockQty"`
	StockValue     string         `db:"StockValue"`
	StockMax       string         `db:"StockMax"`
	StockMin       string         `db:"StockMin"`
	ROP            string         `db:"ROP"`
	EOQ            sql.NullString `db:"EOQ"`
	SafetyQty      sql.NullString `db:"SafetyQty"`
	Weight         sql.NullString  `db:"Weight"`
	Thick          sql.NullString `db:"Thick"`
	LastSellPrice  string         `db:"LastSellPrice"`
	LastSellDate   sql.NullString `db:"LastsellDate"`
	LastBuyPrice   string         `db:"LastBuyPrice"`
	LastBuyDate    sql.NullString `db:"LastBuyDate"`
	Picture        string         `db:"Picture"`
	Description    string         `db:"Description"`
	AddID          string         `db:"AddID"`
	AddDate        sql.NullString `db:"AddDate"`
	EditID         string         `db:"EditID"`
	EditDate       sql.NullString `db:"EditDate"`
}
