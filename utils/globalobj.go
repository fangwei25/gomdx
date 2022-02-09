// Package utils 提供工具类函数
// 包括:
//		全局配置
//		配置文件加载
//
// 当前文件描述:
// @Title  global.go
// @Description  相关配置文件定义及加载方式
// @Author  SoloF - Thu Mar 11 10:32:29 CST 2019
package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/*
	存储一切有关Zinx框架的全局参数，供其他模块使用
	一些参数也可以通过 用户根据 zinx.json来配置
*/
type GlobalObj struct {

	/*
		config file path
	*/
	ConfFilePath string
	RedisAddr    string //redis连接地址 "localhost:6379"
	RedisPWD     string //redis密码 默认为空
	RedisDBIdx   uint32 //redis数据库编号 默认为0
}

var GlobalObject *GlobalObj

//PathExists 判断一个文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//Reload 读取用户的配置文件
func (g *GlobalObj) Reload() {

	if confFileExists, _ := PathExists(g.ConfFilePath); !confFileExists {
		fmt.Println("Config File ", g.ConfFilePath, " is not exist!!")
		return
	}

	data, err := ioutil.ReadFile(g.ConfFilePath)
	if err != nil {
		panic(err)
	}
	fmt.Println("Config File:", g.ConfFilePath)

	//将json数据解析到struct中
	err = json.Unmarshal(data, g)
	if err != nil {
		panic(err)
	}

	//Logger 设置
	// if g.LogFile != "" {
	// 	zglog.SetLogFile(g.LogDir, g.LogFile)
	// }
	// if g.LogDebugClose {
	// 	zglog.CloseDebug()
	// }
}

/*
	提供init方法，默认加载
*/
func init() {
	pwd, err := os.Getwd()
	if err != nil {
		pwd = "."
	}
	//初始化GlobalObject变量，设置一些默认值
	GlobalObject = &GlobalObj{
		ConfFilePath: pwd + "/conf/serverConfig.json",
		RedisAddr:    "localhost:6379",
		RedisPWD:     "123456",
		RedisDBIdx:   0,
	}

	//NOTE: 从配置文件中加载一些用户配置的参数
	GlobalObject.Reload()
}
