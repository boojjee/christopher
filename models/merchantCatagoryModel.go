package models

import (
	"encoding/json"
	// "github.com/elgs/gosqljson"
	"log"
	// "strings"
)

type MerchantCategory struct {
	Id                     int64  `form:"id"`
	Merchant_category_name string `form:"merchant_category_name"`
	Slug                   string `form:"slug"`
	Create_at              int64  `form:"create_at"`
	Update_at              int64  `form:"update_at"`
}

func ListAllCategoriesMerchant(service_name string) (string, string, error) {
	merchantCat_table := service_name + "_merchant_category"

	ConnectDb()

	SQL_SELECT_merchant_category := "SELECT * FROM " + merchantCat_table
	rows, err := DB.Query(SQL_SELECT_merchant_category)
	if err != nil {
		return "", "err", err
	}
	var merchantCat MerchantCategory
	merchants_cat := make([]*MerchantCategory, 0, 17)
	for rows.Next() {
		err := rows.Scan(&merchantCat.Id, &merchantCat.Merchant_category_name, &merchantCat.Slug, &merchantCat.Create_at, &merchantCat.Update_at)
		if err != nil {
			return "", "err", err
		}
		merchants_cat = append(merchants_cat, &MerchantCategory{merchantCat.Id, merchantCat.Merchant_category_name, merchantCat.Slug, merchantCat.Create_at, merchantCat.Update_at})
	}
	log.Println(merchants_cat)
	s, _ := json.Marshal(merchants_cat)
	res_merchants_cat := string(s)
	return res_merchants_cat, "success", err
}

func (catmerchant *MerchantCategory) Save(service_name string) (string, error) {
	merchantCat_table := service_name + "_merchant_category"

	ConnectDb()
	tx, err := DB.Begin()
	if err != nil {
		return "err", err
	}

	SQL_INSERT_merchantCat_table := `INSERT INTO ` + merchantCat_table + `
  (merchant_category_name, slug, create_at, update_at) 
  VALUES (?, ?, ?, ?)
  `
	_, err1 := tx.Exec(SQL_INSERT_merchantCat_table, catmerchant.Merchant_category_name, catmerchant.Slug, catmerchant.Create_at, catmerchant.Update_at)
	if err1 != nil {
		tx.Rollback()
		return "err", err1
	}
	tx.Commit()
	defer CloseDb()
	return "success", nil
}

func (catmerchant *MerchantCategory) Update(service_name string) (string, error) {
	merchantCat_table := service_name + "_merchant_category"

	ConnectDb()
	tx, err := DB.Begin()
	if err != nil {
		return "err", err
	}

	UPDATE_OFFER_CAT := `UPDATE ` + merchantCat_table + ` SET 
  merchant_category_name=? , slug=?, update_at=? 
  WHERE id=?`

	_, err1 := tx.Exec(UPDATE_OFFER_CAT, catmerchant.Merchant_category_name, catmerchant.Slug, catmerchant.Update_at, catmerchant.Id)
	if err1 != nil {
		tx.Rollback()
		return "err", err1
	}
	tx.Commit()
	defer CloseDb()
	return "success", nil
}

func (catmerchant *MerchantCategory) Delete(service_name string) (string, error) {
	merchantCat_table := service_name + "_merchant_category"

	ConnectDb()
	tx, err := DB.Begin()
	if err != nil {
		return "err", err
	}
	SQL_DELETE_OFFERCAT := "DELETE FROM " + merchantCat_table + " WHERE id=?"
	_, err = tx.Exec(SQL_DELETE_OFFERCAT, catmerchant.Id)
	if err != nil {
		tx.Rollback()
		return "err", err
	}

	tx.Commit()
	defer CloseDb()
	return "success", nil
}
