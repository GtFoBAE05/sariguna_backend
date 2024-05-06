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

// FindById implements entity.UserRepositoryInterface.
func (ur *userRepository) FindById(id string) (entity.UserCore, error) {
	dataUser := model.User{}

	query := `SELECT id, fullname, email, password, role FROM users WHERE id = $1`

	err := ur.db.Get(&dataUser, query, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return entity.UserCore{}, fmt.Errorf("no user found with email: %s", id)
		}
		return entity.UserCore{}, fmt.Errorf("error in FindByEmail: %w", err)
	}

	dataResponse := entity.UserModelToUserCore(dataUser)

	return dataResponse, nil
}

// UpdatePassword implements entity.UserRepositoryInterface.
func (ur *userRepository) UpdatePassword(id, password string) error {
	query := `UPDATE  users
		SET password = $1
		WHERE id = $2
		`

	hashedPassword, err := helpers.HashPassword(password)
	if err != nil {
		return err
	}

	_, err = ur.db.Exec(query, hashedPassword, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no user found with id: %s", id)
		}
		return fmt.Errorf("error in findById: %w", err)
	}

	return nil
}
