package ini

import (
	"fmt"
	"io/ioutil"
	"testing"
)

type config struct {
	ServerConf ServerConfig `ini:"Server"`
	MysqlConf  MysqlConfig  `ini:"Mysql"`
}
type ServerConfig struct {
	Ip   string `ini:"ip"`
	Port int    `ini:"port"`
}
type MysqlConfig struct {
	Username string `ini:"username"`
	Passwd   string `ini:"passwd"`
	Database string `ini:"database"`
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
}

func TestIni(t *testing.T) {
	fmt.Println("hello")
	data, err := ioutil.ReadFile("./config.ini")
	_ = data
	if err != nil {
		t.Error("read file dailed")
	}
	var conf config
	err = UnMarshal(data, &conf)
	if err != nil {
		t.Errorf("unmarshal failed err:%v", err)
		return
	}
	t.Logf("unmarshal success ,conf:%#v", conf)
	confData, err := Marshal(conf)
	if err != nil {
		t.Errorf("marshal failed err:%v", err)
		return
	}
	t.Logf("marshal success ,conf:%s", string(confData))

	MarshalFile("D:\\Desktop\\文件夹\\Go\\test.conf", conf)

}

func TestIniConfigFile(t *testing.T) {
	filename := "D:/Desktop/文件夹/Go/config1.conf"
	var conf config
	err := MarshalFile(filename, conf)
	if err != nil {
		t.Errorf("unmarshal failed err:%v", err)
		return
	}
	var conf2 config
	err = UnMarshalFile(filename, &conf)
	if err != nil {
		t.Errorf("unmarshal failed,err:%v", err)
	}
	t.Logf("unmarshal succ,conf:%#v", conf2)
}
