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

func CreateShop() {
	// ConnectDb()
	// result, err := DB.Exec(
	// 	"INSERT INTO shops (title, shop_name) VALUES (?, ?)",
	// 	"gopher",
	// 	27,
	// )
	// if err != nil {
	// 	log.Println(err)
	// }

}
