package models

type RankingStock struct {
	ItemID     string  `db:"ItemID"`
	ItemName   string  `db:"ItemName"`
	Price      float64 `db:"Price"`
	Cost       float64 `db:"Cost"`
	Qty        int     `db:"Qty"`
	StockValue float64 `db:"StockValue"`
}
