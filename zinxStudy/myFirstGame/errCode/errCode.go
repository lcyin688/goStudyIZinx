package errCode

const (
	OK               = 0 // 没问题
	Failed           = 1 // 未知错误
	RegisterSameName = 2 //注册账号已存在
	RegisterName     = 2 //注册名字不合法
	RegisterPassWord = 3 //注册密码不合法
	LoginPassWord    = 4 //登录密码错误
	LoginName        = 4 //登录账号不存在
)
