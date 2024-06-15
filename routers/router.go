package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"outline-importer-confluence/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/ping", &controllers.PingController{})
}
