package datalayer

import (
	dm "pillsreminder/internal/mobileapi/datalayer/postgres/models"
)

type IUserRepo interface {
	Add(user dm.UserDB) (*dm.UserDB, error)
	Update(user dm.UserDB) error
	GetByID(id int) (*dm.UserDB, error)
	DeleteByID(id int) error
	GetByLogin(login string) (*dm.UserDB, error)
}
