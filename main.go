package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type coin struct {
	Symbol    string  `json:symbol`
	Value     float32 `json:value`
	MarketCap float32 `json:marketCap`
}

var coins = []coin{
	{Symbol: "BTC", Value: 100, MarketCap: 50000},
	{Symbol: "ETH", Value: 10, MarketCap: 500},
	{Symbol: "SOL", Value: 1, MarketCap: 5},
}

func getCoins(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, coins)
}

func addCoin(context *gin.Context) {
	var newCoin coin
	if err := context.BindJSON(&newCoin); err != nil {
		return
	}
	coins = append(coins, newCoin)
	context.IndentedJSON(http.StatusCreated, newCoin)
}

func main() {
	router := gin.Default()
	router.GET("/coins", getCoins)
	router.POST("/coins", addCoin)
	router.Run("localhost:9090")
}
