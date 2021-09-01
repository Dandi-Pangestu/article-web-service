package models

import "time"

type Article struct {
	ID      int64 `gorm:"primaryKey"`
	Author  string
	Title   string
	Body    string
	Created time.Time
}
