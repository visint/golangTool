package main

import (
	//	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"

	//	"bytes"
)

func main() {
	//生成client 参数为默认
	s1 := make([]byte, 3)

	client := &http.Client{}
	ip := flag.String("ip", "0", "Ip Addr")
	slot := flag.String("slot", "0", "槽位号")
	infosetId := flag.String("infoid", "0", "信息集id")
	id33 := flag.String("aid", "0", "信息项id")
	index := flag.String("index", "0", "信息集子项")

	xiangvalue1 := flag.String("value", "0", "信息项值")

	flag.Parse() //解析输入的参数

	xiangid, _ := strconv.Atoi(*id33)

	xiangvalue, _ := strconv.Atoi(*xiangvalue1)

	fmt.Println("jiangyibo", *id33)

	s1[0] = (byte)(xiangid) & 0xff
	s1[2] = (byte)(xiangvalue) & 0xff
	s1[1] = (byte)((xiangvalue)>>8) & 0xff

	//	s2 := fmt.Sprintln("%x%x%x%x%x%x", s1[0], s1[1], s1[2], s1[3], s1[4], s1[5])

	s2 := fmt.Sprintf("%x", s1)

	//生成要访问的url
	//url := "http://192.168.3.226/goform/setinfovalue"

	if *ip == "0" {
		fmt.Println("请输入ip")
		return
	}
	if *slot == "0" {
		fmt.Println("请输入slot")
		return
	}

	u, _ := url.Parse("http://" + *ip + "/goform/setinfovalue")
	q := u.Query()
	q.Set("slot", *slot)
	q.Set("infosetId", *infosetId)
	q.Set("index", *index)
	q.Set("infoset", s2)
	u.RawQuery = q.Encode()

	/*
		    var clusterinfo = url.Values{}
			clusterinfo.Add("slot", "8")
			clusterinfo.Add("infosetId", "2")
			clusterinfo.Add("index", "0")
			clusterinfo.Add("infoset", "1")

			data := clusterinfo.Encode()
	*/
	//提交请求
	fmt.Println(u.String())
	reqest, err := http.NewRequest("GET", u.String(), nil)

	if err != nil {

		panic(err)
	}

	//	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	reqest.Header.Set("Authorization", "Basic YWRtaW46YWRtaW4=")

	//处理返回结果
	response, err := client.Do(reqest)

	if err != nil {

		fmt.Println("网络不通")
		return
	}

	//将结果定位到标准输出 也可以直接打印出来 或者定位到其他地方进行相应的处理
	stdout := os.Stdout
	_, err = io.Copy(stdout, response.Body)

	if err != nil {

		fmt.Println("网络不通")
		return
	}

	//返回的状态码
	status := response.StatusCode

	fmt.Println(status)

}

/*
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	v := url.Values{}
	v.Set("huifu", "hello")
	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{}
	req, _ := http.NewRequest("get", "http://192.168.3.226/goform/setinfovalue", body)
	req.Header.Set("Authorization", "YWRtaW46YWRtaW4=")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") //这个一定要加，不加form的值post不过去，被坑了两小时

	fmt.Printf("%+v\n", req) //看下发送的结构

	resp, err := client.Do(req) //发送
	defer resp.Body.Close()     //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data), err)
}
*/
