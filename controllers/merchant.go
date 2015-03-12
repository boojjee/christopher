package controllers

import (
	"christopher/helpers"
	"christopher/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	// "time"
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
type MerchantAuthenForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type Merchants struct {
	Id               int    `json:"id, Number"`
	Username         string `json:"username"`
	Name             string `json:"name"`
	Password         string `json:"password"`
	Email            string `json:"email"`
	Shop_image       string `json:"shop_image"`
	Shop_avatar      string `json:"shop_avatar"`
	Shop_description string `json:"shop_description"`
	Lat              string `json:"lat"`
	Lon              string `json:"lon"`
	Create_at        int    `json:"create_at, Number"`
	Update_at        int    `json:"update_at, Number"`
}
type PublicKey struct {
	Id  int
	Key string
}

func ListMerchant(c *gin.Context) {
	// list all shop
	SERVICE_NAME := c.Params.ByName("service_name")

	mercahnt_list_json, errs := models.GetMerchentLists(SERVICE_NAME)

	if errs == "err" {
		c.JSON(200, gin.H{
			"status":  500,
			"message": "Somting wrong!",
		})
	} else {

		mercahnts := []byte(mercahnt_list_json)
		merchant_slice := make([]Merchants, 0)
		err_unmarshal := json.Unmarshal(mercahnts, &merchant_slice)

		if err_unmarshal != nil {
			log.Fatal(err_unmarshal)
		}

		c.JSON(200, gin.H{
			"status":  200,
			"message": "Sucess!",
			"data":    merchant_slice,
		})
	}

}

// func ViewMerchant(c *gin.Context) {
// 	SERVICE_NAME := c.Params.ByName("service_name")
// 	id_merchant := c.Params.ByName("id")

// 	merchent_info, errs := models.MerchantShowInfo(id_merchant, SERVICE_NAME)
// 	if errs == "err" {
// 		c.JSON(200, gin.H{
// 			"status":  500,
// 			"message": "Somting wrong!",
// 		})
// 	} else {
// 		var merchents MerchantSigle
// 		file := []byte(merchent_info)

// 		err := json.Unmarshal(file, &merchents)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		c.JSON(200, gin.H{
// 			"status":  200,
// 			"message": "Created!",
// 			"data":    merchents,
// 		})
// 	}
// }

func ViewMerchantName(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	merchant_name := c.Params.ByName("name")

	merchent_info, errs := models.MerchantShowInfoByName(merchant_name, SERVICE_NAME)
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
			"message": "Success!",
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
		Shop_avatar:      form.Shop_avatar,
		Shop_description: form.Shop_description,
		Lat:              form.Lat,
		Lon:              form.Lon,
		Create_at:        helpers.Unix_milisec_time_now(),
		Update_at:        helpers.Unix_milisec_time_now(),
	}

	msg, err := merchant.Save(SERVICE_NAME)

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
		Shop_avatar:      form.Shop_avatar,
		Shop_description: form.Shop_description,
		Lat:              form.Lat,
		Lon:              form.Lon,
		Update_at:        helpers.Unix_milisec_time_now(),
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

func AuthenMechant(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	var form MerchantAuthenForm
	c.Bind(&form)
	merchant := &models.Merchant{
		Username: form.Username,
		Password: form.Password,
	}
	msg, err := merchant.Authen(SERVICE_NAME)
	if err != nil {
		log.Println(err)
	}
	if msg == "fail" {
		c.JSON(200, gin.H{
			"status":  500,
			"message": "Authen Fail!",
			"data":    0,
		})
	} else {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "Success",
			"data":    1,
		})
	}
}
