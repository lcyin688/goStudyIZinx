package s_router

import (
	"fmt"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/zlog"
	"github.com/aceld/zinx/znet"
	"google.golang.org/protobuf/proto"
)

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

type RouterLogin struct {
	znet.BaseRouter
}

// RouterLogin Handle
func (t *RouterLogin) Handle(request ziface.IRequest) {
	zlog.Ins().DebugF("Call HelloZinxRouter Handle")
	zlog.Ins().DebugF("recv from client : msgId=%d, data=%+v, len=%d", request.GetMsgID(), string(request.GetData()), len(request.GetData()))

	// msgId := request.GetMsgID()
	// switch msgId {
	// case uint32(msg.MsgId_MSG_CS_Ping):

	// case uint32(msg.MsgId_MSG_CS_Login):

	// default:
	// 	fmt.Println("msgId not found")
	// }

}

type RouterRegister struct {
	znet.BaseRouter
}

// RouterLogin Handle
func (t *RouterRegister) Handle(request ziface.IRequest) {
	zlog.Ins().DebugF("Call RouterRegister Handle")
	zlog.Ins().DebugF("recv from client : msgId=%d, data=%+v, len=%d", request.GetMsgID(), string(request.GetData()), len(request.GetData()))

}
