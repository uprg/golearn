// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	"github.com/jackc/pgx/v5"
// )

// /*
// var can be defined in two types, infered and non-infered
// infered => age := 20
// non-infered =>

// 	var age uint64 = 20
// 	var age int64 = 20
// */
// func main() {
// 	// Connection string
// 	// connStr := "postgres://postgres:12@localhost:5432/postgres"
// 	var connStr string = "postgres://postgres:12@localhost:5432/postgres"

// 	// Connect
// 	conn, err := pgx.Connect(context.Background(), connStr)

// 	if err != nil {
// 		log.Fatalf("Unable to connect to database: %v\n", err)
// 	}

// 	defer conn.Close(context.Background()) // close when done

// 	res, err := conn.Query(context.Background(), "select * from test")

// 	if err != nil {
// 		log.Fatalf("err in query: %v\n", err)
// 	}

// 	defer res.Close()

// 	// fmt.Println(res.Next())

// 	has_more := false

// 	for res.Next() {
// 		var name string
// 		var age float64

// 		has_more = true

// 		if err := res.Scan(&name, &age); err != nil {
// 			log.Fatalf("err is scanning row vales =: %v\n", err)
// 		}

// 		fmt.Println(name, age)
// 	}

// 	if !has_more {
// 		fmt.Println("No rows to fetch")
// 	}

// 	fmt.Println("Connected to PostgreSQL!")
// }

package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	r := gin.Default() // creates a router with default middleware (logger, recovery)

	// Define a route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{"user": name})
	})

	r.GET("/search", func(c *gin.Context) {
		term := c.Query("term")
		c.JSON(200, gin.H{"search": term})
	})

	r.POST("/user", func(c *gin.Context) {
		var user User

		// Bind JSON to struct
		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Do something with user
		c.JSON(200, gin.H{
			"status": "ok",
			"user":   user,
		})
	})

	r.Run() // starts the server on :8080
}
