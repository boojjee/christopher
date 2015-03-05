package controllers

import (
	"encoding/json"
	"github.com/boojjee/christopher/models"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

type MerchantCollection []map[string]string
type MerchantSigle map[string]interface{}
type MerchantForm struct {
	Id               string `form:"id" binding:"required"`
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
	ss := models.GetMerchentLists()
	var merchents MerchantCollection
	file := []byte(ss)

	err := json.Unmarshal(file, &merchents)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{"data": merchents})
}

func ViewMerchant(c *gin.Context) {
	id_merchant := c.Params.ByName("id")
	merchent_info := models.MerchantShowInfo(id_merchant)
	var merchents MerchantSigle
	file := []byte(merchent_info)

	err := json.Unmarshal(file, &merchents)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{"data": merchents})
}

func NewMerchant(c *gin.Context) {
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
	merchant.Save()
	c.JSON(200, gin.H{"status": "Created!"})
}

func UpdateMerchant(c *gin.Context) {
	var form MerchantForm
	c.Bind(&form)
	id_int, _ := strconv.ParseInt(form.Id, 0, 64)
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
	merchant.Update()
	c.JSON(200, gin.H{"status": "Updated!"})
}

func DeleteMerchant(c *gin.Context) {
	id_merchant := c.Params.ByName("id")
	id_int, _ := strconv.ParseInt(id_merchant, 0, 64)
	merchant := &models.Merchant{
		Id: id_int,
	}

	merchant.Delete()
	c.JSON(200, gin.H{"status": "Deleted!"})
}
