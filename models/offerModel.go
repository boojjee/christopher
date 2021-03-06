package models

import (
	"encoding/json"
	"log"
)

type OfferAllContent struct {
	Id                 int64   `json:"id"`
	Offer_uid          string  `json:"offer_uid"`
	Merchant_uid       string  `json:"merchant_uid"`
	Offer_point        float64 `json:"offer_point"`
	Offer_cat_id       int64   `json:"offer_cat_id"`
	Offer_image_banner string  `json:"offer_image_banner"`
	Offer_image_poster string  `json:"offer_image_poster"`
	Used               int64   `json:"used"`
	Quantity           int64   `json:"quantity"`
	Name_en            string  `json:"name_en"`
	Name_th            string  `json:"name_th"`
	Condition_offer_en string  `json:"condition_offer_en"`
	Condition_offer_th string  `json:"condition_offer_th"`
	Description_en     string  `json:"description_en"`
	Description_th     string  `json:"description_th"`
	Status             int64   `json:"status"`
	Value              string  `json:"value"`
	Offer_type         string  `json:"offer_type"`
	Create_at          int64   `json:"create_at"`
	Update_at          int64   `json:"update_at"`
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
	(offer_uid, offer_point, merchant_uid, offer_cat_id, offer_image_banner, offer_image_poster, 
	 used, quantity, status ,value, offer_type, create_at, update_at) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err1 := tx.Exec(SQL_INSERT_OFMETA, offer.Offer_uid, offer.Offer_point, offer.Merchant_uid,
		offer.Offer_cat_id, offer.Offer_image_banner, offer.Offer_image_poster,
		offer.Used, offer.Quantity, offer.Status, offer.Value, offer.Offer_type, offer.Create_at, offer.Update_at)
	log.Println(err1)
	if err1 != nil {
		tx.Rollback()
		return "err", err1
	}

	SQL_INSERT_OFCONTENT_EN := `INSERT INTO ` + offerContent_table + `
	(name, offer_uid, condition_offer, description, lang, create_at, update_at) 
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	_, err2 := tx.Exec(SQL_INSERT_OFCONTENT_EN, offer.Name_en, offer.Offer_uid, offer.Condition_offer_en,
		offer.Description_en, "en", offer.Create_at, offer.Update_at)
	if err2 != nil {
		tx.Rollback()
		return "err", err1
	}

	SQL_INSERT_OFCONTENT_TH := `INSERT INTO ` + offerContent_table + `
	(name, offer_uid, condition_offer, description, lang, create_at, update_at) 
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	_, err3 := tx.Exec(SQL_INSERT_OFCONTENT_TH, offer.Name_th, offer.Offer_uid, offer.Condition_offer_th,
		offer.Description_th, "th", offer.Create_at, offer.Update_at)
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
	offer_image_poster=?, used=?, quantity=?, status=?, value=?, offer_type=?, update_at=? 
	WHERE offer_uid=?`

	_, err1 := tx.Exec(UPDATE_OFFER_META,
		offer.Offer_point, offer.Offer_cat_id, offer.Offer_image_banner,
		offer.Offer_image_poster, offer.Used, offer.Quantity, offer.Status, offer.Value, offer.Offer_type, offer.Update_at, offer.Offer_uid)

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
		offer.Name_th, offer.Condition_offer_th, offer.Description_th, offer.Update_at, offer.Offer_uid)

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
	offerMeta_table := service_name + "_offer_meta"
	offerContent_table := service_name + "_offer_content"
	ConnectDb()
	var oac OfferAllContent
	SQL_DELETE_OFFERMETA := "SELECT * FROM " + offerMeta_table + " WHERE merchant_uid=?"
	rows, err := DB.Query(SQL_DELETE_OFFERMETA, offer.Merchant_uid)
	if err != nil {
		return "", "err", err
	}

	offers := make([]*OfferAllContent, 0, 17)
	for rows.Next() {
		err := rows.Scan(&oac.Id, &oac.Offer_uid, &oac.Offer_point, &oac.Merchant_uid,
			&oac.Offer_cat_id, &oac.Offer_image_banner, &oac.Offer_image_poster,
			&oac.Used, &oac.Quantity, &oac.Status, &oac.Value, &oac.Offer_type, &oac.Create_at, &oac.Update_at)
		if err != nil {
			return "", "err", err
		}

		SQL_SELECT_OFFER_TH := "SELECT name, condition_offer, description FROM " + offerContent_table + " WHERE offer_uid=? AND lang=?"
		rows1, err := DB.Query(SQL_SELECT_OFFER_TH, oac.Offer_uid, "th")
		if err != nil {
			return "", "err", err
		}
		for rows1.Next() {
			err := rows1.Scan(&oac.Name_th, &oac.Condition_offer_th, &oac.Description_th)
			if err != nil {
				return "", "err", err
			}
		}

		SQL_SELECT_OFFER_EN := "SELECT name, condition_offer, description FROM " + offerContent_table + " WHERE offer_uid=? AND lang=?"
		rows2, err := DB.Query(SQL_SELECT_OFFER_EN, oac.Offer_uid, "en")
		if err != nil {
			return "", "err", err
		}
		for rows2.Next() {
			err := rows2.Scan(&oac.Name_en, &oac.Condition_offer_en, &oac.Description_en)
			if err != nil {
				return "", "err", err
			}
		}

		offers = append(offers, &OfferAllContent{
			oac.Id, oac.Offer_uid, oac.Merchant_uid, oac.Offer_point, oac.Offer_cat_id, oac.Offer_image_banner, oac.Offer_image_poster,
			oac.Used, oac.Quantity, oac.Name_en, oac.Name_th, oac.Condition_offer_en, oac.Condition_offer_th,
			oac.Description_en, oac.Description_th, oac.Status, oac.Value, oac.Offer_type, oac.Create_at, oac.Update_at,
		})
	}
	// log.Println(offers)
	s, _ := json.Marshal(offers)
	offers_of_merchant := string(s)
	return offers_of_merchant, "success", err
}

func (offer *OfferAllContent) ListsOffersMerchantStatus(service_name string) (string, string, error) {
	offerMeta_table := service_name + "_offer_meta"
	offerContent_table := service_name + "_offer_content"
	ConnectDb()
	var oac OfferAllContent
	SQL_DELETE_OFFERMETA := "SELECT * FROM " + offerMeta_table + " WHERE merchant_uid=? AND status = 1"
	rows, err := DB.Query(SQL_DELETE_OFFERMETA, offer.Merchant_uid)
	if err != nil {
		return "", "err", err
	}

	offers := make([]*OfferAllContent, 0, 17)
	for rows.Next() {
		err := rows.Scan(&oac.Id, &oac.Offer_uid, &oac.Offer_point, &oac.Merchant_uid,
			&oac.Offer_cat_id, &oac.Offer_image_banner, &oac.Offer_image_poster,
			&oac.Used, &oac.Quantity, &oac.Status, &oac.Value, &oac.Offer_type, &oac.Create_at, &oac.Update_at)
		if err != nil {
			return "", "err", err
		}

		SQL_SELECT_OFFER_TH := "SELECT name, condition_offer, description FROM " + offerContent_table + " WHERE offer_uid=? AND lang=?"
		rows1, err := DB.Query(SQL_SELECT_OFFER_TH, oac.Offer_uid, "th")
		if err != nil {
			return "", "err", err
		}
		for rows1.Next() {
			err := rows1.Scan(&oac.Name_th, &oac.Condition_offer_th, &oac.Description_th)
			if err != nil {
				return "", "err", err
			}
		}

		SQL_SELECT_OFFER_EN := "SELECT name, condition_offer, description FROM " + offerContent_table + " WHERE offer_uid=? AND lang=?"
		rows2, err := DB.Query(SQL_SELECT_OFFER_EN, oac.Offer_uid, "en")
		if err != nil {
			return "", "err", err
		}
		for rows2.Next() {
			err := rows2.Scan(&oac.Name_en, &oac.Condition_offer_en, &oac.Description_en)
			if err != nil {
				return "", "err", err
			}
		}

		offers = append(offers, &OfferAllContent{
			oac.Id, oac.Offer_uid, oac.Merchant_uid, oac.Offer_point, oac.Offer_cat_id, oac.Offer_image_banner, oac.Offer_image_poster,
			oac.Used, oac.Quantity, oac.Name_en, oac.Name_th, oac.Condition_offer_en, oac.Condition_offer_th,
			oac.Description_en, oac.Description_th, oac.Status, oac.Value, oac.Offer_type, oac.Create_at, oac.Update_at,
		})
	}
	// log.Println(offers)
	s, _ := json.Marshal(offers)
	offers_of_merchant := string(s)
	return offers_of_merchant, "success", err
}

func ListsAllOffer(service_name string) (string, string, error) {
	offerMeta_table := service_name + "_offer_meta"
	offerContent_table := service_name + "_offer_content"
	ConnectDb()
	var oac OfferAllContent
	SQL_SELECT_OFFERMETA := "SELECT * FROM " + offerMeta_table
	rows, err := DB.Query(SQL_SELECT_OFFERMETA)
	if err != nil {
		return "", "err", err
	}

	offers := make([]*OfferAllContent, 0, 17)
	for rows.Next() {
		err := rows.Scan(&oac.Id, &oac.Offer_uid, &oac.Offer_point, &oac.Merchant_uid,
			&oac.Offer_cat_id, &oac.Offer_image_banner, &oac.Offer_image_poster,
			&oac.Used, &oac.Quantity, &oac.Status, &oac.Value, &oac.Offer_type, &oac.Create_at, &oac.Update_at)
		if err != nil {
			return "", "err", err
		}

		SQL_SELECT_OFFER_TH := "SELECT name, condition_offer, description FROM " + offerContent_table + " WHERE offer_uid=? AND lang=?"
		rows1, err := DB.Query(SQL_SELECT_OFFER_TH, oac.Offer_uid, "th")
		if err != nil {
			return "", "err", err
		}
		for rows1.Next() {
			err := rows1.Scan(&oac.Name_th, &oac.Condition_offer_th, &oac.Description_th)
			if err != nil {
				return "", "err", err
			}
		}

		SQL_SELECT_OFFER_EN := "SELECT name, condition_offer, description FROM " + offerContent_table + " WHERE offer_uid=? AND lang=?"
		rows2, err := DB.Query(SQL_SELECT_OFFER_EN, oac.Offer_uid, "en")
		if err != nil {
			return "", "err", err
		}
		for rows2.Next() {
			err := rows2.Scan(&oac.Name_en, &oac.Condition_offer_en, &oac.Description_en)
			if err != nil {
				return "", "err", err
			}
		}

		offers = append(offers, &OfferAllContent{
			oac.Id, oac.Offer_uid, oac.Merchant_uid, oac.Offer_point, oac.Offer_cat_id, oac.Offer_image_banner, oac.Offer_image_poster,
			oac.Used, oac.Quantity, oac.Name_en, oac.Name_th, oac.Condition_offer_en, oac.Condition_offer_th,
			oac.Description_en, oac.Description_th, oac.Status, oac.Value, oac.Offer_type, oac.Create_at, oac.Update_at,
		})
	}
	// log.Println(offers)
	s, _ := json.Marshal(offers)
	offers_of_merchant := string(s)
	return offers_of_merchant, "success", err
}
func ListsAllOfferByStatus(service_name string, status string) (string, string, error) {
	offerMeta_table := service_name + "_offer_meta"
	offerContent_table := service_name + "_offer_content"
	ConnectDb()
	var oac OfferAllContent
	SQL_SELECT_OFFERMETA := "SELECT * FROM " + offerMeta_table + " WHERE status = ?"
	rows, err := DB.Query(SQL_SELECT_OFFERMETA, status)
	if err != nil {
		return "", "err", err
	}

	offers := make([]*OfferAllContent, 0, 17)
	for rows.Next() {
		err := rows.Scan(&oac.Id, &oac.Offer_uid, &oac.Offer_point, &oac.Merchant_uid,
			&oac.Offer_cat_id, &oac.Offer_image_banner, &oac.Offer_image_poster,
			&oac.Used, &oac.Quantity, &oac.Status, &oac.Value, &oac.Offer_type, &oac.Create_at, &oac.Update_at)
		if err != nil {
			return "", "err", err
		}

		SQL_SELECT_OFFER_TH := "SELECT name, condition_offer, description FROM " + offerContent_table + " WHERE offer_uid=? AND lang=?"
		rows1, err := DB.Query(SQL_SELECT_OFFER_TH, oac.Offer_uid, "th")
		if err != nil {
			return "", "err", err
		}
		for rows1.Next() {
			err := rows1.Scan(&oac.Name_th, &oac.Condition_offer_th, &oac.Description_th)
			if err != nil {
				return "", "err", err
			}
		}

		SQL_SELECT_OFFER_EN := "SELECT name, condition_offer, description FROM " + offerContent_table + " WHERE offer_uid=? AND lang=?"
		rows2, err := DB.Query(SQL_SELECT_OFFER_EN, oac.Offer_uid, "en")
		if err != nil {
			return "", "err", err
		}
		for rows2.Next() {
			err := rows2.Scan(&oac.Name_en, &oac.Condition_offer_en, &oac.Description_en)
			if err != nil {
				return "", "err", err
			}
		}

		offers = append(offers, &OfferAllContent{
			oac.Id, oac.Offer_uid, oac.Merchant_uid, oac.Offer_point, oac.Offer_cat_id, oac.Offer_image_banner, oac.Offer_image_poster,
			oac.Used, oac.Quantity, oac.Name_en, oac.Name_th, oac.Condition_offer_en, oac.Condition_offer_th,
			oac.Description_en, oac.Description_th, oac.Status, oac.Value, oac.Offer_type, oac.Create_at, oac.Update_at,
		})
	}
	// log.Println(offers)
	s, _ := json.Marshal(offers)
	offers_of_merchant := string(s)
	return offers_of_merchant, "success", err
}

func (offer *OfferAllContent) ShowOfferInfo(service_name string) (string, string, error) {
	offerMeta_table := service_name + "_offer_meta"
	offerContent_table := service_name + "_offer_content"
	ConnectDb()
	var (
		err error
	)

	var oac OfferAllContent
	SQL_SELECT_OFFERMETA := "SELECT * FROM " + offerMeta_table + " WHERE offer_uid=?"
	rows, err := DB.Query(SQL_SELECT_OFFERMETA, offer.Offer_uid)
	if err != nil {
		return "", "err", err
	}

	// offers := make([]*OfferAllContent, 0, 17)
	for rows.Next() {
		err := rows.Scan(&oac.Id, &oac.Offer_uid, &oac.Offer_point, &oac.Merchant_uid,
			&oac.Offer_cat_id, &oac.Offer_image_banner, &oac.Offer_image_poster,
			&oac.Used, &oac.Quantity, &oac.Status, &oac.Value, &oac.Offer_type, &oac.Create_at, &oac.Update_at)
		if err != nil {
			return "", "err", err
		}
	}

	SQL_SELECT_OFFER_TH := "SELECT name, condition_offer, description FROM " + offerContent_table + " WHERE offer_uid=? AND lang=?"
	rows1, err := DB.Query(SQL_SELECT_OFFER_TH, oac.Offer_uid, "th")
	if err != nil {
		return "", "err", err
	}
	for rows1.Next() {
		err := rows1.Scan(&oac.Name_th, &oac.Condition_offer_th, &oac.Description_th)
		if err != nil {
			return "", "err", err
		}
	}

	SQL_SELECT_OFFER_EN := "SELECT name, condition_offer, description FROM " + offerContent_table + " WHERE offer_uid=? AND lang=?"
	rows2, err := DB.Query(SQL_SELECT_OFFER_EN, oac.Offer_uid, "en")
	if err != nil {
		return "", "err", err
	}
	for rows2.Next() {
		err := rows2.Scan(&oac.Name_en, &oac.Condition_offer_en, &oac.Description_en)
		if err != nil {
			return "", "err", err
		}
	}

	result := OfferAllContent{
		oac.Id, oac.Offer_uid, oac.Merchant_uid, oac.Offer_point, oac.Offer_cat_id, oac.Offer_image_banner, oac.Offer_image_poster,
		oac.Used, oac.Quantity, oac.Name_en, oac.Name_th, oac.Condition_offer_en, oac.Condition_offer_th,
		oac.Description_en, oac.Description_th, oac.Status, oac.Value, oac.Offer_type, oac.Create_at, oac.Update_at,
	}
	s, _ := json.Marshal(result)
	offers_info := string(s)
	// log.Println(string(s))
	return offers_info, "success", err
}
