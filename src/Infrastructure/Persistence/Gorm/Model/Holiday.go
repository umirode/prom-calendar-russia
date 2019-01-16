package Model

type Holiday struct {
	ID uint `gorm:"primary_key"`

	Day   uint `gorm:"not null;"`
	Month uint `gorm:"not null;"`
	Year  uint `gorm:"not null;"`

	Shortened bool `gorm:"not null;" sql:"default:false"`
}
