package controllers

import (
	"christopher/helpers"
	"christopher/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
)

type categoriesForm struct {
	Offer_category_name string `form:"offer_category_name"`
	Slug                string `form:"slug"`
}
type categoriesFormEdit struct {
	Id                  int64  `form:"Id"`
	Offer_category_name string `form:"offer_category_name"`
	Slug                string `form:"slug"`
}
type offerCategory struct {
	Id                  int64  `json:"id", Number`
	Offer_category_name string `json:"offer_category_name"`
	Slug                string `json:"slug"`
	Create_at           int64  `json:"create_at"`
	Update_at           int64  `json:"update_at"`
}

func ListAllCategoriesOffer(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	data, msg, err := models.ListAllCategoriesOffer(SERVICE_NAME)
	if msg == "err" {
		c.JSON(200, gin.H{
			"status":  500,
			"message": err,
		})
	} else {
		offers := []byte(data)
		OfferSlice := make([]offerCategory, 0)
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

func NewCategoriesOffer(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	var form categoriesForm
	c.Bind(&form)

	Cat_offer := &models.OfferCategory{
		Offer_category_name: form.Offer_category_name,
		Slug:                form.Slug,
		Create_at:           helpers.Unix_milisec_time_now(),
		Update_at:           helpers.Unix_milisec_time_now(),
	}
	log.Println(Cat_offer)
	msg, err := Cat_offer.Save(SERVICE_NAME)
	if msg == "err" {
		c.JSON(200, gin.H{
			"status":  500,
			"message": err.Error(),
		})
	} else {

		c.JSON(200, gin.H{
			"status":  200,
			"message": "Created!",
		})

	}
}

func UpdateCategoriesOffer(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	offer_cat_id := c.Params.ByName("cat_id")
	var form categoriesFormEdit
	c.Bind(&form)

	Cat_offer := &models.OfferCategory{
		Id:                  helpers.Convert_string_to_int(offer_cat_id),
		Offer_category_name: form.Offer_category_name,
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

func DelelteCategoriesOffer(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	offer_cat_id := c.Params.ByName("cat_id")
	offerCate := &models.OfferCategory{
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
