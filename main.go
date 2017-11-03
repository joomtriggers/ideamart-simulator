package main

import (
	"github.com/astaxie/beego"
	_ "github.com/joomtriggers/ideamart-simulator/routers"
)

func main() {
	beego.Run()
}
