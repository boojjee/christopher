package controllers

import (
	"christopher/helpers"
	"christopher/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type PointSettingObj map[string]interface{}
type PointSettingForm struct {
	Id              string
	PointSettingUID string
	Point_type      string `form:"point_type"`
	Constant_point  string `form:"constant_point"`
	Description     string `form:"description"`
}

func NewPointSetting(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	var form PointSettingForm
	c.Bind(&form)

	formSettingPoint := &models.PointSetting{
		Point_type:      form.Point_type,
		PointSettingUID: helpers.RandomStr(10),
		Constant_point:  helpers.Convert_string_to_float(form.Constant_point),
		Description:     form.Description,
		Create_at:       helpers.Unix_milisec_time_now(),
		Update_at:       helpers.Unix_milisec_time_now(),
	}

	result, err := formSettingPoint.CreatePointSetting(SERVICE_NAME)
	if result == "err" {
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

func UpdatePointSetting(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	setting_uid := c.Params.ByName("uid")

	var form PointSettingForm
	c.Bind(&form)
	formSettingPoint := &models.PointSetting{
		PointSettingUID: setting_uid,
		Point_type:      form.Point_type,
		Constant_point:  helpers.Convert_string_to_float(form.Constant_point),
		Description:     form.Description,
		Update_at:       helpers.Unix_milisec_time_now(),
	}
	result, err := formSettingPoint.UpdatePointSetting(SERVICE_NAME)
	if result == "err" {
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

func DeletePointSetting(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	setting_uid := c.Params.ByName("uid")
	var form PointSettingForm
	c.Bind(&form)
	formSettingPoint := &models.PointSetting{
		PointSettingUID: setting_uid,
	}
	result, err := formSettingPoint.DeletePointSetting(SERVICE_NAME)
	if result == "err" {
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

func ListAllPointSetting(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	data, msg, err := models.ListsAllSettingPoint(SERVICE_NAME)

	if msg == "err" {
		c.JSON(200, gin.H{
			"status":  500,
			"message": err,
		})
	} else {
		settingPoints := []byte(data)
		settingPoints_slice := make([]PointSettingObj, 0)
		err_unmarshal := json.Unmarshal(settingPoints, &settingPoints_slice)

		if err_unmarshal != nil {
			c.JSON(200, gin.H{
				"status":  500,
				"message": "json error",
			})
		}
		c.JSON(200, gin.H{
			"status":  200,
			"message": "Success!",
			"data":    settingPoints_slice,
		})
	}
}
