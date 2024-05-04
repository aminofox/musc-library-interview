package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Album struct {
	ID          primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	ReleaseYear int                `json:"release_year" bson:"release_year"`
	CoverImage  string             `json:"cover_image" bson:"cover_image"`
	Tracks      []string           `json:"tracks" bson:"tracks"`
}

type GetListAlbumOption struct {
	GetListOption
	Title       string `json:"title" bson:"title"`
	ReleaseYear int    `json:"release_year" bson:"release_year"`
	Track       string `json:"track" bson:"track"`
}
