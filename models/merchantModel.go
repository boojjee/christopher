package models

import (
	"encoding/json"
	"errors"
	"log"
)

type MerchantMeta struct {
	Id                   int64  `json:"id"`
	Username             string `json:"username"`
	Password             string `json:"password"`
	Email                string `json:"email"`
	Shop_avatar_s        string `json:"shop_avatar_s"`
	Shop_avatar_l        string `json:"shop_avatar_l"`
	Merchant_uid         string `json:"merchant_uid"`
	Lat                  string `json:"lat"`
	Lon                  string `json:"lon"`
	Province             string `json:"province"`
	Phone_1              string `json:"phone_1"`
	Phone_2              string `json:"phone_2"`
	Fax                  string `json:"fax"`
	Line_id              string `json:"line_id"`
	Facebook_link        string `json:"facebook_link"`
	Website_link         string `json:"website_link"`
	Merchant_status      string `json:"merchant_status"`
	Merchant_category_id int64  `json:"merchant_category_id", Number`
	Merchant_province    string `json:"merchant_province"`
	Name_en              string `json:"name_en"`
	Name_th              string `json:"name_th"`
	Shop_description_en  string `json:"shop_description_en"`
	Shop_description_th  string `json:"shop_description_th"`
	Create_at            int64  `json:"create_at"`
	Update_at            int64  `json:"update_at"`
}
type MerchantMetaGalley struct {
	Id                   int64  `json:"id, Number"`
	Username             string `json:"username"`
	Password             string `json:"password"`
	Email                string `json:"email"`
	Shop_avatar_s        string `json:"shop_avatar_s"`
	Shop_avatar_l        string `json:"shop_avatar_l"`
	Merchant_uid         string `json:"merchant_uid"`
	Lat                  string `json:"lat"`
	Lon                  string `json:"lon"`
	Province             string `json:"province"`
	Phone_1              string `json:"phone_1"`
	Phone_2              string `json:"phone_2"`
	Fax                  string `json:"fax"`
	Line_id              string `json:"line_id"`
	Facebook_link        string `json:"facebook_link"`
	Website_link         string `json:"website_link"`
	Merchant_status      string `json:"merchant_status"`
	Merchant_category_id int64  `json:"merchant_category_id", Number`
	Merchant_province    string `json:"merchant_province"`
	Name_en              string `json:"name_en"`
	Name_th              string `json:"name_th"`
	Shop_description_en  string `json:"shop_description_en"`
	Shop_description_th  string `json:"shop_description_th"`
	Gallery              []*MGallery
	Create_at            int64 `json:"create_at, Number"`
	Update_at            int64 `json:"update_at, Number"`
}
type MerchantContentNoLang struct {
	Username             string `json:"username"`
	Email                string `json:"email"`
	Shop_avatar_s        string `json:"shop_avatar_s"`
	Shop_avatar_l        string `json:"shop_avatar_l"`
	Merchant_uid         string `json:"merchant_uid"`
	Lat                  string `json:"lat"`
	Lon                  string `json:"lon"`
	Province             string `json:"province"`
	Phone_1              string `json:"phone_1"`
	Phone_2              string `json:"phone_2"`
	Fax                  string `json:"fax"`
	Line_id              string `json:"line_id"`
	Facebook_link        string `json:"facebook_link"`
	Website_link         string `json:"website_link"`
	Merchant_status      string `json:"merchant_status"`
	Merchant_category_id int64  `json:"merchant_category_id", Number`
	Merchant_province    string `json:"merchant_province"`
	Name                 string `json:"name"`
	Shop_description     string `json:"shop_description"`
	Create_at            int64  `json:"create_at, Number"`
	Update_at            int64  `json:"update_at, Number"`
}

