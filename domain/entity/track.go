package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Track struct {
	ID          primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Artist      string             `json:"artist" bson:"artist"`
	Album       string             `json:"album" bson:"album"`
	Genre       string             `json:"genre" bson:"genre"`
	ReleaseYear int                `json:"release_year" bson:"release_year"`
	Duration    int                `json:"duration" bson:"duration"`
	Mp3File     string             `json:"mp3_file" bson:"mp3_file"`
}

type GetListTrackOption struct {
	GetListOption
	Title  string `json:"title" bson:"title"`
	Artist string `json:"artist" bson:"artist"`
	Album  string `json:"album" bson:"album"`
	Genre  string `json:"genre" bson:"genre"`
}
