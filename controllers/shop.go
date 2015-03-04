package controllers

import (
	"christopher/models"
	"encoding/json"
	// "fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type ShopCollection []map[string]string
type Shops struct {
	id    int
	title string
}

type ShopForm struct {
	Name        string `form:"name" binding:"required"`
	Description string `form:"description" binding:"required"`
}

func ListShop(c *gin.Context) {
	// list all shop
	ss := models.GetShopLists()
	var shops ShopCollection
	file := []byte(ss)

	err := json.Unmarshal(file, &shops)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{"data": shops})
}

func NewShop(c *gin.Context) {
	var form ShopForm
	c.Bind(&form)
	message := "Hello " + form.Name + form.Description
	c.String(200, message)
}
