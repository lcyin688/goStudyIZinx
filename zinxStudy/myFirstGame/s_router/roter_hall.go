package s_router

import (
	"fmt"

	enumeCode "github.com/aceld/zinx/myFirstGame/EnumeCode"
	msg "github.com/aceld/zinx/myFirstGame/pb"
	"github.com/aceld/zinx/myFirstGame/playerData"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"google.golang.org/protobuf/proto"
)

type RouterHall struct {
	znet.BaseRouter
}

func (t *RouterHall) Handle(req ziface.IRequest) {

	roomList := playerData.GetRoomIdList()
	data := &msg.SC_HallInfo{
		RoomArr: roomList,
	}
	SendMsg(uint32(msg.MsgId_MSG_SC_HallInfo), data, req.GetConnection())
}

type RouterCreateRoom struct {
	znet.BaseRouter
}

func (t *RouterCreateRoom) Handle(req ziface.IRequest) {
	msgTemp := &msg.CS_CreateRoom{}
	err := proto.Unmarshal(req.GetData(), msgTemp)
	if err != nil {
		fmt.Println("Position Unmarshal error ", err, " data = ", req.GetData())
		return
	}

	roomItem := playerData.CreateRoom()
	data := &msg.SC_CreateRoom{
		Code:     int32(1),
		RoomInfo: roomItem,
	}
	SendMsg(uint32(msg.MsgId_MSG_SC_CreateRoom), data, req.GetConnection())
}

type RouterJoinRoom struct {
	znet.BaseRouter
}

func (t *RouterJoinRoom) Handle(req ziface.IRequest) {
	msgTemp := &msg.CS_JoinRoom{}
	err := proto.Unmarshal(req.GetData(), msgTemp)
	if err != nil {
		fmt.Println("Position Unmarshal error ", err, " data = ", req.GetData())
		return
	}
	onEnter(req, msgTemp)
}

// 进入房间
func onEnter(req ziface.IRequest, msgTemp *msg.CS_JoinRoom) {
	rid := msgTemp.RoomId
	roomInfo, ok := playerData.GetPRoom(rid)
	if ok {
		freeSeat := playerData.GetFreeSeat(roomInfo)
		if freeSeat == 0 { //房间已满
			data := &msg.SC_JoinRoom{
				Code: int32(enumeCode.LoginPassWord),
			}
			SendMsg(uint32(msg.MsgId_MSG_SC_JoinRoom), data, req.GetConnection())
		} else {
			pUser := playerData.GetPUser(string(rid))
			playerData.EnterRoom(pUser, roomInfo, int32(freeSeat))
			// 通知客户端进入房间
			data := &msg.SC_JoinRoom{
				Code:     int32(enumeCode.OK),
				RoomInfo: roomInfo,
			}
			SendMsg(uint32(msg.MsgId_MSG_SC_JoinRoom), data, req.GetConnection())

			//通知其他客户端有人加入房间

			roomList := playerData.GetRoomIdList()
			dataHall := &msg.SC_HallInfo{
				RoomArr: roomList,
			}
			// 更新大厅房间信息 推送给所有玩家
			BroadCast(0, uint32(msg.MsgId_MSG_SC_HallInfo), dataHall, "")
		}
	} else {
		data := &msg.SC_JoinRoom{
			Code: int32(enumeCode.NoRoom),
		}
		SendMsg(uint32(msg.MsgId_MSG_SC_JoinRoom), data, req.GetConnection())
	}

}

type RouterMatchRoom struct {
	znet.BaseRouter
}

func (t *RouterMatchRoom) Handle(req ziface.IRequest) {
	// msgTemp := &msg.CS_Register{}
	// err := proto.Unmarshal(req.GetData(), msgTemp)
	// if err != nil {
	// 	fmt.Println("Position Unmarshal error ", err, " data = ", req.GetData())
	// 	return
	// }
	// errCodeType := model.RegisteUserData(msgTemp.Account, msgTemp.Password, msgTemp.HeadId)
	// data := &msg.SC_Register{
	// 	Code: int32(errCodeType),
	// }
	// SendMsg(uint32(msg.MsgId_MSG_SC_Register), data, req.GetConnection())
}

type RouterReady struct {
	znet.BaseRouter
}

func (t *RouterReady) Handle(req ziface.IRequest) {
	msgTemp := &msg.SC_ReadyNHWC{}
	err := proto.Unmarshal(req.GetData(), msgTemp)
	if err != nil {
		fmt.Println("Position Unmarshal error ", err, " data = ", req.GetData())
		return
	}
	onReady(req.GetConnection(), msgTemp)
}

// 客户端请求准备
func onReady(req ziface.IConnection, msgTemp *msg.SC_ReadyNHWC) {
	account := ClientsMapCon[req].Account

	pUser := playerData.GetPUser(account)
	if pUser.IsReady {
		// 	fmt.Println("用户已经准备")
		// 	Send(cid, WsMessage{
		// 		Name: "error",
		// 		Value: struct {
		// 			Code    string `json:"code"`
		// 			Message string `json:"message"`
		// 		}{
		// 			Code:    "2",
		// 			Message: "你已经准备了",
		// 		},
		// 	})
	} else {
		// 	pUser.IsReady = true
		// 	m := WsMessage{
		// 		Name:  "ready",
		// 		Value: map[string]int{"seat": pUser.Seat},
		// 	}
		// 	BroadCast(pUser.Rid, m, 0)
		// 	if CanStartGame(pUser.Rid) {
		// 		StartGame(pUser.Rid)
		// 	}
	}

	// roomInfo, ok := playerData.GetPRoom(rid)
	// if ok {
	// 	freeSeat := playerData.GetFreeSeat(roomInfo)
	// 	if freeSeat == 0 { //房间已满
	// 		data := &msg.SC_JoinRoom{
	// 			Code: int32(enumeCode.LoginPassWord),
	// 		}
	// 		SendMsg(uint32(msg.MsgId_MSG_SC_JoinRoom), data, req.GetConnection())
	// 	} else {
	// 		pUser := playerData.GetPUser(string(rid))
	// 		playerData.EnterRoom(pUser, roomInfo, int32(freeSeat))
	// 		// 通知客户端进入房间
	// 		data := &msg.SC_JoinRoom{
	// 			Code:     int32(enumeCode.OK),
	// 			RoomInfo: roomInfo,
	// 		}
	// 		SendMsg(uint32(msg.MsgId_MSG_SC_JoinRoom), data, req.GetConnection())

	// 		//通知其他客户端有人加入房间

	// 		roomList := playerData.GetRoomIdList()
	// 		dataHall := &msg.SC_HallInfo{
	// 			RoomArr: roomList,
	// 		}
	// 		// 更新大厅房间信息 推送给所有玩家
	// 		BroadCast(0, uint32(msg.MsgId_MSG_SC_HallInfo), dataHall, "")
	// 	}
	// } else {
	// 	data := &msg.SC_JoinRoom{
	// 		Code: int32(enumeCode.NoRoom),
	// 	}
	// 	SendMsg(uint32(msg.MsgId_MSG_SC_JoinRoom), data, req.GetConnection())
	// }

}
