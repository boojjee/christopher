package controllers

import (
	"christopher/helpers"
	"christopher/models"
	// "encoding/json"
	"github.com/gin-gonic/gin"
	// "log"
	// "strconv"
)

type UserForm struct {
	Id          string
	User_uid    string
	Pin         string `form:"pin"`
	Parse_id    string `form:"parse_id"`
	User_status string `form:"user_status"`
}

func NewUser(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	var form UserForm
	c.Bind(&form)

	user := &models.UserContent{
		Pin:         form.Pin,
		User_uid:    helpers.RandomStr(10),
		Parse_id:    form.Parse_id,
		User_status: helpers.Convert_string_to_int(form.User_status),
		Create_at:   helpers.Unix_milisec_time_now(),
		Update_at:   helpers.Unix_milisec_time_now(),
	}

	msg, err := user.Save(SERVICE_NAME)
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

func UpdateUser(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	user_uid := c.Params.ByName("user_uid")
	var form UserForm
	c.Bind(&form)

	user := &models.UserContent{
		Pin:         form.Pin,
		Parse_id:    form.Parse_id,
		User_status: helpers.Convert_string_to_int(form.User_status),
		Update_at:   helpers.Unix_milisec_time_now(),
	}

	msg, err := user.Update(SERVICE_NAME, user_uid)
	if msg == "err" {
		c.JSON(200, gin.H{
			"status": 500,
			"error":  err,
		})
	} else {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "Updated!",
		})
	}
}

func DeleteUser(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	user_uid := c.Params.ByName("user_uid")
	var form UserForm
	c.Bind(&form)

	user := &models.UserContent{
		User_uid: user_uid,
	}
	msg, err := user.Delete(SERVICE_NAME)
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