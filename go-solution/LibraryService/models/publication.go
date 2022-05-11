package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Publication struct {
	ID        primitive.ObjectID
	Author    string
	Title     string
	TimeStamp time.Time
}