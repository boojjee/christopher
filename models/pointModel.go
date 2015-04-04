package models

import (
	// "encoding/json"
	"log"
)

type MyBPoint struct {
	Id           int64
	User_uid     string
	Blance_point float64
	Create_at    int64
	Update_at    int64
}

func (mp *MyBPoint) GetMyCurrentPoint(service_name string) (float64, string, error) {
	blance_point_table := service_name + "_point_balance"
	ConnectDb()

	var blance_point float64
	SQL_SELECT_BPOINT := "SELECT blance_point FROM " + blance_point_table + " WHERE user_uid=?"
	rows, err := DB.Query(SQL_SELECT_BPOINT, mp.User_uid)
	if err != nil {
		return 0, "err", err
	}
	for rows.Next() {
		err := rows.Scan(&blance_point)
		if err != nil {
			log.Fatal(err)
		}
	}
	return blance_point, "success", err
}
