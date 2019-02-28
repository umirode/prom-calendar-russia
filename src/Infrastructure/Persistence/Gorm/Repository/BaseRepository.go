package Repository

import (
	"github.com/jinzhu/gorm"
	"github.com/umirode/prom-calendar-russia/Database/Connection/Gorm"
)

type BaseRepository struct {
}

func (r *BaseRepository) GetGormDB() *gorm.DB {
	return Gorm.NewDatabase()
}
