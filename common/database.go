package common

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/url"
	"os"
)

var DB *sql.DB

func Init() *sql.DB {
	var err error
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Shanghai")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	DB, err = sql.Open(`mysql`, dsn)
	if err != nil && viper.GetBool("debug") {
		fmt.Println(err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return DB
}

func GetDB() *sql.DB {
	return DB
}
