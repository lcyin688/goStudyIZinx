package s_router

import (
	"fmt"

	msg "github.com/aceld/zinx/myFirstGame/pb"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"google.golang.org/protobuf/proto"
)

// ping  自定义路由
type RouterPing struct {
	znet.BaseRouter
}

// Ping Handle
func (t *RouterPing) Handle(request ziface.IRequest) {
	// zlog.Ins().DebugF("Call PingRouter Handle")
	// zlog.Ins().DebugF("recv from client : msgId=%d, data=%+v, len=%d", request.GetMsgID(), string(request.GetData()), len(request.GetData()))

	msg := &msg.CS_Ping{}
	err := proto.Unmarshal(request.GetData(), msg)
	if err != nil {
		fmt.Println("Position Unmarshal error ", err, " data = ", request.GetData())
		return
	}
	// fmt.Printf("recv from client : msgId=%+v, data=%+v\n", request.GetMsgID(), msg)
	sendPong(request)

}

func sendPong(request ziface.IRequest) {
	data := &msg.SC_Pong{
		Timestamp: 1,
	}
	SendMsg(uint32(msg.MsgId_MSG_SC_Pong), data, request.GetConnection())
}
