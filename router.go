package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func getRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	///books

	router.GET("/books", func(c *gin.Context) { getAllBook(c, db) })

	router.GET("/books/:id", func(c *gin.Context) { getBook(c, db) })

	router.POST("/books", func(c *gin.Context) { createBook(c, db) })

	router.DELETE("/books/:id", func(c *gin.Context) { deleteBook(c, db) })

	router.PUT("/books/:id", func(c *gin.Context) { updateBook(c, db) })

	///authors

	router.GET("/authors", func(c *gin.Context) { getAllAuthor(c, db) })

	//router.POST("/authors", func(c *gin.Context) { createService(c, db) })

	///Publishers

	router.GET("/publishers", func(c *gin.Context) { getAllPublish(c, db) })

	//router.POST("/payments", func(c *gin.Context) { createPayment(c, db) })

	///Genre

	router.GET("/genres", func(c *gin.Context) { getAllGenre(c, db) })

	///Хранимые процедуры и доп запросики
	//TODO: Добваить другие запросики
	router.GET("books/info/:id", func(c *gin.Context) { InfoBook(c, db) })
	router.GET("/books/find", func(c *gin.Context) { FindBook(c, db) })

	return router
}
