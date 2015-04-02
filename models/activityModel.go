package models

import (
	"encoding/json"
	"errors"
	"log"
)

type ActivityContentForm struct {
	Id                 int64   `json:"id"`
	Activity_uid       string  `json:"activity_uid"`
	User_uid           string  `json:"user_uid"`
	User_parse_id      string  `json:"user_parse_id"`
	Third_activity_id  string  `json:"third_activity_id"`
	Third_uri          string  `json:"third_uri"`
	Third_token_user   string  `json:"third_token_user"`
	Source             string  `json:"source"`
	Distance           float64 `json:"distance"`
	Duration           float64 `json:"duration"`
	Calories           float64 `json:"calories"`
	Start_activity_lat float64 `json:"start_activity_lat"`
	Start_activity_lon float64 `json:"start_activity_lon"`
	Activity_type      string  `json:"activity_type"`
	Activity_status    int64   `json:"activity_status"`
	MyLocation_lat     float64 `json:"mylocation_lat"`
	MyLocation_lon     float64 `json:"mylocation_lon"`
	Province           string  `json:"province"`
	Create_at          int64   `json:"create_at"`
	Update_at          int64   `json:"update_at"`
}

type ActivityContent struct {
	Id                 int64   `json:"id"`
	Activity_uid       string  `json:"activity_uid"`
	User_uid           string  `json:"user_uid"`
	User_parse_id      string  `json:"user_parse_id"`
	Third_activity_id  string  `json:"third_activity_id"`
	Third_uri          string  `json:"third_uri"`
	Third_token_user   string  `json:"third_token_user"`
	Source             string  `json:"source"`
	Distance           float64 `json:"distance"`
	Duration           float64 `json:"duration"`
	Calories           float64 `json:"calories"`
	Start_activity_lat float64 `json:"start_activity_lat"`
	Start_activity_lon float64 `json:"start_activity_lon"`
	Activity_type      string  `json:"activity_type"`
	Activity_status    int64   `json:"activity_status"`
	Create_at          int64   `json:"create_at"`
	Update_at          int64   `json:"update_at"`
}

type ActivityContentWithPoint struct {
	Id                 int64   `json:"id"`
	Activity_uid       string  `json:"activity_uid"`
	User_uid           string  `json:"user_uid"`
	User_parse_id      string  `json:"user_parse_id"`
	Third_activity_id  string  `json:"third_activity_id"`
	Third_uri          string  `json:"third_uri"`
	Third_token_user   string  `json:"third_token_user"`
	Source             string  `json:"source"`
	Distance           float64 `json:"distance"`
	Duration           float64 `json:"duration"`
	Calories           float64 `json:"calories"`
	G_Point            []*gpointForActivity
	Start_activity_lat float64 `json:"start_activity_lat"`
	Start_activity_lon float64 `json:"start_activity_lon"`
	Activity_type      string  `json:"activity_type"`
	Activity_status    int64   `json:"activity_status"`
	Create_at          int64   `json:"create_at"`
	Update_at          int64   `json:"update_at"`
}
type ActivityInfoWithPoint struct {
	Id                 int64   `json:"id"`
	Activity_uid       string  `json:"activity_uid"`
	User_uid           string  `json:"user_uid"`
	User_parse_id      string  `json:"user_parse_id"`
	Third_activity_id  string  `json:"third_activity_id"`
	Third_uri          string  `json:"third_uri"`
	Third_token_user   string  `json:"third_token_user"`
	Source             string  `json:"source"`
	Distance           float64 `json:"distance"`
	Duration           float64 `json:"duration"`
	Calories           float64 `json:"calories"`
	G_Point            gpointForActivity
	Start_activity_lat float64 `json:"start_activity_lat"`
	Start_activity_lon float64 `json:"start_activity_lon"`
	Activity_type      string  `json:"activity_type"`
	Activity_status    int64   `json:"activity_status"`
	Create_at          int64   `json:"create_at"`
	Update_at          int64   `json:"update_at"`
}

