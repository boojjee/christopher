package models

import ()

type PointSettingContent struct {
	Id          int64
	User_uid    string
	Pin         string
	Parse_id    string
	User_status int64
	Create_at   int64
	Update_at   int64
}

func GetConstantPoint(service_name string, activity_type string) float64 {
	setting_point_table := service_name + "_setting_point"

	ConnectDb()
	var (
		err            error
		constant_point float64
	)

	SQL_SELECT_SPOINT := `
    SELECT constant_point
    FROM ` + setting_point_table + ` 
    WHERE point_type = ?   
  `
	rows, err := DB.Query(SQL_SELECT_SPOINT, activity_type)
	// log.Println(err)
	if err != nil {
		return 0
	}

	for rows.Next() {
		err := rows.Scan(&constant_point)
		if err != nil {
			return 0
		}
	}

	return constant_point

}
