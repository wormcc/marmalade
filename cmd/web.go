package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
	_accountHttpDelivery "github.com/wormcc/marmalade/account/delivery/http"
	_accountRepo "github.com/wormcc/marmalade/account/repository"
	_accountUseCase "github.com/wormcc/marmalade/account/usercase"
	"github.com/wormcc/marmalade/common"
	"time"
)

var Web = cli.Command{
	Name: "web",
	Usage: "Start web server",
	Action: runWeb,
}

func initConfig() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func runWeb(_ *cli.Context) {
	initConfig()
	if viper.GetBool("debug") == false {
		fmt.Println("Server running in production model")
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	dbConn := common.Init()
	defer dbConn.Close()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	accountRepo := _accountRepo.NewMysqlAccountRepository(dbConn)
	au := _accountUseCase.NewAccountUseCase(accountRepo, timeoutContext)
	_accountHttpDelivery.NewAccountHttpHandler(r, au)

	r.Run(viper.GetString("server.address"))
}