type gpointContent struct {
	ID                int64   `json:"id"`
	Activity_uid      string  `json:"activity_uid"`
	Third_activity_id string  `json:"third_activity_id"`
	User_id           string  `json:"user_id"`
	Distance          float64 `json:"distance"`
	G_point           float64 `json:"g_point"`
	Create_at         int64   `json:"create_at"`
	Update_at         int64   `json:"update_at"`
}
type gpointForActivity struct {
	G_point   float64 `json:"g_point"`
	Create_at int64   `json:"create_at"`
	Update_at int64   `json:"update_at"`
}

type ActivityIdPagination struct {
	Next_max_id int64 `json:"next_max_id"`
	Next_min_id int64 `json:"next_min_id"`
}

func (act *ActivityContentForm) Save(service_name string) (string, error) {
	activity_table := service_name + "_activity"
	location_log_table := service_name + "_log_location"
	ConnectDb()
	var (
		err error
	)

	tx, err := DB.Begin()
	if err != nil {
		return "err", err
	}

	SQL_INSERT_ACTIVITY := `INSERT INTO ` + activity_table + `
  (activity_uid, user_uid, user_parse_id, third_activity_id, third_uri, third_token_user,
   source, distance, duration, calories, start_activity_lat, start_activity_lon, activity_type,
   activity_status, create_at, update_at) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)
  `
	_, err1 := tx.Exec(SQL_INSERT_ACTIVITY, act.Activity_uid, act.User_uid, act.User_parse_id, act.Third_activity_id, act.Third_uri,
		act.Third_token_user, act.Source, act.Distance, act.Duration, act.Calories, act.Start_activity_lat, act.Start_activity_lon,
		act.Activity_type, act.Activity_status, act.Create_at, act.Update_at)
	if err1 != nil {
		tx.Rollback()
		return "err", err1
	}

	// log location
	SQL_INSERT_LOGLOCATION := `INSERT INTO ` + location_log_table + `
	(user_uid, mylocation_lat, mylocation_lon, province,tag_log, create_at, update_at) VALUES 
	(?, ?, ?, ?, ?, ?)
	`
	_, err2 := tx.Exec(SQL_INSERT_LOGLOCATION, act.User_uid, act.MyLocation_lat, act.MyLocation_lon, act.Province, act.Create_at, act.Update_at)
	if err2 != nil {
		tx.Rollback()
		return "err", err2
	}

	tx.Commit()
	defer CloseDb()
	return "success", nil
}

func (act *ActivityContent) Update(service_name string) (string, error) {
	activity_table := service_name + "_activity"
	ConnectDb()
	var (
		err error
	)

	tx, err := DB.Begin()
	if err != nil {
		return "err", err
	}
	SQL_UPDATE_ACTIVITY := `UPDATE ` + activity_table + ` SET user_uid=?, user_parse_id=?, 
   third_activity_id=?, third_uri=?, third_token_user=?, source=?, distance=?, duration=?, calories=?, 
   start_activity_lat=?, start_activity_lon=?, activity_type=?, activity_status=?, update_at=?
   WHERE activity_uid =?`
	_, err1 := tx.Exec(SQL_UPDATE_ACTIVITY, act.User_uid, act.User_parse_id, act.Third_activity_id,
		act.Third_uri, act.Third_token_user, act.Source, act.Distance, act.Duration, act.Calories,
		act.Start_activity_lat, act.Start_activity_lon, act.Activity_type, act.Activity_status,
		act.Update_at, act.Activity_uid)
	if err1 != nil {
		return "err", err1
	}

	tx.Commit()
	defer CloseDb()

	return "success", nil

}

func (act *ActivityContent) Delete(service_name string) (string, error) {
	activity_table := service_name + "_activity"
	ConnectDb()
	var (
		err error
	)
	tx, err := DB.Begin()
	if err != nil {
		return "err", err
	}
	SQL_DELETE_ACT := `DELETE from ` + activity_table + ` WHERE activity_uid = ? `
	log.Println(act.Activity_uid)
	log.Println(SQL_DELETE_ACT)
	_, err = tx.Exec(SQL_DELETE_ACT, act.Activity_uid)
	if err != nil {
		tx.Rollback()
		return "err", err
	}

	tx.Commit()
	defer CloseDb()
	return "success", nil
}

