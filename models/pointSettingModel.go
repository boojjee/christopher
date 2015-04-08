package models

import (
	"encoding/json"
	// "log"
)

type PointSetting struct {
	ID              int64   `json:"id"`
	Point_type      string  `json:"point_type"`
	PointSettingUID string  `json:"pointsettinguid"`
	Constant_point  float64 `json:"constant_point"`
	Description     string  `json:"description"`
	Create_at       int64   `json:"create_at"`
	Update_at       int64   `json:"update_at"`
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

func (stp *PointSetting) CreatePointSetting(service_name string) (string, error) {
	point_setting_table := service_name + "_setting_point"
	ConnectDb()
	var (
		err error
	)

	tx, err := DB.Begin()
	if err != nil {
		return "err", err
	}

	SQL_INSERT_STP := `INSERT INTO ` + point_setting_table + `
		(point_setting_uid, point_type, constant_point, description, create_at, update_at) VALUES (?,?,?,?,?,?)
	`
	_, err_insert := tx.Exec(SQL_INSERT_STP, stp.PointSettingUID, stp.Point_type, stp.Constant_point, stp.Description, stp.Create_at, stp.Update_at)
	if err_insert != nil {
		return "err", err_insert
	}
	tx.Commit()
	defer CloseDb()
	return "success", nil
}

func (stp *PointSetting) UpdatePointSetting(service_name string) (string, error) {
	point_setting_table := service_name + "_setting_point"
	ConnectDb()
	var (
		err error
	)

	tx, err := DB.Begin()
	if err != nil {
		return "err", err
	}

	UPDATE_POINT_SETTING := `UPDATE ` + point_setting_table + ` SET 
	point_type=? , constant_point=?, description=?, update_at=? 
	WHERE point_setting_uid=?`

	_, err1 := tx.Exec(UPDATE_POINT_SETTING, stp.Point_type, stp.Constant_point, stp.Description, stp.Update_at, stp.PointSettingUID)
	if err1 != nil {
		tx.Rollback()
		return "err", err1
	}
	tx.Commit()
	defer CloseDb()
	return "success", nil
}

func (stp *PointSetting) DeletePointSetting(service_name string) (string, error) {
	point_setting_table := service_name + "_setting_point"
	ConnectDb()
	var (
		err error
	)

	tx, err := DB.Begin()
	if err != nil {
		return "err", err
	}

	SQL_DELETE_STP := "DELETE from " + point_setting_table + " WHERE point_setting_uid=?"
	_, err = tx.Exec(SQL_DELETE_STP, stp.PointSettingUID)
	if err != nil {
		tx.Rollback()
		return "err", err
	}
	tx.Commit()
	defer CloseDb()
	return "success", nil
}

func ListsAllSettingPoint(service_name string) (string, string, error) {
	point_setting_table := service_name + "_setting_point"
	ConnectDb()
	var (
		err error
		pst PointSetting
	)
	SQL_SELECT := `SELECT * FROM ` + point_setting_table
	rows, err := DB.Query(SQL_SELECT)
	if err != nil {
		return "", "err", err
	}

	PointSettings := make([]*PointSetting, 0, 7)

	for rows.Next() {
		err := rows.Scan(&pst.ID, &pst.PointSettingUID, &pst.Point_type, &pst.Constant_point, &pst.Description, &pst.Create_at, &pst.Update_at)
		if err != nil {
			return "", "err", err
		}
		PointSettings = append(PointSettings, &PointSetting{pst.ID, pst.PointSettingUID, pst.Point_type, pst.Constant_point, pst.Description, pst.Create_at, pst.Update_at})
	}
	s, _ := json.Marshal(PointSettings)
	merchentLists := string(s)
	return merchentLists, "success", nil
}
