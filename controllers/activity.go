package controllers

import (
	"christopher/helpers"
	"christopher/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	// "log"
)

type ActivitiesObj map[string]interface{}
type ActivityForm struct {
	Id                 string
	Activity_uid       string `form:"activity_uid"`
	User_uid           string `form:"user_uid"`
	User_parse_id      string `form:"user_parse_id"`
	Third_activity_id  string `form:"third_activity_id"`
	Third_uri          string `form:"third_uri"`
	Third_token_user   string `form:"third_token_user"`
	Source             string `form:"source"`
	Distance           string `form:"distance"`
	Duration           string `form:"duration"`
	Calories           string `form:"calories"`
	Start_activity_lat string `form:"start_activity_lat"`
	Start_activity_lon string `form:"start_activity_lon"`
	Activity_type      string `form:"activity_type"`
	Activity_status    string `form:"activity_status"`
	MyLocation_lat     string `form:"mylocation_lat"`
	MyLocation_lon     string `form:"mylocation_lon"`
}
type ActivityWithPoint struct {
	Id                 int64             `json:"id, Number"`
	Activity_uid       string            `json:"activity_uid"`
	User_uid           string            `json:"user_uid"`
	User_parse_id      string            `json:"user_parse_id"`
	Third_activity_id  string            `json:"third_activity_id"`
	Third_uri          string            `json:"third_uri"`
	Third_token_user   string            `json:"third_token_user"`
	Source             string            `json:"source"`
	Distance           float64           `json:"distance, Number"`
	Duration           float64           `json:"duration, Number"`
	Calories           float64           `json:"calories, Number"`
	G_Point            gpointForActivity `json:"g_point"`
	Start_activity_lat float64           `json:"start_activity_lat, Number"`
	Start_activity_lon float64           `json:"start_activity_lon, Number"`
	Activity_type      string            `json:"activity_type"`
	Activity_status    int64             `json:"activity_status, Number"`
	MyLocation_lat     string            `json:"mylocation_lat, Number"`
	MyLocation_lon     string            `json:"mylocation_lon, Number"`
	Create_at          int64             `json:"create_at, Number"`
	Update_at          int64             `json:"update_at, Number"`
}
type gpointForActivity struct {
	G_point   float64 `json:"point"`
	Create_at int64   `json:"create_at"`
	Update_at int64   `json:"update_at"`
}

type ActivityIdPagination struct {
	Next_max_id int64 `json:"next_max_id"`
	Next_min_id int64 `json:"next_min_id"`
}

func NewActivity(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	var form ActivityForm
	c.Bind(&form)
	activity := &models.ActivityContentForm{
		Activity_uid:       helpers.RandomStr(10),
		User_uid:           form.User_uid,
		User_parse_id:      form.User_parse_id,
		Third_activity_id:  form.Third_activity_id,
		Third_uri:          form.Third_uri,
		Third_token_user:   form.Third_token_user,
		Source:             form.Source,
		Distance:           helpers.Convert_string_to_float(form.Distance),
		Duration:           helpers.Convert_string_to_float(form.Duration),
		Calories:           helpers.Convert_string_to_float(form.Calories),
		Start_activity_lat: helpers.Convert_string_to_float(form.Start_activity_lat),
		Start_activity_lon: helpers.Convert_string_to_float(form.Start_activity_lon),
		Activity_type:      form.Activity_type,
		Activity_status:    helpers.Convert_string_to_int(form.Activity_status),
		MyLocation_lat:     helpers.Convert_string_to_float(form.MyLocation_lat),
		MyLocation_lon:     helpers.Convert_string_to_float(form.MyLocation_lon),
		Create_at:          helpers.Unix_milisec_time_now(),
		Update_at:          helpers.Unix_milisec_time_now(),
	}

	msg, err := activity.Save(SERVICE_NAME)
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

func UpdateActivity(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	activity_uid := c.Params.ByName("activity_uid")
	var form ActivityForm
	c.Bind(&form)
	activity := &models.ActivityContent{
		Activity_uid:       activity_uid,
		User_uid:           form.User_uid,
		User_parse_id:      form.User_parse_id,
		Third_activity_id:  form.Third_activity_id,
		Third_uri:          form.Third_uri,
		Third_token_user:   form.Third_token_user,
		Source:             form.Source,
		Distance:           helpers.Convert_string_to_float(form.Distance),
		Duration:           helpers.Convert_string_to_float(form.Duration),
		Calories:           helpers.Convert_string_to_float(form.Calories),
		Start_activity_lat: helpers.Convert_string_to_float(form.Start_activity_lat),
		Start_activity_lon: helpers.Convert_string_to_float(form.Start_activity_lon),
		Activity_type:      form.Activity_type,
		Activity_status:    helpers.Convert_string_to_int(form.Activity_status),
		Update_at:          helpers.Unix_milisec_time_now(),
	}
	msg, err := activity.Update(SERVICE_NAME)
	if msg == "err" {
		c.JSON(200, gin.H{
			"status": 500,
			"error":  err,
		})
	} else {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "Update!",
		})
	}
}

