package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Document struct {
	ID       primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Path     string             `json:"path" bson:"path"`
	ParentID string             `json:"parent_id" bson:"parent_id"`
	Category string             `json:"category" bson:"category"`
	Status   DocumentStatus     `json:"status" bson:"status"`
}

type DocumentStatus uint8

const (
	Completed DocumentStatus = iota + 1
	Failed
)
