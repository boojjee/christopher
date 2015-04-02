package main

import (
	"christopher/controllers"
	"christopher/generate"
	"christopher/helpers"
	_ "expvar"
	"flag"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
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
			province, msg, err := helpers.GetProvinceFromBingMapByPoint(helpers.Convert_float_to_string(16.522531), helpers.Convert_float_to_string(102.522819))
			if err != nil {
				log.Println(err)
			}
			if msg == "err" {
				log.Println(err)
				c.JSON(200, gin.H{"status": "OK!", "build": "0.125", "province": err})
			} else {
				c.JSON(200, gin.H{"status": "OK!", "build": "0.125", "province": province})
			}

		})

		//# for Back End
		// User
		v1.POST("/:service_name/user", controllers.NewUser)                // done
		v1.PUT("/:service_name/user/:user_uid", controllers.UpdateUser)    // done
		v1.DELETE("/:service_name/user/:user_uid", controllers.DeleteUser) // done

		// Activity
		v1.GET("/:service_name/activity/:activity_uid", controllers.GetActivity)       // done
		v1.POST("/:service_name/activity", controllers.NewActivity)                    // done
		v1.PUT("/:service_name/activity/:activity_uid", controllers.UpdateActivity)    // done
		v1.DELETE("/:service_name/activity/:activity_uid", controllers.DeleteActivity) // done

		v1.GET("/:service_name/activity_lists_all/:user_uid", controllers.ActivityListsAll)         // done
		v1.GET("/:service_name/latest_activity_list/:user_uid", controllers.LatestActivityList)     // done
		v1.GET("/:service_name/next_activity_list/:user_uid/:max_id", controllers.NextActivityList) // doing + connect with point
		v1.GET("/:service_name/prev_activity_list/:user_uid/:min_id", controllers.PrevActivityList) // doing

		// Merchant
		v1.GET("/:service_name/merchantinfo/:uid", controllers.GetMerchantInfo)                          // fixing
		v1.POST("/:service_name/merchant", controllers.NewMerchant)                                      // fix
		v1.PUT("/:service_name/merchant/:uid", controllers.UpdateMerchant)                               // fix
		v1.DELETE("/:service_name/merchant/:uid", controllers.DeleteMerchant)                            // fix
		v1.POST("/:service_name/merchant_authen", controllers.AuthenMechant)                             // fix
		v1.GET("/:service_name/merchantlists", controllers.GetMerchantsLists_All)                        // fix
		v1.GET("/:service_name/merchantlists_withgallery", controllers.GetMerchantsListsWithGallery_All) // fix

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
