package s_router

import (
	"fmt"

	enumeCode "github.com/aceld/zinx/myFirstGame/EnumeCode"
	playerData "github.com/aceld/zinx/myFirstGame/data"
	"github.com/aceld/zinx/myFirstGame/model"
	msg "github.com/aceld/zinx/myFirstGame/pb"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"google.golang.org/protobuf/proto"
)

type (
	ClientNHWC struct {
		Account string
		Conn    *ziface.IRequest
	}
)

var ClientsMap = make(map[string]ClientNHWC)

func Bind(req ziface.IConnection, account string) {
	c, ok := ClientsMap[account]
	if !ok {
		fmt.Println("没有这个client")
	} else {
		fmt.Println(account, "号用户与", account, "号客户端绑定")
		c.Account = account
	}
}

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
		fmt.Println("Player SendMsg error !")
		return
	}
}

type RouterLogin struct {
	znet.BaseRouter
}

// RouterLogin Handle
func (t *RouterLogin) Handle(request ziface.IRequest) {
	// zlog.Ins().DebugF("Call HelloZinxRouter Handle")
	// zlog.Ins().DebugF("recv from client : msgId=%d, data=%+v, len=%d", request.GetMsgID(), string(request.GetData()), len(request.GetData()))
	msgTemp := &msg.CS_Login{}
	err := proto.Unmarshal(request.GetData(), msgTemp)
	if err != nil {
		fmt.Println("Position Unmarshal error ", err, " data = ", request.GetData())
		return
	}
	pUserData, errCodeType := model.ValidateUserData(msgTemp.Account, msgTemp.Password)
	if errCodeType != enumeCode.OK || pUserData == nil { // 登录失败
		sendLoginErr(request, errCodeType)
	} else {
		data := &msg.SC_Login{
			Code:       int32(errCodeType),
			Token:      "",
			PlayerInfo: pUserData,
		}

		u := playerData.User{
			Account:  pUserData.Account,
			Seat:     0,
			Username: pUserData.NickName,
		}
		playerData.SetPUser(&u)
		Bind(request.GetConnection(), pUserData.Account)

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
