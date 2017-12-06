package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	type Params struct {
		Acmac        string
		Id           int
		Duration     int
		Uptime       int
		Stations     []string
		Conf_version string
		Load         float32
		Freeram      int
		Wan_incoming int
		Wan_outgoing int
	}
	type ReportSats struct {
		Id      string
		Method  string
		Jsonrpc string
		Params  Params
	}

	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}

	tttt := ReportSats{
		Id:     "s",
		Method: "Reds",

		Params: Params{Acmac: "1243244"},
	}

	/*
		group := ColorGroup{
			ID:     1,
			Name:   "Reds",
			Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
		}

			b, err := json.Marshal(group)
			if err != nil {
				fmt.Println("error:", err, b)
			}
	*/
	this.Data["json"] = tttt

	this.ServeJSON()

}

func (this *MainController) Post() {

	type resultT struct {
		dev_id string
	}

	type ReportSats struct {
		Jsonrpc string
		Result  resultT
		Id      int
	}

	var fsd ReportSats

	json.Unmarshal(this.Ctx.Input.RequestBody, &fsd)

	fmt.Println(fsd)

	this.Data["json"] = fsd

	this.ServeJSON()
}

func (this *MainController) Regist() {

	type Result struct {
		pppoe_username string
		pppoe_password string
		route_password string
	}

	type RegistSats struct {
		Id      string
		Method  string
		Jsonrpc string
		Params  Result
	}

	// 插入数据库

	type resultT struct {
		dev_id int
	}

	type ResponseSats struct {
		Jsonrpc string
		Result  resultT
		Id      string
	}

	value := ResponseSats{
		Jsonrpc: "2.0",
		Result:  resultT{dev_id: 800},
		Id:      "81485cb90",
	}
	this.Data["json"] = value
	this.ServeJSON()
}

/* 接收包 */
/*
typedef struct resPack
{
  int id;
  int bufsize;
  int msgType;
} RecvPack;

*/

func (this *MainController) Report() {

	type ApsValue struct {
		Load    string
		Freeram string
		Uptime  string
		Channel string
		Apmac   string
	}

	type ParamsStat struct {
		Id           string
		Acmac        string
		Duration     int
		Uptime       int
		Stations     []string
		Conf_version string
		Load         float32
		Freeram      int
		Wan_incoming int
		Wan_outgoing int
		Aps          ApsValue
	}

	type ResponseSats struct {
		Resulttype string
		Paramsstat ParamsStat
	}

	result := ResponseSats{
		Resulttype: "1",
		Paramsstat: ParamsStat{Id: "dev", Acmac: "010203040506", Duration: 124, Uptime: 112},
	}
	fmt.Println(result)
	//	this.Data["json"] = result
	da, _ := ioutil.ReadFile("/tmp/wan_ip_get.json")
	this.Data["json"] = da
	this.ServeJSON()
}

func (this *MainController) Report1() {

	type clientMach struct {
		Mac          string
		Ipaddr       string
		Wan_incoming int
		Wan_outgoing int
	}

	type DhcpConfig struct {
		Dns string
	}
	type PppoeConfig struct {
		Auto     string
		Dns      string
		Username string
		Password string
	}
	type StaticConfig struct {
		Ipaddr  string
		Mask    string
		Gateway string
		Dns     string
	}
	type RealValue struct {
		Ipaddr  string
		Mask    string
		Gateway string
		Dns     string
		Mac     string
	}
	type RealData struct {
		Wantype      string
		Realvalue    RealValue
		Staticconfig StaticConfig
		Pppoeconfig  PppoeConfig
		Dhcpconfig   DhcpConfig
	}

	type ResponseSats struct {
		Resulttype string
		Realdata   RealData
		Clientmach []clientMach
	}

	result := ResponseSats{
		Resulttype: "2.0",
		Realdata: RealData{Wantype: "pppoe", Pppoeconfig: PppoeConfig{Username: "jiang",
			Password: "boob"},
			Realvalue: RealValue{Ipaddr: "192.168.1.8", Mask: "255.255.255.0"}},
		Clientmach: []clientMach{{Mac: "01:02:03:04:05:06", Ipaddr: "192.168.1.2",
			Wan_incoming: 0, Wan_outgoing: 0}},
	}
	this.Data["json"] = result
	this.ServeJSON()
}

