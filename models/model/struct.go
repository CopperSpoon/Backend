package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//MemberReq -
type MemberReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

//Member -
type Member struct {
	ID       primitive.ObjectID `bson:"_id"`
	Key      string             `bson:"key" json:"key"`
	Account  string             `bson:"account" json:"account"`
	Info     Info               `bson:"info" json:"info"`
	CreateAt string             `bson:"createAt"`
	UpdateAt string             `bson:"updateAt"`
	LoginAt  string             `bson:"loginAt"`
}

//Info -
type Info struct {
	Name  string `bson:"name" json:"name"`
	Email string `bson:"email" json:"email"`
}

//Response -
type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

//Default -
func (m *Member) Default() {
	m.CreateAt = time.Now().String()
}

//Auth -
type Auth struct {
	Account  string `json:"account"`
	Password string `json:"Password"`
}
