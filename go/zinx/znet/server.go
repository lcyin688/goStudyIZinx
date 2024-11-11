package znet

import (
	"fmt"
	"net"

	"github.com/aceld/zinx/zinx/ziface"
)

// IServer的接口实现,定义一个Server服务器模块
type Server struct {
	//服务器名称
	Name string
	//服务器绑定的ip版本
	IPVersion string
	//服务器监听ip
	IP string
	//服务器监听端口
	Port int
}

// 启动服务器
func (s *Server) Start() {
	//1 获取一个TCP的Addr
	fmt.Printf(" [Start] Server Listenner at IP:%s ,Port %d ,is starting \n", s.IP, s.Port)
	net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	//2 监听服务器的地址

	//3 阻塞的等待客户端链接,处理客户端链接业务(读写)

}

// 停止服务器
func (s *Server) Stop() {

}

// 运行服务器
func (s *Server) Serve() {
	s.Start()

}

/**
 * 初始化Server模块的方法
 */
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}
