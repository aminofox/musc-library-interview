package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Playlist struct {
	ID     primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Name   string             `json:"name" bson:"name"`
	Tracks []string           `json:"tracks" bson:"tracks"`
}

type GetListPlaylistOption struct {
	GetListOption
	Name  string `json:"name" bson:"name"`
	Track string `json:"track" bson:"track"`
}
