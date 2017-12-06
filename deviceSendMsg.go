package main

import (
	"fmt"
	"github.com/antlinker/go-mqtt/client"
)

func main() {
	//启动一个http 服务器
	client, err := client.CreateClient(client.MqttOption{
		Addr: "tcp://113.106.98.90:1883",
		//断开连接１秒后自动连接，０不自动重连
		ReconnTimeInterval: 1,
	})
	if err != nil {
		//配置文件解析失败
		panic("配置文件解析失败")
	}
	listener := &MqttRecvPubListener{}
	//注册订阅事件
	client.AddRecvPubListener(listener)
	//建立连接
	err := client.Connect()
	if err != nil {
		//连接失败，不会进入自动重连状态
		panic(fmt.Errorf("连接失败:%v", err))
	}
	mq, err = client.Subscribe("test", QoS2)
	if err != nil {
		//订阅失败
		panic(fmt.Errorf("订阅失败:%v", err))
	}
	//等待订阅成功
	mq.Wait()
	if mq.Err() != nil {
		//订阅失败
		panic(fmt.Errorf("订阅失败:%v", mqt.Err()))
	}
	mq, err := client.Publish("sensor", QoS1, false, []byte("on"))
	if err != nil {
		//发布消息失败
		panic(fmt.Errorf("发布消息失败:%v", err))
	}

	//等待发布消息完成
	mq.Wait()
	if mq.Err() != nil {
		//发布消息失败
		panic(fmt.Errorf("发布消息失败:%v", mqt.Err()))
	}

	fmt.Println("服务器启动成功")
}
