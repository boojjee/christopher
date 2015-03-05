package models

import (
	"github.com/elgs/gosqljson"
	// "log"
)

func GetShopLists() string {
	ConnectDb()
	shopList, _ := gosqljson.QueryDbToMapJson(DB, "lower", "SELECT * FROM shops")
	// log.Println(a)
	defer CloseDb()
	return shopList

}
