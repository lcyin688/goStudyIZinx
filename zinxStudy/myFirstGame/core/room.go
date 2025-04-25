package core

import (
	"fmt"

	enumeCode "github.com/aceld/zinx/myFirstGame/EnumeCode"
	"github.com/aceld/zinx/myFirstGame/config"
	msg "github.com/aceld/zinx/myFirstGame/pb"
	"github.com/aceld/zinx/myFirstGame/tool"
	"github.com/aceld/zinx/ziface"
)

var RoomMap map[int32]*msg.RoomInfo

func init() {
	RoomMap = make(map[int32]*msg.RoomInfo)
	for i := 1; i <= config.RoomSum; i++ {
		createOneTabByRid(int32(i))
	}
}

/**
 * 用房间号创建一个房间
 */
func createOneTabByRid(i int32) {
	arr := []*msg.GameUserItem{}
	RoomMap[i] = &msg.RoomInfo{
		Rid:           i,
		GameNum:       0,
		Max:           int32(config.SeatSum),
		State:         int32(msg.RoomState_None),
		ArrPlayerInfo: arr,
		CreateTime:    0,
		GameTime:      0,
		ResultTime:    0,
		Hint:          "",
		Word:          "",
		WordIndex:     0,
		Painter:       0,
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
		list = append(list, r)
	}
	fmt.Println("房间数据")
	fmt.Println(list)
	return list
}

func GetFreeSeat(pRoom *msg.RoomInfo) int {
	seatMap := pRoom.ArrPlayerInfo
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

func CreateRoom(req ziface.IConnection, gameUserItem *msg.GameUserItem) (int32, *msg.RoomInfo) {
	// rid := GetANewRid()
	// timestamp := time.Now().Unix()
	// fmt.Println("当前时间戳（秒）：", timestamp)
	// arr := []*msg.GameUserItem{}

	var code int32 = int32(enumeCode.OK)

	roomItem := &msg.RoomInfo{}
	//首先 自己如果已经在房间了直接给在的房间
	if gameUserItem.Rid != 0 {
		pRoom, ok := GetPRoom(gameUserItem.Rid)
		if ok {
			code = int32(enumeCode.OK)
			roomItem = pRoom
		}
	} else {
		//如果一个人都没有的话进去占桌子 如果没有空桌子就新建桌子
		for _, v := range RoomMap {
			if (v.State) <= int32(msg.RoomState_Ready) { //还没开打
				//桌子没满
				if len(v.ArrPlayerInfo) == 0 {
					code = int32(enumeCode.OK)
					//匹配成功的时候好自己就坐进去
					gameUserItem.Rid = v.Rid
					v.ArrPlayerInfo = append(v.ArrPlayerInfo, gameUserItem)
					roomItem = v
					break
				}

			}
		}
		//

	}

	SetPRoom(roomItem)
	return code, roomItem
}

func EnterRoom(pUser *msg.GameUserItem, pRoom *msg.RoomInfo, seat int32) {
	pUser.Rid = pRoom.Rid
	pUser.Seat = seat

	pRoom.ArrPlayerInfo = append(pRoom.ArrPlayerInfo, pUser)
	if pRoom.State == int32(msg.RoomState_None) {
		pRoom.State = int32(msg.RoomState_Ready)
	}
}

/***
 * 重置游戏
 */
func ResetGame(rid int32) {
	pRoom, _ := GetPRoom(rid)
	for _, user := range pRoom.ArrPlayerInfo {
		user.Plyer = nil
		user.Rid = 0
		user.Seat = 0
		user.IsReady = false
	}
	pRoom.GameNum = 0
	pRoom.State = int32(msg.RoomState_None)
	pRoom.GameTime = 0
	pRoom.Hint = ""
	pRoom.Painter = 0
}

func MathchRoom(req ziface.IConnection, gameUserItem *msg.GameUserItem) (int32, *msg.RoomInfo) {
	code := int32(enumeCode.NoRoomNotStart)
	roomItem := &msg.RoomInfo{}
	//首先 自己如果已经在房间了直接给在的房间
	if gameUserItem.Rid != 0 {
		pRoom, ok := GetPRoom(gameUserItem.Rid)
		if ok {
			code = int32(enumeCode.OK)
			roomItem = pRoom
		}
	} else {
		isHaveFree := false
		//优先匹配桌子上有人但是还没开打的房间
		for _, v := range RoomMap {
			if (v.State) <= int32(msg.RoomState_Ready) { //还没开打
				//桌子没满
				if len(v.ArrPlayerInfo) < int(v.Max) {
					if len(v.ArrPlayerInfo) > 0 {
						code = int32(enumeCode.OK)
						//匹配成功的时候好自己就坐进去
						gameUserItem.Rid = v.Rid
						v.ArrPlayerInfo = append(v.ArrPlayerInfo, gameUserItem)
						roomItem = v
						isHaveFree = true
						break
					}
				}

			}
		}
		if !isHaveFree {
			for _, v := range RoomMap {
				if (v.State) <= int32(msg.RoomState_Ready) { //还没开打
					//桌子没满
					if len(v.ArrPlayerInfo) == 0 {
						code = int32(enumeCode.OK)
						//匹配成功的时候好自己就坐进去
						gameUserItem.Rid = v.Rid
						v.ArrPlayerInfo = append(v.ArrPlayerInfo, gameUserItem)
						roomItem = v
						isHaveFree = true
						break
					}

				}
			}
		}
		if !isHaveFree {
			newRid := GetANewRid()
			createOneTabByRid(int32(newRid))
		}

	}
	return code, roomItem
}

func ExitRoom(account string) int32 {
	code := int32(enumeCode.OK)
	player := WorldMgrObj.PlayersAcount[account]
	pRoom, ok := GetPRoom(player.GameUserItem.Rid)
	if ok { //如果房间开打了就不能退出了
		allowedStates := []int32{int32(msg.RoomState_None), int32(msg.RoomState_Ready), int32(msg.RoomState_Over)}
		if tool.Contains(allowedStates, pRoom.State) {
			//退出房间的清理掉房间数据
			ClearRoomDataByOnePlayerAccount(account)
		} else {
			code = int32(enumeCode.ExitOnStart)
		}

	} else {
		code = int32(enumeCode.Failed)
	}
	return code
}

// 清理掉玩家房间数据信息
func ClearRoomDataByOnePlayerAccount(account string) {

	player := WorldMgrObj.PlayersAcount[account]
	pRoom, ok := GetPRoom(player.GameUserItem.Rid)
	if ok {
		for i, user := range pRoom.ArrPlayerInfo {
			if user.Plyer.Account == account {
				pRoom.ArrPlayerInfo = append(pRoom.ArrPlayerInfo[:i], pRoom.ArrPlayerInfo[i+1:]...)
			}
		}
	}
	player.GameUserItem.Rid = 0
	player.GameUserItem.Seat = 0
	player.GameUserItem.Score = 0
	player.GameUserItem.IsReady = false

}
