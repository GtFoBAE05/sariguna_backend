package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"sariguna_backend/sariguna/app/user/entity"
	"sariguna_backend/sariguna/app/user/model"
	"sariguna_backend/sariguna/pkg/helpers"

	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) entity.UserRepositoryInterface {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Register(data entity.UserCore) (entity.UserCore, error) {
	request := entity.UserCoreToUserModel(data)

	hashedPassword, err := helpers.HashPassword(request.Password)
	if err != nil {
		return entity.UserCore{}, err
	}

	request.Password = hashedPassword

	query := `INSERT INTO users (id, fullname, email, password, role) 
	VALUES (:id, :fullname, :email, :password, :role)`

	_, err = ur.db.NamedExec(query, request)

	if err != nil {
		return entity.UserCore{}, err
	}

	dataResponse := entity.UserModelToUserCore(request)

	return dataResponse, nil
}

var ErrUserNotFound = errors.New("user not found")

func (ur *userRepository) FindByEmail(email string) (entity.UserCore, error) {
	dataUser := model.User{}

	query := `SELECT id, fullname, email, password, role FROM users WHERE email = $1`

	err := ur.db.Get(&dataUser, query, email)

	if err != nil {
		if err == sql.ErrNoRows {
			return entity.UserCore{}, fmt.Errorf("no user found with email: %s", email)
		}
		return entity.UserCore{}, fmt.Errorf("error in FindByEmail: %w", err)
	}

	dataResponse := entity.UserModelToUserCore(dataUser)

	return dataResponse, nil
}
