package main

import (
	"christopher/controllers"
	// "database/sql"
	"flag"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	// "log"
)

// var DB *sql.DB

func main() {
	flag.Parse()

	router := gin.Default()
	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/", controllers.Home)

		v1.GET("/shops", controllers.ListShop)
		v1.POST("/shop", controllers.NewShop)
	}

	router.Run(":8080")
}

//320
