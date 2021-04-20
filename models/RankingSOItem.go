package models

type RankingSOItem struct {
	ItemID    string  `db:"ItemID"`
	ItemName  string  `db:"ItemName"`
	HighPrice float64 `db:"HighPrice"`
	LowPrice  float64 `db:"LowPrice"`
	Qty       int     `db:"Qty"`
}
