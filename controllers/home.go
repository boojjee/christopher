package controllers

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func Home(c *gin.Context) {
	var data mytype
	file, err := ioutil.ReadFile("test.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{"status": "OK!", "build": data})
}
