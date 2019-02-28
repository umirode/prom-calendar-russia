package Repository

import (
	"github.com/umirode/prom-calendar-russia/src/Domain/Model/Entity"
)

type HolidayRepository struct {
	BaseRepository
}

func NewHolidayRepository() *HolidayRepository {
	return &HolidayRepository{}
}

func (r *HolidayRepository) Save(holiday *Entity.Holiday) error {
	r.GetGormDB().Save(holiday)

	return nil
}

func (r *HolidayRepository) FindAll(withShortened bool) ([]*Entity.Holiday, error) {
	holidays := make([]*Entity.Holiday, 0)

	if withShortened {
		r.GetGormDB().Find(&holidays)
	} else {
		r.GetGormDB().Where("shortened = ?", 0).Find(&holidays)
	}

	return holidays, nil
}

func (r *HolidayRepository) FindAllByYear(year uint, withShortened bool) ([]*Entity.Holiday, error) {
	holidays := make([]*Entity.Holiday, 0)

	if withShortened {
		r.GetGormDB().Where("year = ?", year).Find(&holidays)
	} else {
		r.GetGormDB().Where("year = ? and shortened = ?", year, 0).Find(&holidays)
	}

	return holidays, nil
}

func (r *HolidayRepository) FindAllByYearAndMonth(month uint, year uint, withShortened bool) ([]*Entity.Holiday, error) {
	holidays := make([]*Entity.Holiday, 0)

	if withShortened {
		r.GetGormDB().Where("year = ? and month = ?", year, month).Find(&holidays)
	} else {
		r.GetGormDB().Where("year = ? and month = ? and shortened = ?", year, month, 0).Find(&holidays)
	}

	return holidays, nil
}

func (r *HolidayRepository) FindOneByDayMonthAndYear(day uint, month uint, year uint) (*Entity.Holiday, error) {
	holiday := &Entity.Holiday{}

	r.GetGormDB().Where("day = ? and month = ? and year = ?", day, month, year).First(holiday)
	if r.GetGormDB().NewRecord(holiday) {
		return nil, nil
	}

	return holiday, nil
}
