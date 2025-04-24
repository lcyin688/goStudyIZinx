package main

import (
	"fmt"

	msg "github.com/aceld/zinx/myFirstGame/pb"
	"github.com/aceld/zinx/myFirstGame/s_router"
	"github.com/aceld/zinx/zconf"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

var s ziface.IServer

func main() {
	// Set up as WebSocket before starting. (在启动之前设置为 websocket)
	zconf.GlobalObject.Mode = ""
	zconf.GlobalObject.LogFile = ""

	s = znet.NewServer()
	initAddRouter(s)

	s.SetOnConnStart(DoConnectionBegin)

	// callOnConnStop
	// SetOnConnStop
	s.SetOnConnStop(DoConnectionLost)

	s.Serve()

}

// 关闭服务
func stopServer(s ziface.IServer) {
	s.Stop()
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

	//你画我猜退出房间
	s.AddRouter(uint32(msg.MsgId_MSG_CS_ExitRoom), &s_router.RouterExitRoom{})

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
func DoConnectionBegin(conn ziface.IConnection) {
	fmt.Println("DoConnectionBegin is Called ... ")
	err := conn.SendMsg(2, []byte("DoConnection BEGIN..."))
	if err != nil {
		fmt.Println(err)
	}
}

func DoConnectionLost(conn ziface.IConnection) {
	//断开连接的时候打印下是谁掉线了
	fmt.Println("DoConnectionBegin is Called ... 001 ")

	// Get the "pID" property of the current connection
	// 获取当前连接的PID属性
	pID, _ := conn.GetProperty("pID")
	var playerID int32
	if pID != nil {
		playerID = pID.(int32)
	}
	fmt.Println("DoConnectionBegin is Called ... playerID ", playerID)
	// Get the corresponding player object based on the player ID
	// 根据pID获取对应的玩家对象
	// player := core.WorldMgrObj.GetPlayerByPID(playerID)

	// Trigger the player's disconnection business logic
	// 触发玩家下线业务
	// if player != nil {
	// 	player.LostConnection()
	// }

	if s_router.ClientsMapCon != nil {
		if client, ok := s_router.ClientsMapCon[conn]; ok {
			account := client.Account
			fmt.Println("DoConnectionLost is Called ... ", account)
			// delete(s_router.ClientsMapCon, conn)
		}
	}

}
