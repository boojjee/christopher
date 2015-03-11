package controllers

import (
	"christopher/helpers"
	"christopher/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

type offerCollection []map[string]string
type offerSigle map[string]interface{}
type OfferForm struct {
	Id              string `form:"id"`
	Name            string `form:"name" binding:"required"`
	Offer_point     string `form:"offer_point" binding:"required"`
	Condition_offer string `form:"condition_offer" binding:"required"`
	Cat             string `form:"cat" binding:"required"`
	Merchant_id     string `form:"merchant_id" binding:"required"`
	Description     string `form:"description" binding:"required"`
	Offer_image     string `form:"offer_image" binding:"required"`
	Used            string `form:"used" binding:"required"`
	Qty             string `form:"qty" binding:"required"`
}
type OfferFormEdit struct {
	Name            string `form:"name" binding:"required"`
	Offer_point     string `form:"offer_point" binding:"required"`
	Condition_offer string `form:"condition_offer" binding:"required"`
	Cat             string `form:"cat" binding:"required"`
	Merchant_id     string `form:"merchant_id" binding:"required"`
	Offer_image     string `form:"offer_image" binding:"required"`
	Description     string `form:"description" binding:"required"`
	Used            string `form:"used" binding:"required"`
	Qty             string `form:"qty" binding:"required"`
}

func ListOffersAll(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	listOfferAll := models.GetOfferListAll(SERVICE_NAME)
	var offersAll offerCollection
	file := []byte(listOfferAll)
	err := json.Unmarshal(file, &offersAll)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{
		"status":  200,
		"message": "Sucess!",
		"data":    offersAll,
	})
}

func ListOffersByMerchantID(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	id_merchant := c.Params.ByName("id")
	listOfferAll := models.GetOfferListByMerchantID(id_merchant, SERVICE_NAME)
	var offersList offerCollection
	file := []byte(listOfferAll)
	err := json.Unmarshal(file, &offersList)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{
		"status":  200,
		"message": "Sucess!",
		"data":    offersList,
	})
}
func ViewOffer(c *gin.Context) {
	id_offer := c.Params.ByName("id")
	SERVICE_NAME := c.Params.ByName("service_name")
	offerInfo := models.GetOfferInfo(id_offer, SERVICE_NAME)
	var offerSigle offerSigle
	file := []byte(offerInfo)
	err := json.Unmarshal(file, &offerSigle)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{
		"status":  200,
		"message": "Sucess!",
		"data":    offerSigle,
	})
}

func CreateOffer(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	var form OfferForm
	c.Bind(&form)
	offer := &models.Offer{
		Name:            form.Name,
		Offer_point:     helpers.Convert_string_to_float(form.Offer_point),
		Offer_image:     form.Offer_image,
		Condition_offer: form.Condition_offer,
		Cat:             helpers.Convert_string_to_int(form.Cat),
		Merchant_id:     helpers.Convert_string_to_int(form.Merchant_id),
		Descrtion:       form.Description,
		Used:            helpers.Convert_string_to_int(form.Used),
		Qty:             helpers.Convert_string_to_int(form.Qty),
		Create_at:       time.Now(),
		Update_at:       time.Now(),
	}
	offer.Save(SERVICE_NAME)

	c.JSON(200, gin.H{
		"status":  20,
		"message": "Created!",
	})
}

func UpdateOffer(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	id_offer := c.Params.ByName("id")
	// id_int, _ := strconv.ParseInt(id_offer, 0, 64)
	var form OfferFormEdit
	c.Bind(&form)
	offer := &models.Offer{
		Name:            form.Name,
		Offer_point:     helpers.Convert_string_to_float(form.Offer_point),
		Condition_offer: form.Condition_offer,
		Cat:             helpers.Convert_string_to_int(form.Cat),
		Merchant_id:     helpers.Convert_string_to_int(form.Merchant_id),
		Descrtion:       form.Description,
		Used:            helpers.Convert_string_to_int(form.Used),
		Qty:             helpers.Convert_string_to_int(form.Qty),
		Update_at:       time.Now(),
	}
	offer.Update(SERVICE_NAME, id_offer)

	c.JSON(200, gin.H{
		"status":  200,
		"message": "Updated",
	})
}

func DeleteOffer(c *gin.Context) {
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