func GetLatestActivityList(service_name string, user_uid string) (string, string, string, error) {
	activity_table := service_name + "_activity"
	ConnectDb()
	var (
		err error
	)

	SQL_SELECT_LATESTAC := `
    SELECT *
    FROM ` + activity_table + ` 
    WHERE user_uid =?
    ORDER BY id desc
    LIMIT 20
  `
	rows, err := DB.Query(SQL_SELECT_LATESTAC, user_uid)
	if err != nil {
		return "", "", "err", err
	}

	var (
		actContent    ActivityContent
		actPagination ActivityIdPagination
	)

	ActivitiesContent := make([]*ActivityContent, 0, 20)
	// ActIdPagination := make([]*ActivityIdPagination, 0, 2)

	for rows.Next() {
		err := rows.Scan(&actContent.Id, &actContent.Activity_uid, &actContent.User_uid, &actContent.User_parse_id,
			&actContent.Third_activity_id, &actContent.Third_uri, &actContent.Third_token_user, &actContent.Source,
			&actContent.Distance, &actContent.Duration, &actContent.Calories, &actContent.Start_activity_lat, &actContent.Start_activity_lon,
			&actContent.Activity_type, &actContent.Activity_status, &actContent.Create_at, &actContent.Update_at)
		if err != nil {
			return "", "", "err", err
		}

		ActivitiesContent = append(ActivitiesContent, &ActivityContent{actContent.Id, actContent.Activity_uid, actContent.User_uid, actContent.User_parse_id,
			actContent.Third_activity_id, actContent.Third_uri, actContent.Third_token_user, actContent.Source,
			actContent.Distance, actContent.Duration, actContent.Calories, actContent.Start_activity_lat, actContent.Start_activity_lon,
			actContent.Activity_type, actContent.Activity_status, actContent.Create_at, actContent.Update_at})
	}

	actPagination.Next_max_id = ActivitiesContent[0].Id
	actPagination.Next_min_id = ActivitiesContent[len(ActivitiesContent)-1].Id
	// ActIdPagination = append(ActIdPagination, &ActivityIdPagination{actPagination.Next_max_id, actPagination.Next_min_id})
	p, _ := json.Marshal(actPagination)
	s, _ := json.Marshal(ActivitiesContent)
	activityPaginate := string(p)
	activitylist := string(s)
	return activitylist, activityPaginate, "success", nil
}

func GetNextActivityList(service_name string, user_uid string, max_id string) (string, string, string, error) {
	activity_table := service_name + "_activity"
	ConnectDb()
	var (
		err error
	)
	// helpers.Convert_string_to_int(max_id)
	SQL_SELECT_LATESTAC := `
    SELECT *
    FROM ` + activity_table + ` 
    WHERE user_uid = ? AND  id > ? 
    ORDER BY id desc
    LIMIT 20
  `
	rows, err := DB.Query(SQL_SELECT_LATESTAC, user_uid, max_id)
	log.Println(err)
	if err != nil {
		return "", "", "err", err
	}

	var (
		actContent    ActivityContent
		actPagination ActivityIdPagination
	)

	ActivitiesContent := make([]*ActivityContent, 0, 17)
	// ActIdPagination := make([]*ActivityIdPagination, 0, 2)

	for rows.Next() {
		err := rows.Scan(&actContent.Id, &actContent.Activity_uid, &actContent.User_uid, &actContent.User_parse_id,
			&actContent.Third_activity_id, &actContent.Third_uri, &actContent.Third_token_user, &actContent.Source,
			&actContent.Distance, &actContent.Duration, &actContent.Calories, &actContent.Start_activity_lat, &actContent.Start_activity_lon,
			&actContent.Activity_type, &actContent.Activity_status, &actContent.Create_at, &actContent.Update_at)
		if err != nil {
			return "", "", "err", err
		}

		ActivitiesContent = append(ActivitiesContent, &ActivityContent{actContent.Id, actContent.Activity_uid, actContent.User_uid, actContent.User_parse_id,
			actContent.Third_activity_id, actContent.Third_uri, actContent.Third_token_user, actContent.Source,
			actContent.Distance, actContent.Duration, actContent.Calories, actContent.Start_activity_lat, actContent.Start_activity_lon,
			actContent.Activity_type, actContent.Activity_status, actContent.Create_at, actContent.Update_at})
	}

	actPagination.Next_max_id = ActivitiesContent[0].Id
	actPagination.Next_min_id = ActivitiesContent[len(ActivitiesContent)-1].Id

	p, _ := json.Marshal(actPagination)
	s, _ := json.Marshal(ActivitiesContent)
	activityPaginate := string(p)
	activitylist := string(s)
	return activitylist, activityPaginate, "success", nil
}

