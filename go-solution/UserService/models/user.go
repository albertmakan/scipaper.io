package models

type User struct {
	Username  string
	Password  string
	Salt      string
	FirstName string
	LastName  string
	Email     string
}