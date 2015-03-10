package models

import (
	// "errors"
	"encoding/json"
	"github.com/elgs/gosqljson"
	"log"
	"strings"
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

func GetMerchentLists(service_name string) (string, string) {
	ConnectDb()
	table_name := service_name + "_merchants"
	SELECT_QUERY := "SELECT * FROM " + table_name
	merchentLists, err := gosqljson.QueryDbToMapJson(DB, "lower", SELECT_QUERY)

	if err != nil {
		return "No DB", "err"
	} else {
		defer CloseDb()
		return merchentLists, "ok"

	}

}

// func MerchantShowInfo(id string, service_name string) (string, string) {
// 	ConnectDb()
// 	table_name := service_name + "_merchants"
// 	SELECT_QUERY := "SELECT * FROM " + table_name + " WHERE id=?"
// 	rows, err := DB.Query(SELECT_QUERY, id)
// 	if err != nil {
// 		return "No DB", "err"
// 	}
// 	defer rows.Close()

// 	var m Merchant

// 	for rows.Next() {
// 		err := rows.Scan(&m.Id, &m.Username, &m.Name, &m.Password, &m.Email, &m.Shop_image, &m.Shop_avatar, &m.Shop_description, &m.Lat, &m.Lon, &m.Create_at, &m.Update_at)
// 		if err != nil {
// 			return "row error", "err"
// 		}
// 	}
// 	s, _ := json.Marshal(m)
// 	err = rows.Err()
// 	if err != nil {
// 		return "row error", "err"
// 	}

// 	return strings.ToLower(string(s)), "ok"
// }

func MerchantShowInfoByName(m_name string, service_name string) (string, string) {
	ConnectDb()

	table_name := service_name + "_merchants"
	SELECT_QUERY := "SELECT * FROM " + table_name + " WHERE name=?"

	rows, err := DB.Query(SELECT_QUERY, m_name)
	if err != nil {
		return "No DB", "err"
	}
	defer rows.Close()
	// rowCnt, err := rows.RowsAffected()
	var m Merchant

	for rows.Next() {
		err := rows.Scan(&m.Id, &m.Username, &m.Name, &m.Password, &m.Email, &m.Shop_image, &m.Shop_avatar, &m.Shop_description, &m.Lat, &m.Lon, &m.Create_at, &m.Update_at)
		if err != nil {
			return "row error", "err"
		}
	}

	if m.Id == 0 {
		return "row error", "err"
	} else {
		if err := rows.Err(); err != nil {
			log.Println(err)
			return "row error", "err"
		}
		s, _ := json.Marshal(m)
		err = rows.Err()
		if err != nil {
			return "row error", "err"
		}

		return strings.ToLower(string(s)), "ok"
	}

}

func (m *Merchant) Save(service_name string) error {
	table_name := service_name + "_merchants"
	ConnectDb()
	var (
		err error
	)

	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	SQL_INSERT_POST := "insert into " + table_name + "(username, name, password, email, shop_image, shop_avatar, shop_description, lat, lon, create_at, update_at) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

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

func (m *Merchant) Update(service_name string) error {
	table_name := service_name + "_merchants"
	ConnectDb()
	var (
		err error
	)

	tx, err := DB.Begin()
	if err != nil {
		return err
	}

	SQL_UPDATE_MERCAHANT := "UPDATE " + table_name + " SET username=?, name=?,  email=?, shop_image=?, shop_avatar=?, shop_description=?, lat=?, lon=?, update_at=? WHERE id=?"
	_, err = tx.Exec(SQL_UPDATE_MERCAHANT, m.Username, m.Name, m.Email, m.Shop_image, m.Shop_avatar, m.Shop_description, m.Lat, m.Lon, m.Update_at, m.Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	defer CloseDb()
	return nil
}

func (m *Merchant) Delete(service_name string) error {
	table_name := service_name + "_merchants"
	ConnectDb()
	var (
		err error
	)

	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	SQL_DELETE_POST := "DELETE from " + table_name + " WHERE id=?"
	_, err = tx.Exec(SQL_DELETE_POST, m.Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	defer CloseDb()
	return nil
}

func (m *Merchant) Authen(service_name string) (string, error) {
	table_name := service_name + "_merchants"
	ConnectDb()
	var (
		err error
	)
	SELECT_QUERY := "SELECT id FROM " + table_name + " WHERE username=? AND password=?"
	rows, err := DB.Query(SELECT_QUERY, m.Username, m.Password)
	if err != nil {
		return "fail", err
	}
	for rows.Next() {
		err := rows.Scan(&m.Id)
		if err != nil {
			return "fail", err
		}
	}
	if m.Id == 0 {
		return "fail", err
	} else {
		return "true", nil
	}

}
