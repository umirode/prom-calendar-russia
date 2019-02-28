package Gorm

import (
	"github.com/jinzhu/gorm"
	"github.com/umirode/prom-calendar-russia/src/Domain/Model/Entity"
)

func runMigrations(db *gorm.DB) {
	db.AutoMigrate(
		&Entity.Holiday{},
		&Entity.RefreshToken{},
		&Entity.User{},
	)
}
