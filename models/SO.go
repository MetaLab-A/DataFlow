package models

import "database/sql"

type SO struct {
	DocNo			string	`db:"DocNo"`
	DocDate			sql.NullString	`db:"DocDate"`
	DocType			string	`db:"DocType"`
	RefNo			string	`db:"RefNo"`
	Status			string	`db:"Status"`
	TaxType			string	`db:"TaxType"`
	IsNotify		string	`db:"IsNotify"`
	ArID			string	`db:"ArID"`
	ArName			string	`db:"ArName"`
	AreaID			string	`db:"AreaID"`
	SalemanID		string	`db:"SalemanID"`
	TeamID			string	`db:"TeamID"`
	Credit			string	`db:"Credit"`
	DueDate			string	`db:"DueDate"`
	DeliveryDate	sql.NullString	`db:"DeliveryDate"`
	CompleteDate	sql.NullString	`db:"CompleteDate"`
	PayType			string	`db:"PayType"`
	TotalAmt		string	`db:"TotalAmt"`
	DiscountText	string	`db:"DiscountText"`
	DiscountAmt		string	`db:"DiscountAmt"`
	BefoeTaxAmt		string	`db:"BefoeTaxAmt"`
	TaxAmt			string	`db:"TaxAmt"`
	TaxRate			string	`db:"TaxRate"`
	NetAmt			string	`db:"NetAmt"`
	RemainAmt		string	`db:"RemainAmt"`
	Description		string	`db:"Description"`
	CreateBy		string	`db:"CreateBy"`
	CreateDate		sql.NullString	`db:"CreateDate"`
	CancelBy		string	`db:"CancelBy"`
	CancelDate		sql.NullString	`db:"CancelDate"`
	AddID			string	`db:"AddID"`
	AddDate			sql.NullString	`db:"AddDate"`
	EditID			string	`db:"EditID"`
	EditDate		sql.NullString	`db:"EditDate"`
}