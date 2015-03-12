package models

import (
	"encoding/json"
	// "github.com/elgs/gosqljson"
	"log"
	"strings"
)

type Offer struct {
	Id              int64
	Name            string
	Offer_point     float64
	Condition_offer string
	Cat             int64
	Merchant_id     int64
	Offer_image     string
	Description     string
	Used            int64
	Qty             int64
	Create_at       int64
	Update_at       int64
}

func GetOfferListAll(service_name string) string {
	ConnectDb()
	table_name := service_name + "_offers"
	SELECT_QUERY := "SELECT * FROM " + table_name
	rows, err := DB.Query(SELECT_QUERY)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var o Offer
	Offers := make([]*Offer, 0, 11)
	for rows.Next() {
		err := rows.Scan(&o.Id, &o.Name, &o.Offer_point, &o.Condition_offer, &o.Cat, &o.Merchant_id, &o.Offer_image, &o.Description, &o.Used, &o.Qty, &o.Create_at, &o.Update_at)
		if err != nil {
			log.Fatal(err)
		}
		Offers = append(Offers, &Offer{o.Id, o.Name, o.Offer_point, o.Condition_offer, o.Cat, o.Merchant_id, o.Offer_image, o.Description, o.Used, o.Qty, o.Create_at, o.Update_at})
	}
	log.Println(Offers)
	s, _ := json.Marshal(Offers)
	offerListAll := strings.ToLower(string(s))
	defer CloseDb()

	return offerListAll

}

func GetOfferInfo(offer_id string, service_name string) string {
	ConnectDb()
	// offerInfo, _ := gosqljson.QueryDbToMapJson(DB, "lower", "SELECT * FROM offer WHERE id = ?", offer_id)
	table_name := service_name + "_offers"
	SELECT_QUERY := "SELECT * FROM " + table_name + " WHERE id=?"
	rows, err := DB.Query(SELECT_QUERY, offer_id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var o Offer

	for rows.Next() {
		err := rows.Scan(&o.Id, &o.Name, &o.Offer_point, &o.Condition_offer, &o.Cat, &o.Merchant_id, &o.Offer_image, &o.Description, &o.Used, &o.Qty, &o.Create_at, &o.Update_at)
		if err != nil {
			log.Fatal(err)
		}
	}
	s, _ := json.Marshal(o)
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	defer CloseDb()
	return string(s)

}

func GetOfferListByMerchantID(merchant_id string, service_name string) string {
	ConnectDb()
	table_name := service_name + "_offers"
	SQL_SELECT_OFFER := "SELECT * FROM " + table_name + " WHERE merchant_id = ?"

	rows, err := DB.Query(SQL_SELECT_OFFER, merchant_id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var o Offer
	Offers := make([]*Offer, 0, 11)
	for rows.Next() {
		err := rows.Scan(&o.Id, &o.Name, &o.Offer_point, &o.Condition_offer, &o.Cat, &o.Merchant_id, &o.Offer_image, &o.Description, &o.Used, &o.Qty, &o.Create_at, &o.Update_at)
		if err != nil {
			log.Fatal(err)
		}
		Offers = append(Offers, &Offer{o.Id, o.Name, o.Offer_point, o.Condition_offer, o.Cat, o.Merchant_id, o.Offer_image, o.Description, o.Used, o.Qty, o.Create_at, o.Update_at})
	}
	result_json, _ := json.Marshal(Offers)
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	defer CloseDb()
	return string(result_json)
}

func (o *Offer) Save(service_name string) error {
	ConnectDb()
	table_name := service_name + "_offers"
	var (
		err error
	)

	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	SQL_INSERT_POST := "insert into " + table_name + " (name, condition_offer, merchant_id, offer_image , description, used, offer_point, cat, qty, create_at, update_at) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	_, err = tx.Exec(SQL_INSERT_POST, o.Name, o.Condition_offer, o.Merchant_id, o.Offer_image, o.Description, o.Used, o.Offer_point, o.Cat, o.Qty, o.Create_at, o.Update_at)

	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	defer CloseDb()
	return nil
}

func (o *Offer) Update(service_name string, offer_id string) error {
	ConnectDb()
	table_name := service_name + "_offers"
	var (
		err error
	)

	tx, err := DB.Begin()
	if err != nil {
		return err
	}

	SQL_UPDATE_OFFER := "UPDATE " + table_name + " SET name=?, condition_offer=?,  offer_point=?, cat=?, qty=?, update_at=? WHERE id=?"
	_, err = tx.Exec(SQL_UPDATE_OFFER, o.Name, o.Condition_offer, o.Offer_point, o.Cat, o.Qty, o.Update_at, offer_id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	defer CloseDb()
	return nil
}

func (m *Offer) Delete(service_name string) error {
	ConnectDb()
	table_name := service_name + "_offers"
	var (
		err error
	)
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	SQL_DELETE := "DELETE from " + table_name + " WHERE id=?"
	_, err = tx.Exec(SQL_DELETE, m.Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	defer CloseDb()
	return nil
}
