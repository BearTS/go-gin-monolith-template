package database

import (
	"github.com/BearTS/go-gin-monolith/database/seeds"
	"gorm.io/gorm"
)

type Seed struct {
	TableName string
	Run       func(*gorm.DB) error
}

func Seeder(db *gorm.DB) []Seed {
	users := Seed{TableName: "users", Run: func(d *gorm.DB) error { return seeds.Users(db) }}

	return []Seed{
		users,
	}
}
