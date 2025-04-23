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
)

type UserData struct {
	UID      string `bson:"uid"`      // 明确映射到 "uid"
	UserName string `bson:"userName"` // 映射到 "userName"（驼峰命名）
	Password string `bson:"password"`
	HeadID   int32  `bson:"headId"` // 注意字段名与标签匹配
	Money    int64  `bson:"money"`
}

const initMoney = 1000

var mongoClient *mongo.Client = nil

/**通过 账号密码 验证 获取玩家信息 **/
func ValidateUserData(uid string, password string) (*msg.PlayerInfo, enumeCode.ErrCodeType) {
	code := enumeCode.Failed
	//向数据库中写入数据 mongodb://localhost:27017 uid password headId
	if uid == "" || password == "" {
		return nil, code
	} else {
		initMongod()
		c := mongoClient.Database("mydb").Collection("users")

		user := &UserData{}
		// 5. 执行查询
		err := c.FindOne(
			context.TODO(),
			bson.M{"uid": uid},
		).Decode(user)

		// 6. 错误处理

		if err == mongo.ErrNoDocuments {
			code = enumeCode.LoginName
			return nil, code
		} else if err != nil {
			// 处理其他错误（如网络中断、权限不足等）
			return nil, code
		} else {
			err = c.FindOne(
				context.TODO(),
				bson.M{"uid": uid, "password": password},
			).Decode(user)
			if err != nil {
				code = enumeCode.LoginPassWord
				return nil, code
			}
			userOut := &msg.PlayerInfo{}
			userOut.HeadId = int32(user.HeadID)
			userOut.Account = user.UID
			userOut.NickName = user.UserName
			// 使用user对象数据
			return userOut, enumeCode.OK
		}
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
		//首先看下账号是否已经存在了
		var result UserData
		err := collection.FindOne(context.TODO(), bson.M{"uid": uid}).Decode(&result)
		if err == nil {
			code = enumeCode.RegisterSameName
			// ClearUserData() //测试删除数据库
		} else {
			InsertUser(collection, uid, password, headId)
			code = enumeCode.OK
		}
	}
	return code
}

func InsertUser(collection *mongo.Collection, uid string, password string, headId int32) error {
	user := &UserData{
		UID:      uid,
		UserName: uid,
		Password: password,
		HeadID:   headId,
		Money:    initMoney,
	}

	// 执行插入
	_, err := collection.InsertOne(context.TODO(), user)
	return err
}

func initMongod() {
	if mongoClient == nil {

		clientOpts := options.Client().
			ApplyURI("mongodb://localhost:27017").
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

/** 清空数据库 **/
func ClearUserData() {
	initMongod()
	collection := mongoClient.Database("mydb").Collection("users")
	// 清空数据库
	collection.Drop(context.TODO())
}
