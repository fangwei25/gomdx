# gomdx

[![Source graph](https://sourcegraph.com/github.com/fangwei25/gomdx/-/badge.svg?style=flat-square)](https://sourcegraph.com/github.com/fangwei25/gomdx?badge)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/fangwei25/gomdx)
[![Go Report Card](https://goreportcard.com/badge/github.com/fangwei25/gomdx?style=flat-square)](https://goreportcard.com/report/github.com/fangwei25/gomdx)
[![Codecov](https://img.shields.io/codecov/c/github/fangwei25/gomdx.svg?style=flat-square)](https://codecov.io/gh/fangwei25/gomdx)
[![License](https://img.shields.io/github/license/fangwei25/gomdx)](https://raw.githubusercontent.com/fangwei25/gomdx/master/LICENSE)

gomdx 是一个基于golang的多维度数据统计和查询库

它提供基于配置的数据统计功能

    在时间维度上，支持粒度上至永久，年，下至分钟的数据统计
    在统计方式上，支持次数累加，数值累加，最大值记录，最小值记录

它支持自定义数据源

    库中计划提供三种数据源：
        fake:   用于开发、演示的数据源，直接将函数调用参数打印到标准输出，不做数据统计
        memery: 内存记录，数据直接记录于进程内存中，进程停止即销毁
        redis:  redis数据源，数据记录于指定的redis中
当然，你还可以自定自己的数据源，只要满足 [datasource/idatasource.go](./datasource/idatasource.go) 中的DataSourcer接口即可

## 如何使用

可参见example/main.go

## 开发计划

+ （已完成）数据统计逻辑
+ （已完成）数据查询接口
+ （已完成）支持维度配置
+ （已完成）可配置的数据源
+ （已完成）维度配置的动态更新
+ （计划中）缓存机制

## 其它项目推荐
 
[盘古](https://github.com/pangum/pangu) 一个Golang应用程序快速开发框架
[gex](https://github.com/fangwei25/gomdx) Golang外部命令执行扩展库
