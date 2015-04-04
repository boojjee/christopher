package models

import (
	// "encoding/json"
	"log"
)

type UserContent struct {
	Id          int64
	User_uid    string
	Pin         string
	Parse_id    string
	User_status int64
	Create_at   int64
	Update_at   int64
}

func (u *UserContent) Save(service_name string) (string, error) {
	user_table := service_name + "_user"
	ConnectDb()
	var (
		err error
	)

	tx, err := DB.Begin()
	if err != nil {
		return "err", err
	}
	SQL_INSERT_MMETA := `INSERT INTO ` + user_table + ` 
  (user_uid, pin, parse_id, user_status, create_at, update_at) VALUES (?, ?, ?, ?, ?, ?)
  `
	_, err1 := tx.Exec(SQL_INSERT_MMETA, u.User_uid, u.Pin, u.Parse_id, u.User_status, u.Create_at, u.Update_at)
	if err1 != nil {
		tx.Rollback()
		return "err", err1
	}
	tx.Commit()
	defer CloseDb()
	return "success", nil
}

func (u *UserContent) UpdatePin(service_name string, user_uid string) (string, error) {
	user_table := service_name + "_user"
	ConnectDb()
	var (
		err error
	)

	SQL_UPDATE_USER := `UPDATE ` + user_table + ` SET pin=?, update_at=? WHERE user_uid = ?`
	tx, err := DB.Begin()
	if err != nil {
		return "err", err
	}
	log.Println(SQL_UPDATE_USER)
	_, err1 := tx.Exec(SQL_UPDATE_USER, u.Pin, u.Update_at, user_uid)

	log.Println(err1)
	if err != nil {
		return "err", err1
	}

	tx.Commit()
	defer CloseDb()

	return "success", nil
}

func (u *UserContent) Update(service_name string, user_uid string) (string, error) {
	user_table := service_name + "_user"
	ConnectDb()
	var (
		err error
	)

	SQL_UPDATE_USER := `UPDATE ` + user_table + ` SET pin=? , parse_id=?, user_status=?, update_at=?
  WHERE user_uid = ?`
	tx, err := DB.Begin()
	if err != nil {
		return "err", err
	}

	_, err1 := tx.Exec(SQL_UPDATE_USER, u.Pin, u.Parse_id, u.User_status, u.Update_at, user_uid)

	if err != nil {
		return "err", err1
	}

	tx.Commit()
	defer CloseDb()

	return "success", nil
}

func (u *UserContent) Delete(service_name string) (string, error) {
	user_table := service_name + "_user"
	ConnectDb()
	var (
		err error
	)
	tx, err := DB.Begin()
	if err != nil {
		return "err", err
	}
	SQL_DELETE_USER := `DELETE from ` + user_table + ` WHERE user_uid = ? `

	_, err = tx.Exec(SQL_DELETE_USER, u.Id)
	if err != nil {
		tx.Rollback()
		return "err", err
	}

	tx.Commit()
	defer CloseDb()
	return "success", nil
}
