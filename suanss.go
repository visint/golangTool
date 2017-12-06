package main

import "fmt"
import "time"
import "encoding/json"
import "os"
import "os/exec"
import "bufio"

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
			t1.Reset(5 * time.Minute)
			fmt.Println("ok")
		}
	}
}

func fazhan(cn net.Conn) {
	var zhi = make([]byte, 1024)
	for {
		cn.Read(zhi)
		fmt.Println(zhi)
	}
}

func main() {

	//	http.ListenAndServe()
	//time.After()

	//	a1, _ := net.ResolveTCPAddr("tcp", "192.168.11.2:9090")
	f22, err := net.Dial("udp", "222.45.234.6:9998")
	if err != nil {
		fmt.Println("jiang")
	}

	m1 := bufio.NewReader(os.Stdin)
	for {
		n1, _, _ := m1.ReadLine()
		fmt.Println(n1)
		for i := 0; i < 10; i++ {
			f22.Write(n1)
		}
	}

	a := A{Sfds: "jiang", Sftt: 2}
	b, _ := json.Marshal(a)
	exec.Command("/bin/sh", "ps")
	fmt.Println(b)
}
