package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uint          `json:"id" gorm:"not null; primaryKey; autoIncrement"`
	Username    string        `json:"username" gorm:"size:25; not null; unique; index" validate:"required"`
	Email       string        `json:"email" gorm:"size:100; not null;unique;index" validate:"required,email"`
	Age         uint8         `json:"age"`
	Permissions []*Permission `json:"permissions" gorm:"many2many:user_permissions;"`
	Admin       *bool         `json:"admin" gorm:"not null;"`
}

type Permission struct {
	gorm.Model
	ID    int     `json:"id" gorm:"not null;primaryKey"`
	Name  string  `json:"name" gorm:"size:25;not null;unique;index" validate:"required"`
	Users []*User `json:"users" gorm:"many2many:user_permissions;"`
}
