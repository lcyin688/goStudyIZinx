package playerData

import (
	"fmt"
	"time"

	enumeCode "github.com/aceld/zinx/myFirstGame/EnumeCode"
	"github.com/aceld/zinx/myFirstGame/config"
	msg "github.com/aceld/zinx/myFirstGame/pb"
	"github.com/aceld/zinx/ziface"
)

var RoomMap map[int32]*msg.RoomInfo

func init() {

	RoomMap = make(map[int32]*msg.RoomInfo)
	for i := 1; i <= config.RoomSum; i++ {
		seatMap := make(map[int32]*msg.GameUserItem)
		for j := 1; j <= config.SeatSum; j++ {
			seatMap[int32(j)] = nil
		}
		RoomMap[int32(i)] = &msg.RoomInfo{
			Rid:           int32(i),
			GameNum:       0,
			Max:           int32(config.SeatSum),
			State:         int32(msg.RoomState_None),
			MapPlayerInfo: seatMap,
			CreateTime:    0,
			StartTime:     0,
			ResultTime:    0,
			Hint:          "",
			Word:          "",
			WordIndex:     0,
			Painter:       0,
		}
	}
}

func SetPRoom(pRoom *msg.RoomInfo) {
	RoomMap[pRoom.Rid] = pRoom
}

func GetPRoom(rid int32) (*msg.RoomInfo, bool) {
	pRoom, ok := RoomMap[rid]
	if !ok {
		fmt.Println("没有该房间")
	}
	return pRoom, ok
}

func GetRoomIdList() []*msg.RoomInfo {
	list := []*msg.RoomInfo{}
	for _, r := range RoomMap {
		// l := 0
		// for _, s := range r.MapPlayerInfo {
		// 	if s != nil {
		// 		l++
		// 	}
		// }
		list = append(list, r)
	}
	fmt.Println("房间数据")
	fmt.Println(list)
	return list
}

func GetFreeSeat(pRoom *msg.RoomInfo) int {
	seatMap := pRoom.MapPlayerInfo
	for i, _ := range seatMap {
		if seatMap[i] == nil {
			return int(i)
		}
	}
	return 0
}

func GetANewRid() int32 {
	var maxRid int32 = 0
	for _, v := range RoomMap {
		if (v.Rid) > maxRid {
			maxRid = v.Rid
		}
	}
	maxRid++
	return maxRid
}

func CreateRoom() *msg.RoomInfo {
	mapGameUserItem := make(map[int32]*msg.GameUserItem)
	rid := GetANewRid()
	for i := 1; i < config.SeatSum; i++ {
		mapGameUserItem[int32(i)] = nil
	}
	timestamp := time.Now().Unix()
	fmt.Println("当前时间戳（秒）：", timestamp)
	r := msg.RoomInfo{
		Rid:           rid,
		GameNum:       0,
		Max:           int32(config.SeatSum),
		State:         int32(msg.RoomState_None),
		CreateTime:    timestamp,
		StartTime:     0,
		ResultTime:    0,
		Hint:          "",
		Word:          "",
		WordIndex:     0,
		Painter:       0,
		MapPlayerInfo: mapGameUserItem,
	}
	SetPRoom(&r)
	return &r
}

func EnterRoom(pUser *msg.GameUserItem, pRoom *msg.RoomInfo, seat int32) {
	pUser.Rid = pRoom.Rid
	pUser.Seat = seat
	pRoom.MapPlayerInfo[seat] = pUser
	if pRoom.State == int32(msg.RoomState_None) {
		pRoom.State = int32(msg.RoomState_Ready)
	}
}

/***
 * 重置游戏
 */
func ResetGame(rid int32) {
	pRoom, _ := GetPRoom(rid)
	for _, user := range pRoom.MapPlayerInfo {
		user.Plyer = nil
		user.Rid = 0
		user.Seat = 0
		user.IsReady = false
	}
	pRoom.GameNum = 0
	pRoom.State = int32(msg.RoomState_None)
	pRoom.StartTime = 0
	pRoom.Hint = ""
	pRoom.Painter = 0
}

func MathchRoom(req ziface.IConnection) (int32, *msg.RoomInfo) {
	code := int32(enumeCode.NoRoomNotStart)
	roomItem := &msg.RoomInfo{}
	for _, v := range RoomMap {
		if (v.State) <= int32(msg.RoomState_Ready) { //还没开打
			//桌子没满
			if len(v.MapPlayerInfo) < int(v.Max) {
				code = int32(enumeCode.OK)
				roomItem = v
				break
			}

		}
	}
	return code, roomItem
}
