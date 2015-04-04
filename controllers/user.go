package controllers

import (
	"christopher/helpers"
	"christopher/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	// "strconv"
)

type UserForm struct {
	Id          string
	User_uid    string
	Pin         string `form:"pin"`
	Parse_id    string `form:"parse_id"`
	User_status string `form:"user_status"`
}

type UserRJSON struct {
	User_uid string `json:"user_uid"`
}

func NewUser(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	var form UserForm
	c.Bind(&form)

	user := &models.UserContent{
		Pin:         form.Pin,
		User_uid:    helpers.RandomStr(10),
		Parse_id:    form.Parse_id,
		User_status: 1,
		Create_at:   helpers.Unix_milisec_time_now(),
		Update_at:   helpers.Unix_milisec_time_now(),
	}
	data, msg, err := user.Save(SERVICE_NAME)
	jsondata := `{ "user_uid": "` + data + `"}`
	res := &UserRJSON{}
	json.Unmarshal([]byte(jsondata), &res)

	if msg == "err" {
		c.JSON(200, gin.H{
			"status": 500,
			"error":  err,
		})
	} else {
		c.JSON(200, gin.H{
			"status":  200,
			"data":    res,
			"message": "Created!",
		})
	}
}

func UpdateUserPin(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	user_uid := c.Params.ByName("user_uid")
	var form UserForm
	c.Bind(&form)

	log.Println(user_uid)
	log.Println("DDDDD")
	user := &models.UserContent{
		Pin:       form.Pin,
		Update_at: helpers.Unix_milisec_time_now(),
	}
	msg, err := user.UpdatePin(SERVICE_NAME, user_uid)
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

func GetUserUID(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	parse_id := c.Params.ByName("parse_id")
	var form UserForm
	c.Bind(&form)

	user := &models.UserContent{
		Parse_id: parse_id,
	}
	data, msg, err := user.GetUIDByParseID(SERVICE_NAME)
	jsondata := `{ "user_uid": "` + data + `"}`
	res := &UserRJSON{}
	json.Unmarshal([]byte(jsondata), &res)
	if msg == "err" {
		c.JSON(200, gin.H{
			"status": 500,
			"error":  err,
		})
	} else {
		c.JSON(200, gin.H{
			"status":  200,
			"data":    res,
			"message": "Created!",
		})
	}
}
