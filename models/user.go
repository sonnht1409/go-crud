package models

import (
	"../db"
	"../forms"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Username string        `json:username`
	Password string        `json:password`
}

type UserModel struct{}

var server = "127.0.0.1"

var dbConnect = db.NewConnection(server)

func (m *UserModel) Create(data forms.User) error {
	collection := dbConnect.Use("user", "users")
	err := collection.Insert(bson.M{"username": data.Username, "password": data.Password})
	return err
}

func (m *UserModel) Find() (list []User, err error) {
	collection := dbConnect.Use("user", "users")
	err = collection.Find(bson.M{}).All(&list)
	return list, err
}

func (m *UserModel) Get(id string) (user User, err error) {
	collection := dbConnect.Use("user", "users")
	err = collection.FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

func (m *UserModel) FindByUsername(username string) (user User, err error) {
	collection := dbConnect.Use("user", "users")
	err = collection.Find(bson.M{"username": username}).One(&user)
	return user, err
}
