package main

import (
	"fmt"
	"github.com/go-ini/ini"
)

type ConfigList struct {
	Port int
	DbName string
	SQLDriver string
}

var Config ConfigList

func init() {
	//config.iniはwebサーバーで普及している設定ファイル
	//os.Openなどで開くこともできるが、サードパーティのものを使ったほうが効率的
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		Port: cfg.Section("Web").Key("port").MustInt(),
		DbName: cfg.Section("db").Key("name").String(),
		SQLDriver: cfg.Section("db").Key("driver").String(),
	}
}

func main() {
	fmt.Printf("%T %v\n", Config.Port, Config.Port) //int 8080
	fmt.Printf("%T %v\n", Config.DbName, Config.DbName) //string stockdata.sql
	fmt.Printf("%T %v\n", Config.SQLDriver, Config.SQLDriver) //string sqlite3
}