package s_router

import (
	"fmt"

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
	// errCodeType := model.RegisteUserData(msgTemp.Account, msgTemp.Password, msgTemp.HeadId)
	// data := &msg.SC_Register{
	// 	Code: int32(errCodeType),
	// }
	// SendMsg(uint32(msg.MsgId_MSG_SC_Register), data, req.GetConnection())
}

type RouterJoinRoom struct {
	znet.BaseRouter
}

func (t *RouterJoinRoom) Handle(req ziface.IRequest) {
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
