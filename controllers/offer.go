package controllers

import (
	// "christopher/helpers"
	// "christopher/models"
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
	Qty                string `form:"qty"`
	Name_en            string `from:"name_en"`
	Name_th            string `from:"name_th" `
	Condition_offer_en string `from:"condition_offer_en"`
	Condition_offer_th string `from:"condition_offer_th"`
	Description_en     string `from:"description_en" `
	Description_th     string `from:"description_en"`
	Lang               string `from:"lang"`
}

func NewOffer(c *gin.Context) {
	// SERVICE_NAME := c.Params.ByName("service_name")
	// var form offerFormAllLang
	// c.Bind(&form)
	// offer := &models.OfferAllContent{
	// 	Offer_uid:          helpers.RandomStr(10),
	// 	Merchant_uid:       form.Merchant_uid,
	// 	Offer_cat_id:       helpers.Convert_string_to_int(form.Offer_cat_id),
	// 	Offer_image_banner: form.Offer_image_banner,
	// 	Offer_image_poster: form.Offer_image_poster,
	// 	Used:               helpers.Convert_string_to_int(form.Used),
	// 	Qty:                helpers.Convert_string_to_int(form.Qty),
	// 	Name_en:            form.Name_en,
	// 	Name_th:            form.Name_th,
	// 	Condition_offer_en: form.Condition_offer_en,
	// 	Condition_offer_th: form.Condition_offer_th,
	// 	Description_en:     form.Description_en,
	// 	Description_th:     form.Description_th,
	// 	Lang:               form.Lang,
	// 	Create_at:          helpers.Unix_milisec_time_now(),
	// 	Update_at:          helpers.Unix_milisec_time_now(),
	// }

}
