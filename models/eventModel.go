package models

import (
// "errors"
)

type EventMeta struct {
	Name_en              string `json:"name_en"`
	Name_th              string `json:"name_th"`
	Detail_th            string `json:"detail_th"`
	Detail_en            string `json:"detail_en"`
	Term_of_use_th       string `json:"term_of_use_th"`
	Term_of_use_en       string `json:"term_of_use_en"`
	Event_condition_th   string `json:"event_condition_th"`
	Event_condition_en   string `json:"event_condition_en"`
	Event_image_banner   string `json:"event_image_banner"`
	Event_image_poster   string `json:"event_image_poster"`
	Province             string `json:"province"`
	Event_cat_id         string `json:"event_cat_id"`
	Event_start_datetime string `json:"event_start_datetime"`
	Event_end_datetime   string `json:"event_end_datetime"`
	Spacial_value        string `json:"spacial_value"`
	Spacial_value_type   string `json:"spacial_value_type "`
	Used                 string `json:"used"`
	Quantity             string `json:"quantity"`
	Status               string `json:"status"`
	Event_location_lat   string `json:"event_location_lat"`
	Event_location_lon   string `json:"event_location_lon"`
	Create_at            int64  `json:"create_at"`
	Update_at            int64  `json:"update_at"`
}

func (evMeta *EventMeta) Save(service_name string) (string, string, error) {
	table_event_meta := service_name + "_event_meta"
	// table_event_content := service_name + "_event_content"
	ConnectDb()

	tx, err := DB.Begin()
	if err != nil {
		return "ddd", "err", err
	}

	SQL_INSERT_ACTIVITY := `INSERT INTO ` + table_event_meta + ` ( ) VALUES ( ) `
	_, err1 := tx.Exec(SQL_INSERT_ACTIVITY)
	if err1 != nil {
		return "ddd", "err", err1
	}
	return "ddd", "ddd", err
}
