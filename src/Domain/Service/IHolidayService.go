package Service

import (
	"github.com/umirode/prom-calendar-russia/src/Domain/Model/Entity"
	"github.com/umirode/prom-calendar-russia/src/Domain/Service/DTO"
)

type IHolidayService interface {
	GetOneByYearMonthAndDay(year uint, month uint, day uint) (*Entity.Holiday, error)
	GetAllByYearAndMonth(year uint, month uint) ([]*Entity.Holiday, error)
	GetAllByYear(year uint) ([]*Entity.Holiday, error)

	CreateIfNotExists(holidayDTO *DTO.HolidayDTO) error
}
