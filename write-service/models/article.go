package models

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID      int64 `gorm:"primaryKey"`
	Author  string
	Title   string
	Body    string
	Created time.Time
}

func (m *Article) BeforeCreate(tx *gorm.DB) error {
	m.Created = time.Now()

	return nil
}
