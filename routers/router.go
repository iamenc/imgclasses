package routers

import (
	"imgclasses/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/classes", &controllers.ClassesImg{})
}
