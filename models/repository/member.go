package repository

import (
	"BBlog/models/model"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Create -
func Create(in *model.Member) string {
	in.ID = primitive.NewObjectID()
	db := ConnectDB()
	collection := db.Database("BBlog").Collection("Member")
	_, err := collection.InsertOne(context.TODO(), in)
	if err != nil {
		return err.Error()
	}
	return ""
}
