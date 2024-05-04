package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	FullName string             `json:"full_name" bson:"full_name"`
	Email    string             `json:"email" bson:"email"`
	Password *string            `json:"password,omitempty"`
}

type GetListUserOption struct {
	GetListOption
}

func (u *User) IsValid() bool {
	return u.ID.Hex() != ""
}
