package models

type QtySummary struct {
	ItemID   string
	ItemName string
	VSSOQty  int
	RRPOQty  int
	StockQty int
	SOQty    int
	VSQty    int
	POQty    int
	RRQty    int
	TotalAmt float64
}
