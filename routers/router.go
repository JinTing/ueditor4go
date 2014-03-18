package routers

import (
	"beeblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.SetStaticPath("/ueditor", "ueditor")
	beego.Router("/ue", &controllers.UEController{})
	beego.AutoRouter(&controllers.UEController{})
}