func (this *MainController) Report2() {
	type wifiConf struct {
		Name string
		Pass string
	}

	type macFilter struct {
		Mac   string
		Stats string
	}
	type urlFilter struct {
		Url   string
		Stats string
	}

	type ResponseSats struct {
		Resulttype string
		Wificonf   []wifiConf
		Macfilter  []macFilter
		Urlfilter  []urlFilter
	}

	fmt.Println(string(this.Ctx.Input.RequestBody))

	result := ResponseSats{
		Resulttype: "2.0",
		Wificonf:   []wifiConf{{Name: "andWiFi-pifii", Pass: "jiang"}},
		Macfilter:  []macFilter{{Mac: "01:02:03:04:05:06", Stats: "enable"}},
		Urlfilter:  []urlFilter{{Url: "www.qq.com", Stats: "enable"}},
	}
	this.Data["json"] = result
	this.ServeJSON()
}

func (this *MainController) Report3() {
	type StatsValue struct {
		Mac string
	}

	type ApsValue struct {
		Load     float32
		Freeram  int
		Uptime   int
		Duration int
		Channel  string
		Apmac    string
		//		Stats    []StatsValue
	}

	type ParamsValue struct {
		Acmac        string
		Id           int
		Duration     int
		Uptime       int
		Conf_version string
		Load         float32
		Freeram      int
		Aps          []ApsValue
	}
	type ReportSats struct {
		Id      string
		Jsonrpc string
		Method  string

		Params ParamsValue
	}

	var fsd ReportSats

	fmt.Println(string(this.Ctx.Input.RequestBody))

	json.Unmarshal(this.Ctx.Input.RequestBody, &fsd)

	fmt.Println(fsd)

	type RealDate struct {
		Next_report_delay int
		ConfChange        int
		Reboot            int
	}

	type RealResult struct {
		Jsonrpc string   `json:"jsonrpc"`
		Result  RealDate `json:"result"`
		Id      string   `json:"id"`
	}

	result := RealResult{
		Jsonrpc: "2.0",
		Id:      "s",
		Result:  RealDate{},
	}

	this.Data["json"] = result
	this.ServeJSON()
}

func (this *MainController) Config() {

	type ParamsT struct {
		conf_version string
		acmac        string
	}

	type ConfigT struct {
		Id      string
		Method  string
		Jsonrpc string
		Params  []ParamsT
	}

	type WanInter struct {
		Ds_bandwidth int
		Us_bandwidth int
	}

	type Apconfstt struct {
		Apmac   string
		Channel string
	}

	type acconfT struct {
		White_list    []string
		Black_list    []string
		Office_ssid   string
		Password      string
		Lan_ssid      string
		Wan_interface WanInter
		Apconfs       Apconfstt
	}

	type resultT struct {
		Delay   int
		Version string
		Acconf  acconfT
		Id      string
	}

	type ReportSats struct {
		Jsonrpc string
		Result  resultT
	}

	tttt := ReportSats{
		Jsonrpc: "2.0",
		Result: resultT{Delay: 1, Version: "Reds",
			Acconf: acconfT{
				Office_ssid: "andWiFi",
				Password:    "bbb",
				Lan_ssid:    "dsf",
				Wan_interface: WanInter{Ds_bandwidth: 10,
					Us_bandwidth: 10},
				Apconfs: Apconfstt{Apmac: "dsa",
					Channel: "dsf"},
			},
			Id: "1234214"},
	}

	this.Data["json"] = tttt
	this.ServeJSON()
}

func (this *MainController) Firmware() {

	type Firmwareupdate struct {
		Id      string
		Method  string
		Jsonrpc string
		Params  []string
	}

	type ResultTT struct {
		Delay     int
		Swver     string
		Image_url string
	}

	type Firmwareresult struct {
		Jsonrpc string
		Result  ResultTT
		Id      string
	}

	tttt := Firmwareresult{
		Jsonrpc: "2.0",
		Result: ResultTT{Delay: 1,
			Swver:     "11234",
			Image_url: "fsddf"},
		//Image_url:"http:\/\/www.22wifi.cn\/upload\/firmware\/563c7d7fc039d.bin",},
		Id: "3ba4aa0",
	}

	this.Data["json"] = tttt
	this.ServeJSON()
}

func (this *MainController) FirmwareT() {
	/*
		type Firmwareupdate struct {
			Id      string
			Method  string
			Jsonrpc string
			Params  []string
		}

		type Firmwareresult struct {
			Jsonrpc string
			Result  string
			Id      string
		}

		tttt := Firmwareupdate{
			Jsonrpc: "2.0",
			Id:      "s",
			Params:  []string{"bb"},
		}
	*/

	type Firmwareupdate struct {
		Factory string
	}

	tttt := Firmwareupdate{
		Factory: "pifii",
	}

	this.Data["json"] = tttt

	this.ServeJSON()

}
