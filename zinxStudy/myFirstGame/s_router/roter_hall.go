package s_router

import (
	"fmt"
	"time"

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

	account := ClientsMapCon[req.GetConnection()].Account
	pUser := playerData.GetPUser(account)
	code, roomItem := playerData.CreateRoom(req.GetConnection(), pUser)
	data := &msg.SC_CreateRoom{
		Code:     code,
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
	msgTemp := &msg.CS_MatchRoom{}
	err := proto.Unmarshal(req.GetData(), msgTemp)
	if err != nil {
		fmt.Println("Position Unmarshal error ", err, " data = ", req.GetData())
		return
	}
	account := ClientsMapCon[req.GetConnection()].Account
	pUser := playerData.GetPUser(account)
	code, roomItem := playerData.MathchRoom(req.GetConnection(), pUser)
	data := &msg.SC_CreateRoom{
		Code:     code,
		RoomInfo: roomItem,
	}
	SendMsg(uint32(msg.MsgId_MSG_SC_MatchRoom), data, req.GetConnection())
	//匹配房间广播通知其他玩家
	BroadCast(roomItem.Rid, uint32(msg.MsgId_MSG_SC_MatchRoom), data, pUser.Plyer.Account)

}

type RouterReady struct {
	znet.BaseRouter
}

func (t *RouterReady) Handle(req ziface.IRequest) {
	msgTemp := &msg.SC_NHWCReady{}
	err := proto.Unmarshal(req.GetData(), msgTemp)
	if err != nil {
		fmt.Println("Position Unmarshal error ", err, " data = ", req.GetData())
		return
	}
	onReady(req.GetConnection(), msgTemp)
}

// 客户端请求准备
func onReady(req ziface.IConnection, msgTemp *msg.SC_NHWCReady) {
	account := ClientsMapCon[req].Account
	pUser := playerData.GetPUser(account)
	rid := pUser.Rid
	roomInfo, _ := playerData.GetPRoom(rid)
	if pUser.IsReady {
		// 	fmt.Println("用户已经准备")
		data := &msg.SC_NHWCReady{
			Code:     int32(enumeCode.PlayerReadyed),
			RoomInfo: roomInfo,
		}
		SendMsg(uint32(msg.MsgId_MSG_SC_NHWCReady), data, req)
	} else {
		pUser.IsReady = true
		data := &msg.SC_NHWCReady{
			Code:     int32(enumeCode.OK),
			RoomInfo: roomInfo,
		}
		BroadCast(rid, uint32(msg.MsgId_MSG_SC_NHWCReady), data, "")
		if CanStartGame(pUser.Rid) {
			StartGame(pUser.Rid)
		}
	}

}

/***
 * 判断是否可以开始游戏
 */
func CanStartGame(rid int32) bool {
	seatSum := 0
	readySum := 0
	seatMap := playerData.RoomMap[rid].ArrPlayerInfo
	for _, seat := range seatMap {
		if seat != nil {
			seatSum++
			if seat.IsReady {
				readySum++
			}
		}
	}
	if seatSum == readySum && seatSum >= 2 {
		return true
	} else {
		return false
	}
}

/***
 * 开始游戏
 */
func StartGame(rid int32) {
	pRoom, _ := playerData.GetPRoom(rid)
	pRoom.Painter = getNextSeat(rid)
	pRoom.WordIndex++
	pRoom.Word = playerData.WordsList[pRoom.WordIndex].Word
	pRoom.Hint = playerData.WordsList[pRoom.WordIndex].D
	pRoom.State = int32(msg.RoomState_Draw)
	pRoom.GameNum++
	pRoom.GameTime = time.Now().Unix()
	timer := time.NewTimer(time.Second * time.Duration(3))
	pUser := getUserBySeat(pRoom, pRoom.Painter)

	pUser.IsReady = true
	data := &msg.SC_NHWCStart{
		RoomInfo: pRoom,
	}
	BroadCast(rid, uint32(msg.MsgId_MSG_SC_NHWCStart), data, "")

	go func() {
		<-timer.C
		if pRoom != nil && pRoom.State == int32(msg.RoomState_Draw) {
			showAnswer(rid)
		}
	}()
}

/***
 * 根据座位号获取用户
 */
func getUserBySeat(pRoom *msg.RoomInfo, seat int32) *msg.GameUserItem {
	for _, v := range pRoom.ArrPlayerInfo {
		if v.Seat == seat {
			return v
		}
	}
	return nil
}

func getNextSeat(rid int32) int32 {
	pRoom, _ := playerData.GetPRoom(rid)
	currSeat := pRoom.Painter
	i := currSeat
	for {
		if i == int32(len(pRoom.ArrPlayerInfo)) {
			i = 1
		} else {
			i++
		}
		if pRoom.ArrPlayerInfo[i] != nil {
			return i
		}
	}

}

/***
 * 显示答案
 */
func showAnswer(rid int32) {

	data := &msg.SC_NHWCResult{
		Word: playerData.RoomMap[rid].Word,
	}
	BroadCast(rid, uint32(msg.MsgId_MSG_SC_NHWCStart), data, "")

	timer2 := time.NewTimer(time.Second * time.Duration(3))
	go func() {
		<-timer2.C
		if playerData.RoomMap[rid] == nil || playerData.RoomMap[rid].State != int32(msg.RoomState_Result) {
			return
		}
		if playerData.RoomMap[rid].GameNum >= 3 {
			OverGame(rid)
		} else {
			StartGame(rid)
		}
	}()
}

/***
 * 游戏结束
 */
func OverGame(rid int32) {
	data := &msg.SC_NHWCOver{
		Word: playerData.RoomMap[rid].Word,
	}
	BroadCast(rid, uint32(msg.MsgId_MSG_SC_NHWCStart), data, "")

	playerData.ResetGame(rid)
}

type RouterDrawClear struct {
	znet.BaseRouter
}

func (t *RouterDrawClear) Handle(req ziface.IRequest) {
	msgTemp := &msg.CS_NHWCDrawClear{}
	err := proto.Unmarshal(req.GetData(), msgTemp)
	if err != nil {
		fmt.Println("Position Unmarshal error ", err, " data = ", req.GetData())
		return
	}

	data := &msg.SC_NHWCDrawClear{}
	account := ClientsMapCon[req.GetConnection()].Account
	pUser := playerData.GetPUser(account)
	rid := pUser.Rid

	BroadCast(rid, uint32(msg.MsgId_MSG_SC_NHWCDrawClear), data, "")

}

type RouterDrawWidth struct {
	znet.BaseRouter
}

func (t *RouterDrawWidth) Handle(req ziface.IRequest) {
	msgTemp := &msg.CS_NHWCDrawWidth{}
	err := proto.Unmarshal(req.GetData(), msgTemp)
	if err != nil {
		fmt.Println("Position Unmarshal error ", err, " data = ", req.GetData())
		return
	}

	account := ClientsMapCon[req.GetConnection()].Account
	pUser := playerData.GetPUser(account)
	rid := pUser.Rid
	data := &msg.SC_NHWCDrawWidth{}
	data.Width = msgTemp.Width
	BroadCast(rid, uint32(msg.MsgId_MSG_SC_NHWCDrawWidth), data, "")

}

type RouterDrawColor struct {
	znet.BaseRouter
}

func (t *RouterDrawColor) Handle(req ziface.IRequest) {
	msgTemp := &msg.CS_NHWCDrawColor{}
	err := proto.Unmarshal(req.GetData(), msgTemp)
	if err != nil {
		fmt.Println("Position Unmarshal error ", err, " data = ", req.GetData())
		return
	}

	account := ClientsMapCon[req.GetConnection()].Account
	pUser := playerData.GetPUser(account)
	rid := pUser.Rid
	data := &msg.SC_NHWCDrawColor{}
	data.Color = msgTemp.Color
	BroadCast(rid, uint32(msg.MsgId_MSG_SC_NHWCDrawColor), data, "")

}

type RouterDrawPath struct {
	znet.BaseRouter
}

func (t *RouterDrawPath) Handle(req ziface.IRequest) {
	msgTemp := &msg.CS_NHWCDrawPath{}
	err := proto.Unmarshal(req.GetData(), msgTemp)
	if err != nil {
		fmt.Println("Position Unmarshal error ", err, " data = ", req.GetData())
		return
	}
	account := ClientsMapCon[req.GetConnection()].Account
	pUser := playerData.GetPUser(account)
	rid := pUser.Rid
	data := &msg.SC_NHWCDrawPath{}
	data.PointArr = msgTemp.PointArr
	BroadCast(rid, uint32(msg.MsgId_MSG_SC_NHWCStart), data, "")

}

type RouterExitRoom struct {
	znet.BaseRouter
}

func (t *RouterExitRoom) Handle(req ziface.IRequest) {
	msgTemp := &msg.CS_ExitRoom{}
	err := proto.Unmarshal(req.GetData(), msgTemp)
	if err != nil {
		fmt.Println("Position Unmarshal error ", err, " data = ", req.GetData())
		return
	}

	account := ClientsMapCon[req.GetConnection()].Account
	pUser := playerData.GetPUser(account)
	rid := pUser.Rid
	code := playerData.ExitRoom(account)
	data := &msg.SC_ExitRoom{
		Code:    code,
		Account: account,
	}
	SendMsg(uint32(msg.MsgId_MSG_SC_ExitRoom), data, req.GetConnection())
	//匹配房间广播通知其他玩家
	BroadCast(rid, uint32(msg.MsgId_MSG_SC_ExitRoom), data, account)
}
