package controllers

import (
	"encoding/json"
	"github.com/boojjee/christopher/helpers"
	"github.com/boojjee/christopher/models"
	"github.com/gin-gonic/gin"
	"log"
	// "strconv"
	"time"
)

type offerCollection []map[string]string
type offerSigle map[string]interface{}
type OfferForm struct {
	Id          string `form:"id" binding:"required"`
	Name        string `form:"name" binding:"required"`
	Offer_point string `form:"offer_point" binding:"required"`
	Condition   string `form:"condition" binding:"required"`
	Cat         string `form:"cat" binding:"required"`
	Merchant_id string `form:"merchant_id" binding:"required"`
	Description string `form:"description" binding:"required"`
	Used        string `form:"used" binding:"required"`
	Qty         string `form:"qty" binding:"required"`
}

func ListOffersAll(c *gin.Context) {
	listOfferAll := models.GetOfferListAll()
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
	id_merchant := c.Params.ByName("id")
	listOfferAll := models.GetOfferListByMerchantID(id_merchant)
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
	offerInfo := models.GetOfferInfo(id_offer)
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
	var form OfferForm
	c.Bind(&form)
	offer := &models.Offer{
		Name:        form.Name,
		Offer_point: helpers.Convert_string_to_float(form.Offer_point),
		Condition:   form.Condition,
		Cat:         helpers.Convert_string_to_int(form.Cat),
		Merchant_id: helpers.Convert_string_to_int(form.Merchant_id),
		Descrtion:   form.Description,
		Used:        helpers.Convert_string_to_int(form.Used),
		Qty:         helpers.Convert_string_to_int(form.Qty),
		Create_at:   time.Now(),
		Update_at:   time.Now(),
	}
	offer.Save()

	c.JSON(200, gin.H{
		"status":  20,
		"message": "Created!",
	})
}

// func UpdateOffer(c *gin.Context) {
// 	id_offer := c.Params.ByName("id")
// 	id_int, _ := strconv.ParseInt(id_offer, 0, 64)
// 	var form MerchantFormEdit
// 	c.Bind(&form)
// 	merchant := &models.Merchant{
// 		Id:               id_int,
// 		Username:         form.Username,
// 		Name:             form.Name,
// 		Email:            form.Email,
// 		Shop_image:       form.Shop_image,
// 		Shop_avatar:      form.Shop_avatar,
// 		Shop_description: form.Shop_description,
// 		Lat:              form.Lat,
// 		Lon:              form.Lon,
// 		Update_at:        time.Now(),
// 	}
// 	merchant.Update()

// 	c.JSON(200, gin.H{
// 		"status":  200,
// 		"message": "Updated",
// 	})
// }
