package controllers

import (
	"christopher/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

type MerchantCollection []map[string]string
type MerchantSigle map[string]interface{}
type MerchantForm struct {
	Username         string `form:"username" binding:"required"`
	Name             string `form:"name" binding:"required"`
	Password         string `form:"password" binding:"required"`
	Email            string `form:"email" binding:"required"`
	Shop_image       string `form:"shop_image" binding:"required"`
	Shop_avatar      string `form:"shop_avatar" binding:"required"`
	Shop_description string `form:"shop_description" binding:"required"`
	Lat              string `form:"lat" binding:"required"`
	Lon              string `form:"lon" binding:"required"`
}
type MerchantFormEdit struct {
	Username         string `form:"username" binding:"required"`
	Name             string `form:"name" binding:"required"`
	Password         string `form:"password" binding:"required"`
	Email            string `form:"email" binding:"required"`
	Shop_image       string `form:"shop_image" binding:"required"`
	Shop_avatar      string `form:"shop_avatar" binding:"required"`
	Shop_description string `form:"shop_description" binding:"required"`
	Lat              string `form:"lat" binding:"required"`
	Lon              string `form:"lon" binding:"required"`
}

func ListMerchant(c *gin.Context) {
	// list all shop
	SERVICE_NAME := c.Params.ByName("service_name")

	ss, errs := models.GetMerchentLists(SERVICE_NAME)
	log.Println(ss)
	if errs == "err" {
		c.JSON(200, gin.H{
			"status":  500,
			"message": "Somting wrong!",
		})
	} else {
		var merchents MerchantCollection
		file := []byte(ss)

		err := json.Unmarshal(file, &merchents)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(200, gin.H{
			"status":  200,
			"message": "Sucess!",
			"data":    merchents,
		})
	}

}

func ViewMerchant(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	id_merchant := c.Params.ByName("id")

	merchent_info, errs := models.MerchantShowInfo(id_merchant, SERVICE_NAME)
	if errs == "err" {
		c.JSON(200, gin.H{
			"status":  500,
			"message": "Somting wrong!",
		})
	} else {
		var merchents MerchantSigle
		file := []byte(merchent_info)

		err := json.Unmarshal(file, &merchents)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(200, gin.H{
			"status":  200,
			"message": "Created!",
			"data":    merchents,
		})
	}

}

func NewMerchant(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	var form MerchantForm
	c.Bind(&form)
	merchant := &models.Merchant{
		Username:         form.Username,
		Name:             form.Name,
		Password:         form.Password,
		Email:            form.Email,
		Shop_image:       form.Shop_image,
		Shop_avatar:      form.Shop_avatar,
		Shop_description: form.Shop_description,
		Lat:              form.Lat,
		Lon:              form.Lon,
		Create_at:        time.Now(),
		Update_at:        time.Now(),
	}
	merchant.Save(SERVICE_NAME)

	c.JSON(200, gin.H{
		"status":  20,
		"message": "Created!",
	})
}

func UpdateMerchant(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	id_merchant := c.Params.ByName("id")
	id_int, _ := strconv.ParseInt(id_merchant, 0, 64)
	var form MerchantFormEdit
	c.Bind(&form)
	merchant := &models.Merchant{
		Id:               id_int,
		Username:         form.Username,
		Name:             form.Name,
		Email:            form.Email,
		Shop_image:       form.Shop_image,
		Shop_avatar:      form.Shop_avatar,
		Shop_description: form.Shop_description,
		Lat:              form.Lat,
		Lon:              form.Lon,
		Update_at:        time.Now(),
	}
	err := merchant.Update(SERVICE_NAME)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  500,
			"message": "Somting wrong!",
		})
	} else {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "Updated",
		})
	}

}

func DeleteMerchant(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	id_merchant := c.Params.ByName("id")
	id_int, _ := strconv.ParseInt(id_merchant, 0, 64)
	merchant := &models.Merchant{
		Id: id_int,
	}

	err := merchant.Delete(SERVICE_NAME)
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
