package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Represents a user
type User struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
	Tasks    []Task             `bson:"tasks"`
}
