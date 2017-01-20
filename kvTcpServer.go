package main

import "fmt"
import "time"

import "encoding/json"
import "bufio"
import "os"

import "net"

type A struct {
	Sfds string
	Sftt int
}

var t1 *time.Timer

func dingshi() {
	for {
		select {
		case <-t1.C:
			t1.Reset(5 * time.Second)
			fmt.Println("ok")
		}
	}
}

var kvvalue map[string]net.Conn = make(map[string]net.Conn)

func insertKv(key string, con net.Conn) {
	_, err := kvvalue[key]
	if !err {
		//	v1.Close()
		kvvalue[key] = con

	} else {
		kvvalue[key] = con
	}
}

func fazhan_ok(client string, mtype string) {
	var zhi = make([]byte, 1024)

	type CommandStats struct {
		Jsonrpc string
		Id      string
		Method  string
	}

	type CommandData struct {
		Etype string
		Edata CommandStats
	}

	result := CommandData{
		Etype: "1",
		Edata: CommandStats{
			Id:     "1",
			Method: "Reds",
		},
	}

	val, _ := json.Marshal(result)
	cn, err := kvvalue[client]
	if !err {

		fmt.Println(string(val))
	} else {
		cn.Write(val)
		cn.Read(zhi)
		fmt.Println("jiang_ok")
	}

}

func qidongServer() {
	c1, _ := net.ResolveTCPAddr("tcp", ":9527")
	d1, _ := net.ListenTCP("tcp", c1)
	for {
		f1, _ := d1.Accept()
		fmt.Println("ok")
		insertKv("1", f1)
	}

}

var inputReader *bufio.Reader
var input string
var err error

func main() {
	go qidongServer()

	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	for {
		input, err = inputReader.ReadString('\n') //func (b *Reader) ReadString(delim byte) (line string, err error) ,‘S’ 这个例子里使用S表示结束符，也可以用其它，如'\n'
		if err == nil {
			fmt.Println("The input was:" + input)
			go fazhan_ok("1", input)
		}
	}

}