type MerchantContentNoLangGallery struct {
	Username             string `json:"username"`
	Email                string `json:"email"`
	Shop_avatar_s        string `json:"shop_avatar_s"`
	Shop_avatar_l        string `json:"shop_avatar_l"`
	Merchant_uid         string `json:"merchant_uid"`
	Lat                  string `json:"lat"`
	Lon                  string `json:"lon"`
	Province             string `json:"province"`
	Phone_1              string `json:"phone_1"`
	Phone_2              string `json:"phone_2"`
	Fax                  string `json:"fax"`
	Line_id              string `json:"line_id"`
	Facebook_link        string `json:"facebook_link"`
	Website_link         string `json:"website_link"`
	Merchant_status      string `json:"merchant_status"`
	Merchant_category_id int64  `json:"merchant_category_id", Number`
	Merchant_province    string `json:"merchant_province"`
	Name                 string `json:"name"`
	Shop_description     string `json:"shop_description"`
	Gallery              []*MGallery
	Create_at            int64 `json:"create_at"`
	Update_at            int64 `json:"update_at"`
}

type MerchantContentTH struct {
	Name_th             string
	Shop_description_th string
}
type MerchantContentEN struct {
	Name_en             string
	Shop_description_en string
}

type MerchantContent struct {
	Name             string
	Shop_description string
}

type MGallery struct {
	Photo_url string `json:"photo_url"`
	Create_at int64  `json:"create_at"`
	Update_at int64  `json:"update_at"`
}
type MerAllGallery struct {
	Id          int64  `json:"photo_url, Number"`
	Photo_url   string `json:"photo_url"`
	Merchant_id string `json:"merchant_id"`
	Create_at   int64  `json:"create_at, Number"`
	Update_at   int64  `json:"update_at, Number"`
}

