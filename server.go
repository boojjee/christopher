package main

import (
	"./controllers"
	"./generate"
	"flag"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// var DB *sql.DB

func main() {
	flag.Parse()

	router := gin.Default()
	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/", controllers.Home)

		// Merchant API RESTful
		v1.GET("/:service_name/merchants", controllers.ListMerchant)
		v1.GET("/:service_name/merchant/:id", controllers.ViewMerchant)
		v1.POST("/:service_name/merchant", controllers.NewMerchant)
		v1.PUT("/:service_name/merchant/:id", controllers.UpdateMerchant)
		v1.DELETE("/:service_name/merchant/:id", controllers.DeleteMerchant)

		// Offers API RESTful
		v1.GET("/:service_name/offer_all", controllers.ListOffersAll)
		v1.GET("/:service_name/offer_of_merchant/:id", controllers.ListOffersByMerchantID)
		v1.GET("/:service_name/offer/:id", controllers.ViewOffer)
		v1.POST("/:service_name/offer", controllers.CreateOffer)
		// v1.PUT("/offer/:id", controllers.UpdateMerchant)
		// v1.DELETE("/offer/:id", controllers.DeleteMerchant)

	}

	gen := router.Group("/generate")
	{
		gen.PUT("/service/:service_name", generate.Gen_table)
	}
	// CORS middleware

	router.Use(CORSMiddleware())
	router.Run(":8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.Abort()
			return
		}
		c.Next()
	}
}

//320
