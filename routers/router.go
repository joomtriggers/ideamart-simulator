package routers

import (
	"github.com/astaxie/beego"
	"github.com/joomtriggers/ideamart-simulator/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
