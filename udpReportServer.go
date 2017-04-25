package main

import (
	"fmt"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: %s", err.Error())
		os.Exit(1)
	}
}

func recvUDPMsg(conn *net.UDPConn) {
	var buf []byte = make([]byte, 1500)

	var sendbuf []byte = make([]byte, 128)

	n, raddr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}

	fmt.Println("msg is ", string(buf[0:n]))
	sendbuf[16] = 8
	sendbuf[24] = 3
	//WriteToUDP
	//func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error)
	_, err = conn.WriteToUDP(sendbuf, raddr)
	checkError(err)
}

func main() {
	udp_addr, err := net.ResolveUDPAddr("udp", ":9998")
	checkError(err)

	conn, err := net.ListenUDP("udp", udp_addr)
	defer conn.Close()
	checkError(err)

	//go recvUDPMsg(conn)
	recvUDPMsg(conn)
}
