package main

import (
	"github.com/astaxie/beego/orm"
	_ "lianyun/routers"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func init()  {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:111111@tcp(39.104.142.193:3306)/xiuzhen?charset=utf8")
	orm.RegisterDataBase("lianyun", "mysql", "root:111111@tcp(39.104.142.193:3306)/lianyun?charset=utf8")
}
func main() {

	//if beego.BConfig.RunMode == "dev" {
	//	beego.BConfig.WebConfig.DirectoryIndex = true
	//	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	//}
	beego.Run()
}
