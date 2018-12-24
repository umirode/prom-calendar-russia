package Converter

import (
	"github.com/umirode/prom-calendar-russia/src/Domain/Model/Entity"
	"github.com/umirode/prom-calendar-russia/src/Infrastructure/Persistence/Gorm/Model"
)

type HolidayConverter struct {
}

func (*HolidayConverter) ToDatabaseEntity(entity interface{}) (interface{}, error) {
	if entity == nil {
		return nil, nil
	}

	holiday := entity.(*Entity.Holiday)

	model := &Model.Holiday{
		ID:    holiday.ID,
		Day:   holiday.Day,
		Month: holiday.Month,
		Year:  holiday.Year,
	}

	return model, nil
}

func (*HolidayConverter) ToAppEntity(entity interface{}) (interface{}, error) {
	if entity == nil {
		return nil, nil
	}

	holiday := entity.(*Model.Holiday)

	model := &Entity.Holiday{
		ID:    holiday.ID,
		Day:   holiday.Day,
		Month: holiday.Month,
		Year:  holiday.Year,
	}

	return model, nil
}
