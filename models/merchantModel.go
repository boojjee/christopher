package models

import (
	// "errors"
	"encoding/json"
	"github.com/elgs/gosqljson"
	"log"
	"time"
)

type Merchant struct {
	Id               int64
	Username         string
	Name             string
	Password         string
	Email            string
	Shop_image       string
	Shop_avatar      string
	Shop_description string
	Lat              string
	Lon              string
	Create_at        time.Time
	Update_at        time.Time
}

func GetMerchentLists() string {
	ConnectDb()
	merchentLists, _ := gosqljson.QueryDbToMapJson(DB, "lower", "SELECT * FROM merchants")
	defer CloseDb()
	return merchentLists

}

func MerchantShowInfo(id string) string {
	ConnectDb()

	rows, err := DB.Query("SELECT * FROM merchants WHERE id=?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var m Merchant

	for rows.Next() {
		err := rows.Scan(&m.Id, &m.Username, &m.Name, &m.Password, &m.Email, &m.Shop_image, &m.Shop_avatar, &m.Shop_description, &m.Lat, &m.Lon, &m.Create_at, &m.Update_at)
		if err != nil {
			log.Fatal(err)
		}
	}
	s, _ := json.Marshal(m)
	// log.Println(string(s))
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return string(s)
}

func (m *Merchant) Save() error {
	ConnectDb()
	var (
		err error
	)

	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	SQL_INSERT_POST := "insert into merchants(username, name, password, email, shop_image, shop_avatar, shop_description, lat, lon, create_at, update_at) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := tx.Exec(SQL_INSERT_POST, m.Username, m.Name, m.Password, m.Email, m.Shop_image, m.Shop_avatar, m.Shop_description, m.Lat, m.Lon, m.Create_at, m.Update_at)
	if err != nil {
		tx.Rollback()
		return err
	}
	log.Println(result)
	tx.Commit()
	defer CloseDb()
	return nil
}

func (m *Merchant) Update() error {
	ConnectDb()
	var (
		err error
	)

	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	SQL_UPDATE_MERCAHANT := "UPDATE merchants SET username=?, name=?,  email=?, shop_image=?, shop_avatar=?, shop_description=?, lat=?, lon=?, update_at=? WHERE id=?"
	result, err := tx.Exec(SQL_UPDATE_MERCAHANT, m.Username, m.Name, m.Email, m.Shop_image, m.Shop_avatar, m.Shop_description, m.Lat, m.Lon, m.Update_at, m.Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	log.Println(result)
	defer CloseDb()
	return nil
}

func (m *Merchant) Delete() error {
	ConnectDb()
	var (
		err error
	)

	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	SQL_DELETE_POST := "DELETE from merchants WHERE id=?"
	_, err = tx.Exec(SQL_DELETE_POST, m.Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	defer CloseDb()
	return nil
}
