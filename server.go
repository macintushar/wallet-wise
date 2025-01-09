package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.StaticFS("/app", http.Dir("./ui/dist"))

	// Health Check Endpoint
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	// Transactions route
	r.GET("/api/transactions", getAllTransactions)
	r.POST("/api/transactions/add", addTransaction)
	r.PUT("/api/transactions/:id", updateTransaction)
	r.DELETE("/api/transactions/:id", deleteTransaction)

	r.Run(":3000")
}
