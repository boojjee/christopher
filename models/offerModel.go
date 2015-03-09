package models

import (
	"encoding/json"
	"github.com/elgs/gosqljson"
	"log"
	"time"
)

type Offer struct {
	Id          int64
	Name        string
	Offer_point float64
	Condition   string
	Cat         int64
	Merchant_id int64
	Descrtion   string
	Used        int64
	Qty         int64
	Create_at   time.Time
	Update_at   time.Time
}

func GetOfferListAll() string {
	ConnectDb()
	offerListAll, _ := gosqljson.QueryDbToMapJson(DB, "lower", "SELECT * FROM offer")
	defer CloseDb()
	return offerListAll
}

func GetOfferInfo(offer_id string) string {
	ConnectDb()
	// offerInfo, _ := gosqljson.QueryDbToMapJson(DB, "lower", "SELECT * FROM offer WHERE id = ?", offer_id)
	rows, err := DB.Query("SELECT * FROM offer WHERE id=?", offer_id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var o Offer

	for rows.Next() {
		err := rows.Scan(&o.Id, &o.Name, &o.Offer_point, &o.Condition, &o.Cat, &o.Merchant_id, &o.Descrtion, &o.Used, &o.Qty, &o.Create_at, &o.Update_at)
		if err != nil {
			log.Fatal(err)
		}
	}
	s, _ := json.Marshal(o)
	// log.Println(string(s))
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	defer CloseDb()
	return string(s)

}

func GetOfferListByMerchantID(merchant_id string) string {
	ConnectDb()
	offerListByMerchantID, _ := gosqljson.QueryDbToMapJson(DB, "lower", "SELECT * FROM offer WHERE merchant_id = ?", merchant_id)
	defer CloseDb()
	return offerListByMerchantID
}

func (o *Offer) Save() error {
	ConnectDb()
	var (
		err error
	)

	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	SQL_INSERT_POST := "insert into offer(name, condition, offer_point, cat, qty, create_at, update_at) values(?, ?, ?, ?, ?, ?, ?)"
	_, err = tx.Exec(SQL_INSERT_POST, o.Name, o.Condition, o.Offer_point, o.Cat, o.Qty, o.Create_at, o.Update_at)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	defer CloseDb()
	return nil
}

func (o *Offer) Update() error {
	ConnectDb()
	var (
		err error
	)

	tx, err := DB.Begin()
	if err != nil {
		return err
	}

	SQL_UPDATE_OFFER := "UPDATE offer SET name=?, condition=?,  offer_point=?, cat=?, qty=?, update_at=? WHERE id=?"
	_, err = tx.Exec(SQL_UPDATE_OFFER, o.Name, o.Condition, o.Offer_point, o.Cat, o.Qty, o.Update_at)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	defer CloseDb()
	return nil
}

func (m *Offer) Delete() error {
	ConnectDb()
	var (
		err error
	)

	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	SQL_DELETE := "DELETE from offer WHERE id=?"
	_, err = tx.Exec(SQL_DELETE, m.Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	defer CloseDb()
	return nil
}
