package znet

import (
	"net"

	"github.com/aceld/zinx/zinx/ziface"
)

/** 链接模块 */
type Connection struct {
	//当前连接 socket TCP 套接字
	Conn *net.TCPConn
	//链接的ID
	ConnID uint32
	//当前链接状态
	isClosed bool
	//链接绑定的处理业务 方法API
	handleAPI ziface.HandlerFunc
	//当前链接的关闭状态
	ExitChan chan bool
}

// 初始化链接模块的方法
func NewConnection(coon *net.TCPConn, coonID uint32, callback_api ziface.HandlerFunc) *Connection {
	c := &Connection{
		Conn:      coon,
		ConnID:    coonID,
		handleAPI: callback_api,
		isClosed:  false,
		ExitChan:  make(chan bool, 1)}
	return c
}


// //启动链接
// Start()
// //停止链接
// Stop()
// //获取当前链接绑定 socket Conn
// GetTCPConnection() *net.TCPConn
// //获取当前链接ID
// GetConnID() uint32
// //获取远程客户端的  TCP 状态 IP Port
// RemoteAddr() net.Addr
// //发送数据 ,将数据发送给客户端
// SendMsg(data []byte) error
