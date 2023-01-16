package controllers

import (
	"CandleShopBackend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateCandleInput struct {
	Picture       string `json:"Picture" binding:"required"`
	Title         string `json:"Title" binding:"required"`
	FullName      string `json:"FullName" binding:"required"`
	Description   string `json:"Description" binding:"required"`
	PriceMassive  string `json:"PriceMassive" binding:"required"`
	TypeFragrance string `json:"TypeFragrance" binding:"required"`
	TopNotes      string `json:"TopNotes" binding:"required"`
	MiddleNotes   string `json:"MiddleNotes" binding:"required"`
	BasicNotes    string `json:"BasicNotes" binding:"required"`
}

type UpdateCandleInput struct {
	Picture       string `json:"Picture"`
	Title         string `json:"Title"`
	FullName      string `json:"FullName"`
	Description   string `json:"Description"`
	PriceMassive  string `json:"PriceMassive"`
	TypeFragrance string `json:"TypeFragrance"`
	TopNotes      string `json:"TopNotes"`
	MiddleNotes   string `json:"MiddleNotes""`
	BasicNotes    string `json:"BasicNotes""`
}

// GET /Candles
// Get all Candels
func FindCandles(c *gin.Context) {
	var candles []models.Candle
	models.DB.Find(&candles)

	c.JSON(http.StatusOK, gin.H{"data": candles})
}

// POST /books
// Create new book
func CreateCandle(c *gin.Context) {
	// Validate input
	var input CreateCandleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	candle := models.Candle{Picture: input.Picture, FullName: input.FullName, Title: input.Title, Description: input.Description, PriceMassive: input.PriceMassive, TypeFragrance: input.TypeFragrance, TopNotes: input.TopNotes, MiddleNotes: input.MiddleNotes, BasicNotes: input.BasicNotes}
	models.DB.Create(&candle)

	c.JSON(http.StatusOK, gin.H{"data": candle})
}

// GET /books/:id
// Find a book
func FindCandle(c *gin.Context) { // Get model if exist
	var candle models.Candle

	if err := models.DB.Where("id = ?", c.Param("id")).First(&candle).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": candle})
}

// PATCH /books/:id
// Update a book
func UpdateCandle(c *gin.Context) {
	// Get model if exist
	var candle models.Candle
	if err := models.DB.Where("id = ?", c.Param("id")).First(&candle).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateCandleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&candle).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": candle})
}

// DELETE /books/:id
// Delete a book
func DeleteCandle(c *gin.Context) {
	// Get model if exist
	var candle models.Candle
	if err := models.DB.Where("id = ?", c.Param("id")).First(&candle).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&candle)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
