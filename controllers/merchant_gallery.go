package controllers

import (
	"christopher/helpers"
	"christopher/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type MerchantGallerySigle map[string]interface{}
type MerchantGallery struct {
	Id           int64  `json:"id, Number"`
	Photo_url    string `json:"photo_url"`
	Merchant_uid string `json:"merchant_uid"`
	Create_at    int64  `json:"create_at, Number"`
	Update_at    int64  `json:"update_at, Number"`
}
type MerchantGalleryForm struct {
	Id           int64  `form:"id"`
	Photo_url    string `form:"photo_url"`
	Merchant_uid string `form:"merchant_uid"`
	Create_at    int64  `form:"create_at"`
	Update_at    int64  `form:"update_at"`
}

// get merchant info with image gallery

// create
func NewMerchantGaller(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	var form MerchantGalleryForm
	c.Bind(&form)

	merchantGallery := &models.MerchantGallery{
		Photo_url:    form.Photo_url,
		Merchant_uid: form.Merchant_uid,
		Create_at:    helpers.Unix_milisec_time_now(),
		Update_at:    helpers.Unix_milisec_time_now(),
	}
	msg, err := merchantGallery.Save(SERVICE_NAME)
	if msg == "err" {
		c.JSON(200, gin.H{
			"status": 500,
			"error":  err,
		})
	} else {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "Created!",
		})
	}
}

// update
func UpdateMerchantGaller(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	photo_id := c.Params.ByName("id")
	var form MerchantGalleryForm
	c.Bind(&form)

	merchantGallery := &models.MerchantGallery{
		Photo_url:    form.Photo_url,
		Merchant_uid: form.Merchant_uid,
		Create_at:    helpers.Unix_milisec_time_now(),
		Update_at:    helpers.Unix_milisec_time_now(),
	}
	msg, err := merchantGallery.Update(SERVICE_NAME, photo_id)
	if msg == "err" {
		c.JSON(200, gin.H{
			"status": 500,
			"error":  err,
		})
	} else {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "Updated!",
		})
	}
}

// delete
func DeleteMerchantGaller(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	id_photo := c.Params.ByName("uid")
	id_int, _ := strconv.ParseInt(id_photo, 0, 64)
	merchantGallery := &models.MerchantGallery{
		Id: id_int,
	}

	err := merchantGallery.Delete(SERVICE_NAME)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  500,
			"message": "Somting wrong!",
		})
	} else {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "Deleted",
		})
	}
}

func GetMerchantsGallery(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	merchant_uid := c.Params.ByName("uid")
	merchant := &models.MerchantMeta{
		Merchant_uid: merchant_uid,
	}

	data, msg, err := merchant.MerchantShowGallery(SERVICE_NAME)

	if msg == "err" {
		c.JSON(200, gin.H{
			"status":  500,
			"message": err,
		})
	} else {
		log.Println(data)

		mercahnts := []byte(data)
		// res := &MerchantGallery{}
		merchant_slice := make([]MerchantGallerySigle, 0)
		err_unmarshal := json.Unmarshal(mercahnts, &merchant_slice)

		if err_unmarshal != nil {
			c.JSON(200, gin.H{
				"status":  500,
				"message": "json error",
			})
		}
		c.JSON(200, gin.H{
			"status":  200,
			"message": "Success!",
			"data":    merchant_slice,
		})
	}
}
