package controllers

import (
	"christopher/helpers"
	"christopher/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	// "strconv"
)

type MerchantCollection []map[string]string
type MerchantSigle map[string]interface{}
type MerchantForm struct {
	Username             string `form:"username"`
	Name                 string `form:"name"`
	Password             string `form:"password"`
	Email                string `form:"email"`
	Shop_avatar_s        string `form:"shop_avatar_s"`
	Shop_avatar_l        string `form:"shop_avatar_l"`
	Shop_description     string `form:"shop_description"`
	Lat                  string `form:"lat"`
	Lon                  string `form:"lon"`
	Merchant_category_id string `form:"merchant_category_id" `
	Merchant_province    string `form:"merchant_province"`
}

type MerchantFormAllLang struct {
	Username             string `form:"username"`
	Password             string `form:"password"`
	Email                string `form:"email"`
	Shop_avatar_s        string `form:"shop_avatar_s"`
	Shop_avatar_l        string `form:"shop_avatar_l"`
	Lat                  string `form:"lat"`
	Lon                  string `form:"lon"`
	Province             string `form:"province"`
	Phone_1              string `form:"phone_1"`
	Phone_2              string `form:"phone_2"`
	Fax                  string `form:"fax"`
	Line_id              string `form:"line_id"`
	Facebook_link        string `form:"facebook_link"`
	Website_link         string `form:"website_link"`
	Name_en              string `form:"name_en"`
	Name_th              string `form:"name_th"`
	Shop_description_en  string `form:"shop_description_en"`
	Shop_description_th  string `form:"shop_description_th"`
	Merchant_status      string `form:"merchant_status"`
	Merchant_category_id string `form:"merchant_category_id"`
	Merchant_province    string `form:"merchant_province"`
}

type MerchantFormEdit struct {
	Username             string `form:"username"`
	Name                 string `form:"name"`
	Password             string `form:"password"`
	Email                string `form:"email"`
	Shop_avatar_s        string `form:"shop_avatar_s"`
	Shop_avatar_l        string `form:"shop_avatar_l"`
	Shop_description     string `form:"shop_description"`
	Lat                  string `form:"lat"`
	Lon                  string `form:"lon"`
	Province             string `form:"province"`
	Merchant_category_id string `form:"merchant_category_id"`
	Merchant_province    string `form:"merchant_province"`
}

type MerchantAuthenForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type Merchants struct {
	Id                   int    `json:"id, Number"`
	Username             string `json:"username"`
	Name                 string `json:"name"`
	Password             string `json:"password"`
	Email                string `json:"email"`
	Shop_avatar_s        string `json:"shop_avatar_s"`
	Shop_avatar_l        string `json:"shop_avatar_s"`
	Shop_description     string `json:"shop_description"`
	Lat                  string `json:"lat"`
	Lon                  string `json:"lon"`
	Province             string `form:"province"`
	Create_at            int    `json:"create_at, Number"`
	Update_at            int    `json:"update_at, Number"`
	Merchant_category_id int    `json:"merchant_category_id, Number"`
	Merchant_province    string `form:"merchant_province"`
}
type MerchantFormAllLangJson struct {
	Id                   int    `json:"id, Number"`
	Username             string `json:"username"`
	Password             string `json:"password"`
	Email                string `json:"email"`
	Shop_avatar_s        string `json:"shop_avatar_s"`
	Shop_avatar_l        string `json:"shop_avatar_l"`
	Merchant_uid         string `json:"merchant_uid"`
	Lat                  string `json:"lat"`
	Lon                  string `json:"lon"`
	Province             string `form:"province"`
	Phone_1              string `json:"phone_1"`
	Phone_2              string `json:"phone_2"`
	Fax                  string `json:"fax"`
	Line_id              string `json:"line_id"`
	Facebook_link        string `json:"facebook_link"`
	Website_link         string `json:"website_link"`
	Name_en              string `json:"name_en"`
	Name_th              string `json:"name_th"`
	Shop_description_en  string `json:"shop_description_en"`
	Shop_description_th  string `json:"shop_description_th"`
	Merchant_status      string `json:"merchant_status"`
	Merchant_category_id int    `json:"merchant_category_id, Number"`
	Merchant_province    string `form:"merchant_province"`
	Create_at            int    `json:"create_at, Number"`
	Update_at            int    `json:"update_at, Number"`
}

type MerchantFormByLangJson struct {
	Id                   int    `json:"id, Number"`
	Username             string `json:"username"`
	Password             string `json:"password"`
	Email                string `json:"email"`
	Shop_avatar_s        string `json:"shop_avatar_s"`
	Shop_avatar_l        string `json:"shop_avatar_l"`
	Lat                  string `json:"lat"`
	Lon                  string `json:"lon"`
	Province             string `form:"province"`
	Phone_1              string `json:"phone_1"`
	Phone_2              string `json:"phone_2"`
	Fax                  string `json:"fax"`
	Line_id              string `json:"line_id"`
	Facebook_link        string `json:"facebook_link"`
	Website_link         string `json:"website_link"`
	Name                 string `json:"name"`
	Shop_description     string `json:"shop_description"`
	Merchant_status      string `json:"merchant_status"`
	Merchant_category_id int    `json:"merchant_category_id_id, Number"`
	Merchant_province    string `form:"merchant_province"`
	Create_at            int    `json:"create_at, Number"`
	Update_at            int    `json:"update_at, Number"`
}

