package main

import (
	msg "github.com/aceld/zinx/myFirstGame/pb"
	"github.com/aceld/zinx/myFirstGame/s_router"
	"github.com/aceld/zinx/zconf"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

func main() {
	// Set up as WebSocket before starting. (在启动之前设置为 websocket)
	zconf.GlobalObject.Mode = ""
	zconf.GlobalObject.LogFile = ""

	s := znet.NewServer()
	initAddRouter(s)
	s.Serve()
}

/** 添加所有客户端的请求接收 */
func initAddRouter(s ziface.IServer) {
	// 心跳
	s.AddRouter(uint32(msg.MsgId_MSG_CS_Ping), &s_router.RouterPing{})
	//登录
	s.AddRouter(uint32(msg.MsgId_MSG_CS_Login), &s_router.RouterLogin{})
	//注册
	s.AddRouter(uint32(msg.MsgId_MSG_CS_Register), &s_router.RouterRegister{})

}
