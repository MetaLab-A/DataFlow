package metaapis

import (
	models "DataFlow/models"
	"fmt"
	"strconv"
)

func CalInvItem2RankingItem(store map[string]models.InvoiceItem) map[string]*models.RankingItem {
	repeatedID := make(map[string]*models.RankingItem)

	for _, s := range store {
		tempObj := repeatedID[s.ItemID]
		fTotal, _ := strconv.ParseFloat(s.TotalAmt, 64)
		fProfit, _ := strconv.ParseFloat(s.ProfitAmt, 64)
		fQty, _ := strconv.Atoi(s.Qty)
		fPrice, _ := strconv.ParseFloat(s.Price, 64)
		fCost, _ := strconv.ParseFloat(s.Cost, 64)
		fMargin, _ := strconv.ParseFloat(s.Margin, 64)

		// Create new data in map if it found first time
		if tempObj == nil {
			tempObj = &models.RankingItem{ItemID: s.ItemID, ItemName: s.ItemName}
		}

		// High - Low decision making
		tempObj.HighPrice = calNewHigh(fPrice, tempObj.HighPrice)
		tempObj.HighCost = calNewHigh(fCost, tempObj.HighCost)
		tempObj.HighMargin = calNewHigh(fMargin, tempObj.HighMargin)
		tempObj.LowPrice = calNewLow(fPrice, tempObj.LowPrice)
		tempObj.LowCost = calNewLow(fCost, tempObj.LowCost)
		tempObj.LowMargin = calNewLow(fMargin, tempObj.LowMargin)

		tempObj.TotalAmt += fTotal
		tempObj.ProfitAmt += fProfit
		tempObj.Qty += fQty
		repeatedID[s.ItemID] = tempObj
	}

	return repeatedID
}

func CalSOItem2RankingItem(store map[string]models.SOItem) map[string]*models.RankingSOItem {
	repeatedID := make(map[string]*models.RankingSOItem)

	for _, s := range store {
		tempObj := repeatedID[s.ItemID]
		fQty, _ := strconv.Atoi(s.Qty)
		fPrice, _ := strconv.ParseFloat(s.Price, 64)

		// Create new data in map if it found first time
		if tempObj == nil {
			tempObj = &models.RankingSOItem{ItemID: s.ItemID, ItemName: s.ItemName}
		}

		// High - Low decision making
		tempObj.HighPrice = calNewHigh(fPrice, tempObj.HighPrice)
		tempObj.LowPrice = calNewLow(fPrice, tempObj.LowPrice)

		tempObj.Qty += fQty
		repeatedID[s.ItemID] = tempObj
	}

	return repeatedID
}

func CalStockItem2RankingItem(store map[string]models.Stock) map[string]*models.RankingStock {
	repeatedID := make(map[string]*models.RankingStock)

	for _, s := range store {
		tempObj := repeatedID[s.ID]
		fQty, _ := strconv.Atoi(s.StockQty)
		fPrice, _ := strconv.ParseFloat(s.Price, 64)
		fCost, _ := strconv.ParseFloat(s.Cost, 64)
		fStockVal, _ := strconv.ParseFloat(s.StockValue, 64)

		// Create new data in map if it found first time
		if tempObj == nil {
			tempObj = &models.RankingStock{ItemID: s.ID, ItemName: s.Name}
		}

		tempObj.Price = fPrice
		tempObj.Cost = fCost
		tempObj.Qty = fQty
		tempObj.StockValue = fStockVal
		repeatedID[s.ID] = tempObj
	}

	return repeatedID
}

func calNewHigh(incoming float64, record float64) float64 {
	if record <= 0 || incoming > record {
		return incoming
	}
	return record
}

func calNewLow(incoming float64, record float64) float64 {
	if record <= 0 || incoming < record {
		return incoming
	}
	return record
}

func CalPrintRanking(store map[string]*models.RankingItem) {
	for k, v := range store {
		fmt.Println("=====", k, "=====")
		fmt.Println("Price(H, L)", v.HighPrice, v.LowPrice)
		fmt.Println("Cost(H, L)", v.HighCost, v.LowCost)
		fmt.Println("Margin(H, L)", v.HighMargin, v.LowMargin)
		fmt.Println("TotalAmt:", v.TotalAmt)
		fmt.Println("ProfitAmt:", v.ProfitAmt)
		fmt.Println("Qty:", v.Qty)
	}
}

func CalPrintSORanking(store map[string]*models.RankingSOItem) {
	for k, v := range store {
		fmt.Println("=====", k, "SO Item =====")
		fmt.Println("Price(H, L)", v.HighPrice, v.LowPrice)
		fmt.Println("Qty:", v.Qty)
	}
}

func CalPrintStockRanking(store map[string]*models.RankingStock) {
	for k, v := range store {
		fmt.Println("=====", k, "Stock Item =====")
		fmt.Println("Price", v.Price)
		fmt.Println("Cost", v.Cost)
		fmt.Println("Qty:", v.Qty)
		fmt.Println("StockValue:", v.StockValue)
	}
}