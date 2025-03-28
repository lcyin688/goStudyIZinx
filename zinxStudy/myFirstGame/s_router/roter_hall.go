package s_router

import (
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

type RouterHall struct {
	znet.BaseRouter
}

func (t *RouterHall) Handle(req ziface.IRequest) {
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