type MerchantMetaJSON struct {
	Username             string `json:"username"`
	Name                 string `json:"name"`
	Password             string `json:"password"`
	Email                string `json:"email"`
	Shop_avatar_s        string `json:"shop_avatar_s"`
	Shop_avatar_l        string `json:"shop_avatar_l"`
	Shop_description     string `json:"shop_description"`
	Lat                  string `json:"lat"`
	Lon                  string `json:"lon"`
	Province             string `form:"province"`
	Create_at            int    `json:"create_at, Number"`
	Update_at            int    `json:"update_at, Number"`
	Merchant_category_id int    `json:"merchant_category_id_id, Number"`
	Merchant_province    string `form:"merchant_province"`
}
type MerchantContentJSON struct {
	Username             string `json:"username"`
	Name                 string `json:"name"`
	Password             string `json:"password"`
	Email                string `json:"email"`
	Shop_avatar_s        string `json:"shop_avatar_s"`
	Shop_avatar_l        string `json:"shop_avatar_l"`
	Shop_description     string `json:"shop_description"`
	Lat                  string `json:"lat"`
	Lon                  string `json:"lon"`
	Province             string `form:"province"`
	Create_at            int    `json:"create_at, Number"`
	Update_at            int    `json:"update_at, Number"`
	Merchant_category_id int    `json:"merchant_category_id_id, Number"`
	Merchant_province    string `form:"merchant_province"`
}

type MerchantAllJSON struct {
	Username             string `json:"username"`
	Name                 string `json:"name"`
	Password             string `json:"password"`
	Email                string `json:"email"`
	Shop_avatar_s        string `json:"shop_avatar_s"`
	Shop_avatar_l        string `json:"shop_avatar_l"`
	Shop_description     string `json:"shop_description"`
	Lat                  string `json:"lat"`
	Lon                  string `json:"lon"`
	Province             string `form:"province"`
	Create_at            int    `json:"create_at, Number"`
	Update_at            int    `json:"update_at, Number"`
	Merchant_category_id int    `json:"merchant_category_id_id, Number"`
	Merchant_province    string `form:"merchant_province"`
}

// Action CRUD

