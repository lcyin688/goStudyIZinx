package core

import (
	"fmt"
	"sync"

	msg "github.com/aceld/zinx/myFirstGame/pb"
	"github.com/aceld/zinx/ziface"
)

// Player object
type Player struct {
	PID          int32              // Player ID
	Conn         ziface.IConnection // Current player's connection
	GameUserItem *msg.GameUserItem
}

// Player ID Generator
var PIDGen int32 = 1  // Counter for generating player IDs(用来生成玩家ID的计数器)
var IDLock sync.Mutex // Mutex for protecting PIDGen(保护PIDGen的互斥机制)

// NewPlayer Create a player object
func NewPlayer(conn ziface.IConnection, playerInfo *msg.PlayerInfo) *Player {
	IDLock.Lock()
	ID := PIDGen
	PIDGen++
	IDLock.Unlock()

	gameUserItem := &msg.GameUserItem{
		Plyer:    playerInfo,
		Rid:      0,
		Seat:     0,
		Score:    0,
		IsReady:  false,
		IsOnline: true,
	}

	p := &Player{
		PID:          ID,
		Conn:         conn,
		GameUserItem: gameUserItem,
	}

	return p
}

// Player logs off
// 玩家下线
func (p *Player) LostConnection() {
	fmt.Println("玩家下线了 ... name ", p.GameUserItem.Plyer.NickName)
	p.GameUserItem.IsOnline = false

}

func (p *Player) DeletePlayer() {
	WorldMgrObj.RemovePlayerByPID(p.PID)
}
