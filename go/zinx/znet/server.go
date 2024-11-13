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

	fmt.Printf(" [Start] Server Listenner at IP:%s ,Port %d ,is starting \n", s.IP, s.Port)
	go func() {
		//1 获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error : ", err)
			return
		}
		//2 监听服务器的地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("TCPListener : ", s.IPVersion, " start error : ", err)
			return
		}
		fmt.Println("start Zinx server  success , now listenning ...", s.Name)
		//3 阻塞的等待客户端链接,处理客户端链接业务(读写)
		for {
			// 通过Accept阻塞等待客户端链接，成功返回链接对象
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept error : ", err)
				continue
			}
			//客户端已经与服务器建立连接,做一些业务,做一个 最大 512 字节长度的回显业务
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println(" recv buf  err : ", err)
						continue
					}
					//回显功能
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println(" write back buf  err : ", err)
						continue
					}

				}
			}()
		}
	}()
}

// 停止服务器
func (s *Server) Stop() {
	//TODO 将一些服务器的资源,状态或者已经开启的链接信息 进行停止或者回收
}

// 运行服务器
func (s *Server) Serve() {
	s.Start()
	//TODO 做一些启动服务器之后的额外业务

	//阻塞状态 防止程序退出
	select {}
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
