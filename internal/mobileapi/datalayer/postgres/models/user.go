package models

import "time"

type UserDB struct {
	ID       int
	Name     string
	Email    string
	Login    string
	Password string
	Created  time.Time
	Deleted  time.Time
}
