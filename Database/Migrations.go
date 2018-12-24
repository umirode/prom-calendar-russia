package Database

import (
	"github.com/jinzhu/gorm"
	"github.com/umirode/prom-calendar-russia/src/Infrastructure/Persistence/Gorm/Model"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(
		&Model.Holiday{},
	)
}