// Get all my activity
func GetActivityListsAll(service_name string, user_uid string) (string, string, error) {
	activity_table := service_name + "_activity"
	point_table := service_name + "_point"
	ConnectDb()
	var (
		err error
	)
	// helpers.Convert_string_to_int(max_id)
	SQL_SELECT_ACLIST := `
    SELECT *
    FROM ` + activity_table + ` 
    WHERE user_uid = ?  
    ORDER BY id desc 
  `
	rows, err := DB.Query(SQL_SELECT_ACLIST, user_uid)
	// log.Println(err)
	if err != nil {
		return "", "err", err
	}
	var actContent ActivityContentWithPoint
	var gpointContents gpointContent

	ActivitiesContents := make([]*ActivityContentWithPoint, 0, 18)

	for rows.Next() {
		err := rows.Scan(&actContent.Id, &actContent.Activity_uid, &actContent.User_uid, &actContent.User_parse_id,
			&actContent.Third_activity_id, &actContent.Third_uri, &actContent.Third_token_user, &actContent.Source,
			&actContent.Distance, &actContent.Duration, &actContent.Calories, &actContent.Start_activity_lat, &actContent.Start_activity_lon,
			&actContent.Activity_type, &actContent.Activity_status, &actContent.Create_at, &actContent.Update_at)
		if err != nil {
			return "", "err", err
		}
		gPointSlice := make([]*gpointForActivity, 0, 3)
		SQL_SELECT_POINT := `
	    SELECT g_point, create_at, update_at
	    FROM ` + point_table + ` 
	    WHERE activity_uid = ?
	  `
		// log.Println(actContent.Activity_uid)
		rows, err := DB.Query(SQL_SELECT_POINT, actContent.Activity_uid)
		if err != nil {
			return "", "err", err
		}
		for rows.Next() {
			err := rows.Scan(&gpointContents.G_point, &gpointContents.Create_at, &gpointContents.Update_at)
			if err != nil {
				return "", "err", err
			}
			gPointSlice = append(gPointSlice, &gpointForActivity{gpointContents.G_point, gpointContents.Create_at, gpointContents.Update_at})

		}

		ActivitiesContents = append(ActivitiesContents, &ActivityContentWithPoint{actContent.Id, actContent.Activity_uid, actContent.User_uid, actContent.User_parse_id,
			actContent.Third_activity_id, actContent.Third_uri, actContent.Third_token_user, actContent.Source,
			actContent.Distance, actContent.Duration, actContent.Calories, gPointSlice, actContent.Start_activity_lat, actContent.Start_activity_lon,
			actContent.Activity_type, actContent.Activity_status, actContent.Create_at, actContent.Update_at})

	}
	s, _ := json.Marshal(ActivitiesContents)
	activitylist := string(s)
	return activitylist, "success", nil
}

