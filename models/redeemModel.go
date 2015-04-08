package models

import (
// "errors"
// "log"
)

type RedeemContent struct {
	Offer_uid    string
	User_uid     string
	Redeem_point float64
	Code         string
	Expiry_date  int64
	Status       int64
	Create_at    int64
	Update_at    int64
}

func (r *RedeemContent) GetCodeRedeem(service_name string) (string, string, error) {
	// redeem_table := service_name + "_redeem"
	user_table := service_name + "_user"
	// offer_table := service_name + "_offer_meta"
	ConnectDb()
	var (
		err      error
		user_pin string
	)

	// tx, err := DB.Begin()
	// if err != nil {
	// 	return "", "err", err
	// }

	SQL_SELECTPIN := `SELECT pin From ` + user_table + ` WHERE user_uid = ?`
	rows, err := DB.Query(SQL_SELECTPIN, r.User_uid)
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

	// if user_pin == ""|user_pin == 0 {
	// 	return "", "err", errors.New("the pin code is not correct")
	// } else {

	// 	SQL_SELECT_OFFER := `SELECT * FROM `

	// }
	return "", "err", err
	// validate pin
	// get offer detail
	// check offer condition
	// gen code

}
