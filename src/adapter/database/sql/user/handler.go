package userdb

import (
	"gofiber-example/src/adapter/database/sql"
	"gofiber-example/src/internal/model/dbmodel"
)

type CreateFn func(obj *dbmodel.User) error

func Create(user *dbmodel.User) error {
	db, err := sql.GetGormDB()

	if err != nil {
		return err
	}

	return db.Create(user).Error
}

type GetUserByEmailFn func(email string) (*dbmodel.User, error)

func GetUserByEmail(email string) (*dbmodel.User, error) {
	db, err := sql.GetGormDB()
	var user dbmodel.User

	if err != nil {
		return nil, err
	}

	err = db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil

}
