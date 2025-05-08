package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ToDoList struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"id, omitempty"`
	Task   string             `jsom:"task, omitempty"`
	Status bool               `json:"status,omitempty"`
}
