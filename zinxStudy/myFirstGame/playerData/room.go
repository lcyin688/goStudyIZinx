package playerData

import (
	"fmt"

	"github.com/aceld/zinx/myFirstGame/config"
	msg "github.com/aceld/zinx/myFirstGame/pb"
)

var RoomMap map[int32]*msg.RoomInfo

func init() {

	RoomMap = make(map[int32]*msg.RoomInfo)
	for i := 1; i <= config.RoomSum; i++ {
		seatMap := make(map[int32]*msg.PlayerInfo)
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

func GetRoomIdList() []interface{} {
	var list []interface{}
	for _, r := range RoomMap {
		l := 0
		for _, s := range r.MapPlayerInfo {
			if s != nil {
				l++
			}
		}
		list = append(list, struct {
			Rid int32 `json:"rid"`
			Num int32 `json:"num"` //人数
			Max int   `json:"max"` //房间座位数
		}{
			Rid: r.Rid,
			Num: (int32(l)),
			Max: len(r.MapPlayerInfo),
		})
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

func GetANewRid() int {
	maxRid := 0
	for _, v := range RoomMap {
		if (int(v.Rid)) > maxRid {
			maxRid = int(v.Rid)
		}
	}
	maxRid++
	return maxRid
}
