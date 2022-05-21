package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Publication struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	PaperID   primitive.ObjectID  `bson:"paperId" json:"paperId"`
	Author    string  						`bson:"author" json:"author"`
	Title     string  						`bson:"title" json:"title"`
}