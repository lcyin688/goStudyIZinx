package core

import (
	"sync"
)

// WorldManager The overall management module of the current game world 当前游戏世界的总管理模块
type WorldManager struct {
	Players       map[int32]*Player  // (当前在线的玩家集合)
	PlayersAcount map[string]*Player // (当前在线的玩家集合)
	pLock         sync.RWMutex       // Mutual exclusion mechanism to protect Players(保护Players的互斥读写机制)
}

// Provide an external handle to the world management module
// 提供一个对外的世界管理模块句柄
var WorldMgrObj *WorldManager

// Provide an initialization method for WorldManager
// 提供WorldManager 初始化方法
func init() {
	WorldMgrObj = &WorldManager{
		Players:       make(map[int32]*Player),
		PlayersAcount: make(map[string]*Player),
	}
}

// AddPlayer Provide the ability to add a player, adding the player to the player information table Players
// (提供添加一个玩家的的功能，将玩家添加进玩家信息表Players)
func (wm *WorldManager) AddPlayer(player *Player) {
	// Add the player to the world manager
	// 将player添加到 世界管理器中
	wm.pLock.Lock()
	//如果这个玩家之前存在过了,则删除之前的
	if wm.PlayersAcount[player.GameUserItem.Plyer.Account] != nil {
		oldPlayer := wm.PlayersAcount[player.GameUserItem.Plyer.Account]
		oldPlayer.GameUserItem.IsOnline = true
		oldPlayer.Conn = player.Conn
		oldPlayer.PID = player.PID

		wm.Players[player.PID] = oldPlayer
		wm.PlayersAcount[player.GameUserItem.Plyer.Account] = oldPlayer
	} else {
		wm.Players[player.PID] = player
		wm.PlayersAcount[player.GameUserItem.Plyer.Account] = player
	}

	wm.pLock.Unlock()

	// Add the player to the AOI network planning

}

// RemovePlayerByPID Remove a player from the player information table by player ID
// 从玩家信息表中移除一个玩家
func (wm *WorldManager) RemovePlayerByPID(pID int32) {
	wm.pLock.Lock()
	delete(wm.Players, pID)
	wm.pLock.Unlock()
}

// GetPlayerByPID Get corresponding player information by player ID
// 通过玩家ID 获取对应玩家信息
func (wm *WorldManager) GetPlayerByPID(pID int32) *Player {
	wm.pLock.RLock()
	defer wm.pLock.RUnlock()

	return wm.Players[pID]
}

// GetAllPlayers Get information of all players
// 获取所有玩家的信息
func (wm *WorldManager) GetAllPlayers() []*Player {
	wm.pLock.RLock()
	defer wm.pLock.RUnlock()

	// Create a slice to return player collection
	// 创建返回的player集合切片
	players := make([]*Player, 0)

	// Add to the slice
	// 添加切片
	for _, v := range wm.Players {
		players = append(players, v)
	}

	return players
}
