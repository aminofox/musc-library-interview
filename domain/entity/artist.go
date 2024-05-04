package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Artist struct {
	ID      primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Name    string             `json:"name" bson:"name"`
	Country string             `json:"country" bson:"country"`
	Avatar  string             `json:"avatar" bson:"avatar"`
	Tracks  []string           `json:"tracks" bson:"tracks"`
}

type GetListArtistOption struct {
	GetListOption
	Name    string `json:"name" bson:"name"`
	Country string `json:"country" bson:"country"`
	Track   string `json:"track" bson:"track"`
}