func NewMerchant(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	var form MerchantFormAllLang
	c.Bind(&form)

	merchant := &models.MerchantMeta{
		Username:             form.Username,
		Password:             form.Password,
		Email:                form.Email,
		Merchant_uid:         helpers.RandomStr(10),
		Shop_avatar_s:        form.Shop_avatar_s,
		Shop_avatar_l:        form.Shop_avatar_l,
		Lat:                  form.Lat,
		Lon:                  form.Lon,
		Province:             form.Province,
		Phone_1:              form.Phone_1,
		Phone_2:              form.Phone_2,
		Fax:                  form.Fax,
		Line_id:              form.Line_id,
		Facebook_link:        form.Facebook_link,
		Website_link:         form.Website_link,
		Name_en:              form.Name_en,
		Name_th:              form.Name_th,
		Shop_description_en:  form.Shop_description_en,
		Shop_description_th:  form.Shop_description_th,
		Merchant_status:      "0",
		Merchant_category_id: helpers.Convert_string_to_int(form.Merchant_category_id),
		Merchant_province:    form.Merchant_province,
		Create_at:            helpers.Unix_milisec_time_now(),
		Update_at:            helpers.Unix_milisec_time_now(),
	}

	msg, err := merchant.Save(SERVICE_NAME)

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

func UpdateMerchant(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	uid_merchant := c.Params.ByName("uid")
	var form MerchantFormAllLang
	c.Bind(&form)
	merchant := &models.MerchantMeta{
		Username:             form.Username,
		Email:                form.Email,
		Merchant_uid:         uid_merchant,
		Shop_avatar_s:        form.Shop_avatar_s,
		Shop_avatar_l:        form.Shop_avatar_l,
		Lat:                  form.Lat,
		Lon:                  form.Lon,
		Province:             form.Province,
		Phone_1:              form.Phone_1,
		Phone_2:              form.Phone_2,
		Fax:                  form.Fax,
		Line_id:              form.Line_id,
		Facebook_link:        form.Facebook_link,
		Website_link:         form.Website_link,
		Name_en:              form.Name_en,
		Name_th:              form.Name_th,
		Merchant_status:      form.Merchant_status,
		Merchant_category_id: helpers.Convert_string_to_int(form.Merchant_category_id),
		Merchant_province:    form.Merchant_province,
		Shop_description_en:  form.Shop_description_en,
		Shop_description_th:  form.Shop_description_th,
		Create_at:            helpers.Unix_milisec_time_now(),
		Update_at:            helpers.Unix_milisec_time_now(),
	}
	msg, err := merchant.Update(SERVICE_NAME)

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

func DeleteMerchant(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	uid_merchant := c.Params.ByName("uid")
	merchant := &models.MerchantMeta{
		Merchant_uid: uid_merchant,
	}

	msg, err := merchant.Delete(SERVICE_NAME)
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

func AuthenMechant(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	var form MerchantAuthenForm
	c.Bind(&form)
	merchant := &models.MerchantMeta{
		Username: form.Username,
		Password: form.Password,
	}
	msg, err := merchant.Authen(SERVICE_NAME)
	if err != nil {
		log.Println(err)
	}
	if msg == "fail" {
		c.JSON(200, gin.H{
			"status":  500,
			"message": "Authen Fail!",
			"data":    0,
		})
	} else {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "Success",
			"data":    1,
		})
	}
}

// ----------------- ready

func GetMerchantInfo(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	uid := c.Params.ByName("uid")

	merchant := &models.MerchantMeta{
		Merchant_uid: uid,
	}

	data, msg, err := merchant.MerchantShowInfoAllLang(SERVICE_NAME)

	if msg == "err" {
		c.JSON(200, gin.H{
			"status":  500,
			"message": err,
		})
	} else {
		// log.Println(data)
		mercahnts := []byte(data)
		res := &MerchantFormAllLangJson{}
		err_unmarshal := json.Unmarshal(mercahnts, &res)

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

func GetMerchantInfoByLang(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	uid := c.Params.ByName("uid")
	lang := c.Params.ByName("lang")

	merchant := &models.MerchantMeta{
		Merchant_uid: uid,
	}

	data, msg, err := merchant.MerchantShowInfoByLang(SERVICE_NAME, lang)

	if msg == "err" {
		c.JSON(200, gin.H{
			"status":  500,
			"message": err,
		})
	} else {
		// log.Println(data)
		mercahnts := []byte(data)
		res := &MerchantFormByLangJson{}
		err_unmarshal := json.Unmarshal(mercahnts, &res)

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

// -------

func GetMerchantsLists_All(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")

	data, msg, err := models.MerchantListAllLang(SERVICE_NAME)

	if msg == "err" {
		c.JSON(200, gin.H{
			"status":  500,
			"message": err,
		})
	} else {
		// log.Println(data)
		mercahnts := []byte(data)
		merchant_slice := make([]MerchantSigle, 0)
		err_unmarshal := json.Unmarshal(mercahnts, &merchant_slice)

		if err_unmarshal != nil {
			c.JSON(200, gin.H{
				"status":  500,
				"message": "json error",
			})
		}
		c.JSON(200, gin.H{
			"status":  200,
			"message": "Success!",
			"data":    merchant_slice,
		})
	}
}

func GetMerchantsListsByLang(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	lang := c.Params.ByName("lang")
	data, msg, err := models.MerchantListByLang(SERVICE_NAME, lang)

	if msg == "err" {
		c.JSON(200, gin.H{
			"status":  500,
			"message": err,
		})
	} else {
		// log.Println(data)
		mercahnts := []byte(data)
		merchant_slice := make([]MerchantSigle, 0)
		err_unmarshal := json.Unmarshal(mercahnts, &merchant_slice)

		if err_unmarshal != nil {
			c.JSON(200, gin.H{
				"status":  500,
				"message": "json error",
			})
		}
		c.JSON(200, gin.H{
			"status":  200,
			"message": "Success!",
			"data":    merchant_slice,
		})
	}
}

// -------
func GetMerchantsListsWithGallery_All(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	data, msg, err := models.MerchantListWithGallery(SERVICE_NAME)

	if msg == "err" {
		c.JSON(200, gin.H{
			"status":  500,
			"message": err,
		})
	} else {
		// log.Println(data)
		mercahnts := []byte(data)
		merchant_slice := make([]MerchantSigle, 0)
		err_unmarshal := json.Unmarshal(mercahnts, &merchant_slice)

		if err_unmarshal != nil {
			c.JSON(200, gin.H{
				"status":  500,
				"message": "json error",
			})
		}
		c.JSON(200, gin.H{
			"status":  200,
			"message": "Success!",
			"data":    merchant_slice,
		})
	}
}

func GetMerchantsListsWithGalleryByLang(c *gin.Context) {
	SERVICE_NAME := c.Params.ByName("service_name")
	lang := c.Params.ByName("lang")
	data, msg, err := models.MerchantListWithGalleryByLang(SERVICE_NAME, lang)

	if msg == "err" {
		c.JSON(200, gin.H{
			"status":  500,
			"message": err,
		})
	} else {
		// log.Println(data)
		mercahnts := []byte(data)
		merchant_slice := make([]MerchantSigle, 0)
		err_unmarshal := json.Unmarshal(mercahnts, &merchant_slice)

		if err_unmarshal != nil {
			c.JSON(200, gin.H{
				"status":  500,
				"message": "json error",
			})
		}
		c.JSON(200, gin.H{
			"status":  200,
			"message": "Success!",
			"data":    merchant_slice,
		})
	}
}
