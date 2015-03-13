package controllers

import (
	// "christopher/models"
	"github.com/gin-gonic/gin"
	"log"
)

// get merchant info with image gallery
// create
func NewMerchantGaller(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	log.Println(SERVICE_NAME)
}
func UpdateMerchantGaller(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	log.Println(SERVICE_NAME)
}
func DeleteMerchantGaller(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	log.Println(SERVICE_NAME)
}

// update
// delete
