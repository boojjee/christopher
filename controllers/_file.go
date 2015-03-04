package controllers

import (
	"encoding/json"
	"fit/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type ShopCollection []map[string]string
type Shops struct {
	id    int
	title string
}

func ShopIndex(c *gin.Context) {
	ss := models.GetShopLists()
	var shops ShopCollection
	file := []byte(ss)

	err := json.Unmarshal(file, &shops)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{"data": shops})
}
