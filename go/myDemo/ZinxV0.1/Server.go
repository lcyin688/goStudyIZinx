package main

import "github.com/aceld/zinx/zinx/znet"

/************************************************************************************
	> File Name: main.go
	> Author: 李春印 @China
	> Mail: 184905792@qq.com
	> Created Time: 2018年09月03日 星期三 15时01分03秒
 **********************************************************************************/

func main() {
	//1. 创建一个Server 对象句柄,使用Zinx的api
	s := znet.NewServer("[zinxV0.1]")
	//2 启动 server
	s.Serve()

}
