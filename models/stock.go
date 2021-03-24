package models

type Stock struct {
	ID           string
	Name         string
	GroupID      string
	Cost         string
	Price        string
	StockQty     string
	StockValue   string
	LastBuyDate  sql.NullString
	LastSellDate sql.NullString
	EditDate     string
}