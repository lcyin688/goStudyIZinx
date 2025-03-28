package playerData

import (
	"fmt"
)

type User struct {
	Account  string
	Rid      int
	Seat     int
	Score    int
	IsReady  bool
	Username string
}

var UserMap map[string]*User

func init() {
	fmt.Println("init   Data.user.go  ")
	UserMap = make(map[string]*User)
}

func SetPUser(pUser *User) {
	UserMap[pUser.Account] = pUser
}

func GetPUser(uid string) *User {
	pUser, ok := UserMap[uid]
	if !ok {
		fmt.Println("没有该用户")
	}
	return pUser
}
