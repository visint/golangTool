package routers

import (
	"github.com/astaxie/beego"
	"helloworld/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/device/register", &controllers.MainController{}, "*:Regist")
	beego.Router("/device/report", &controllers.MainController{}, "*:Report")
	beego.Router("/device/report1", &controllers.MainController{}, "*:Report1")
	beego.Router("/device/report2", &controllers.MainController{}, "*:Report2")
	beego.Router("/device/Report_Interface", &controllers.MainController{}, "*:Report3")
	beego.Router("/device/config", &controllers.MainController{}, "*:Config")
	beego.Router("/device/firmware", &controllers.MainController{}, "*:Firmware")

	beego.Router("/router/command1", &controllers.RouterCommController{}, "*:Command1")
	beego.Router("/router/command2", &controllers.RouterCommController{}, "*:Command2")
	beego.Router("/router/command3", &controllers.RouterCommController{}, "*:Command3")
	beego.Router("/router/command4", &controllers.RouterCommController{}, "*:Command4")
	beego.Router("/router/command5", &controllers.RouterCommController{}, "*:Command5")
	beego.Router("/router/command6", &controllers.RouterCommController{}, "*:Command6")
	beego.Router("/router/command7", &controllers.RouterCommController{}, "*:Command7")
	beego.Router("/router/command8", &controllers.RouterCommController{}, "*:Command8")
	beego.Router("/router/command9", &controllers.RouterCommController{}, "*:Command9")
	beego.Router("/router/command10", &controllers.RouterCommController{}, "*:Command10")

	beego.Router("business/parents/login", &controllers.PeerController{}, "*:Login")
	beego.Router("business/parents/save", &controllers.PeerController{}, "*:Save")
	beego.Router("business/parents/message", &controllers.PeerController{}, "*:Message")
}
