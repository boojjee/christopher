package main

import (
	"christopher/controllers"
	"christopher/generate"
	// "christopher/helpers"
	_ "expvar"
	"flag"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	// "log"
	// "time"
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
		v1.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "OK!", "build": "0.131"})
		})

		//# for Back End
		// User
		v1.POST("/:service_name/user", controllers.NewUser)                    // done
		v1.PUT("/:service_name/user/:user_uid", controllers.UpdateUser)        // done
		v1.PUT("/:service_name/user/:user_uid/pin", controllers.UpdateUserPin) // done
		v1.DELETE("/:service_name/user/:user_uid", controllers.DeleteUser)     // done
		v1.GET("/:service_name/user/:parse_id/parse", controllers.GetUserUID)  // done
		v1.GET("/:service_name/checkpin/:user_uid", controllers.CheckHadPin)   // done

		// DashBoard
		v1.GET("/:service_name/my_point/:user_uid", controllers.GetMyPoint) // done

		// For frontEnd
		v1.GET("/:service_name/workout", controllers.GetWorkOut) // done
		v1.GET("/:service_name/gpoint", controllers.GetPoints)   // done

		// Activity
		v1.GET("/:service_name/activity/:activity_uid", controllers.GetActivity)       // done
		v1.POST("/:service_name/activity", controllers.NewActivity)                    // done
		v1.PUT("/:service_name/activity/:activity_uid", controllers.UpdateActivity)    // done
		v1.DELETE("/:service_name/activity/:activity_uid", controllers.DeleteActivity) // done

		v1.GET("/:service_name/activity_lists_all/:user_uid", controllers.ActivityListsAll)         // done
		v1.GET("/:service_name/latest_activity_list/:user_uid", controllers.LatestActivityList)     // done
		v1.GET("/:service_name/next_activity_list/:user_uid/:max_id", controllers.NextActivityList) // done
		v1.GET("/:service_name/prev_activity_list/:user_uid/:min_id", controllers.PrevActivityList) // done

		// Merchant
		v1.GET("/:service_name/merchantinfo/:uid", controllers.GetMerchantInfo)                          // done
		v1.POST("/:service_name/merchant", controllers.NewMerchant)                                      // done
		v1.PUT("/:service_name/merchant/:uid", controllers.UpdateMerchant)                               // done
		v1.DELETE("/:service_name/merchant/:uid", controllers.DeleteMerchant)                            // done
		v1.POST("/:service_name/merchant_authen", controllers.AuthenMechant)                             // done
		v1.GET("/:service_name/merchantlists", controllers.GetMerchantsLists_All)                        // done
		v1.GET("/:service_name/merchantlists_withgallery", controllers.GetMerchantsListsWithGallery_All) // done

		// Merchant Gallery
		v1.GET("/:service_name/merchant_gallery/:uid", controllers.GetMerchantsGallery)     // done
		v1.POST("/:service_name/merchant_gallery", controllers.NewMerchantGaller)           // done
		v1.PUT("/:service_name/merchant_gallery/:uid", controllers.UpdateMerchantGaller)    // done
		v1.DELETE("/:service_name/merchant_gallery/:uid", controllers.DeleteMerchantGaller) // done

		// Merchant  Category
		v1.GET("/:service_name/merchant_categories", controllers.ListAllCategoriesMerchant)          // done
		v1.POST("/:service_name/merchant_category", controllers.NewCategoriesMerchant)               // done
		v1.PUT("/:service_name/merchant_category/:cat_id", controllers.UpdateCategoriesMerchant)     // done
		v1.DELETE("/:service_name/merchant_category/:cat_id", controllers.DelelteCategoriesMerchant) // done

		//# for Client iOS, Android
		v1.GET("/:service_name/merchantinfo/:uid/:lang", controllers.GetMerchantInfoByLang)                             // done
		v1.GET("/:service_name/merchantlistsbylang/:lang", controllers.GetMerchantsListsByLang)                         // done
		v1.GET("/:service_name/merchantlists_withgallery_bylang/:lang", controllers.GetMerchantsListsWithGalleryByLang) // done

		//# Offers API RESTful
		v1.GET("/:service_name/offers", controllers.ListOffersAll)                               // done
		v1.GET("/:service_name/offer/:uid", controllers.ViewOffer)                               // done
		v1.GET("/:service_name/offermerchant/:merchant_uid", controllers.ListOffersByMerchantID) // done
		v1.POST("/:service_name/offer", controllers.NewOffer)                                    // done
		v1.PUT("/:service_name/offer/:uid", controllers.UpdateOffer)                             // done
		v1.DELETE("/:service_name/offer/:uid", controllers.DeleteOffer)                          // done

		v1.GET("/:service_name/offer_categories", controllers.ListAllCategoriesOffer)          // done
		v1.POST("/:service_name/ ", controllers.NewCategoriesOffer)                            // done
		v1.PUT("/:service_name/offer_category/:cat_id", controllers.UpdateCategoriesOffer)     // done
		v1.DELETE("/:service_name/offer_category/:cat_id", controllers.DelelteCategoriesOffer) // done

		// Point Setting
		v1.GET("/:service_name/setting_point", controllers.ListAllPointSetting)        // done
		v1.POST("/:service_name/setting_point", controllers.NewPointSetting)           // done
		v1.PUT("/:service_name/setting_point/:uid", controllers.UpdatePointSetting)    // done
		v1.DELETE("/:service_name/setting_point/:uid", controllers.DeletePointSetting) // done

		// Redeem
		v1.POST("/:service_name/redeem", controllers.RedeemOffer) // doing
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
