package Repository

import (
	"github.com/jinzhu/gorm"
	"github.com/umirode/prom-calendar-russia/src/Domain/Model/Entity"
	"github.com/umirode/prom-calendar-russia/src/Infrastructure/Persistence/Gorm/Converter"
	"github.com/umirode/prom-calendar-russia/src/Infrastructure/Persistence/Gorm/Model"
)

type HolidayRepository struct {
	db               *gorm.DB
	holidayConverter *Converter.HolidayConverter
}

func NewHolidayRepository(db *gorm.DB) *HolidayRepository {
	return &HolidayRepository{
		holidayConverter: &Converter.HolidayConverter{},
		db:               db,
	}
}

func (r *HolidayRepository) Save(holiday *Entity.Holiday) error {
	model, _ := r.holidayConverter.ToDatabaseEntity(holiday)
	if model == nil {
		return nil
	}

	r.db.Save(model)

	return nil
}

func (r *HolidayRepository) FindAll() ([]*Entity.Holiday, error) {
	models := make([]*Model.Holiday, 0)

	r.db.Find(&models)

	entities := make([]*Entity.Holiday, 0)
	for _, model := range models {
		entity, _ := r.holidayConverter.ToAppEntity(model)

		entities = append(entities, entity.(*Entity.Holiday))
	}

	return entities, nil
}

func (r *HolidayRepository) FindAllByYear(year uint) ([]*Entity.Holiday, error) {
	models := make([]*Model.Holiday, 0)

	r.db.Where("year = ?", year).Find(&models)

	entities := make([]*Entity.Holiday, 0)
	for _, model := range models {
		entity, _ := r.holidayConverter.ToAppEntity(model)

		entities = append(entities, entity.(*Entity.Holiday))
	}

	return entities, nil
}

func (r *HolidayRepository) FindAllByYearAndMonth(month uint, year uint) ([]*Entity.Holiday, error) {
	models := make([]*Model.Holiday, 0)

	r.db.Where("year = ? and month = ?", year, month).Find(&models)

	entities := make([]*Entity.Holiday, 0)
	for _, model := range models {
		entity, _ := r.holidayConverter.ToAppEntity(model)

		entities = append(entities, entity.(*Entity.Holiday))
	}

	return entities, nil
}

func (r *HolidayRepository) FindOneByDayMonthAndYear(day uint, month uint, year uint) (*Entity.Holiday, error) {
	model := &Model.Holiday{}

	r.db.Where("day = ? and month = ? and year = ?", day, month, year).First(model)
	if r.db.NewRecord(model) {
		return nil, nil
	}

	entity, _ := r.holidayConverter.ToAppEntity(model)
	if entity == nil {
		return nil, nil
	}

	return entity.(*Entity.Holiday), nil
}
