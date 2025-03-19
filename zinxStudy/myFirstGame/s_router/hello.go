package s_router

import (
	"fmt"

	msg "github.com/aceld/zinx/myFirstGame/pb"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/zlog"
	"github.com/aceld/zinx/znet"
	"google.golang.org/protobuf/proto"
)

type HelloZinxRouter struct {
	znet.BaseRouter
}

// HelloZinxRouter Handle
func (this *HelloZinxRouter) Handle(request ziface.IRequest) {
	zlog.Ins().DebugF("Call HelloZinxRouter Handle")
	zlog.Ins().DebugF("recv from client : msgId=%d, data=%+v, len=%d", request.GetMsgID(), string(request.GetData()), len(request.GetData()))

	msg := &msg.CS_Ping{}
	err := proto.Unmarshal(request.GetData(), msg)
	if err != nil {
		fmt.Println("Position Unmarshal error ", err, " data = ", request.GetData())
		return
	}
	fmt.Printf("recv from client : msgId=%+v, data=%+v\n", request.GetMsgID(), msg)
	sendPong(request)

}

func sendPong(request ziface.IRequest) {
	data := &msg.SC_Pong{
		Timestamp: 1,
	}
	SendMsg(uint32(msg.MsgId_MSG_SC_Pong), data, request)
}

func SendMsg(msgID uint32, data proto.Message, req ziface.IRequest) {

	if req.GetConnection() == nil {
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
	if err := req.GetConnection().SendMsg(msgID, msg); err != nil {
		fmt.Println("Player SendMsg error !")
		return
	}
}