func (m *MerchantMeta) Save(service_name string) (string, error) {
	merchant_meta := service_name + "_merchant_meta"
	merchant_content := service_name + "_merchant_content"
	ConnectDb()
	var (
		err error
	)

	tx, err := DB.Begin()
	if err != nil {
		return "err", err
	}
	SQL_INSERT_MMETA := `INSERT INTO ` + merchant_meta + `
											(merchant_uid, username,  password, email , shop_avatar_s, shop_avatar_l, lat, lon, province,
											 phone_1, phone_2, fax, line_id, facebook_link, website_link, merchant_status, 
											 merchant_category, merchant_province, create_at, update_at) 
											VALUES(?, ?, ?,  ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err1 := tx.Exec(SQL_INSERT_MMETA, m.Merchant_uid, m.Username, m.Password, m.Email, m.Shop_avatar_s, m.Shop_avatar_l, m.Lat, m.Lon, m.Province,
		m.Phone_1, m.Phone_2, m.Fax, m.Line_id, m.Facebook_link, m.Website_link, m.Merchant_status, m.Merchant_category_id, m.Merchant_province, m.Create_at, m.Update_at)

	if err1 != nil {
		tx.Rollback()
		return "err", err1
	}
	SQL_INSERT_MCONTENT_EN := `INSERT INTO ` + merchant_content + `
											(merchant_uid, name , shop_description, lang, create_at, update_at) 
											VALUES(?, ?, ?,  ?, ?, ?)`
	_, err2 := tx.Exec(SQL_INSERT_MCONTENT_EN, m.Merchant_uid, m.Name_en, m.Shop_description_en, "en", m.Create_at, m.Update_at)
	if err2 != nil {
		tx.Rollback()
		return "err", err2
	}
	SQL_INSERT_MCONTENT_TH := `INSERT INTO ` + merchant_content + `
											(merchant_uid, name , shop_description, lang, create_at, update_at) 
											VALUES(?, ?, ?,  ?, ?, ?)`
	_, err3 := tx.Exec(SQL_INSERT_MCONTENT_TH, m.Merchant_uid, m.Name_th, m.Shop_description_th, "th", m.Create_at, m.Update_at)
	if err3 != nil {
		tx.Rollback()
		return "err", err3
	}

	tx.Commit()
	defer CloseDb()
	return "success", nil
}

func (m *MerchantMeta) Update(service_name string) (string, error) {
	merchant_meta := service_name + "_merchant_meta"
	merchant_content := service_name + "_merchant_content"

	ConnectDb()
	var (
		err error
	)

	tx, err := DB.Begin()
	if err != nil {
		return "err", err
	}

	SQL_UPDATE_MERCAHANT_META := `UPDATE ` + merchant_meta + ` SET username=?, email=? , shop_avatar_s=?, shop_avatar_l=?,  lat=?, lon=?, province=?
		phone_1=?, phone_2=?, fax=?, line_id=?, facebook_link=?, website_link=?, merchant_status=?, merchant_category=?, merchant_province=?, update_at=? WHERE merchant_uid=?`
	_, err3 := tx.Exec(SQL_UPDATE_MERCAHANT_META,
		m.Username, m.Email, m.Shop_avatar_s, m.Shop_avatar_l, m.Lat, m.Lon, m.Province,
		m.Phone_1, m.Phone_2, m.Fax, m.Line_id, m.Facebook_link,
		m.Website_link, m.Merchant_status, m.Merchant_category_id, m.Merchant_province, m.Update_at, m.Merchant_uid)

	// log.Println(err3)
	if err3 != nil {
		tx.Rollback()
		return "err", err3
	}

	SQL_UPDATE_MERCAHANT_CONTENT_EN := `UPDATE ` + merchant_content + ` SET
	 name=? , shop_description=?, update_at=? WHERE merchant_uid=? AND lang='en' `

	_, err5 := tx.Exec(SQL_UPDATE_MERCAHANT_CONTENT_EN,
		m.Name_en, m.Shop_description_en, m.Update_at, m.Merchant_uid)
	if err5 != nil {
		tx.Rollback()
		return "err", err5
	}
	SQL_UPDATE_MERCAHANT_CONTENT_TH := `UPDATE ` + merchant_content + ` SET
	 name=? , shop_description=?, update_at=? WHERE merchant_uid=? AND lang='th' `

	_, err7 := tx.Exec(SQL_UPDATE_MERCAHANT_CONTENT_TH,
		m.Name_th, m.Shop_description_th, m.Update_at, m.Merchant_uid)
	if err7 != nil {
		tx.Rollback()
		return "err", err7
	}
	tx.Commit()
	defer CloseDb()

	return "success", nil
}

func (m *MerchantMeta) Delete(service_name string) (string, error) {
	merchant_meta := service_name + "_merchant_meta"
	merchant_content := service_name + "_merchant_content"

	ConnectDb()
	var (
		err error
	)

	tx, err := DB.Begin()
	if err != nil {
		return "err", err
	}
	SQL_DELETE_MERCHANTMETA := "DELETE from " + merchant_meta + " WHERE merchant_uid=?"
	_, err = tx.Exec(SQL_DELETE_MERCHANTMETA, m.Merchant_uid)
	if err != nil {
		tx.Rollback()
		return "err", err
	}
	SQL_DELETE_MERCHANT_CONTENT := "DELETE from " + merchant_content + " WHERE merchant_uid=?"
	_, err = tx.Exec(SQL_DELETE_MERCHANT_CONTENT, m.Merchant_uid)
	if err != nil {
		tx.Rollback()
		return "err", err
	}

	tx.Commit()
	defer CloseDb()
	return "success", nil
}

func (m *MerchantMeta) Authen(service_name string) (string, error) {
	merchant_meta := service_name + "_merchant_meta"
	ConnectDb()
	var (
		err error
	)
	SELECT_QUERY := "SELECT merchant_uid FROM " + merchant_meta + " WHERE username=? AND password=?"
	rows, err := DB.Query(SELECT_QUERY, m.Username, m.Password)
	if err != nil {
		return "fail", err
	}
	for rows.Next() {
		err := rows.Scan(&m.Merchant_uid)
		if err != nil {
			return "fail", err
		}
	}

	if err := rows.Err(); err != nil {
		return "fail", err
	}

	if m.Merchant_uid == "" {
		return "fail", err
	} else {
		return "true", nil
	}
}

func (m *MerchantMeta) MerchantShowInfoAllLang(service_name string) (string, string, error) {
	merchant_meta := service_name + "_merchant_meta"
	merchant_content := service_name + "_merchant_content"
	ConnectDb()
	var (
		err error
	)
	SELECT_QUERY := "SELECT * FROM " + merchant_meta + " WHERE merchant_uid = '" + m.Merchant_uid + "'"

	rows, err := DB.Query(SELECT_QUERY)
	if err != nil {
		return "", "err", err
	}
	var m_meta MerchantMeta
	var m_meta_content_en MerchantContentEN
	var m_meta_content_th MerchantContentTH
	// Merchants := make([]*MerchantMeta, 0, 19)
	for rows.Next() {
		err := rows.Scan(&m_meta.Id, &m_meta.Merchant_uid, &m_meta.Username, &m_meta.Password, &m_meta.Email, &m_meta.Shop_avatar_s, &m_meta.Shop_avatar_l,
			&m_meta.Lat, &m_meta.Lon, &m_meta.Province, &m_meta.Phone_1, &m_meta.Phone_2, &m_meta.Fax, &m_meta.Line_id, &m_meta.Facebook_link,
			&m_meta.Website_link, &m_meta.Merchant_status, &m_meta.Merchant_category_id, &m_meta.Merchant_province, &m_meta.Create_at, &m_meta.Update_at)
		if err != nil {
			return "", "err", err
		}
	}

	if m_meta.Id == 0 {
		return "", "err", errors.New("No Data")
	}

	SELECT_QUERY_MCONTENT_TH := "SELECT name, shop_description FROM " + merchant_content + " WHERE lang='th' AND merchant_uid = '" + m.Merchant_uid + "'"
	// log.Println(SELECT_QUERY_MCONTENT_TH)
	rows2, err := DB.Query(SELECT_QUERY_MCONTENT_TH)
	if err != nil {
		return "", "err", err
	}
	for rows2.Next() {
		err := rows2.Scan(&m_meta_content_th.Name_th, &m_meta_content_th.Shop_description_th)
		if err != nil {
			return "", "err", err
		}
	}

	SELECT_QUERY_MCONTENT_EN := "SELECT name, shop_description FROM " + merchant_content + " WHERE lang='en' AND merchant_uid = '" + m.Merchant_uid + "'"
	rows3, err := DB.Query(SELECT_QUERY_MCONTENT_EN)
	if err != nil {
		return "", "err", err
	}
	for rows3.Next() {
		err := rows3.Scan(&m_meta_content_en.Name_en, &m_meta_content_en.Shop_description_en)
		if err != nil {
			return "", "err", err
		}
	}
	mm := MerchantMeta{
		m_meta.Id, m_meta.Username, m_meta.Password, m_meta.Email, m_meta.Shop_avatar_s, m_meta.Shop_avatar_l, m_meta.Merchant_uid, m_meta.Lat, m_meta.Lon, m_meta.Province, m_meta.Phone_1, m_meta.Phone_2,
		m_meta.Fax, m_meta.Line_id, m_meta.Facebook_link, m_meta.Website_link, m_meta.Merchant_status, m_meta.Merchant_category_id, m_meta.Merchant_province, m_meta_content_en.Name_en, m_meta_content_th.Name_th,
		m_meta_content_en.Shop_description_en, m_meta_content_th.Shop_description_th, m_meta.Create_at, m_meta.Update_at,
	}

	b, err := json.Marshal(mm)

	log.Println("")
	merchentLists := string(b)
	return merchentLists, "success", nil
}

func (m *MerchantMeta) MerchantShowInfoByLang(service_name string, lang string) (string, string, error) {
	merchant_meta := service_name + "_merchant_meta"
	merchant_content := service_name + "_merchant_content"
	ConnectDb()
	var (
		err error
	)
	SELECT_QUERY := "SELECT * FROM " + merchant_meta + " WHERE merchant_uid = ?"
	rows, err := DB.Query(SELECT_QUERY, m.Merchant_uid)
	if err != nil {
		return "", "err", err
	}

	var m_meta MerchantMeta
	var mContent MerchantContent

	for rows.Next() {
		err := rows.Scan(&m_meta.Id, &m_meta.Merchant_uid, &m_meta.Username, &m_meta.Password, &m_meta.Email, &m_meta.Shop_avatar_s, &m_meta.Shop_avatar_l,
			&m_meta.Lat, &m_meta.Lon, &m_meta.Province, &m_meta.Phone_1, &m_meta.Phone_2, &m_meta.Fax, &m_meta.Line_id, &m_meta.Facebook_link,
			&m_meta.Website_link, &m_meta.Merchant_status, &m_meta.Merchant_category_id, &m_meta.Merchant_province, &m_meta.Create_at, &m_meta.Update_at)
		if err != nil {
			return "", "err", err
		}
	}
	SELECT_QUERY_MCONTENT := "SELECT name, shop_description FROM " + merchant_content + " WHERE lang=?"
	rows2, err := DB.Query(SELECT_QUERY_MCONTENT, lang)
	if err != nil {
		return "", "err", err
	}
	for rows2.Next() {
		err := rows2.Scan(&mContent.Name, &mContent.Shop_description)
		if err != nil {
			return "", "err", err
		}
	}

	result := MerchantContentNoLang{
		m_meta.Username, m_meta.Email, m_meta.Shop_avatar_s, m_meta.Shop_avatar_l, m_meta.Merchant_uid, m_meta.Lat, m_meta.Lon, m_meta.Province,
		m_meta.Phone_1, m_meta.Phone_2, m_meta.Fax, m_meta.Line_id, m_meta.Facebook_link, m_meta.Website_link,
		m_meta.Merchant_status, m_meta.Merchant_category_id, m_meta.Merchant_province, mContent.Name, mContent.Shop_description, m_meta.Create_at, m_meta.Update_at,
	}

	s, _ := json.Marshal(result)
	merchentLists := string(s)
	return merchentLists, "success", nil
}

func MerchantListAllLang(service_name string) (string, string, error) {
	merchant_meta := service_name + "_merchant_meta"
	merchant_content := service_name + "_merchant_content"
	ConnectDb()
	var (
		err error
	)
	SELECT_QUERY := "SELECT * FROM " + merchant_meta
	rows, err := DB.Query(SELECT_QUERY)
	if err != nil {
		return "", "err", err
	}
	var m_meta MerchantMeta
	var m_meta_content_en MerchantContentEN
	var m_meta_content_th MerchantContentTH
	Merchants := make([]*MerchantMeta, 0, 19)
	for rows.Next() {
		err := rows.Scan(&m_meta.Id, &m_meta.Merchant_uid, &m_meta.Username, &m_meta.Password, &m_meta.Email, &m_meta.Shop_avatar_s, &m_meta.Shop_avatar_l,
			&m_meta.Lat, &m_meta.Lon, &m_meta.Province, &m_meta.Phone_1, &m_meta.Phone_2, &m_meta.Fax, &m_meta.Line_id, &m_meta.Facebook_link,
			&m_meta.Website_link, &m_meta.Merchant_status, &m_meta.Merchant_category_id, &m_meta.Merchant_province, &m_meta.Create_at, &m_meta.Update_at)
		if err != nil {
			return "", "err", err
		}

		SELECT_QUERY_MCONTENT_TH := "SELECT name, shop_description FROM " + merchant_content + " WHERE lang='th' AND merchant_uid = '" + m_meta.Merchant_uid + "'"
		rows2, err := DB.Query(SELECT_QUERY_MCONTENT_TH)
		if err != nil {
			return "", "err", err
		}
		for rows2.Next() {
			err := rows2.Scan(&m_meta_content_th.Name_th, &m_meta_content_th.Shop_description_th)
			if err != nil {
				return "", "err", err
			}
		}

		SELECT_QUERY_MCONTENT_EN := "SELECT name, shop_description FROM " + merchant_content + " WHERE lang='en' AND merchant_uid = '" + m_meta.Merchant_uid + "'"
		rows3, err := DB.Query(SELECT_QUERY_MCONTENT_EN)
		if err != nil {
			return "", "err", err
		}
		for rows3.Next() {
			err := rows3.Scan(&m_meta_content_en.Name_en, &m_meta_content_en.Shop_description_en)
			if err != nil {
				return "", "err", err
			}
		}

		Merchants = append(Merchants, &MerchantMeta{
			m_meta.Id, m_meta.Username, m_meta.Password, m_meta.Email, m_meta.Shop_avatar_s, m_meta.Shop_avatar_l, m_meta.Merchant_uid, m_meta.Lat, m_meta.Lon, m_meta.Province, m_meta.Phone_1, m_meta.Phone_2,
			m_meta.Fax, m_meta.Line_id, m_meta.Facebook_link, m_meta.Website_link, m_meta.Merchant_status, m_meta.Merchant_category_id, m_meta.Merchant_province, m_meta_content_en.Name_en, m_meta_content_th.Name_th,
			m_meta_content_en.Shop_description_en, m_meta_content_th.Shop_description_th, m_meta.Create_at, m_meta.Update_at,
		})

	}

	s, _ := json.Marshal(Merchants)
	// log.Println(string(s))
	merchentLists := string(s)
	return merchentLists, "success", nil
}

func MerchantListByLang(service_name string, lang string) (string, string, error) {
	merchant_meta := service_name + "_merchant_meta"
	merchant_content := service_name + "_merchant_content"
	ConnectDb()
	var (
		err error
	)
	SELECT_QUERY := "SELECT * FROM " + merchant_meta
	rows, err := DB.Query(SELECT_QUERY)
	if err != nil {
		return "", "err", err
	}

	var m_meta MerchantMeta
	var mContent MerchantContent
	// Merchants := make([]*MerchantMeta, 0, 19)
	MerchantContentNoLangs := make([]*MerchantContentNoLang, 0, 16)
	for rows.Next() {
		err := rows.Scan(&m_meta.Id, &m_meta.Merchant_uid, &m_meta.Username, &m_meta.Password, &m_meta.Email, &m_meta.Shop_avatar_s, &m_meta.Shop_avatar_l,
			&m_meta.Lat, &m_meta.Lon, &m_meta.Province, &m_meta.Phone_1, &m_meta.Phone_2, &m_meta.Fax, &m_meta.Line_id, &m_meta.Facebook_link,
			&m_meta.Website_link, &m_meta.Merchant_status, &m_meta.Merchant_category_id, &m_meta.Merchant_province, &m_meta.Create_at, &m_meta.Update_at)
		if err != nil {
			return "", "err", err
		}

		SELECT_QUERY_MCONTENT := "SELECT name, shop_description FROM " + merchant_content + " WHERE lang='" + lang + "' AND merchant_uid = '" + m_meta.Merchant_uid + "'"
		rows2, err := DB.Query(SELECT_QUERY_MCONTENT)
		if err != nil {
			return "", "err", err
		}

		for rows2.Next() {
			err := rows2.Scan(&mContent.Name, &mContent.Shop_description)
			if err != nil {
				return "", "err", err
			}
		}

		MerchantContentNoLangs = append(MerchantContentNoLangs, &MerchantContentNoLang{
			m_meta.Username, m_meta.Email, m_meta.Shop_avatar_s, m_meta.Shop_avatar_l, m_meta.Merchant_uid, m_meta.Lat, m_meta.Lon, m_meta.Province,
			m_meta.Phone_1, m_meta.Phone_2, m_meta.Fax, m_meta.Line_id, m_meta.Facebook_link, m_meta.Website_link,
			m_meta.Merchant_status, m_meta.Merchant_category_id, m_meta.Merchant_province, mContent.Name, mContent.Shop_description, m_meta.Create_at, m_meta.Update_at,
		})

	}

	s, _ := json.Marshal(MerchantContentNoLangs)
	// log.Println(string(s))
	merchentLists := string(s)
	return merchentLists, "success", nil
}

func MerchantListWithGallery(service_name string) (string, string, error) {
	merchant_meta := service_name + "_merchant_meta"
	merchant_content := service_name + "_merchant_content"
	merchant_photo_gallery := service_name + "_merchants_photo_gallery"

	ConnectDb()
	var (
		err error
	)
	SELECT_QUERY := "SELECT * FROM " + merchant_meta
	rows, err := DB.Query(SELECT_QUERY)
	if err != nil {
		return "", "err", err
	}
	var m_meta MerchantMetaGalley
	var m_meta_content_en MerchantContentEN
	var m_meta_content_th MerchantContentTH
	Merchants := make([]*MerchantMetaGalley, 0, 20)

	for rows.Next() {
		err := rows.Scan(&m_meta.Id, &m_meta.Merchant_uid, &m_meta.Username, &m_meta.Password, &m_meta.Email, &m_meta.Shop_avatar_s, &m_meta.Shop_avatar_l,
			&m_meta.Lat, &m_meta.Lon, &m_meta.Province, &m_meta.Phone_1, &m_meta.Phone_2, &m_meta.Fax, &m_meta.Line_id, &m_meta.Facebook_link,
			&m_meta.Website_link, &m_meta.Merchant_status, &m_meta.Merchant_category_id, &m_meta.Merchant_province, &m_meta.Create_at, &m_meta.Update_at)
		if err != nil {
			return "", "err", err
		}

		SELECT_QUERY_MCONTENT_TH := "SELECT name, shop_description FROM " + merchant_content + " WHERE lang='th' AND merchant_uid = '" + m_meta.Merchant_uid + "'"
		rows2, err := DB.Query(SELECT_QUERY_MCONTENT_TH)

		if err != nil {
			return "", "err", err
		}
		for rows2.Next() {
			err := rows2.Scan(&m_meta_content_th.Name_th, &m_meta_content_th.Shop_description_th)
			if err != nil {
				return "", "err", err
			}
		}

		SELECT_QUERY_MCONTENT_EN := "SELECT name, shop_description FROM " + merchant_content + " WHERE lang='en' AND merchant_uid = '" + m_meta.Merchant_uid + "'"
		rows3, err := DB.Query(SELECT_QUERY_MCONTENT_EN)

		if err != nil {
			return "", "err", err
		}
		for rows3.Next() {
			err := rows3.Scan(&m_meta_content_en.Name_en, &m_meta_content_en.Shop_description_en)
			if err != nil {
				return "", "err", err
			}
		}

		SELECT_QUERY_MGALLERY := "SELECT photo_url, create_at, update_at FROM " + merchant_photo_gallery + " WHERE merchant_uid = ?"
		rows4, err := DB.Query(SELECT_QUERY_MGALLERY, m_meta.Merchant_uid)

		if err != nil {
			return "", "err", err
		}
		var g MGallery
		mGallery := make([]*MGallery, 0, 4)
		for rows4.Next() {
			err := rows4.Scan(&g.Photo_url, &g.Create_at, &g.Update_at)
			if err != nil {
				return "", "err", err
			}
			mGallery = append(mGallery, &MGallery{g.Photo_url, g.Create_at, g.Update_at})
		}

		Merchants = append(Merchants, &MerchantMetaGalley{
			m_meta.Id, m_meta.Username, m_meta.Password, m_meta.Email, m_meta.Shop_avatar_s, m_meta.Shop_avatar_l, m_meta.Merchant_uid, m_meta.Lat, m_meta.Lon, m_meta.Province,
			m_meta.Phone_1, m_meta.Phone_2, m_meta.Fax, m_meta.Line_id, m_meta.Facebook_link, m_meta.Website_link, m_meta.Merchant_status, m_meta.Merchant_category_id,
			m_meta.Merchant_province, m_meta_content_en.Name_en, m_meta_content_th.Name_th, m_meta_content_en.Shop_description_en, m_meta_content_th.Shop_description_th, mGallery, m_meta.Create_at, m_meta.Update_at,
		})

	}
	s, _ := json.Marshal(Merchants)
	merchentLists := string(s)
	return merchentLists, "success", nil
}

func MerchantListWithGalleryByLang(service_name string, lang string) (string, string, error) {
	merchant_meta := service_name + "_merchant_meta"
	merchant_content := service_name + "_merchant_content"
	merchant_photo_gallery := service_name + "_merchants_photo_gallery"

	ConnectDb()
	var (
		err error
	)
	SELECT_QUERY := "SELECT * FROM " + merchant_meta
	rows, err := DB.Query(SELECT_QUERY)
	if err != nil {
		return "", "err", err
	}
	var m_meta MerchantMeta
	var mContent MerchantContent
	Merchants := make([]*MerchantContentNoLangGallery, 0, 19)

	for rows.Next() {
		err := rows.Scan(&m_meta.Id, &m_meta.Merchant_uid, &m_meta.Username, &m_meta.Password, &m_meta.Email, &m_meta.Shop_avatar_s, &m_meta.Shop_avatar_l,
			&m_meta.Lat, &m_meta.Lon, &m_meta.Province, &m_meta.Phone_1, &m_meta.Phone_2, &m_meta.Fax, &m_meta.Line_id, &m_meta.Facebook_link,
			&m_meta.Website_link, &m_meta.Merchant_status, &m_meta.Merchant_category_id, &m_meta.Merchant_province, &m_meta.Create_at, &m_meta.Update_at)
		if err != nil {
			return "", "err", err
		}

		SELECT_QUERY_MCONTENT := "SELECT name, shop_description FROM " + merchant_content + " WHERE lang='" + lang + "' AND merchant_uid = '" + m_meta.Merchant_uid + "'"
		rows2, err := DB.Query(SELECT_QUERY_MCONTENT)

		if err != nil {
			return "", "err", err
		}
		for rows2.Next() {
			err := rows2.Scan(&mContent.Name, &mContent.Shop_description)
			if err != nil {
				return "", "err", err
			}
		}

		SELECT_QUERY_MGALLERY := "SELECT photo_url, create_at, update_at FROM " + merchant_photo_gallery + " WHERE merchant_uid = ?"
		rows4, err := DB.Query(SELECT_QUERY_MGALLERY, m_meta.Merchant_uid)

		if err != nil {
			return "", "err", err
		}
		var g MGallery
		mGallery := make([]*MGallery, 0, 4)
		for rows4.Next() {
			err := rows4.Scan(&g.Photo_url, &g.Create_at, &g.Update_at)
			if err != nil {
				return "", "err", err
			}
			mGallery = append(mGallery, &MGallery{g.Photo_url, g.Create_at, g.Update_at})
		}

		Merchants = append(Merchants, &MerchantContentNoLangGallery{
			m_meta.Username, m_meta.Email, m_meta.Shop_avatar_s, m_meta.Shop_avatar_l, m_meta.Merchant_uid, m_meta.Lat, m_meta.Lon, m_meta.Province, m_meta.Phone_1, m_meta.Phone_2,
			m_meta.Fax, m_meta.Line_id, m_meta.Facebook_link, m_meta.Website_link, m_meta.Merchant_status, m_meta.Merchant_category_id, m_meta.Merchant_province, mContent.Name,
			mContent.Shop_description, mGallery, m_meta.Create_at, m_meta.Update_at,
		})

	}
	s, _ := json.Marshal(Merchants)
	merchentLists := string(s)
	return merchentLists, "success", nil
}
