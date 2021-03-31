package models

import "database/sql"

type Stock struct {
	ID				string	`db:"ID"`
	Status			string	`db:"Status"`
	IsShow			string	`db:"IsShow"`
	Name			string	`db:"Name"`
	Name2			string	`db:"Name2"`
	ShortName		string	`db:"ShortName"`
	StockID			string	`db:"StockID"`
	GroupID			string	`db:"GroupID"`
	UnitID			string	`db:"UnitID"`
	PurchaseUnitID	string	`db:"PurchaseUnitID"`
	ApID			string	`db:"ApID"`
	Cost			string	`db:"Cost"`
	Price			string	`db:"Price"`
	StockQty		string	`db:"StockQty"`
	StockValue		string	`db:"StockValue"`
	StockMax		string	`db:"StockMax"`
	StockMin		string	`db:"StockMin"`
	ROP				string	`db:"ROP"`
	EOQ				string	`db:"EOQ"`
	SafetyQty		string	`db:"SafetyQty"`
	Weight			string	`db:"Weight"`
	Thick			string	`db:"Thick"`
	LastSellPrice	string	`db:"LastSellPrice"`
	LastSellDate	sql.NullString	`db:"LastSellDate"`
	LastBuyPrice	string	`db:"LastBuyPrice"`
	LastBuyDate		sql.NullString	`db:"LastBuyDate"`
	Description		string	`db:"Description"`
	AddID			string	`db:"AddID"`
	AddDate			sql.NullString	`db:"AddDate"`
	EditID			string	`db:"EditID"`
	EditDate		sql.NullString	`db:"EditDate"`
}