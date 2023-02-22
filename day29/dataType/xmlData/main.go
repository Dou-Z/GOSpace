package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Servers struct {
	Name    xml.Name `xml:"servers"`
	Version string   `xml:"version",attr`
	Servers []Server `xml:"server"`
}

type Server struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIp"`
}

func main() {
	data, err := ioutil.ReadFile("./config.xml")
	if err != nil {
		fmt.Printf("read config.xml failed,err:%s\n", err)
		return
	}
	var servers Servers
	err = xml.Unmarshal(data, &servers)
	if err != nil {
		fmt.Printf("Unmarshal failed ,err:%s\n", err)
		return
	}
	fmt.Printf("xml:%#v\n", servers)
}
