package models

import (
	"christopher/helpers"
	"encoding/json"
	"errors"
	"log"
)

type RedeemContent struct {
	Offer_uid      string
	User_uid       string
	Redeem_uid     string
	Redeem_point   float64
	Code           string
	Expiry_date    int64
	Status         int64
	Pin            string
	MyLocation_lat string
	MyLocation_lon string
	Province       string
	Create_at      int64
	Update_at      int64
}

type RedeemJSON struct {
	User_uid     string  `json:"user_uid"`
	RedeemCode   string  `json:"redeem_code"`
	Point_used   float64 `json:"point_used"`
	Blance_point float64 `json:"blance_point"`
	Expire_date  int64   `json:"expire_date"`
	Redeem_date  int64   `json:"redeem_date"`
}

func (r *RedeemContent) GetCodeRedeem(service_name string) (string, string, error) {
	redeem_table := service_name + "_redeem"
	user_table := service_name + "_user"
	offer_table := service_name + "_offer_meta"
	point_balance_table := service_name + "_point_balance"
	location_log_table := service_name + "_log_location"

	ConnectDb()
	var (
		err error
		// user_pin       string
		user_pin_count       int64
		offer_point          float64
		quantity             int64
		used                 int64
		totalQuantityBalance int64
		myPoint              float64
		current_blance       float64
	)
	// check PIN
	SQL_SELECTPIN := `SELECT COUNT(pin) AS pin FROM  ` + user_table + ` WHERE user_uid = ? AND pin = ?`
	rows, err := DB.Query(SQL_SELECTPIN, r.User_uid, r.Pin)
	if err != nil {
		return "", "err", err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user_pin_count)
		if err != nil {
			return "", "err", err
		}
	}

	if user_pin_count == 0 {
		return "", "err", errors.New("the pin code is not correct")
	} else {
		log.Println("Debug: pin is ok")
		// GET OFFER Data
		SQL_SELECT_OFFER := `SELECT offer_point, quantity, used FROM ` + offer_table + ` 
			WHERE offer_uid =? AND status =? `
		rows2, err2 := DB.Query(SQL_SELECT_OFFER, r.Offer_uid, 1)
		if err2 != nil {
			return "", "err", err2
		}
		defer rows2.Close()
		for rows2.Next() {
			err := rows2.Scan(&offer_point, &quantity, &used)
			if err != nil {
				return "", "err", err
			}
		}

		log.Println(offer_point)
		log.Println(quantity)
		log.Println(used)

		totalQuantityBalance = quantity - used
		log.Println(totalQuantityBalance)
		// check offer condition
		if totalQuantityBalance > 0 {
			log.Println("Debug: totalQuantityBalance > 0  ")
			// GET MY Current point
			SQL_SELECT_MYPOINT := `SELECT blance_point FROM ` + point_balance_table + ` 
			WHERE user_uid =?`
			//myPoint
			rows3, err3 := DB.Query(SQL_SELECT_MYPOINT, r.User_uid)
			if err3 != nil {
				return "", "err", err3
			}
			defer rows3.Close()
			for rows3.Next() {
				err := rows3.Scan(&myPoint)
				if err != nil {
					return "", "err", err
				}
			}
			log.Println(myPoint)
			expr_date := helpers.UnixTimeAddMinFromNow(30)
			log.Println("Redeem EXPire : ")
			log.Println(expr_date)
			// insert to redeem
			tx, err := DB.Begin()
			if err != nil {
				return "", "err", err
			}

			SQL_INSERT_REDEEM := `INSERT INTO ` + redeem_table + ` 
			(redeem_uid, offer_uid, user_uid, redeem_point, code, expiry_date, status, create_at, update_at) 
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
			`
			_, err_insertRD := tx.Exec(SQL_INSERT_REDEEM, r.Redeem_uid, r.Offer_uid, r.User_uid, offer_point, r.Code, expr_date, 1, r.Create_at, r.Update_at)
			log.Println(err_insertRD)
			if err_insertRD != nil {
				tx.Rollback()
				return "", "err", err_insertRD
			}

			// update offer usage
			SQL_UPDATE_OFFER := `UPDATE ` + offer_table + ` SET used = used + 1 WHERE offer_uid = ?`
			_, err_update_offer := tx.Exec(SQL_UPDATE_OFFER, r.Offer_uid)
			log.Println(err_update_offer)
			if err_update_offer != nil {
				tx.Rollback()
				return "", "err", err_update_offer
			}

			// update point balance
			SQL_UPDATE_PBL := `UPDATE ` + point_balance_table + ` SET blance_point = blance_point - ? WHERE user_uid = ?`
			_, err_update_pbl := tx.Exec(SQL_UPDATE_PBL, offer_point, r.User_uid)
			if err_update_pbl != nil {
				tx.Rollback()
				return "", "err", err_update_pbl
			}

			// log location send go routine
			SQL_INSERT_LOGLOCATION := `INSERT INTO ` + location_log_table + `
			(user_uid, mylocation_lat, mylocation_lon, province, tag_log, create_at, update_at) VALUES 
			(?, ?, ?, ?, ?, ?, ?)
			`
			_, err_loglocation := tx.Exec(SQL_INSERT_LOGLOCATION, r.User_uid, r.MyLocation_lat, r.MyLocation_lon, r.Province, "Redeem", r.Create_at, r.Update_at)
			if err_loglocation != nil {
				tx.Rollback()
				return "", "err", err_loglocation
			}

			SQL_SELECT_MYPOINT2 := `SELECT blance_point FROM ` + point_balance_table + ` WHERE user_uid =?`
			rows_cr, err_cr := DB.Query(SQL_SELECT_MYPOINT2, r.User_uid)
			if err_cr != nil {

				return "", "err", err_cr
			}
			defer rows_cr.Close()
			for rows_cr.Next() {
				err := rows_cr.Scan(&current_blance)
				if err != nil {
					return "", "err", err
				}
			}

			my_redeem := RedeemJSON{r.User_uid, r.Code, offer_point, current_blance, expr_date, r.Create_at}

			s, _ := json.Marshal(my_redeem)
			result := string(s)
			tx.Commit()
			log.Println(result)
			defer CloseDb()
			return result, "Success", nil
		} else {
			return "", "err", errors.New("Offer is finish")
		}

	}

	// gen code

}
