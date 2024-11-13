package ziface

import "net"

//定义链接模块的抽象层
type IConnection interface {
	//启动链接
	Start()
	//停止链接
	Stop()
	//获取当前链接绑定 socket Conn
	GetTCPConnection() *net.TCPConn
	//获取当前链接ID
	GetConnID() uint32
	//获取远程客户端的  TCP 状态 IP Port
	RemoteAddr() net.Addr
	//发送数据 ,将数据发送给客户端
	SendMsg(data []byte) error
}

//定义一个消息处理模块的抽象层
type HandlerFunc func(*net.TCPConn, []byte, int) error
