package main

import (
	"CandleShopBackend/controllers"
	"CandleShopBackend/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase() // new

	r.GET("/candles", controllers.FindCandles)    // new
	r.POST("/candles", controllers.CreateCandle)  // new
	r.GET("/candles/:id", controllers.FindCandle) // new
	r.PATCH("/candles/:id", controllers.UpdateCandle)
	r.DELETE("/candles/:id", controllers.DeleteCandle)

	r.Run()
}
