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

	//大厅信息
	s.AddRouter(uint32(msg.MsgId_MSG_CS_HallInfo), &s_router.RouterHall{})

	//创建房间信息
	s.AddRouter(uint32(msg.MsgId_MSG_CS_CreateRoom), &s_router.RouterCreateRoom{})

	//进入房间
	s.AddRouter(uint32(msg.MsgId_MSG_CS_JoinRoom), &s_router.RouterJoinRoom{})

	//匹配房间信息
	s.AddRouter(uint32(msg.MsgId_MSG_CS_MatchRoom), &s_router.RouterMatchRoom{})

	//你画我猜准备
	s.AddRouter(uint32(msg.MsgId_MSG_CS_NHWCReady), &s_router.RouterReady{})

	//你画我猜清理
	s.AddRouter(uint32(msg.MsgId_MSG_CS_NHWCDrawClear), &s_router.RouterDrawClear{})

	//你画我猜Width
	s.AddRouter(uint32(msg.MsgId_MSG_CS_NHWCDrawWidth), &s_router.RouterDrawWidth{})

	//你画我猜Color
	s.AddRouter(uint32(msg.MsgId_MSG_CS_NHWCDrawColor), &s_router.RouterDrawColor{})

	//你画我猜Path
	s.AddRouter(uint32(msg.MsgId_MSG_CS_NHWCDrawPath), &s_router.RouterDrawPath{})

}
