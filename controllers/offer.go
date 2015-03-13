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

type offerCollection []map[string]string
type offerSigle map[string]interface{}
type OfferForm struct {
	Id                 string `form:"id"`
	Name               string `form:"name" `
	Offer_point        string `form:"offer_point" `
	Condition_offer    string `form:"condition_offer" `
	Cat                string `form:"cat" `
	Merchant_id        string `form:"merchant_id" `
	Description        string `form:"description" `
	Offer_image_banner string `form:"offer_image_banner" `
	Offer_image_poster string `form:"offer_image_poster" `
	Used               string `form:"used" `
	Qty                string `form:"qty" `
}
type OfferFormEdit struct {
	Name               string `form:"name"`
	Offer_point        string `form:"offer_point"`
	Condition_offer    string `form:"condition_offer"`
	Cat                string `form:"cat"`
	Merchant_id        string `form:"merchant_id"`
	Description        string `form:"description"`
	Offer_image_banner string `form:"offer_image_banner"`
	Offer_image_poster string `form:"offer_image_poster"`
	Used               string `form:"used"`
	Qty                string `form:"qty"`
}

type Offers struct {
	Id                 int     `json:"id, Number"`
	Name               string  `json:"name, string"`
	Offer_point        float64 `json:"offer_point, Number"`
	Condition_offer    string  `json:"condition_offer, string"`
	Offer_image_banner string  `json:"offer_image_banner, string"`
	Offer_image_poster string  `json:"offer_image_poster, string"`
	Cat                int     `json:"cat, Number"`
	Merchant_id        int     `json:"merchant_id, Number"`
	Description        string  `json:"merchant_id, string"`
	Used               int     `json:"used, Number"`
	Qty                int     `json:"qty, Number"`
	Create_at          int     `json:"create_at, Number"`
	Update_at          int     `json:"update_at, Number"`
}

func ListOffersAll(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	listOfferAll := models.GetOfferListAll(SERVICE_NAME)
	offerall := []byte(listOfferAll)
	offer_slice := make([]Offers, 0)
	err := json.Unmarshal(offerall, &offer_slice)

	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{
		"status":  200,
		"message": "Sucess!",
		"data":    offer_slice,
	})
}

func ListOffersByMerchantID(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	id_merchant := c.Params.ByName("id")

	listOfferAll := models.GetOfferListByMerchantID(id_merchant, SERVICE_NAME)
	offerall := []byte(listOfferAll)
	offer_slice := make([]Offers, 0)
	err := json.Unmarshal(offerall, &offer_slice)

	if err != nil {
		c.JSON(200, gin.H{
			"status":  500,
			"message": "Somting wrong!",
		})
	} else {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "Sucess!",
			"data":    offer_slice,
		})
	}

}
func ViewOffer(c *gin.Context) {
	id_offer := c.Params.ByName("id")
	SERVICE_NAME := c.Params.ByName("service_name")
	offerInfo := models.GetOfferInfo(id_offer, SERVICE_NAME)
	var offerSigle offerSigle
	file := []byte(offerInfo)
	err := json.Unmarshal(file, &offerSigle)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  500,
			"message": "Somting wrong!",
		})
	} else {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "Sucess!",
			"data":    offerSigle,
		})
	}

}

func CreateOffer(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	var form OfferForm
	c.Bind(&form)
	offer := &models.Offer{
		Name:               form.Name,
		Offer_point:        helpers.Convert_string_to_float(form.Offer_point),
		Offer_image_banner: form.Offer_image_banner,
		Offer_image_poster: form.Offer_image_poster,
		Condition_offer:    form.Condition_offer,
		Cat:                helpers.Convert_string_to_int(form.Cat),
		Merchant_id:        helpers.Convert_string_to_int(form.Merchant_id),
		Description:        form.Description,
		Used:               helpers.Convert_string_to_int(form.Used),
		Qty:                helpers.Convert_string_to_int(form.Qty),
		Create_at:          helpers.Unix_milisec_time_now(),
		Update_at:          helpers.Unix_milisec_time_now(),
	}
	err := offer.Save(SERVICE_NAME)
	log.Println(err)

	if err != nil {
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

func UpdateOffer(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	id_offer := c.Params.ByName("id")
	// id_int, _ := strconv.ParseInt(id_offer, 0, 64)
	var form OfferFormEdit
	c.Bind(&form)
	offer := &models.Offer{
		Name:               form.Name,
		Offer_point:        helpers.Convert_string_to_float(form.Offer_point),
		Condition_offer:    form.Condition_offer,
		Offer_image_banner: form.Offer_image_banner,
		Offer_image_poster: form.Offer_image_poster,
		Cat:                helpers.Convert_string_to_int(form.Cat),
		Merchant_id:        helpers.Convert_string_to_int(form.Merchant_id),
		Description:        form.Description,
		Used:               helpers.Convert_string_to_int(form.Used),
		Qty:                helpers.Convert_string_to_int(form.Qty),
		Update_at:          helpers.Unix_milisec_time_now(),
	}
	err := offer.Update(SERVICE_NAME, id_offer)
	if err != nil {
		c.JSON(200, gin.H{
			"status": 500,
			"error":  err,
		})
	} else {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "Updated",
		})
	}

}

func DeleteOffer(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	id_offer := c.Params.ByName("id")
	id_int, _ := strconv.ParseInt(id_offer, 0, 64)
	offer := &models.Offer{
		Id: id_int,
	}
	err := offer.Delete(SERVICE_NAME)

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