// Get all my activity
func GetPrevActivityList(service_name string, user_uid string, min_id string) (string, string, string, error) {
	activity_table := service_name + "_activity"
	ConnectDb()
	var (
		err error
	)
	// helpers.Convert_string_to_int(max_id)
	SQL_SELECT_LATESTAC := `
    SELECT *
    FROM ` + activity_table + ` 
    WHERE user_uid = ? AND  id < ? 
    ORDER BY id desc
    LIMIT 20
  `
	rows, err := DB.Query(SQL_SELECT_LATESTAC, user_uid, min_id)
	log.Println(err)
	if err != nil {
		return "", "", "err", err
	}

	var (
		actContent    ActivityContent
		actPagination ActivityIdPagination
	)

	ActivitiesContent := make([]*ActivityContent, 0, 17)
	// ActIdPagination := make([]*ActivityIdPagination, 0, 2)

	for rows.Next() {
		err := rows.Scan(&actContent.Id, &actContent.Activity_uid, &actContent.User_uid, &actContent.User_parse_id,
			&actContent.Third_activity_id, &actContent.Third_uri, &actContent.Third_token_user, &actContent.Source,
			&actContent.Distance, &actContent.Duration, &actContent.Calories, &actContent.Start_activity_lat, &actContent.Start_activity_lon,
			&actContent.Activity_type, &actContent.Activity_status, &actContent.Create_at, &actContent.Update_at)
		if err != nil {
			return "", "", "err", err
		}

		ActivitiesContent = append(ActivitiesContent, &ActivityContent{actContent.Id, actContent.Activity_uid, actContent.User_uid, actContent.User_parse_id,
			actContent.Third_activity_id, actContent.Third_uri, actContent.Third_token_user, actContent.Source,
			actContent.Distance, actContent.Duration, actContent.Calories, actContent.Start_activity_lat, actContent.Start_activity_lon,
			actContent.Activity_type, actContent.Activity_status, actContent.Create_at, actContent.Update_at})
	}

	actPagination.Next_max_id = ActivitiesContent[0].Id
	actPagination.Next_min_id = ActivitiesContent[len(ActivitiesContent)-1].Id

	p, _ := json.Marshal(actPagination)
	s, _ := json.Marshal(ActivitiesContent)
	activityPaginate := string(p)
	activitylist := string(s)
	return activitylist, activityPaginate, "success", nil
}

func GetActivityByAcUID(service_name string, activity_uid string) (string, string, error) {
	activity_table := service_name + "_activity"
	point_table := service_name + "_point"
	ConnectDb()
	var (
		err error
	)
	// helpers.Convert_string_to_int(max_id)
	SQL_SELECT_ACLIST := `
    SELECT *
    FROM ` + activity_table + ` 
    WHERE activity_uid = ?  
    LIMIT 1
  `
	rows, err := DB.Query(SQL_SELECT_ACLIST, activity_uid)
	// log.Println(err)
	if err != nil {
		return "", "err", err
	}
	var actContent ActivityContentWithPoint
	var gpointContents gpointContent

	// ActivitiesContents := make([]*ActivityContentWithPoint, 0, 18)
	// gPointSlice := make([]*gpointForActivity, 0, 3)
	for rows.Next() {
		err := rows.Scan(&actContent.Id, &actContent.Activity_uid, &actContent.User_uid, &actContent.User_parse_id,
			&actContent.Third_activity_id, &actContent.Third_uri, &actContent.Third_token_user, &actContent.Source,
			&actContent.Distance, &actContent.Duration, &actContent.Calories, &actContent.Start_activity_lat, &actContent.Start_activity_lon,
			&actContent.Activity_type, &actContent.Activity_status, &actContent.Create_at, &actContent.Update_at)
		if err != nil {
			return "", "err", err
		}

		SQL_SELECT_POINT := `
	    SELECT g_point, create_at, update_at
	    FROM ` + point_table + ` 
	    WHERE activity_uid = ?
	  `
		// log.Println(actContent.Activity_uid)
		rows, err := DB.Query(SQL_SELECT_POINT, actContent.Activity_uid)
		if err != nil {
			return "", "err", err
		}
		for rows.Next() {
			err := rows.Scan(&gpointContents.G_point, &gpointContents.Create_at, &gpointContents.Update_at)
			if err != nil {
				return "", "err", err
			}
		}

	}

	if actContent.Id == 0 {
		return "", "err", errors.New(`{ Message :"No Data" }`)
	}

	myAc := ActivityInfoWithPoint{actContent.Id, actContent.Activity_uid, actContent.User_uid, actContent.User_parse_id,
		actContent.Third_activity_id, actContent.Third_uri, actContent.Third_token_user, actContent.Source,
		actContent.Distance, actContent.Duration, actContent.Calories, gpointForActivity{gpointContents.G_point, gpointContents.Create_at, gpointContents.Update_at}, actContent.Start_activity_lat, actContent.Start_activity_lon,
		actContent.Activity_type, actContent.Activity_status, actContent.Create_at, actContent.Update_at}

	s, _ := json.Marshal(myAc)
	activitylist := string(s)
	return activitylist, "success", nil
}
