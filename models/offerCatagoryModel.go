package models

import (
	"encoding/json"
	// "github.com/elgs/gosqljson"
	// "log"
	// "strings"
)

type OfferCategory struct {
	Id                  int64  `form:"id"`
	Offer_category_name string `form:"offer_category_name"`
	Slug                string `form:"slug"`
	Value               string `form:"value"`
	Offer_type          string `form:"offer_type"`
	Create_at           int64  `form:"create_at"`
	Update_at           int64  `form:"update_at"`
}

func ListAllCategoriesOffer(service_name string) (string, string, error) {
	offerCat_table := service_name + "_offer_category"

	ConnectDb()

	SQL_SELECT_offer_category := "SELECT * FROM " + offerCat_table
	rows, err := DB.Query(SQL_SELECT_offer_category)
	if err != nil {
		return "", "err", err
	}
	var offcat OfferCategory
	offers_cat := make([]*OfferCategory, 0, 17)
	for rows.Next() {
		err := rows.Scan(&offcat.Id, &offcat.Offer_category_name, &offcat.Slug, &offcat.Value, &offcat.Offer_type, &offcat.Create_at, &offcat.Update_at)
		if err != nil {
			return "", "err", err
		}
		offers_cat = append(offers_cat, &OfferCategory{offcat.Id, offcat.Offer_category_name, offcat.Slug, offcat.Value, offcat.Offer_type, offcat.Create_at, offcat.Update_at})
	}
	// log.Println(offers_cat)
	s, _ := json.Marshal(offers_cat)
	res_offers_cat := string(s)
	return res_offers_cat, "success", err
}

func (catoffer *OfferCategory) Save(service_name string) (string, error) {
	offerCat_table := service_name + "_offer_category"

	ConnectDb()
	tx, err := DB.Begin()
	if err != nil {
		return "err", err
	}

	SQL_INSERT_offerCat_table := `INSERT INTO ` + offerCat_table + `
  (offer_category_name, slug, value, offer_type, create_at, update_at) 
  VALUES (?, ?, ?, ?, ?, ?)
  `
	_, err1 := tx.Exec(SQL_INSERT_offerCat_table, catoffer.Offer_category_name, catoffer.Slug, catoffer.Value, catoffer.Offer_type, catoffer.Create_at, catoffer.Update_at)
	if err1 != nil {
		tx.Rollback()
		return "err", err1
	}
	tx.Commit()
	defer CloseDb()
	return "success", nil
}

func (catoffer *OfferCategory) Update(service_name string) (string, error) {
	offerCat_table := service_name + "_offer_category"

	ConnectDb()
	tx, err := DB.Begin()
	if err != nil {
		return "err", err
	}

	UPDATE_OFFER_CAT := `UPDATE ` + offerCat_table + ` SET 
  offer_category_name=? , slug=?, value=?, offer_type=?,  update_at=? 
  WHERE id=?`

	_, err1 := tx.Exec(UPDATE_OFFER_CAT, catoffer.Offer_category_name, catoffer.Slug, catoffer.Value, catoffer.Offer_type, catoffer.Update_at, catoffer.Id)
	if err1 != nil {
		tx.Rollback()
		return "err", err1
	}
	tx.Commit()
	defer CloseDb()
	return "success", nil
}

func (catoffer *OfferCategory) Delete(service_name string) (string, error) {
	offerCat_table := service_name + "_offer_category"

	ConnectDb()
	tx, err := DB.Begin()
	if err != nil {
		return "err", err
	}
	SQL_DELETE_OFFERCAT := "DELETE FROM " + offerCat_table + " WHERE id=?"
	_, err = tx.Exec(SQL_DELETE_OFFERCAT, catoffer.Id)
	if err != nil {
		tx.Rollback()
		return "err", err
	}

	tx.Commit()
	defer CloseDb()
	return "success", nil
}
