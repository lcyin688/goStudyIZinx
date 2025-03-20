package model

import (
	"context"
	"log"
	"time"

	enumeCode "github.com/aceld/zinx/myFirstGame/EnumeCode"
	msg "github.com/aceld/zinx/myFirstGame/pb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type UserData struct {
	uid      string
	userName string
	password string
	headId   int32
	money    int64
}

const initMoney = 1000

var mongoClient *mongo.Client = nil

/**通过 账号密码 验证 获取玩家信息 **/
func ValidateUserData(uid string, password string) (*msg.PlayerInfo, enumeCode.ErrCodeType) {
	clientOpts := options.Client().
		ApplyURI("mongodb://localhost:27017").
		SetAuth(options.Credential{
			Username:   "adminLcy",
			Password:   "xiuxiu520",
			AuthSource: "adminLcy",
		}).
		SetConnectTimeout(10 * time.Second).
		SetMaxPoolSize(100)

	// 2. 建立连接
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal("连接失败:", enumeCode.Failed)
		return nil, enumeCode.Failed
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal("关闭连接异常:", enumeCode.Failed)
		}
	}()

	// 3. 验证连接状态
	// if err = client.Ping(context.TODO(), nil); err != nil {
	// 	log.Fatal("心跳检测失败:", enumeCode.Failed)
	// 	return nil, enumeCode.Failed
	// }
	log.Println("连接成功")
	user := &msg.PlayerInfo{}
	// 4. 获取集合
	c := client.Database("lcyFirst").Collection("users")

	// 5. 执行查询
	err = c.FindOne(
		context.TODO(),
		bson.M{"uid": uid, "password": password},
	).Decode(user)

	// 6. 错误处理

	if err == mongo.ErrNoDocuments {
		// 处理文档不存在逻辑
		return nil, enumeCode.Failed
	} else if err != nil {
		// 处理其他错误（如网络中断、权限不足等）
		return nil, enumeCode.Failed
	} else {
		// 使用user对象数据
		return user, enumeCode.OK
	}
}

func RegisteUserData(uid string, password string, headId int32) enumeCode.ErrCodeType {
	code := enumeCode.Failed
	//向数据库中写入数据 mongodb://localhost:27017 uid password headId
	if uid == "" || password == "" {
		code = enumeCode.Failed
	} else {
		initMongod()
		collection := mongoClient.Database("mydb").Collection("users")
		//如果之前没注册过
		if !IsUserExist(collection, uid) {
			code = enumeCode.OK
		} else {
			InsertUser(collection, uid, password, headId)
		}
	}
	return code
}

func IsUserExist(collection *mongo.Collection, uid string) bool {
	return false
}

func InsertUser(collection *mongo.Collection, uid string, password string, headId int32) error {
	user := &UserData{
		uid:      uid,
		userName: uid,
		password: password,
		headId:   headId,
		money:    initMoney,
	}

	// 执行插入
	_, err := collection.InsertOne(context.TODO(), user)
	return err
}

func initMongod() {
	if mongoClient == nil {

		clientOpts := options.Client().
			ApplyURI("mongodb://localhost:27017").
			SetAuth(options.Credential{
				Username:   "adminLcy",
				Password:   "xiuxiu520",
				AuthSource: "adminLcy",
			}).
			SetConnectTimeout(10 * time.Second).
			SetMaxPoolSize(100)

		// 2. 建立连接
		client, err := mongo.Connect(context.TODO(), clientOpts)
		if err != nil {
			log.Fatal("连接失败:", err)
			return
		}
		mongoClient = client
	}

}
func test() {
	// 使用更安全的连接字符串
	uri := "mongodb://appuser:AppPassword2023@localhost:27017/myappdb?" +
		"authSource=admin&" +
		"authMechanism=SCRAM-SHA-256&" +
		"tls=true&" + // 强制 TLS
		"maxPoolSize=100&" +
		"serverSelectionTimeoutMS=5000"

	// 客户端配置
	clientOpts := options.Client().
		ApplyURI(uri).
		SetMinPoolSize(10).
		SetRetryReads(true).
		SetConnectTimeout(5 * time.Second).
		SetSocketTimeout(3 * time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Printf("Disconnect error: %v", err)
		}
	}()

	// 增强的健康检查
	if err = client.Ping(ctx, readpref.PrimaryPreferred()); err != nil {
		log.Fatalf("Ping failed: %v", err)
	}
}
