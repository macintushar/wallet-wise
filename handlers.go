package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	Amount      string   `json:"amount"`
	Description string   `json:"description"`
	Currency    string   `json:"currency"`
	Tags        []string `json:"tags"`
	Date        string   `json:"date"`
}

func getAllTransactions(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello from Transaction",
	})

}

func addTransaction(c *gin.Context) {
	var body = Transaction{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	amount := body.Amount
	description := body.Description
	currency := body.Currency
	tags := body.Tags
	date := body.Date

	fmt.Println(amount, description, currency, tags, date)

	c.JSON(200, gin.H{
		"message": "Transaction added successfully",
		"data":    body,
	})
}

func updateTransaction(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Transaction deleted successfully",
	})
	fmt.Println(c.Params.ByName("id"))
}

func deleteTransaction(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Transaction deleted successfully",
	})
	c.Request.URL.Query()
}
