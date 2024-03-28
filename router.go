package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func getRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	///

	router.GET("/clients", func(c *gin.Context) { getAllClient(c, db) })

	router.GET("/clients/:id", func(c *gin.Context) { getClient(c, db) })

	router.POST("/clients", func(c *gin.Context) { createClient(c, db) })

	router.DELETE("/clients/:id", func(c *gin.Context) { deleteClient(c, db) })

	router.PUT("/clients/:id", func(c *gin.Context) { updateClient(c, db) })

	///

	router.GET("/services", func(c *gin.Context) { getAllServices(c, db) })

	router.POST("/services", func(c *gin.Context) { createService(c, db) })

	router.GET("/payments", func(c *gin.Context) { getAllPayments(c, db) })

	router.POST("/payments", func(c *gin.Context) { createPayment(c, db) })

	///

	router.GET("clients/orders/:id", func(c *gin.Context) { StoredProcedure(c, db) })

	return router
}
