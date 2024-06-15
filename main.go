package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "outline-importer-confluence/routers"
)

func main() {
	beego.Run()
}
