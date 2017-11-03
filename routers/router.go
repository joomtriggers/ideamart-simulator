package routers

import (
	"github.com/joomtriggers/ideamart-simulator/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
