package models

type Paper struct {
	ID       string
	Author   string
	Title    string
	Sections []Section
}

type Section struct {
	Name    string
	Content string
}