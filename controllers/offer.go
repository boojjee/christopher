package controllers

import (
	"christopher/helpers"
	"christopher/models"
	// "encoding/json"
	"github.com/gin-gonic/gin"
	// "log"

	// "time"
)

type OfferSingle map[string]interface{}
type offerFormAllLang struct {
	Offer_point        string `form:"offer_point"`
	Merchant_uid       string `form:"merchant_uid"`
	Offer_cat_id       string `form:"offer_cat_id"`
	Offer_image_banner string `form:"offer_image_banner"`
	Offer_image_poster string `form:"offer_image_poster"`
	Used               string `form:"used"`
	Quantity           string `form:"quantity"`
	Name_en            string `form:"name_en"`
	Name_th            string `form:"name_th" `
	Condition_offer_en string `form:"condition_offer_en"`
	Condition_offer_th string `form:"condition_offer_th"`
	Description_en     string `form:"description_en" `
	Description_th     string `form:"description_en"`
	Lang               string `form:"lang"`
}

func NewOffer(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	var form offerFormAllLang
	c.Bind(&form)

	offer := &models.OfferAllContent{
		Offer_uid:          helpers.RandomStr(10),
		Merchant_uid:       form.Merchant_uid,
		Offer_point:        helpers.Convert_string_to_float(form.Offer_point),
		Offer_cat_id:       helpers.Convert_string_to_int(form.Offer_cat_id),
		Offer_image_banner: form.Offer_image_banner,
		Offer_image_poster: form.Offer_image_poster,
		Used:               0,
		Quantity:           helpers.Convert_string_to_int(form.Quantity),
		Name_en:            form.Name_en,
		Name_th:            form.Name_th,
		Condition_offer_en: form.Condition_offer_en,
		Condition_offer_th: form.Condition_offer_th,
		Description_en:     form.Description_en,
		Description_th:     form.Description_th,
		Create_at:          helpers.Unix_milisec_time_now(),
		Update_at:          helpers.Unix_milisec_time_now(),
	}

	msg, err := offer.Save(SERVICE_NAME)
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

func UpdateOffer(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	OFFER_UID := c.Params.ByName("uid")
	var form offerFormAllLang
	c.Bind(&form)
	offer := &models.OfferAllContent{
		Offer_uid:          OFFER_UID,
		Merchant_uid:       form.Merchant_uid,
		Offer_point:        helpers.Convert_string_to_float(form.Offer_point),
		Offer_cat_id:       helpers.Convert_string_to_int(form.Offer_cat_id),
		Offer_image_banner: form.Offer_image_banner,
		Offer_image_poster: form.Offer_image_poster,
		Used:               helpers.Convert_string_to_int(form.Used),
		Quantity:           helpers.Convert_string_to_int(form.Quantity),
		Name_en:            form.Name_en,
		Name_th:            form.Name_th,
		Condition_offer_en: form.Condition_offer_en,
		Condition_offer_th: form.Condition_offer_th,
		Description_en:     form.Description_en,
		Description_th:     form.Description_th,
		Create_at:          helpers.Unix_milisec_time_now(),
		Update_at:          helpers.Unix_milisec_time_now(),
	}
	msg, err := offer.Update(SERVICE_NAME)
	if msg == "err" {
		c.JSON(200, gin.H{
			"status": 500,
			"error":  err,
		})
	} else {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "Update!",
		})
	}
}

func DeleteOffer(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	OFFER_UID := c.Params.ByName("uid")
	var form offerFormAllLang
	c.Bind(&form)
	offer := &models.OfferAllContent{
		Offer_uid: OFFER_UID,
	}
	msg, err := offer.Delete(SERVICE_NAME)
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

func ListOffersByMerchantID(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	Merchant_uid := c.Params.ByName("merchant_uid")
	offer := &models.OfferAllContent{
		Merchant_uid: Merchant_uid,
	}
	data, msg, err := offer.ListsOfferByMerchant(SERVICE_NAME)

	if msg == "err" {
		c.JSON(200, gin.H{
			"status": 500,
			"error":  err,
		})
	} else {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "Deleted!",
			"data":    data,
		})
	}
}
