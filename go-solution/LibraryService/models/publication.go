package models

import "time"

type Publication struct {
	ID        string
	Author    string
	Title     string
	TimeStamp time.Time
}