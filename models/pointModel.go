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

func GetWorkOut(service_name string) (float64, string, error) {
	// SELECT  SUM(distance) as total FROM ginkgo_point
	activity_table := service_name + "_activity"
	var totalPoint float64

	ConnectDb()
	SQL_SELECT_BPOINT := "SELECT SUM(distance) FROM " + activity_table
	rows, err := DB.Query(SQL_SELECT_BPOINT)
	if err != nil {
		return 0, "err", err
	}
	for rows.Next() {
		err := rows.Scan(&totalPoint)
		if err != nil {
			return 0, "err", err
		}
	}
	return totalPoint, "success", err

}

func GetPoints(service_name string) (float64, string, error) {
	// SELECT  SUM(distance) as total FROM ginkgo_point
	point_table := service_name + "_point"
	var totalPoint float64

	ConnectDb()
	SQL_SELECT_BPOINT := "SELECT SUM(g_point) FROM " + point_table
	rows, err := DB.Query(SQL_SELECT_BPOINT)
	if err != nil {
		return 0, "err", err
	}
	for rows.Next() {
		err := rows.Scan(&totalPoint)
		if err != nil {
			return 0, "err", err
		}
	}
	return totalPoint, "success", err

}
