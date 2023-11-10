package main

import (
	"intikom-interview/dal"
	"intikom-interview/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Initialize the database connection
	var err error
	db, err := gorm.Open(mysql.Open("root:amintaba12345@tcp(127.0.0.1:3306)/intikom_test?charset=utf8mb4&parseTime=True&loc=Local"))

	if err != nil {
		panic("Failed to connect to the database")
	}

	dal.SetDefault(db)

	r := gin.Default()

	routes.PublicRoutes(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
