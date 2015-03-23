package models

import (
	"encoding/json"
	// "github.com/elgs/gosqljson"
	"log"
	// "strings"
)

type OfferCatagory struct {
	Id                  int64  `form:"id"`
	Offer_catagory_name string `form:"offer_catagory_name"`
	Slug                string `form:"slug"`
	Create_at           int64  `form:"create_at"`
	Update_at           int64  `form:"update_at"`
}

func ListAllCatagoriesOffer(service_name string) (string, string, error) {
	offerCat_table := service_name + "_offer_catagory"

	ConnectDb()

	SQL_SELECT_offer_catagory := "SELECT * FROM " + offerCat_table
	rows, err := DB.Query(SQL_SELECT_offer_catagory)
	if err != nil {
		return "", "err", err
	}
	var offcat OfferCatagory
	offers_cat := make([]*OfferCatagory, 0, 17)
	for rows.Next() {
		err := rows.Scan(&offcat.Id, &offcat.Offer_catagory_name, &offcat.Slug, &offcat.Create_at, &offcat.Update_at)
		if err != nil {
			return "", "err", err
		}
		offers_cat = append(offers_cat, &OfferCatagory{offcat.Id, offcat.Offer_catagory_name, offcat.Slug, offcat.Create_at, offcat.Update_at})
	}
	log.Println(offers_cat)
	s, _ := json.Marshal(offers_cat)
	res_offers_cat := string(s)
	return res_offers_cat, "success", err
}

func (catoffer *OfferCatagory) Save(service_name string) (string, error) {
	offerCat_table := service_name + "_offer_catagory"

	ConnectDb()
	tx, err := DB.Begin()
	if err != nil {
		return "err", err
	}

	SQL_INSERT_offerCat_table := `INSERT INTO ` + offerCat_table + `
  (offer_catagory_name, slug, create_at, update_at) 
  VALUES (?, ?, ?, ?)
  `
	_, err1 := tx.Exec(SQL_INSERT_offerCat_table, catoffer.Offer_catagory_name, catoffer.Slug, catoffer.Create_at, catoffer.Update_at)
	if err1 != nil {
		tx.Rollback()
		return "err", err1
	}
	tx.Commit()
	defer CloseDb()
	return "success", nil
}

func (catoffer *OfferCatagory) Update(service_name string) (string, error) {
	offerCat_table := service_name + "_offer_catagory"

	ConnectDb()
	tx, err := DB.Begin()
	if err != nil {
		return "err", err
	}

	UPDATE_OFFER_CAT := `UPDATE ` + offerCat_table + ` SET 
  offer_catagory_name=? , slug=?, update_at=? 
  WHERE id=?`

	_, err1 := tx.Exec(UPDATE_OFFER_CAT, catoffer.Offer_catagory_name, catoffer.Slug, catoffer.Update_at, catoffer.Id)
	if err1 != nil {
		tx.Rollback()
		return "err", err1
	}
	tx.Commit()
	defer CloseDb()
	return "success", nil
}

func (catoffer *OfferCatagory) Delete(service_name string) (string, error) {
	offerCat_table := service_name + "_offer_catagory"

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
