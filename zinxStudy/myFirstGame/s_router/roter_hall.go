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
	roomInfo, ok := playerData.GetPRoom(msgTemp.Id)
	if ok {
		freeSeat := playerData.GetFreeSeat(roomInfo)
		if freeSeat == 0 { //房间已满
			data := &msg.SC_Register{
				Code: int32(enumeCode.LoginPassWord),
			}
			SendMsg(uint32(msg.MsgId_MSG_SC_JoinRoom), data, req.GetConnection())
		} else {
			pUser := playerData.GetPUser(string(msgTemp.Id))
			playerData.EnterRoom(pUser, roomInfo, freeSeat)
			// 通知客户端进入房间
			m := WsMessage{
				Name:  "enter",
				Value: roomInfo,
			}
			Send(cid, m)

			//通知其他客户端有人加入房间
			pUser := data.GetPUser(uid)
			BroadCast(rid, WsMessage{
				Name:  "message",
				Value: "【系统】" + pUser.Username + "进入房间",
			}, cid)
			BroadCast(rid, WsMessage{
				Name:  "room",
				Value: roomInfo,
			}, cid)

			// 更新大厅房间信息
			roomList := data.GetRoomIdList()
			BroadCast(0, WsMessage{
				Name: "hall",
				Value: struct {
					RoomList []interface{} `json:"roomlist"`
				}{
					RoomList: roomList,
				},
			}, cid)
		}
	} else {
		// data := &msg.SC_Register{
		// 	Code: int32(errCodeType),
		// }
		// SendMsg(uint32(msg.MsgId_MSG_SC_Register), data, req.GetConnection())
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
