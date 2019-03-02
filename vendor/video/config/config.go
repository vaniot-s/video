package config

import (
	"encoding/json"
	"os"
)

type Configuration struct{
	LBAddr string `json:"lb_addr"` //负载均衡地址
	OssAddr string `json:"oss_addr"` //oss存储
}

var configuration *Configuration

func init(){
	file,_:=os.Open("./conf.json") //读取配置文件
	defer file.Close() //关闭
	decoder:=json.NewDecoder(file)
	configuration=&Configuration{}
	//解析文件
	err:=decoder.Decode(configuration)
	if err!=nil{
		panic(err)
	}
}

func GetLbAddr()string {
	return configuration.LBAddr
}
func GetOssAddr()string {
	return configuration.OssAddr
}