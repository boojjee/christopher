package models

import (
// "encoding/json"
// "github.com/elgs/gosqljson"
// "log"
// "strings"
)

type OfferAllContent struct {
	Offer_uid          string
	Merchant_uid       string
	Offer_cat_id       int64
	Offer_image_banner string
	Offer_image_poster string
	Used               int64
	Qty                int64
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
