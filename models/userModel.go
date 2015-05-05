package models

import (
	// "encoding/json"
	"errors"
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

func (u *UserContent) CheckHadPin(service_name string) (string, string, error) {
	user_table := service_name + "_user"
	ConnectDb()
	var (
		err      error
		user_pin string
	)
	SQL_SELECTPIN := `SELECT pin FROM ` + user_table + ` WHERE user_uid = ?`
	rows, err := DB.Query(SQL_SELECTPIN, u.User_uid)
	if err != nil {
		return "", "err", err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user_pin)
		if err != nil {
			return "", "err", err
		}
	}
	log.Println(user_pin)
	if user_pin == "" {
		return "no", "err", errors.New("no data ")
	} else {
		return user_pin, "success", err
	}

}

func (u *UserContent) Save(service_name string) (string, string, error) {
	user_table := service_name + "_user"
	ConnectDb()
	var (
		err      error
		user_uid string
	)

	tx, err := DB.Begin()
	if err != nil {
		return "", "err", err
	}
	defer tx.Rollback()

	SQL_SELECT_USRID := `SELECT user_uid FROM ` + user_table + ` where parse_id = ? LIMIT 1`
	rows, err := DB.Query(SQL_SELECT_USRID, u.Parse_id)

	if err != nil {
		return "", "err", err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user_uid)
		if err != nil {
			return "", "err", err
		}
	}
	if user_uid == "" {
		SQL_INSERT_MMETA := `INSERT INTO ` + user_table + ` 
  (user_uid, pin, parse_id, user_status, create_at, update_at) VALUES (?, ?, ?, ?, ?, ?)
  `
		_, err1 := tx.Exec(SQL_INSERT_MMETA, u.User_uid, u.Pin, u.Parse_id, u.User_status, u.Create_at, u.Update_at)
		if err1 != nil {
			tx.Rollback()
			return "", "err", err1
		}
		// log.Println(u.User_uid)
		tx.Commit()
		defer CloseDb()
		return u.User_uid, "success to created", nil
	} else {
		// log.Println("had")
		return user_uid, "success", nil
	}

}

func (u *UserContent) GetUIDByParseID(service_name string) (string, string, error) {
	user_table := service_name + "_user"
	ConnectDb()
	var (
		err      error
		user_uid string
	)
	SQL_SELECT_USRID := `SELECT user_uid FROM ` + user_table + ` where parse_id = ? LIMIT 1`
	rows, err := DB.Query(SQL_SELECT_USRID, u.Parse_id)
	if err != nil {
		return "", "err", err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user_uid)
		if err != nil {
			return "", "err", err
		}
	}
	return user_uid, "success", nil
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
