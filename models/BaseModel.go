package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type BaseModel struct {
	Id       uuid.UUID `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}