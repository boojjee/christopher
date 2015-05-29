package controllers

import (
	// "christopher/helpers"
	// "christopher/models"
	"github.com/gin-gonic/gin"
)

type EventFormAllLang struct {
	Name_en            string  `form:"name_en"`
	Name_th            string  `form:"name_th"`
	Detail_th          string  `form:"detail_th"`
	Detail_en          string  `form:"detail_en"`
	Term_of_use_th     string  `form:"term_of_use_th"`
	Term_of_use_en     string  `form:"term_of_use_en"`
	Event_condition_th string  `form:"event_condition_th"`
	Event_condition_en string  `form:"event_condition_en"`
	Event_image_banner string  `form:"event_image_banner"`
	All_province       int64   `form:"all_province"`
	Province           string  `form:"province"`
	Event_start_date   int64   `form:"event_start_date"`
	Event_end_date     int64   `form:"event_end_date"`
	Spacial_value      int64   `form:"spacial_value"`
	Spacial_value_type int64   `form:"spacial_value_type "`
	Quantity           int64   `form:"quantity"`
	Status             int64   `form:"status"`
	Location_lat       float64 `form:"location_lat"`
	Location_lon       float64 `form:"location_lon"`
}

func ListAllEvents(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	c.JSON(200, gin.H{
		"status": 200,
		"data":   SERVICE_NAME,
	})
}

func NewEvent(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	var form EventFormAllLang
	c.Bind(&form)
	c.JSON(200, gin.H{
		"status": 200,
		"data":   SERVICE_NAME,
	})
}
