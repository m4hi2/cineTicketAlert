package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DefaultModel struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primarykey;default:public.uuid_generate_v4()"`
	CreatedAt time.Time      `json:"created_at" gorm:"index:,type:brin"`
	UpdatedAt time.Time      `json:"-" gorm:"index:,type:brin"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index:,type:brin"`
}
