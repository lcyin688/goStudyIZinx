package s_router

import (
	"fmt"

	"github.com/aceld/zinx/myFirstGame/core"
	enumeCode "github.com/aceld/zinx/myFirstGame/enumCode"
	"github.com/aceld/zinx/myFirstGame/model"
	msg "github.com/aceld/zinx/myFirstGame/pb"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"google.golang.org/protobuf/proto"
)

func SendMsg(msgID uint32, data proto.Message, req ziface.IConnection) {
	if req == nil {
		fmt.Println("connection in player is nil")
		return
	}
	// 将proto Message结构体序列化
	msg, err := proto.Marshal(data)
	if err != nil {
		fmt.Println("marshal msg err: ", err)
		return
	}
	// 调用Zinx框架的SendMsg发包
	if err := req.SendMsg(msgID, msg); err != nil {
		fmt.Println("Player SendMsg error !", err)
		return
	}
}

func SendBuffMsg(msgID uint32, data proto.Message, req ziface.IConnection) {
	if req == nil {
		fmt.Println("connection in player is nil")
		return
	}
	// 将proto Message结构体序列化
	msg, err := proto.Marshal(data)
	if err != nil {
		fmt.Println("marshal msg err: ", err)
		return
	}
	// 调用Zinx框架的SendMsg发包
	if err := req.SendBuffMsg(msgID, msg); err != nil {
		fmt.Println("Player SendBuffMsg error ! ", err)
		return
	}
}

/**
 * 广播消息
 */
func BroadCast(roomId int32, msgId uint32, msg proto.Message, exclude string) {
	fmt.Println("广播消息 roomId msgId ", roomId, msgId, msg)
	if roomId == 0 {
		for _, user := range core.WorldMgrObj.Players {
			if user.GameUserItem.IsOnline {
				SendBuffMsg(msgId, msg, user.Conn)
			}
		}
	} else { //只给某个房间内的人
		pRoom, _ := core.GetPRoom(roomId)
		for _, userItem := range pRoom.ArrPlayerInfo {
			if userItem != nil && userItem.Plyer.Account != exclude {
				player := core.WorldMgrObj.PlayersAcount[userItem.Plyer.Account]
				fmt.Println("广播消息 player", player.GameUserItem)
				if player.GameUserItem.IsOnline {
					SendBuffMsg(msgId, msg, player.Conn)
				}
			}
		}
	}
}

type RouterLogin struct {
	znet.BaseRouter
}

// RouterLogin Handle
func (t *RouterLogin) Handle(request ziface.IRequest) {
	// zlog.Ins().DebugF("Call HelloZinxRouter Handle")
	// zlog.Ins().DebugF("recv from client : msgId=%d, data=%+v, len=%d", request.GetMsgID(), string(request.GetData()), len(request.GetData()))
	// (1. 将客户端传来的proto协议解码)
	msgTemp := &msg.CS_Login{}
	err := proto.Unmarshal(request.GetData(), msgTemp)
	if err != nil {
		fmt.Println("Position Unmarshal error ", err, " data = ", request.GetData())
		return
	}
	playerInfo, errCodeType := model.ValidateUserData(msgTemp.Account, msgTemp.Password)
	if errCodeType != enumeCode.OK || playerInfo == nil { // 登录失败
		sendLoginErr(request, errCodeType)
	} else {
		// 登录成功
		conn := request.GetConnection()
		// 创建一个玩家
		player := core.NewPlayer(conn, playerInfo)
		// 将当前新上线玩家添加到worldManager中 如果玩家之前上过线线，则删除
		core.WorldMgrObj.AddPlayer(player)
		// 将该连接绑定属性PID
		conn.SetProperty("pID", player.PID)
		data := &msg.SC_Login{
			Code:       int32(errCodeType),
			Token:      "",
			PlayerInfo: player.GameUserItem.Plyer,
		}
		SendMsg(uint32(msg.MsgId_MSG_SC_Login), data, request.GetConnection())
	}
}

func sendLoginErr(req ziface.IRequest, code enumeCode.ErrCodeType) {
	data := &msg.SC_Login{
		Code:       int32(code),
		Token:      "",
		PlayerInfo: nil,
	}
	SendMsg(uint32(msg.MsgId_MSG_SC_Login), data, req.GetConnection())
}

type RouterRegister struct {
	znet.BaseRouter
}

func (t *RouterRegister) Handle(req ziface.IRequest) {
	msgTemp := &msg.CS_Register{}
	err := proto.Unmarshal(req.GetData(), msgTemp)
	if err != nil {
		fmt.Println("Position Unmarshal error ", err, " data = ", req.GetData())
		return
	}
	errCodeType := model.RegisteUserData(msgTemp.Account, msgTemp.Password, msgTemp.HeadId)
	data := &msg.SC_Register{
		Code: int32(errCodeType),
	}
	SendMsg(uint32(msg.MsgId_MSG_SC_Register), data, req.GetConnection())
}
