package Repository

import (
	"github.com/umirode/prom-calendar-russia/src/Domain/Model/Entity"
)

type IHolidayRepository interface {
	Save(holiday *Entity.Holiday) error

	FindAll(withShortened bool) ([]*Entity.Holiday, error)
	FindAllByYear(year uint, withShortened bool) ([]*Entity.Holiday, error)
	FindAllByYearAndMonth(month uint, year uint, withShortened bool) ([]*Entity.Holiday, error)
	FindOneByDayMonthAndYear(day uint, month uint, year uint) (*Entity.Holiday, error)
}
