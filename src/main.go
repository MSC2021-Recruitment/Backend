package main

import (
	"MSC2021/src/global"
	"MSC2021/src/initialize"
	"log"
)

func main() {
	if err := initialize.InitConfig(); err != nil {
		log.Fatalln("Could not init configuration, exit...")
	}
	global.LOGGER.Info("\n" +
		"===================================================================================\n" +
		"   MSC Recruitment 2021 Version 0.1-2021.09.06+Canary\n" +
		"   Powered By Reverier from MSC in XDU with love.\n\n" +
		"   GitHub Repo: https://github.com/MSC2021-Recruitment/Backend\n\n" +
		"   If you have some trouble in using it, plz feel free to create an issue.\n" +
		"   you can also contact author by emailing to reverier.xu@outlook.com.\n" +
		"===================================================================================")
	global.DATABASE = initialize.GormMysql()
	initialize.InitTables(global.DATABASE)
	global.REDIS = initialize.Redis()
	engine, err := initialize.InitRouter()
	if err != nil {
		log.Fatalln("Could not init router, exit...")
	}
	engine.Run(":8080")
}
