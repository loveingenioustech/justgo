package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	demoJsonUnmarshal()

	demoJsonMarshal()

	demoJsonMarshal2()
}

func demoJsonUnmarshal() {
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
}

func demoJsonMarshal() {
	var s Serverslice
	s.Servers = append(s.Servers, Server{ServerName: "Hangzhou_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Hongkong_VPN", ServerIP: "127.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}

func demoJsonMarshal2(){
	type Server2 struct {
		// ID 不会导出到JSON中
		ID int `json:"-"`

		// ServerName2 的值会进行二次JSON编码
		ServerName  string `json:"serverName"`
		ServerName2 string `json:"serverName2,string"`

		// 如果 ServerIP 为空，则不输出到JSON串中
		ServerIP   string `json:"serverIP,omitempty"`
	}

	s := Server2 {
		ID:         3,
		ServerName:  `Go "1.0" `,
		ServerName2: `Go "1.0" `,
		ServerIP:   ``,
	}
	b, _ := json.Marshal(s)
	fmt.Println(string(b))
}
