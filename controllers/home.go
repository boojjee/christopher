package controllers

import (
	// "encoding/json"
	"github.com/gin-gonic/gin"
	// "io/ioutil"
	// "log"
)

// type mytype []map[string]string

func Home(c *gin.Context) {
	// var data mytype
	// file, err := ioutil.ReadFile("version.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = json.Unmarshal(file, &data)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	c.JSON(200, gin.H{"status": "OK!", "build": "0.124"})
}
