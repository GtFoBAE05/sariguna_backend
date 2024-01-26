package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserCore struct {
	Id        uuid.UUID
	Fullname  string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
