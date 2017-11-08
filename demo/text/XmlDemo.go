package main

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"fmt"
)

type Recurlyservers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

func main() {
	demoUnmarshal()

	demoMarshal()
}

func demoUnmarshal(){
	file, err := os.Open("text/servers.xml") // For read access.
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(v)
}

func demoMarshal()  {
	v := &Recurlyservers{Version: "1"}
	v.Svs = append(v.Svs, server{ServerName:"Shanghai_VPN", ServerIP:"127.0.0.1"})
	v.Svs = append(v.Svs, server{ServerName:"Beijing_VPN", ServerIP:"127.0.0.2"})
	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	// xml.MarshalIndent或者xml.Marshal输出的信息都是不带XML头
	// 为了生成正确的xml文件，我们使用了xml包预定义的Header变量
	os.Stdout.Write([]byte(xml.Header))

	os.Stdout.Write(output)
}