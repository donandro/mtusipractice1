package postgres

import (
	"strings"

	dm "pillsreminder/internal/mobileapi/datalayer/postgres/models"
	db "pillsreminder/pkg/db/postgres"
)

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (rtr *UserRepo) Add(user dm.UserDB) (*dm.UserDB, error) {
	conn := db.DefaultDatabase.GetDB()

	var newUser dm.UserDB
	result := conn.Raw(`
		INSERT INTO "users" ("name", login, password, email)
		VALUES(?, ?, ?, ?)
		RETURNING id`,
		user.Name,
		user.Login,
		user.Password,
		user.Email,
	).Scan(&newUser)

	if result.Error != nil {
		return nil, result.Error
	}

	return &newUser, nil
}

func (rtr *UserRepo) Update(user dm.UserDB) error {
	conn := db.DefaultDatabase.GetDB()

	result := conn.Exec(`
		UPDATE "users" SET "name" = ?, password = ?, login = ?, email = ? WHERE id = ?;`,
		user.Name,
		user.Password,
		user.Login,
		user.Email,
		user.ID,
	)

	return result.Error
}

func (rtr *UserRepo) GetByID(id int) (*dm.UserDB, error) {
	conn := db.DefaultDatabase.GetDB()

	var user dm.UserDB
	result := conn.Raw(`
		SELECT id, "name", login, password, email, created, deleted
		FROM "users"
		WHERE id = ?`,
		id,
	).Scan(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}

	return &user, nil
}

func (rtr *UserRepo) GetByLogin(login string) (*dm.UserDB, error) {
	conn := db.DefaultDatabase.GetDB()

	var user dm.UserDB
	result := conn.Raw(`
		SELECT id, "name", login, password, email, created, deleted
		FROM "users"
		WHERE lower(login) = ?`,
		strings.ToLower(login),
	).Scan(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}

	return &user, nil
}

func (rtr *UserRepo) DeleteByID(id int) error {
	conn := db.DefaultDatabase.GetDB()

	result := conn.Exec(`
		UPDATE "users" SET deleted = now() WHERE id = ?`,
		id,
	)

	return result.Error
}
