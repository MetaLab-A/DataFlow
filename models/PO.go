package models

import "database/sql"

type PO struct {
	RowOrder     string         `db:"RowOrder"`
	DocNo        string         `db:"DocNo"`
	DocDate      sql.NullString `db:"DocDate"`
	DocType      string         `db:"DocType"`
	RefNo        string         `db:"RefNo"`
	RefDate      sql.NullString `db:"RefDate"`
	PoNo         string         `db:"PoNo"`
	Status       string         `db:"Status"`
	TaxType      string         `db:"TaxType"`
	ApID         string         `db:"ApID"`
	ApName       string         `db:"ApName"`
	Credit       string         `db:"Credit"`
	DueDate      sql.NullString `db:"DueDate"`
	CompleteDate sql.NullString `db:"CompleteDate"`
	PayType      string         `db:"PayType"`
	TotalAmt     string         `db:"TotalAmt"`
	DiscountText string         `db:"DiscountText"`
	DiscountAmt  string         `db:"DiscountAmt"`
	BefoeTaxAmt  string         `db:"BefoeTaxAmt"`
	TaxAmt       string         `db:"TaxAmt"`
	TaxRate      string         `db:"TaxRate"`
	NetAmt       string         `db:"NetAmt"`
	RemainAmt    string         `db:"RemainAmt"`
	Description  string         `db:"Description"`
	CreateBy     string         `db:"CreateBy"`
	CreateDate   sql.NullString `db:"CreateDate"`
	CancelBy     sql.NullString `db:"CancelBy"`
	CancelDate   sql.NullString `db:"CancelDate"`
	AddID        string         `db:"AddID"`
	AddDate      sql.NullString `db:"AddDate"`
	EditID       string         `db:"EditID"`
	EditDate     sql.NullString `db:"EditDate"`
}
