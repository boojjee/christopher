package controllers

import (
	"christopher/helpers"
	"christopher/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	// "log"

	// "time"
)

type catagoriesForm struct {
	Offer_category_name string `form:"offer_catagory_name"`
	Slug                string `form:"slug"`
}
type catagoriesFormEdit struct {
	Id                  int64  `form:"Id"`
	Offer_category_name string `form:"offer_catagory_name"`
	Slug                string `form:"slug"`
}
type offerCatagory struct {
	Id                  int64  `json:"id", Number`
	Offer_catagory_name string `json:"offer_catagory_name`
	Slug                string `json:"slug`
	Create_at           int64  `json:"Create_at"`
	Update_at           int64  `json:"Update_at"`
}

func ListAllCatagoriesOffer(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	data, msg, err := models.ListAllCatagoriesOffer(SERVICE_NAME)
	if msg == "err" {
		c.JSON(200, gin.H{
			"status":  500,
			"message": err,
		})
	} else {
		offers := []byte(data)
		OfferSlice := make([]offerCatagory, 0)
		err_unmarshal := json.Unmarshal(offers, &OfferSlice)
		if err_unmarshal != nil {
			c.JSON(200, gin.H{
				"status": 500,
				"error":  err,
			})
		} else {
			c.JSON(200, gin.H{
				"status":  200,
				"message": "Success!",
				"data":    OfferSlice,
			})
		}
	}
}

func NewCatagoriesOffer(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	var form catagoriesForm
	c.Bind(&form)

	Cat_offer := &models.OfferCatagory{
		Offer_catagory_name: form.Offer_category_name,
		Slug:                form.Slug,
		Create_at:           helpers.Unix_milisec_time_now(),
		Update_at:           helpers.Unix_milisec_time_now(),
	}

	msg, err := Cat_offer.Save(SERVICE_NAME)
	if msg == "err" {
		c.JSON(200, gin.H{
			"status":  500,
			"message": err,
		})
	} else {

		c.JSON(200, gin.H{
			"status":  200,
			"message": "Created!",
		})

	}
}

func UpdateCatagoriesOffer(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	offer_cat_id := c.Params.ByName("cat_id")
	var form catagoriesFormEdit
	c.Bind(&form)

	Cat_offer := &models.OfferCatagory{
		Id:                  helpers.Convert_string_to_int(offer_cat_id),
		Offer_catagory_name: form.Offer_category_name,
		Slug:                form.Slug,
		Create_at:           helpers.Unix_milisec_time_now(),
		Update_at:           helpers.Unix_milisec_time_now(),
	}

	msg, err := Cat_offer.Update(SERVICE_NAME)
	if msg == "err" {
		c.JSON(200, gin.H{
			"status":  500,
			"message": err,
		})
	} else {

		c.JSON(200, gin.H{
			"status":  200,
			"message": "Update!",
		})

	}
}

func DelelteCatagoriesOffer(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	offer_cat_id := c.Params.ByName("cat_id")
	offerCate := &models.OfferCatagory{
		Id: helpers.Convert_string_to_int(offer_cat_id),
	}

	msg, err := offerCate.Delete(SERVICE_NAME)
	if msg == "err" {
		c.JSON(200, gin.H{
			"status": 500,
			"error":  err,
		})
	} else {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "Deleted!",
		})
	}
}
