package controllers

import (
	//	"encoding/json"
	//	"fmt"
	"github.com/astaxie/beego"
)

type RouterCommController struct {
	beego.Controller
}

func (this *RouterCommController) Get() {
	type DhcpParams struct {
		Dns string
	}
	type PppoeParams struct {
		Usrname string
		Usrpass string
		Dns     string
	}
	type CommandStats struct {
		Jsonrpc     string
		Id          string
		Method      string
		Dhcpparams  DhcpParams
		Pppoeparams PppoeParams
	}

	type CommandData struct {
		Etype string
		Edata CommandStats
	}

	result := CommandData{
		Etype: "1",
		Edata: CommandStats{
			Id:          "1",
			Method:      "Reds",
			Pppoeparams: PppoeParams{Usrname: "1243244", Usrpass: "jiang"}},
	}

	this.Data["json"] = result

	this.ServeJSON()

}

func (this *RouterCommController) Command1() {

	type CommandStats struct {
		Jsonrpc string
		Id      string
		Params  string
	}

	type CommandData struct {
		Etype string
		Edata CommandStats
	}

	result := CommandData{
		Etype: "1",
		Edata: CommandStats{
			Id:     "1",
			Params: "1",
		},
	}

	this.Data["json"] = result

	this.ServeJSON()

}

func (this *RouterCommController) Command3() {
	type DhcpParams struct {
		Dns string
	}
	type PppoeParams struct {
		Usrname string
		Usrpass string
		Dns     string
	}
	type StaticParams struct {
		Ipaddr  string
		Network string
		Gateway string
		Dns1    string
		Dns2    string
	}
	type CommandStats struct {
		Jsonrpc      string
		Id           string
		Params       string
		Dhcpparams   DhcpParams
		Pppoeparams  PppoeParams
		Staticparams StaticParams
	}

	type CommandData struct {
		Etype string
		Edata CommandStats
	}

	result := CommandData{
		Etype: "3",
		Edata: CommandStats{
			Id:         "1",
			Params:     "1",
			Dhcpparams: DhcpParams{Dns: "8.8.8.8"}},
	}

	result1 := CommandData{
		Etype: "2",
		Edata: CommandStats{
			Id:          "1",
			Params:      "2",
			Pppoeparams: PppoeParams{Usrname: "1243244", Usrpass: "jiang", Dns: "4.4.4.4"}},
	}

	result2 := CommandData{
		Etype: "2",
		Edata: CommandStats{
			Id:     "1",
			Params: "3",
			Staticparams: StaticParams{Ipaddr: "192.168.2.33", Network: "255.255.255.0",
				Gateway: "192.168.1.1", Dns1: "8.8.8.8", Dns2: "4.4.4.4"}},
	}

	this.Data["json"] = result
	this.Data["json1"] = result1
	this.Data["json2"] = result2

	this.ServeJSON()

}

func (this *RouterCommController) Command2() {

	type CommandStats struct {
		Jsonrpc string
		Id      string
		Params  string
	}

	type CommandData struct {
		Etype string
		Edata CommandStats
	}

	result := CommandData{
		Etype: "2",
		Edata: CommandStats{
			Id:     "1",
			Params: "password",
		},
	}

	this.Data["json"] = result

	this.ServeJSON()

}

func (this *RouterCommController) Command4() {

	type WifiParams struct {
		Wifiname string
		Wifipass string
	}

	type CommandStats struct {
		Jsonrpc    string
		Id         string
		Params     string
		Wifiparams []WifiParams
	}

	type CommandData struct {
		Etype string
		Edata CommandStats
	}

	result := CommandData{
		Etype: "4",
		Edata: CommandStats{
			Id:         "1",
			Params:     "1",
			Wifiparams: []WifiParams{{Wifiname: "jiang", Wifipass: "jiang"}, {Wifiname: "bbbbb", Wifipass: "jiang"}},
		},
	}

	this.Data["json"] = result

	this.ServeJSON()

}
func (this *RouterCommController) Command5() {

	type MacParams struct {
		Macname   string
		Macaccess string
	}

	type CommandStats struct {
		Jsonrpc   string
		Id        string
		Params    string
		Macparams []MacParams
	}

	type CommandData struct {
		Etype string
		Edata CommandStats
	}

	result := CommandData{
		Etype: "5",
		Edata: CommandStats{
			Id:     "1",
			Params: "1",
			Macparams: []MacParams{{Macname: "01:02:03:04:05:06", Macaccess: "enable"},
				{Macname: "01:02:03:04:05:06", Macaccess: "disable"}},
		},
	}

	this.Data["json"] = result

	this.ServeJSON()

}

func (this *RouterCommController) Command6() {

	type CommandStats struct {
		Jsonrpc string
		Id      string
		Params  string
	}

	type CommandData struct {
		Etype string
		Edata CommandStats
	}

	result := CommandData{
		Etype: "6",
		Edata: CommandStats{
			Id:     "1",
			Params: "1",
		},
	}

	this.Data["json"] = result

	this.ServeJSON()

}
func (this *RouterCommController) Command7() {

	type CommandStats struct {
		Jsonrpc string
		Id      string
		Params  string
	}

	type CommandData struct {
		Etype string
		Edata CommandStats
	}

	result := CommandData{
		Etype: "7",
		Edata: CommandStats{
			Id:     "1",
			Params: "1",
		},
	}

	this.Data["json"] = result

	this.ServeJSON()

}

func (this *RouterCommController) Command8() {

	type UrlParams struct {
		Urlname   string
		Urlaccess string
	}

	type CommandStats struct {
		Jsonrpc   string
		Id        string
		Params    string
		Urlparams []UrlParams
	}

	type CommandData struct {
		Etype string
		Edata CommandStats
	}

	result := CommandData{
		Etype: "8",
		Edata: CommandStats{
			Id:     "1",
			Params: "1",
			Urlparams: []UrlParams{{Urlname: "www.qq.com", Urlaccess: "enable"},
				{Urlname: "www.baidu.com", Urlaccess: "disable"}},
		},
	}

	this.Data["json"] = result

	this.ServeJSON()

}

func (this *RouterCommController) Command9() {

	type CommandStats struct {
		Jsonrpc string
		Id      string
		Params  string
	}

	type CommandData struct {
		Etype string
		Edata CommandStats
	}

	result := CommandData{
		Etype: "9",
		Edata: CommandStats{
			Id:     "1",
			Params: "ps",
		},
	}

	this.Data["json"] = result

	this.ServeJSON()

}

func (this *RouterCommController) Command10() {

	type CommandStats struct {
		Jsonrpc string
		Id      string
		Params  string
	}

	type CommandData struct {
		Etype string
		Edata CommandStats
	}

	result := CommandData{
		Etype: "1",
		Edata: CommandStats{
			Id:     "1",
			Params: "1",
		},
	}

	this.Data["json"] = result

	this.ServeJSON()

}
