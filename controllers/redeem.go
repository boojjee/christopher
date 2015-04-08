package controllers

import (
	"christopher/helpers"
	"christopher/models"
	"github.com/gin-gonic/gin"
)

type RedeemForm struct {
}

func RedeemOffer(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	redeem := &models.RedeemContent{
		Offer_uid: "ddddd",
		Code:      helpers.RandomStr(8),
	}

	data, msg, err := redeem.GetCodeRedeem(SERVICE_NAME)

	if msg == "err" {
		c.JSON(200, gin.H{
			"status": 500,
			"error":  err,
		})
	} else {
		c.JSON(200, gin.H{
			"status":  200,
			"data":    data,
			"message": "Success!",
		})
	}

}
