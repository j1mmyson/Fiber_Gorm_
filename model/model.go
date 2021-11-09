package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid"`
	Password string
}

type UserInfo struct {
	UserID       string
	Name         string
	Email        string
	School       string
	Major        string
	Introduction string
}

type Article struct {
	gorm.Model
	UserID  uuid.UUID `gorm:"type:uuid"`
	Title   string
	Content string
	Tags    []Tag
}

type Tag struct {
	gorm.Model
	Name string
}

type UserTag struct {
	gorm.Model
	UserID uuid.UUID `gorm:"type:uuid"`
	TagID  uuid.UUID `gorm:"type:uuid"`
}

type ArticleTag struct {
	gorm.Model
	ArticleID uuid.UUID `gorm:"type:uuid"`
	TagID     uuid.UUID `gorm:"type:uuid"`
}
