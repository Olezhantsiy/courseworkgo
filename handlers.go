package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getAllBook(c *gin.Context, db *sql.DB) {
	var books []book
	rows, err := db.Query("SELECT id, Title ,Code, YearPublish, CountPage, Price, Hardcover, Abstract, Status, AuthorId, PublishId, GenreId  FROM book")
	if err != nil {
		c.JSON(500, gin.H{"message": "Error Get All book"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var b book
		//TODO:повторить
		err := rows.Scan(&b.id, &b.Title, &b.Code, &b.YearPublish, &b.CountPage, &b.Price, &b.Hardcover, &b.Abstract, &b.Status, &b.AuthorId, &b.PublishId, &b.GenreId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		books = append(books, b)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, books)
}

func getBook(c *gin.Context, db *sql.DB) {
	var b book
	id := c.Param("id")

	err := db.QueryRow("SELECT id, Title ,Code, YearPublish, CountPage, Price, Hardcover, Abstract, Status, AuthorId, PublishId, GenreId FROM book WHERE Id = ?",
		id).Scan(&b.id, &b.Title, &b.Code, &b.YearPublish, &b.CountPage, &b.Price, &b.Hardcover, &b.Abstract, &b.Status, &b.AuthorId, &b.PublishId, &b.GenreId)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "user lost brooo"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something wrong mann..."})
		return
	}
	c.JSON(http.StatusOK, b)
}

func createBook(c *gin.Context, db *sql.DB) {

	var nb book //newUser
	if err := c.ShouldBindJSON(&nb); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something wrong mann..."})
		return
	}

	_, err := db.Exec("INSERT INTO book (Title,Code, YearPublish,CountPage, Price, Hardcover, Abstract, Status) VALUES (?,?,?,?,?,?,?,?)",
		nb.Title, nb.Code, nb.YearPublish, nb.CountPage, nb.Price, nb.Hardcover, nb.Abstract, nb.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//c.JSON(http.StatusOK, gin.H{"massage": "it's all good man"})
	getAllBook(c, db)
}

func deleteBook(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	_, err := db.Exec("DELETE FROM book WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	getAllBook(c, db)
	//c.JSON(http.StatusOK, gin.H{"massage": "it's all good man"})
}

func updateBook(c *gin.Context, db *sql.DB) {
	var nb book
	id := c.Param("id")
	if err := c.ShouldBindJSON(&nb); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec(
		"UPDATE book SET Title = ?, Code = ?, YearPublish = ?,CountPage= ?,Price= ?, Hardcover= ?, Abstract= ?, Status = ? WHERE id = ?",
		nb.Title, nb.Code, nb.YearPublish, nb.CountPage, nb.Price, nb.Hardcover, nb.Abstract, nb.Status, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	getAllBook(c, db)
	//c.JSON(http.StatusOK, gin.H{"massage": "it's all good man"})
}

///

func getAllGenre(c *gin.Context, db *sql.DB) {
	var genres []genre
	rows, err := db.Query("SELECT Id, Name FROM genre")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error Get All genres"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var g genre
		err := rows.Scan(&g.id, &g.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "something wrong..."})
			return
		}
		genres = append(genres, g)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, genres)
}

func getAllAuthor(c *gin.Context, db *sql.DB) {
	var authors []author
	rows, err := db.Query("SELECT id, FirstName, LastName FROM author")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error Get All Payments"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var a author
		err := rows.Scan(&a.id, &a.FirstName, &a.LastName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		authors = append(authors, a)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, authors)
}

func getAllPublish(c *gin.Context, db *sql.DB) {
	var publishs []publish
	rows, err := db.Query("SELECT id, NamePublish, Address, Site FROM publish")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error Get All Payments"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var p publish
		err := rows.Scan(&p.id, &p.NamePublish, &p.Address, &p.Site)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		publishs = append(publishs, p)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, publishs)
}

///

func InfoBook(c *gin.Context, db *sql.DB) {
	id := c.Param("id")

	var bookInfo []SP_InfoBook

	rows, err := db.Query("call bookstorage.bookInfo(?)", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var i SP_InfoBook
		err := rows.Scan(&i.FirstName, &i.LastName, &i.Title, &i.PublishYear, &i.Publish)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		bookInfo = append(bookInfo, i)
	}
	c.JSON(http.StatusOK, bookInfo)
}

func FindBook(c *gin.Context, db *sql.DB) {
	var books []book
	search := c.Query("search")
	rows, err := db.Query("SELECT id, Title, Code, YearPublish, CountPage, Price, Hardcover, Abstract, Status, AuthorId, PublishId, GenreId FROM book WHERE Title LIKE ?",
		"%"+search+"%")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	defer rows.Close()

	for rows.Next() {
		var b book
		err := rows.Scan(&b.id, &b.Title, &b.Code, &b.YearPublish, &b.CountPage, &b.Price, &b.Hardcover, &b.Abstract, &b.Status, &b.AuthorId, &b.PublishId, &b.GenreId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		books = append(books, b)
	}
	c.JSON(http.StatusOK, books)
}

func SearchBook(c *gin.Context, db *sql.DB) {
	var books []book
	var args []interface{}
	//var quaryStr string

	title := c.Query("title")
	author := c.Query("author")
	genre := c.Query("genre")
	publish := c.Query("publish")

	quaryStr := "SELECT id, Title, Code, YearPublish, CountPage, Price, Hardcover, Abstract, Status, AuthorId, PublishId, GenreId FROM book WHERE 1=1"

	if title != "" {
		quaryStr += " AND Title Like ?"
		args = append(args, "%"+title+"%")
		fmt.Println(title)
		fmt.Println(args)
	}

	if author != "" {
		quaryStr += " AND AuthorId Like ?"
		args = append(args, "%"+author+"%")
		fmt.Println(author)
		fmt.Println(args)
	}

	if genre != "" {
		quaryStr += " AND GenreId Like ?"
		args = append(args, "%"+genre+"%")
		fmt.Println(genre)
		fmt.Println(args)
	}

	if publish != "" {
		quaryStr += " AND PublishId Like ?"
		args = append(args, "%"+publish+"%")
		fmt.Println(publish)
		fmt.Println(args)
	}

	quaryStr += ""

	rows, err := db.Query(quaryStr, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var b book
		err := rows.Scan(&b.id, &b.Title, &b.Code, &b.YearPublish, &b.CountPage, &b.Price, &b.Hardcover, &b.Abstract, &b.Status, &b.AuthorId, &b.PublishId, &b.GenreId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		books = append(books, b)
	}
	c.JSON(http.StatusOK, books)

}
