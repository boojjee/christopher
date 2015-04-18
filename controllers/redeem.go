package controllers

import (
	"christopher/helpers"
	"christopher/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
)

type RedeemMap map[string]interface{}
type RedeemForm struct {
	Pin            string `form:"pin"`
	Offer_uid      string `form:"offer_uid"`
	User_uid       string `form:"user_uid"`
	MyLocation_lat string `form:"myLocation_lat"`
	MyLocation_lon string `form:"myLocation_lon"`
}
type RDJSON struct {
	UserUid     string  `json:"user_uid"`
	RedeemCode  string  `json:"redeem_code"`
	PointUsed   float64 `json:"point_used, Number"`
	BlancePoint float64 `json:"blance_point, Number"`
	ExpireDate  float64 `json:"expire_date, Number"`
	RedeemDate  float64 `json:"redeem_date, Number"`
}

func MyRedeemHistory(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	MyUid := c.Params.ByName("uid")
	redeem := &models.RedeemContent{
		User_uid: MyUid,
	}
	data, msg, err := redeem.GetHistoryRedeem(SERVICE_NAME)
	if msg == "err" {
		c.JSON(200, gin.H{
			"status":  500,
			"message": err,
		})
	} else {
		redeemHis := []byte(data)
		redeemHisMap := make([]RedeemMap, 0)
		err_unmarshal := json.Unmarshal(redeemHis, &redeemHisMap)
		if err_unmarshal != nil {
			c.JSON(200, gin.H{
				"status": 500,
				"error":  err,
			})
		} else {
			c.JSON(200, gin.H{
				"status":  200,
				"message": "Success!",
				"data":    redeemHisMap,
			})
		}
	}
}

func RedeemList(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")

	data, msg, err := models.GetRedeemLists(SERVICE_NAME)
	if msg == "err" {
		c.JSON(200, gin.H{
			"status":  500,
			"message": err,
		})
	} else {
		redeemHis := []byte(data)
		redeemHisMap := make([]RedeemMap, 0)
		err_unmarshal := json.Unmarshal(redeemHis, &redeemHisMap)
		if err_unmarshal != nil {
			c.JSON(200, gin.H{
				"status": 500,
				"error":  err,
			})
		} else {
			c.JSON(200, gin.H{
				"status":  200,
				"message": "Success!",
				"data":    redeemHisMap,
			})
		}
	}
}

func RedeemOffer(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	var form RedeemForm
	c.Bind(&form)

	province, msg1, err := helpers.GetProvinceFromBingMapByPoint(form.MyLocation_lat, form.MyLocation_lon)

	if err != nil || msg1 == "err" {
		log.Println(err)
	}

	redeem := &models.RedeemContent{
		Offer_uid:      form.Offer_uid,
		User_uid:       form.User_uid,
		Redeem_uid:     helpers.RandomStr(10),
		Code:           helpers.RandomStr(8),
		MyLocation_lat: form.MyLocation_lat,
		MyLocation_lon: form.MyLocation_lon,
		Province:       province,
		Pin:            form.Pin,
		Create_at:      helpers.Unix_milisec_time_now(),
		Update_at:      helpers.Unix_milisec_time_now(),
	}

	data, msg, err := redeem.GetCodeRedeem(SERVICE_NAME)

	res := &RDJSON{}
	json.Unmarshal([]byte(data), &res)
	log.Println(res)

	if msg == "err" {
		c.JSON(200, gin.H{
			"status": 500,
			"error":  err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"status":  200,
			"data":    res,
			"message": "Success!",
		})
	}

}
