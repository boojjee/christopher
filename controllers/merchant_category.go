package controllers

import (
	"christopher/helpers"
	"christopher/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type categoriesFormMerchant struct {
	Merchant_category_name string `form:"merchant_category_name"`
	Slug                   string `form:"slug"`
}
type categoriesFormMerchantEdit struct {
	Id                     int64  `form:"Id"`
	Merchant_category_name string `form:"merchant_category_name"`
	Slug                   string `form:"slug"`
}
type merchantCategory struct {
	Id                     int64  `json:"id", Number`
	Merchant_category_name string `json:"merchant_category_name"`
	Slug                   string `json:"slug"`
	Create_at              int64  `json:"create_at"`
	Update_at              int64  `json:"update_at"`
}

func ListAllCategoriesMerchant(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	data, msg, err := models.ListAllCategoriesMerchant(SERVICE_NAME)
	if msg == "err" {
		c.JSON(200, gin.H{
			"status":  500,
			"message": err,
		})
	} else {
		merchants := []byte(data)
		MerchantSlice := make([]merchantCategory, 0)
		err_unmarshal := json.Unmarshal(merchants, &MerchantSlice)
		if err_unmarshal != nil {
			c.JSON(200, gin.H{
				"status": 500,
				"error":  err,
			})
		} else {
			c.JSON(200, gin.H{
				"status":  200,
				"message": "Success!",
				"data":    MerchantSlice,
			})
		}
	}
}

func NewCategoriesMerchant(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	var form categoriesFormMerchant
	c.Bind(&form)

	Cat_merchant := &models.MerchantCategory{
		Merchant_category_name: form.Merchant_category_name,
		Slug:      form.Slug,
		Create_at: helpers.Unix_milisec_time_now(),
		Update_at: helpers.Unix_milisec_time_now(),
	}

	msg, err := Cat_merchant.Save(SERVICE_NAME)
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

func UpdateCategoriesMerchant(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	merchant_cat_id := c.Params.ByName("cat_id")
	var form categoriesFormMerchantEdit
	c.Bind(&form)

	Cat_merchant := &models.MerchantCategory{
		Id: helpers.Convert_string_to_int(merchant_cat_id),
		Merchant_category_name: form.Merchant_category_name,
		Slug:      form.Slug,
		Create_at: helpers.Unix_milisec_time_now(),
		Update_at: helpers.Unix_milisec_time_now(),
	}

	msg, err := Cat_merchant.Update(SERVICE_NAME)
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

func DelelteCategoriesMerchant(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	merchant_cat_id := c.Params.ByName("cat_id")
	merchantCate := &models.MerchantCategory{
		Id: helpers.Convert_string_to_int(merchant_cat_id),
	}

	msg, err := merchantCate.Delete(SERVICE_NAME)
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
