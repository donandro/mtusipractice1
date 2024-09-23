package models

import "time"

type User struct {
	ID       int
	Name     string
	Email    string
	Login    string
	Password string
	Created  time.Time
	Deleted  time.Time
}
