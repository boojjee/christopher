package models

import (
	"encoding/json"
	// "github.com/elgs/gosqljson"
	"log"
	// "strings"
)

type OfferAllContent struct {
	Offer_uid          string
	Merchant_uid       string
	Offer_point        float64
	Offer_cat_id       int64
	Offer_image_banner string
	Offer_image_poster string
	Used               int64
	Quantity           int64
	Name_en            string
	Name_th            string
	Condition_offer_en string
	Condition_offer_th string
	Description_en     string
	Description_th     string
	Lang               string
	Create_at          int64
	Update_at          int64
}

func (offer *OfferAllContent) Save(service_name string) (string, error) {
	offerMeta_table := service_name + "_offer_meta"
	offerContent_table := service_name + "_offer_content"
	ConnectDb()
	var (
		err error
	)

	tx, err := DB.Begin()
	if err != nil {
		return "err", err
	}

	SQL_INSERT_OFMETA := `INSERT INTO ` + offerMeta_table + `
	(offer_uid, offer_point, merchant_uid, offer_cat_id, offer_image_banner, offer_image_poster, used, quantity, create_at, update_at) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err1 := tx.Exec(SQL_INSERT_OFMETA, offer.Offer_uid, offer.Offer_point, offer.Merchant_uid,
		offer.Offer_cat_id, offer.Offer_image_banner, offer.Offer_image_poster,
		offer.Used, offer.Quantity, offer.Create_at, offer.Update_at)
	log.Println(err1)
	if err1 != nil {
		tx.Rollback()
		return "err", err1
	}

	log.Println(offer)

	SQL_INSERT_OFCONTENT_EN := `INSERT INTO ` + offerContent_table + `
	(name, offer_uid, condition_offer, description, lang, create_at, update_at) 
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	log.Println(offer.Name_th)
	_, err2 := tx.Exec(SQL_INSERT_OFCONTENT_EN, offer.Name_en, offer.Offer_uid, offer.Condition_offer_en,
		offer.Description_en, "en", offer.Create_at, offer.Update_at)
	log.Println(err2)
	if err2 != nil {
		tx.Rollback()
		return "err", err1
	}

	SQL_INSERT_OFCONTENT_TH := `INSERT INTO ` + offerContent_table + `
	(name, offer_uid, condition_offer, description, lang, create_at, update_at) 
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	_, err3 := tx.Exec(SQL_INSERT_OFCONTENT_TH, offer.Name_en, offer.Offer_uid, offer.Condition_offer_th,
		offer.Description_th, "th", offer.Create_at, offer.Update_at)
	log.Println(err3)
	if err3 != nil {
		tx.Rollback()
		return "err", err1
	}

	tx.Commit()
	defer CloseDb()
	return "success", nil

}

func (offer *OfferAllContent) Update(service_name string) (string, error) {
	offerMeta_table := service_name + "_offer_meta"
	offerContent_table := service_name + "_offer_content"
	ConnectDb()
	var (
		err error
	)

	tx, err := DB.Begin()
	if err != nil {
		return "err", err
	}

	UPDATE_OFFER_META := `UPDATE ` + offerMeta_table + ` SET 
	offer_point=? , offer_cat_id=?, offer_image_banner=?,
	offer_image_poster=?, used=?, quantity=?, update_at=? 
	WHERE offer_uid=?`

	_, err1 := tx.Exec(UPDATE_OFFER_META,
		offer.Offer_point, offer.Offer_cat_id, offer.Offer_image_banner,
		offer.Offer_image_poster, offer.Used, offer.Quantity, offer.Update_at, offer.Offer_uid)

	if err1 != nil {
		tx.Rollback()
		return "err", err1
	}

	UPDATE_OFFER_CONTENT_EN := `UPDATE ` + offerContent_table + ` SET 
	name=?, condition_offer=?, description=?, update_at=? WHERE offer_uid=? AND lang='en'`

	_, err2 := tx.Exec(UPDATE_OFFER_CONTENT_EN,
		offer.Name_en, offer.Condition_offer_en, offer.Description_en, offer.Update_at, offer.Offer_uid)

	if err2 != nil {
		tx.Rollback()
		return "err", err2
	}

	UPDATE_OFFER_CONTENT_TH := `UPDATE ` + offerContent_table + ` SET 
	name=?, condition_offer=?, description=?, update_at=? WHERE offer_uid=? AND lang='th'`

	_, err3 := tx.Exec(UPDATE_OFFER_CONTENT_TH,
		offer.Name_en, offer.Condition_offer_en, offer.Description_en, offer.Update_at, offer.Offer_uid)

	if err3 != nil {
		tx.Rollback()
		return "err", err3
	}

	tx.Commit()
	defer CloseDb()
	return "success", nil
}

func (offer *OfferAllContent) Delete(service_name string) (string, error) {
	offerMeta_table := service_name + "_offer_meta"
	offerContent_table := service_name + "_offer_content"
	ConnectDb()
	var (
		err error
	)

	tx, err := DB.Begin()
	if err != nil {
		return "err", err
	}

	SQL_DELETE_OFFERMETA := "DELETE from " + offerMeta_table + " WHERE offer_uid=?"
	_, err = tx.Exec(SQL_DELETE_OFFERMETA, offer.Offer_uid)
	if err != nil {
		tx.Rollback()
		return "err", err
	}

	SQL_DELETE_OFFERCONTENT := "DELETE from " + offerContent_table + " WHERE offer_uid=?"
	_, err = tx.Exec(SQL_DELETE_OFFERCONTENT, offer.Offer_uid)
	if err != nil {
		tx.Rollback()
		return "err", err
	}

	tx.Commit()
	defer CloseDb()
	return "success", nil

}

func (offer *OfferAllContent) ListsOfferByMerchant(service_name string) (string, string, error) {
	// offerMeta_table := service_name + "_offer_meta"
	// offerContent_table := service_name + "_offer_content"
	merchant_meta := service_name + "_merchant_meta"
	merchant_content := service_name + "_merchant_content"
	ConnectDb()
	// var (
	// 	err error
	// )

	// tx, err := DB.Begin()
	// if err != nil {
	// 	return "", "err", err
	// }
	SQL_SELECT_MERCHANT := "SELECT * FROM " + merchant_meta + " WHERE merchant_uid=?"
	rows, err := DB.Query(SQL_SELECT_MERCHANT, offer.Merchant_uid)
	if err != nil {
		return "", "err", err
	}
	var m_meta MerchantMeta
	var m_meta_content_en MerchantContentEN
	var m_meta_content_th MerchantContentTH
	Merchants := make([]*MerchantMeta, 0, 19)
	for rows.Next() {
		err := rows.Scan(&m_meta.Id, &m_meta.Merchant_uid, &m_meta.Username, &m_meta.Password, &m_meta.Email, &m_meta.Shop_avatar,
			&m_meta.Lat, &m_meta.Lon, &m_meta.Phone_1, &m_meta.Phone_2, &m_meta.Fax, &m_meta.Line_id, &m_meta.Facebook_link,
			&m_meta.Website_link, &m_meta.Merchant_status, &m_meta.Create_at, &m_meta.Update_at)
		if err != nil {
			return "", "err", err
		}

		SELECT_QUERY_MCONTENT_TH := "SELECT name, shop_description FROM " + merchant_content + " WHERE lang='th' AND merchant_uid = '" + offer.Merchant_uid + "'"
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

		SELECT_QUERY_MCONTENT_EN := "SELECT name, shop_description FROM " + merchant_content + " WHERE lang='en' AND merchant_uid = '" + offer.Merchant_uid + "'"
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
			m_meta.Id, m_meta.Username, m_meta.Password, m_meta.Email, m_meta.Shop_avatar, m_meta.Merchant_uid, m_meta.Lat, m_meta.Lon, m_meta.Phone_1, m_meta.Phone_2,
			m_meta.Fax, m_meta.Line_id, m_meta.Facebook_link, m_meta.Website_link, m_meta.Merchant_status, m_meta_content_en.Name_en, m_meta_content_th.Name_th,
			m_meta_content_en.Shop_description_en, m_meta_content_th.Shop_description_th, m_meta.Create_at, m_meta.Update_at,
		})

	}

	s, _ := json.Marshal(Merchants)
	log.Println(s)
	log.Println(Merchants)
	return "", "success", nil
}
