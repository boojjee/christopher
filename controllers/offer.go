package controllers

import (
	"christopher/helpers"
	"christopher/models"
	"encoding/json"
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
	Description_th     string `form:"description_th"`
	Lang               string `form:"lang"`
	Value              string `form:"value"`
	Offer_type         string `form:"offer_type"`
}
type offerFormAllLang2 struct {
	Id                 int    `json:"id", Number`
	Offer_uid          string `json:"offer_uid"`
	Merchant_uid       string `json:"merchant_uid"`
	Offer_point        int    `json:"offer_point", Number`
	Offer_cat_id       int    `json:"offer_cat_id", Number`
	Offer_image_banner string `json:"offer_image_banner"`
	Offer_image_poster string `json:"offer_image_poster"`
	Used               int    `json:"used", Number`
	Quantity           int    `json:"quantity", Number`
	Name_en            string `json:"name_en"`
	Name_th            string `json:"name_th" `
	Condition_offer_en string `json:"condition_offer_en"`
	Condition_offer_th string `json:"condition_offer_th"`
	Description_en     string `json:"description_en" `
	Description_th     string `json:"description_th"`
	Value              string `json:"value"`
	Offer_type         string `json:"offer_type"`
	Create_at          int    `json:"Create_at"`
	Update_at          int    `json:"Update_at"`
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
		Status:             1,
		Quantity:           helpers.Convert_string_to_int(form.Quantity),
		Name_en:            form.Name_en,
		Name_th:            form.Name_th,
		Condition_offer_en: form.Condition_offer_en,
		Condition_offer_th: form.Condition_offer_th,
		Description_en:     form.Description_en,
		Description_th:     form.Description_th,
		Value:              form.Value,
		Offer_type:         form.Offer_type,
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
		Value:              form.Value,
		Offer_type:         form.Offer_type,
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
			"status":  500,
			"message": err,
		})
	} else {
		offers := []byte(data)
		OfferSlice := make([]OfferSingle, 0)
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
func ListOffersMerchantIDWithStatus(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	Merchant_uid := c.Params.ByName("merchant_uid")
	offer := &models.OfferAllContent{
		Merchant_uid: Merchant_uid,
	}
	data, msg, err := offer.ListsOffersMerchantStatus(SERVICE_NAME)
	if msg == "err" {
		c.JSON(200, gin.H{
			"status":  500,
			"message": err,
		})
	} else {
		offers := []byte(data)
		OfferSlice := make([]OfferSingle, 0)
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

func ListOffersAll(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	// offer := &models.OfferAllContent{}
	data, msg, err := models.ListsAllOffer(SERVICE_NAME)
	if msg == "err" {
		c.JSON(200, gin.H{
			"status":  500,
			"message": err,
		})
	} else {
		offers := []byte(data)
		OfferSlice := make([]OfferSingle, 0)
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
func ListOffersAllByStatus(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	status := c.Params.ByName("status")
	data, msg, err := models.ListsAllOfferByStatus(SERVICE_NAME, status)
	if msg == "err" {
		c.JSON(200, gin.H{
			"status":  500,
			"message": err,
		})
	} else {
		offers := []byte(data)
		OfferSlice := make([]OfferSingle, 0)
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

func ViewOffer(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	offer_uid := c.Params.ByName("uid")
	offer := &models.OfferAllContent{
		Offer_uid: offer_uid,
	}
	data, msg, err := offer.ShowOfferInfo(SERVICE_NAME)
	if msg == "err" {
		c.JSON(200, gin.H{
			"status":  500,
			"message": err,
		})
	} else {
		offers := []byte(data)
		res := &offerFormAllLang2{}
		err_unmarshal := json.Unmarshal(offers, &res)

		if err_unmarshal != nil {
			c.JSON(200, gin.H{
				"status": 500,
				"error":  err,
			})
		} else {
			c.JSON(200, gin.H{
				"status":  200,
				"message": "Success!",
				"data":    res,
			})
		}
	}
}