func DeleteActivity(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	activity_uid := c.Params.ByName("activity_uid")
	var form UserForm
	c.Bind(&form)

	activity := &models.ActivityContent{
		Activity_uid: activity_uid,
	}
	msg, err := activity.Delete(SERVICE_NAME)
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

func GetActivity(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	activity_uid := c.Params.ByName("activity_uid")

	data, msg, err := models.GetActivityByAcUID(SERVICE_NAME, activity_uid)

	if msg == "err" {

		c.JSON(200, gin.H{
			"status":  500,
			"message": err,
		})

	} else {
		activities := []byte(data)
		res := &ActivityWithPoint{}
		err_unmarshal := json.Unmarshal(activities, &res)

		if err_unmarshal != nil {
			c.JSON(200, gin.H{
				"status":  500,
				"message": "json error",
			})
		}

		c.JSON(200, gin.H{
			"status":  200,
			"message": "Success!",
			"data":    res,
		})
	}
}

func ActivityListsAll(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	user_uid := c.Params.ByName("user_uid")
	data, msg, err := models.GetActivityListsAll(SERVICE_NAME, user_uid)
	if msg == "err" {
		c.JSON(200, gin.H{
			"status":  500,
			"message": err,
		})
	} else {
		// log.Println(data)
		activities := []byte(data)
		activities_slice := make([]ActivitiesObj, 0)
		err_unmarshal := json.Unmarshal(activities, &activities_slice)

		// log.Println(actpage)

		if err_unmarshal != nil {
			c.JSON(200, gin.H{
				"status":  500,
				"message": "json error",
			})
		}

		c.JSON(200, gin.H{
			"status":  200,
			"message": "Success!",
			"data":    activities_slice,
		})
	}
}

func LatestActivityList(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	user_uid := c.Params.ByName("user_uid")
	data, pagination, msg, err := models.GetLatestActivityList(SERVICE_NAME, user_uid)
	if msg == "err" {
		c.JSON(200, gin.H{
			"status":  500,
			"message": err,
		})
	} else {
		// log.Println(data)
		activities := []byte(data)
		activities_slice := make([]ActivitiesObj, 0)
		err_unmarshal := json.Unmarshal(activities, &activities_slice)

		// log.Println(pagination)
		actpagination := []byte(pagination)
		actpage := &ActivityIdPagination{}
		err_unmarshal1 := json.Unmarshal(actpagination, &actpage)

		// log.Println(actpage)

		if err_unmarshal1 != nil {
			c.JSON(200, gin.H{
				"status":  500,
				"message": "json error",
			})
		}

		if err_unmarshal != nil {
			c.JSON(200, gin.H{
				"status":  500,
				"message": "json error",
			})
		}
		c.JSON(200, gin.H{
			"status":     200,
			"message":    "Success!",
			"pagination": actpage,
			"data":       activities_slice,
		})
	}
}

func NextActivityList(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	Max_id := c.Params.ByName("max_id")
	user_uid := c.Params.ByName("user_uid")
	data, pagination, msg, err := models.GetNextActivityList(SERVICE_NAME, user_uid, Max_id)
	if msg == "err" {
		c.JSON(200, gin.H{
			"status":  500,
			"message": err,
		})
	} else {
		// log.Println(data)
		activities := []byte(data)
		activities_slice := make([]ActivitiesObj, 0)
		err_unmarshal := json.Unmarshal(activities, &activities_slice)

		// log.Println(pagination)
		actpagination := []byte(pagination)
		actpage := &ActivityIdPagination{}
		err_unmarshal1 := json.Unmarshal(actpagination, &actpage)

		// log.Println(actpage)

		if err_unmarshal1 != nil {
			c.JSON(200, gin.H{
				"status":  500,
				"message": "json error",
			})
		}

		if err_unmarshal != nil {
			c.JSON(200, gin.H{
				"status":  500,
				"message": "json error",
			})
		}
		c.JSON(200, gin.H{
			"status":     200,
			"message":    "Success!",
			"pagination": actpage,
			"data":       activities_slice,
		})
	}
}

func PrevActivityList(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	Min_id := c.Params.ByName("min_id")
	user_uid := c.Params.ByName("user_uid")
	data, pagination, msg, err := models.GetPrevActivityList(SERVICE_NAME, user_uid, Min_id)
	if msg == "err" {
		c.JSON(200, gin.H{
			"status":  500,
			"message": err,
		})
	} else {
		// log.Println(data)
		activities := []byte(data)
		activities_slice := make([]ActivitiesObj, 0)
		err_unmarshal := json.Unmarshal(activities, &activities_slice)

		// log.Println(pagination)
		actpagination := []byte(pagination)
		actpage := &ActivityIdPagination{}
		err_unmarshal1 := json.Unmarshal(actpagination, &actpage)

		// log.Println(actpage)

		if err_unmarshal1 != nil {
			c.JSON(200, gin.H{
				"status":  500,
				"message": "json error",
			})
		}

		if err_unmarshal != nil {
			c.JSON(200, gin.H{
				"status":  500,
				"message": "json error",
			})
		}
		c.JSON(200, gin.H{
			"status":     200,
			"message":    "Success!",
			"pagination": actpage,
			"data":       activities_slice,
		})
	}
}
