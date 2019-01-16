package Hydrator

import (
	"github.com/umirode/prom-calendar-russia/src/Domain/Model/Entity"
)

type HolidayHydrator struct {
}

func (*HolidayHydrator) Create(data map[string]interface{}) (interface{}, error) {
	panic("implement me")
}

func (*HolidayHydrator) Hydrate(data map[string]interface{}, object interface{}) (interface{}, error) {
	panic("implement me")
}

func (*HolidayHydrator) Extract(object interface{}) (map[string]interface{}, error) {
	holiday := object.(*Entity.Holiday)

	return map[string]interface{}{
		"day":       holiday.Day,
		"month":     holiday.Month,
		"year":      holiday.Year,
		"shortened": holiday.Shortened,
	}, nil
}
