package playerData

import (
	"fmt"

	msg "github.com/aceld/zinx/myFirstGame/pb"
)

var UserMap map[string]*msg.GameUserItem

func init() {
	UserMap = make(map[string]*msg.GameUserItem)
}

func SetPUser(pUser *msg.GameUserItem) {
	UserMap[pUser.Plyer.Account] = pUser
}

func GetPUser(uid string) *msg.GameUserItem {
	pUser, ok := UserMap[uid]
	if !ok {
		fmt.Println("没有该用户")
	}
	return pUser
}
