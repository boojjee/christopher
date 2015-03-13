package models

import (
	// "strings"
	"log"
)

type MerchantGallery struct {
	Id          int64
	Photo_url   string
	Merchant_id int64
	Create_at   int64
	Update_at   int64
}

func (m *MerchantGallery) Save(service_name string) (string, error) {
	table_name := service_name + "_merchants_photo_gallery"
	ConnectDb()
	var (
		err error
	)

	tx, err := DB.Begin()
	if err != nil {
		return "err", err
	}
	log.Println(m)
	SQL_INSERT_POST := "insert into " + table_name + "(photo_url, merchant_id, create_at, update_at) values(?, ?, ?, ?)"
	result, err := tx.Exec(SQL_INSERT_POST, m.Photo_url, m.Merchant_id, m.Create_at, m.Update_at)
	log.Println(result)
	if err != nil {
		tx.Rollback()
		return "err", err
	}
	tx.Commit()
	defer CloseDb()
	return "success", nil
}

func (m *MerchantGallery) Update(service_name string, photo_id string) (string, error) {
	table_name := service_name + "_merchants_photo_gallery"
	ConnectDb()
	var (
		err error
	)

	tx, err := DB.Begin()
	if err != nil {
		return "err", err
	}
	log.Println(m)
	SQL_UDPATE_POST := "UPDATE  " + table_name + " SET photo_url=?, merchant_id=?, update_at=? WHERE id=?"
	result, err := tx.Exec(SQL_UDPATE_POST, m.Photo_url, m.Merchant_id, m.Update_at, photo_id)
	log.Println(result)
	if err != nil {
		tx.Rollback()
		return "err", err
	}
	tx.Commit()
	defer CloseDb()
	return "success", nil
}

func (m *MerchantGallery) Delete(service_name string) error {
	ConnectDb()
	table_name := service_name + "_merchants_photo_gallery"
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
