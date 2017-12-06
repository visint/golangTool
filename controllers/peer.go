package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jxufeliujj/blog/models"
)

type PeerController struct {
	beego.Controller
}

func (this *PeerController) Get() {

	phone := this.GetString("phone")
	password := this.GetString("password")
	validate := this.GetString("validate")
	fmt.Println(phone, password, validate)

	type ReportSats struct {
		ReturnCode int
		Data       string
		desc       string
	}

	result := ReportSats{
		ReturnCode: 200,
		Data:       "Reds",
		desc:       "",
	}
	this.Data["json"] = result

	this.ServeJSON()

}

func (this *PeerController) Save() {

	phone := this.GetString("phone")
	password := this.GetString("password")
	validate := this.GetString("validate")

	fmt.Println(phone, password, validate)

	type ReportSats struct {
		ReturnCode int
		Data       string
		desc       string
	}

	result := ReportSats{
		ReturnCode: 200,
		Data:       "Reds",
		desc:       "",
	}
	this.Data["json"] = result
	this.ServeJSON()
}

func (this *PeerController) Login() {

	phone := this.GetString("phone")
	password := this.GetString("password")

	fmt.Println(phone, password)

	type ReportSats struct {
		ReturnCode int
		Data       string
		desc       string
	}

	result := ReportSats{
		ReturnCode: 200,
		Data:       "Reds",
		desc:       "成功",
	}
	this.Data["json"] = result

	this.ServeJSON()
}

func (this *PeerController) Message() {
	var list []*models.User
	var user models.User

	count, _ := user.Query().Count()

	if count > 0 {
		user.Query().OrderBy("-id").All(&list)
	}

	fmt.Println("jiang", count)

	type ReportSats struct {
		Type int
		Data string
		desc string
	}

	result := ReportSats{
		Type: 200,
		Data: "Reds",
		desc: "成功",
	}
	this.Data["json"] = result

	this.ServeJSON()

}

/*
{
    "type": 1,
"msg: {
 "aps": {
            "action": 4,
            "alert": "绑定成功",
            "badge": 0,
  "unlocktime": "12:00"
  "child": "abcdefg"
        }
 }，
"device":1,
"token":"123456789",
"topic":"action"
}
*/

// /business/parents/save?phone=12345678910&password=123456&validate=1234

// /business/parents/login?phone=12345678910&password=123456
/*
type StudentS struct {
	phone    string
	password string
	validate string
}

type StudentS struct {
	nick_name       string
	photo           string
	age             int
	sex             int
	status          int
	today_ban_begin string
	today_ban_end   string
	eye_protect     string
	Id              int
	cid             int
	clock_name      string
	begintime       string
	endtime         string
	mon             int
	tues            int
	wed             int
	thur            int
	fri             int
	sat             int
	sun             int
	create_time     int
}
*/
