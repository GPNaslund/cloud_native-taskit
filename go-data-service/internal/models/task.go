package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Represents a task
type Task struct {
	Id      primitive.ObjectID `bson:"_id,omitempty"`
	Title   string             `bson:"title"`
	Details string             `bson:"details"`
	IsDone  bool               `bson:"is_done"`
}
