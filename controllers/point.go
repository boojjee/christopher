package controllers

import (
	// "christopher/helpers"
	"christopher/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	// "log"
)

type BalancePointForm struct {
	Id           string
	User_uid     string
	Blance_point string `form:"blance_point"`
}
type Mygpoint struct {
	G_Point float64 `json:"g_point"`
}
type WorkOutAll struct {
	Workout float64 `json:"workout"`
}

func GetMyPoint(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	user_uid := c.Params.ByName("user_uid")
	var form BalancePointForm
	c.Bind(&form)

	myBPoint := &models.MyBPoint{
		User_uid: user_uid,
	}

	result, msg, err := myBPoint.GetMyCurrentPoint(SERVICE_NAME)
	mapD := map[string]float64{"g_point": result}
	mapB, _ := json.Marshal(mapD)
	res := &Mygpoint{}
	json.Unmarshal(mapB, &res)

	if msg == "err" {
		c.JSON(200, gin.H{
			"status": 500,
			"error":  err,
		})
	} else {
		c.JSON(200, gin.H{
			"status":  200,
			"data":    res,
			"message": "Success!",
		})
	}
}

func GetWorkOut(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")

	result, msg, err := models.GetWorkOut(SERVICE_NAME)
	mapD := map[string]float64{"workout": result}
	mapB, _ := json.Marshal(mapD)
	res := &WorkOutAll{}
	json.Unmarshal(mapB, &res)

	if msg == "err" {
		c.JSON(200, gin.H{
			"status": 500,
			"error":  err,
		})
	} else {
		c.JSON(200, gin.H{
			"status":  200,
			"data":    res,
			"message": "Success!",
		})
	}
}

func GetPoints(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")

	result, msg, err := models.GetPoints(SERVICE_NAME)
	mapD := map[string]float64{"g_point": result}
	mapB, _ := json.Marshal(mapD)
	res := &Mygpoint{}
	json.Unmarshal(mapB, &res)

	if msg == "err" {
		c.JSON(200, gin.H{
			"status": 500,
			"error":  err,
		})
	} else {
		c.JSON(200, gin.H{
			"status":  200,
			"data":    res,
			"message": "Success!",
		})
	}
}
