package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Paper struct {
	ID       primitive.ObjectID	`bson:"_id,omitempty" json:"id"`
	Author   string							`bson:"author" json:"author"`
	AuthorID string							`bson:"authorId" json:"authorId"`
	Title    string							`bson:"title" json:"title"`
	Sections []Section					`bson:"sections" json:"sections"`
}

type Section struct {
	Name    string	`bson:"name" json:"name"`
	Content string	`bson:"content" json:"content"`
}