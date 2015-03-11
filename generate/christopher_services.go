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
	// defer DB.Close()
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

		//Merchant
		var merchant = prefix + "_" + "merchants"
		var merchant_table = `CREATE TABLE ` + merchant + ` (
        id int(11) unsigned NOT NULL AUTO_INCREMENT,
        username varchar(200) DEFAULT NULL,
        name varchar(200) DEFAULT NULL UNIQUE,
        password varchar(100) DEFAULT NULL,
        email varchar(100) DEFAULT NULL,
        shop_image varchar(200) DEFAULT NULL,
        shop_avatar varchar(200) DEFAULT NULL,
        shop_description text,
        lat varchar(100) DEFAULT NULL,
        lon varchar(100) DEFAULT NULL,
        create_at timestamp NULL DEFAULT NULL,
        update_at timestamp NULL DEFAULT '0000-00-00 00:00:00',
        PRIMARY KEY (id)
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8;`

		result_merchant, err := DB.Exec(merchant_table)
		log.Println(result_merchant)
		checkErr(err)
		//------------------------------------------------------

		var offer = prefix + "_" + "offers"
		var offer_table = `CREATE TABLE ` + offer + ` (
    id int(11) unsigned NOT NULL AUTO_INCREMENT,
    name varchar(200) DEFAULT NULL,
    offer_point double(100,10) DEFAULT NULL,
    condition_offer varchar(200) DEFAULT NULL,
    cat int(11) DEFAULT NULL,
    merchant_id int(11) DEFAULT NULL,
    offer_image varchar(200) DEFAULT NULL,
    description text,
    used int(11) DEFAULT NULL,
    qty int(11) DEFAULT NULL,
    create_at timestamp NULL DEFAULT NULL,
    update_at timestamp NULL DEFAULT NULL,
    PRIMARY KEY (id)
    ) ENGINE=InnoDB  DEFAULT CHARSET=utf8;`

		result_offer, err := DB.Exec(offer_table)
		log.Println(result_offer)
		checkErr(err)
		//------------------------------------------------------

		var point = prefix + "_" + "point"
		var point_table = `CREATE TABLE ` + point + ` (
    id int(11) unsigned NOT NULL AUTO_INCREMENT,
    user_id int(11) DEFAULT NULL,
    point double(100,10) DEFAULT NULL,
    create_at timestamp NULL DEFAULT NULL,
    update_at timestamp NULL DEFAULT NULL,
    PRIMARY KEY (id)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;`

		result_point, err := DB.Exec(point_table)
		log.Println(result_point)
		checkErr(err)
		//------------------------------------------------------

		var redeem = prefix + "_" + "redeem"
		var redeem_table = `CREATE TABLE ` + redeem + ` (
    id int(11) unsigned NOT NULL AUTO_INCREMENT,
    offer_id int(11) DEFAULT NULL,
    user_id int(11) DEFAULT NULL,
    redeem_point double(100,10) DEFAULT NULL,
    code varchar(100) DEFAULT NULL,
    expiry_date datetime DEFAULT NULL,
    status int(11) DEFAULT NULL,
    create_at timestamp NULL DEFAULT NULL,
    update_at timestamp NULL DEFAULT NULL,
    PRIMARY KEY (id)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;`
		result_redeem, err := DB.Exec(redeem_table)
		log.Println(result_redeem)
		checkErr(err)
		//------------------------------------------------------

		var activity = prefix + "_" + "activity"
		var activity_table = `CREATE TABLE ` + activity + ` ( 
    id int(11) unsigned NOT NULL AUTO_INCREMENT,
    activity_id varchar(200) DEFAULT NULL,
    distance double(100,10) DEFAULT NULL,
    status int(11) DEFAULT NULL,
    user_id int(11) DEFAULT NULL,
    create_at timestamp NULL DEFAULT NULL,
    update_at timestamp NULL DEFAULT NULL,
    PRIMARY KEY (id)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;`
		result_activity, err := DB.Exec(activity_table)
		log.Println(result_activity)
		checkErr(err)
		//------------------------------------------------------

		var user = prefix + "_" + "user"
		var user_table = `CREATE TABLE ` + user + ` (  
    id int(11) unsigned NOT NULL AUTO_INCREMENT,
    pin varchar(100) DEFAULT NULL,
    parse_id varchar(200) DEFAULT NULL,
    create_at timestamp NULL DEFAULT NULL,
    update_at timestamp NULL DEFAULT NULL,
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
