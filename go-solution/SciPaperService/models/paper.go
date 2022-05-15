package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Paper struct {
	ID       primitive.ObjectID  `bson:"_id,omitempty"`
	Author   string
	AuthorID string
	Title    string
	Sections []Section
}

type Section struct {
	Name    string
	Content string
}