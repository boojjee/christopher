package generate

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var DB *sql.DB

// Creates and tests database connection
func ConnectDb() {
	var err error
	ds := "root:AdminAdmin@/christopher?parseTime=true"
	DB, err = sql.Open("mysql", ds)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	// defer DB.Close()]
}

// Closes database connection
func CloseDb() {
	DB.Close()
}

func Gen_table(c *gin.Context) {
	var prefix = c.Params.ByName("service_name")
	ConnectDb()
	var (
		id        string
		name      string
		details   string
		create_at time.Time
		update_at time.Time
	)
	var CHECK_IS_HAVE_SERVICE = `SELECT * FROM christopher_service WHERE name = ?`
	result_check_service, err := DB.Query(CHECK_IS_HAVE_SERVICE, prefix)
	for result_check_service.Next() {
		err := result_check_service.Scan(&id, &name, &details, &create_at, &update_at)
		if err != nil {
			log.Println(err)
		}
	}
	if err != nil {
		log.Println(err)
	}

	if name == "" {
		// create DB
		var INSERT_SERVICE = `INSERT INTO christopher_service (name, details, create_at, update_at) VALUES (?, ?, ?, ?)`
		result_insert_service, err := DB.Exec(INSERT_SERVICE, prefix, "null", time.Now(), time.Now())
		log.Println(result_insert_service)
		checkErr(err)

		//Merchant Meta
		var mmerchant_metachant = prefix + "_" + "merchant_meta"
		var merchant_meta_table = `CREATE TABLE ` + mmerchant_metachant + ` (
        id BIGINT(20) unsigned NOT NULL AUTO_INCREMENT,
        merchant_uid VARCHAR(255) DEFAULT NULL UNIQUE,  
        username VARCHAR(255) DEFAULT NULL,
        password VARCHAR(255) DEFAULT NULL,
        email VARCHAR(255) DEFAULT NULL,
        shop_avatar_s VARCHAR(200) DEFAULT NULL,
        shop_avatar_l VARCHAR(200) DEFAULT NULL,
        lat VARCHAR(255) DEFAULT NULL,
        lon VARCHAR(255) DEFAULT NULL,
        phone_1 VARCHAR(255) DEFAULT NULL,
        phone_2 VARCHAR(255) DEFAULT NULL,
        fax VARCHAR(255) DEFAULT NULL,
        line_id VARCHAR(255) DEFAULT NULL,
        facebook_link VARCHAR(255) DEFAULT NULL,
        website_link VARCHAR(255) DEFAULT NULL, 
        merchant_status VARCHAR(255) DEFAULT NULL, 
        create_at BIGINT(20) NULL DEFAULT NULL,
        update_at BIGINT(20) NULL DEFAULT NULL,
        PRIMARY KEY (id)
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8;`

		result_merchant_meta, err := DB.Exec(merchant_meta_table)
		log.Println(result_merchant_meta)
		checkErr(err)
		//------------------------------------------------------
		//Merchant
		var merchant_content = prefix + "_" + "merchant_content"
		var merchant_content_table = `CREATE TABLE ` + merchant_content + ` (
        id BIGINT(20) unsigned NOT NULL AUTO_INCREMENT,
        merchant_uid VARCHAR(255) DEFAULT NULL, 
        name VARCHAR(255) DEFAULT NULL ,
        shop_description LONGTEXT,  
        lang VARCHAR(10) DEFAULT NULL,
        create_at BIGINT(20) NULL DEFAULT NULL,
        update_at BIGINT(20) NULL DEFAULT NULL,
        PRIMARY KEY (id)
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8;`

		result_merchant_content, err := DB.Exec(merchant_content_table)
		log.Println(result_merchant_content)
		checkErr(err)
		//------------------------------------------------------

		//Merchant
		var merchant_gallery = prefix + "_" + "merchants_photo_gallery"
		var merchants_photo_gallery_table = `CREATE TABLE ` + merchant_gallery + ` (
        id BIGINT(20) unsigned NOT NULL AUTO_INCREMENT,
        photo_url VARCHAR(255) DEFAULT NULL UNIQUE,
        merchant_uid VARCHAR(255) DEFAULT NULL , 
        create_at BIGINT(20) NULL DEFAULT NULL,
        update_at BIGINT(20) NULL DEFAULT NULL,
        PRIMARY KEY (id)
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8;`

		result_merchants_photo_gallery_table, err := DB.Exec(merchants_photo_gallery_table)
		log.Println(result_merchants_photo_gallery_table)
		checkErr(err)
		//------------------------------------------------------
		// Offer Catagory
		var merchant_category = prefix + "_" + "merchant_category"
		var merchant_category_table = `CREATE TABLE ` + merchant_category + ` (
        id BIGINT(20) unsigned NOT NULL AUTO_INCREMENT, 
        merchant_category_name VARCHAR(200) DEFAULT NULL, 
        slug VARCHAR(200) DEFAULT NULL, 
        create_at BIGINT(20) NULL DEFAULT NULL,
        update_at BIGINT(20) NULL DEFAULT NULL,
        PRIMARY KEY (id)
        ) ENGINE=InnoDB  DEFAULT CHARSET=utf8;`

		result_merchant_category, err := DB.Exec(merchant_category_table)
		log.Println(result_merchant_category)
		checkErr(err)
		//------------------------------------------------------
		// Offer Meta
		var offer_meta = prefix + "_" + "offer_meta"
		var offer_meta_table = `CREATE TABLE ` + offer_meta + ` (
        id BIGINT(20) unsigned NOT NULL AUTO_INCREMENT, 
        offer_uid VARCHAR(255) DEFAULT NULL UNIQUE, 
        offer_point double(100,5) DEFAULT NULL,
        merchant_uid VARCHAR(255) DEFAULT NULL,
        offer_cat_id BIGINT(20) DEFAULT NULL,
        offer_image_banner VARCHAR(200) DEFAULT NULL,
        offer_image_poster VARCHAR(200) DEFAULT NULL,
        used BIGINT(20) DEFAULT NULL,
        quantity BIGINT(20) DEFAULT NULL,
        total BIGINT(20) DEFAULT NULL,
        create_at BIGINT(20) NULL DEFAULT NULL,
        update_at BIGINT(20) NULL DEFAULT NULL,
        PRIMARY KEY (id)
        ) ENGINE=InnoDB  DEFAULT CHARSET=utf8;`

		result_offer_meta, err := DB.Exec(offer_meta_table)
		log.Println(result_offer_meta)
		checkErr(err)
		//------------------------------------------------------

		// Offer Content
		var offer = prefix + "_" + "offer_content"
		var offer_table = `CREATE TABLE ` + offer + ` (
        id BIGINT(20) unsigned NOT NULL AUTO_INCREMENT, 
        offer_uid VARCHAR(255) DEFAULT NULL,
        name VARCHAR(200) DEFAULT NULL,
        condition_offer VARCHAR(200) DEFAULT NULL, 
        description LONGTEXT, 
        lang VARCHAR(10) DEFAULT NULL,
        create_at BIGINT(20) NULL DEFAULT NULL,
        update_at BIGINT(20) NULL DEFAULT NULL,
        PRIMARY KEY (id)
        ) ENGINE=InnoDB  DEFAULT CHARSET=utf8;`

		result_offer, err := DB.Exec(offer_table)
		log.Println(result_offer)
		checkErr(err)
		//------------------------------------------------------
		// Offer Catagory
		var offer_category = prefix + "_" + "offer_category"
		var offer_category_table = `CREATE TABLE ` + offer_category + ` (
        id BIGINT(20) unsigned NOT NULL AUTO_INCREMENT, 
        offer_category_name VARCHAR(200) DEFAULT NULL, 
        slug VARCHAR(200) DEFAULT NULL, 
        create_at BIGINT(20) NULL DEFAULT NULL,
        update_at BIGINT(20) NULL DEFAULT NULL,
        PRIMARY KEY (id)
        ) ENGINE=InnoDB  DEFAULT CHARSET=utf8;`

		result_offer_category, err := DB.Exec(offer_category_table)
		log.Println(result_offer_category)
		checkErr(err)
		//------------------------------------------------------
		// Point
		var point = prefix + "_" + "point"
		var point_table = `CREATE TABLE ` + point + ` (
        id BIGINT(20) unsigned NOT NULL AUTO_INCREMENT,
        activity_uid VARCHAR(200) DEFAULT NULL,
        parse_uid VARCHAR(200) DEFAULT NULL,
        user_uid VARCHAR(200) DEFAULT NULL,
        distance double(100,5) DEFAULT NULL,
        g_point double(100,5) DEFAULT NULL,
        g_point_status double(100,5) DEFAULT NULL,
        g_point_expire BIGINT(20) DEFAULT NULL,
        create_at BIGINT(20) NULL DEFAULT NULL,
        update_at BIGINT(20) NULL DEFAULT NULL,
        PRIMARY KEY (id)
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8;`

		result_point, err := DB.Exec(point_table)
		log.Println(result_point)
		checkErr(err)
		//------------------------------------------------------

		var redeem = prefix + "_" + "redeem"
		var redeem_table = `CREATE TABLE ` + redeem + ` (
        id BIGINT(20) unsigned NOT NULL AUTO_INCREMENT,
        offer_id BIGINT(20) DEFAULT NULL,
        user_id BIGINT(20) DEFAULT NULL,
        redeem_point double(100,5) DEFAULT NULL,
        code VARCHAR(100) DEFAULT NULL,
        expiry_date datetime DEFAULT NULL,
        status BIGINT(20) DEFAULT NULL,
        create_at BIGINT(20) NULL DEFAULT NULL,
        update_at BIGINT(20) NULL DEFAULT NULL,
        PRIMARY KEY (id)
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8;`
		result_redeem, err := DB.Exec(redeem_table)
		log.Println(result_redeem)
		checkErr(err)
		//------------------------------------------------------

		var activity = prefix + "_" + "activity"
		var activity_table = `CREATE TABLE ` + activity + ` ( 
        id BIGINT(20) unsigned NOT NULL AUTO_INCREMENT,
        activity_uid VARCHAR(200) DEFAULT NULL,
        user_uid VARCHAR(200) DEFAULT NULL,
        user_parse_id VARCHAR(200) DEFAULT NULL,
        third_activity_id VARCHAR(200) DEFAULT NULL,
        third_uri VARCHAR(200) DEFAULT NULL,
        third_token_user VARCHAR(200) DEFAULT NULL,
        source VARCHAR(200) DEFAULT NULL,
        distance double(100,5) DEFAULT NULL,
        duration double(100,5) DEFAULT NULL,
        calories double(100,5) DEFAULT NULL,
        start_activity_lat double(100,7) DEFAULT NULL,
        start_activity_lon double(100,7) DEFAULT NULL,
        activity_type BIGINT(20) DEFAULT NULL, 
        activity_status BIGINT(20) DEFAULT NULL, 
        create_at BIGINT(20) NULL DEFAULT NULL,
        update_at BIGINT(20) NULL DEFAULT NULL,
        PRIMARY KEY (id)
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8;`
		result_activity, err := DB.Exec(activity_table)
		log.Println(result_activity)
		checkErr(err)
		//------------------------------------------------------

		var user = prefix + "_" + "user"
		var user_table = `CREATE TABLE ` + user + ` (  
        id BIGINT(20) unsigned NOT NULL AUTO_INCREMENT,
        pin VARCHAR(100) DEFAULT NULL,
        parse_id VARCHAR(200) DEFAULT NULL,
        user_status BIGINT(20) NULL DEFAULT NULL,
        create_at BIGINT(20) NULL DEFAULT NULL,
        update_at BIGINT(20) NULL DEFAULT NULL,
        PRIMARY KEY (id)
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8;`
		result_user, err := DB.Exec(user_table)
		log.Println(result_user)

		CloseDb()
		c.JSON(200, gin.H{"status": "Service Generated!"})

	} else {

		c.JSON(200, gin.H{"status": "Service is already!"})
	}
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
