package main

import (
	_ "expvar"
	"flag"
	// "fmt"
	// "github.com/boojjee/christopher/controllers"
	"christopher/controllers"
	"christopher/generate"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// var DB *sql.DB
var minversion string

func main() {
	flag.Parse()
	// fmt.Println(minversion)
	router := gin.Default()
	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/", controllers.Home)

		//# for Back End
		// Ready
		v1.GET("/:service_name/merchantinfo/:uid", controllers.GetMerchantInfo)                          // done
		v1.POST("/:service_name/merchant", controllers.NewMerchant)                                      // done
		v1.PUT("/:service_name/merchant/:uid", controllers.UpdateMerchant)                               // done
		v1.DELETE("/:service_name/merchant/:uid", controllers.DeleteMerchant)                            // done
		v1.POST("/:service_name/merchant_authen", controllers.AuthenMechant)                             // done
		v1.GET("/:service_name/merchantlists", controllers.GetMerchantsLists_All)                        // done
		v1.GET("/:service_name/merchantlists_withgallery", controllers.GetMerchantsListsWithGallery_All) // done

		v1.GET("/:service_name/merchant_gallery/:uid", controllers.GetMerchantsGallery)     // done
		v1.POST("/:service_name/merchant_gallery", controllers.NewMerchantGaller)           // done
		v1.PUT("/:service_name/merchant_gallery/:uid", controllers.UpdateMerchantGaller)    // done
		v1.DELETE("/:service_name/merchant_gallery/:uid", controllers.DeleteMerchantGaller) // done

		//# for Client iOS, Android
		v1.GET("/:service_name/merchantinfo/:uid/:lang", controllers.GetMerchantInfoByLang)                             // done
		v1.GET("/:service_name/merchantlistsbylang/:lang", controllers.GetMerchantsListsByLang)                         // done
		v1.GET("/:service_name/merchantlists_withgallery_bylang/:lang", controllers.GetMerchantsListsWithGalleryByLang) // done

		//# Offers API RESTful
		// v1.GET("/:service_name/offers", controllers.ListOffersAll)
		// v1.GET("/:service_name/offer/:id", controllers.ViewOffer)
		v1.GET("/:service_name/offermerchant/:merchant_uid", controllers.ListOffersByMerchantID) //doing
		v1.POST("/:service_name/offer", controllers.NewOffer)                                    // done
		v1.PUT("/:service_name/offer/:uid", controllers.UpdateOffer)                             // done
		v1.DELETE("/:service_name/offer/:uid", controllers.DeleteOffer)                          // done

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
