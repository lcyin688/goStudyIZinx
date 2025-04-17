package enumeCode

type ErrCodeType int

const (
	OK               ErrCodeType = 0 // 没问题
	Failed           ErrCodeType = 1 // 未知错误
	RegisterSameName ErrCodeType = 2 //注册账号已存在
	RegisterName     ErrCodeType = 3 //注册名字不合法
	RegisterPassWord ErrCodeType = 4 //注册密码不合法
	LoginPassWord    ErrCodeType = 5 //登录密码错误
	LoginName        ErrCodeType = 6 //登录账号不存在
	EnterRoomFull    ErrCodeType = 7 //房间已满
	NoRoom           ErrCodeType = 8 //不存在该房间
	PlayerReadyed    ErrCodeType = 9 //玩家已经准备过了
)
