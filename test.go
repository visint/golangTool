package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
	"os"
	"strings"
)

func main() {
	arg_num := len(os.Args)
	fmt.Printf("the num of input is %d\n", arg_num)
	if arg_num < 2 {
		return
	}
	fmt.Println(os.Args[1])
	SSH("root", "root", strings.TrimSpace(os.Args[1])+":22")
}

func SSH(user, password, ip_port string) {
	PassWd := []ssh.AuthMethod{ssh.Password(password)}
	Conf := ssh.ClientConfig{User: user, Auth: PassWd}
	Client, _ := ssh.Dial("tcp", ip_port, &Conf)
	defer Client.Close()
	if session, err := Client.NewSession(); err == nil {
		defer session.Close()
		session.Stdout = os.Stdout
		session.Stderr = os.Stderr
		session.Run("killall -9 watchdogd")
	}

	a := bufio.NewReader(os.Stdin)
	for {
		b, _, z := a.ReadLine()
		if z != nil {
			return
		}
		command := string(b)
		if session, err := Client.NewSession(); err == nil {
			defer session.Close()
			session.Stdout = os.Stdout
			session.Stderr = os.Stderr
			session.Run(command)
		}
	}
}
