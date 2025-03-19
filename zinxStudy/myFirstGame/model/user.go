package model

func ValidateUserData(uid string, password string) (*UserData, error) {
	session, err := mgo.Dial("119.29.40.244:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("users")
	user := UserData{}
	err = c.Find(bson.M{"uid": uid, "password": password}).One(&user)
	return &user, err
}
