package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getAllClient(c *gin.Context, db *sql.DB) {
	var clients []client
	rows, err := db.Query("SELECT Id, FirstName,LastName, Email FROM client")
	if err != nil {
		c.JSON(500, gin.H{"message": "Error Get All User"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var u client
		//TODO:повторить
		err := rows.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "something wrong mann..."})
			return
		}
		clients = append(clients, u)
	}

	//TODO:разобраться тут поподробней
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, clients)
}

func getClient(c *gin.Context, db *sql.DB) {
	var u client
	id := c.Param("id")

	err := db.QueryRow("SELECT Id, FirstName,LastName, Email FROM client WHERE Id = ?", id).Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "user lost brooo"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something wrong mann..."})
		return
	}
	c.JSON(http.StatusOK, u)
}

func createClient(c *gin.Context, db *sql.DB) {

	var nu client //newUser
	if err := c.ShouldBindJSON(&nu); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something wrong mann..."})
		return
	}

	_, err := db.Exec("INSERT INTO client (FirstName,LastName, Email) VALUES (? , ?, ?)", nu.FirstName, nu.LastName, nu.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//c.JSON(http.StatusOK, gin.H{"massage": "it's all good man"})
	getAllClient(c, db)
}

func deleteClient(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	_, err := db.Exec("DELETE FROM client WHERE Id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	getAllClient(c, db)
	//c.JSON(http.StatusOK, gin.H{"massage": "it's all good man"})
}

func updateClient(c *gin.Context, db *sql.DB) {
	var nu client
	id := c.Param("id")
	if err := c.ShouldBindJSON(&nu); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec("UPDATE client SET FirstName = ?, LastName = ?, Email = ? WHERE Id = ?", nu.FirstName, nu.LastName, nu.Email, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	getAllClient(c, db)
	//c.JSON(http.StatusOK, gin.H{"massage": "it's all good man"})
}

///

func getAllServices(c *gin.Context, db *sql.DB) {
	var services []service
	rows, err := db.Query("SELECT Id, ServiceName, Price FROM service")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error Get All Services"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var s service
		err := rows.Scan(&s.Id, &s.Name, &s.Price)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "something wrong..."})
			return
		}
		services = append(services, s)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, services)
}

func createService(c *gin.Context, db *sql.DB) {
	var ns service

	if err := c.ShouldBindJSON(&ns); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something wrong..."})
		return
	}

	_, err := db.Exec("INSERT INTO service (Name, Price) VALUES (?, ?)", ns.Name, ns.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	getAllServices(c, db)
}

func getAllPayments(c *gin.Context, db *sql.DB) {
	var payments []payment
	rows, err := db.Query("SELECT Id, ClientId, ServiceId,Date, Quantity, Amount FROM payment")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error Get All Payments"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var p payment
		err := rows.Scan(&p.Id, &p.ClientId, &p.ServiceId, &p.Date, &p.Quantity, &p.Amount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		payments = append(payments, p)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, payments)
}

func createPayment(c *gin.Context, db *sql.DB) {
	var np payment

	if err := c.ShouldBindJSON(&np); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something wrong..."})
		return
	}

	_, err := db.Exec("INSERT INTO payment (ClientId, ServiceId, Date, Quantity, Amount) VALUES (?, ?, ?, ?, ?)",
		np.ClientId, np.ServiceId, np.Quantity, np.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	getAllPayments(c, db)
}

///

func StoredProcedure(c *gin.Context, db *sql.DB) {
	id := c.Param("id")

	var info []InfoDemo

	rows, err := db.Query("call mydb.infoDemo(?)", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var i InfoDemo
		err := rows.Scan(&i.FirstName, &i.LastName, &i.Phone, &i.ServiceName, &i.Quantity, &i.Amount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		info = append(info, i)
	}
	c.JSON(http.StatusOK, info)
}
