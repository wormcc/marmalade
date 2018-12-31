package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	_accountHttpDelivery "github.com/wormcc/marmalade/account/delivery/http"
	_accountRepo "github.com/wormcc/marmalade/account/repository"
	_accountUseCase "github.com/wormcc/marmalade/account/usercase"
	"log"
	"net/url"
	"os"
	"time"
)

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initDB() *sql.DB {
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
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil && viper.GetBool("debug") {
		fmt.Println(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return dbConn
}

func main() {
	if viper.GetBool("debug") == false {
		fmt.Println("Server running in production model")
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	dbConn := initDB()
	defer dbConn.Close()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	accountRepo := _accountRepo.NewMysqlAccountRepository(dbConn)
	au := _accountUseCase.NewAccountUseCase(accountRepo, timeoutContext)
	_accountHttpDelivery.NewAccountHttpHandler(r, au)

	r.Run(viper.GetString("server.address"))
}
