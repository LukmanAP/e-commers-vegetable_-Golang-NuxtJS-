package model

import "time"

type Role struct {
	ID        int64 `gorm:"primaryKey"`
	Name      string
	User      []User `gorm:"many2many:user_role;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
