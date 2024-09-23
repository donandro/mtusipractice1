package models

import "time"

type RefreshToken struct {
	Token   string
	Expires time.Time
	UserID  int
}
