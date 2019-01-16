package Service

import (
	"github.com/umirode/prom-calendar-russia/src/Domain/Model/Entity"
	"github.com/umirode/prom-calendar-russia/src/Domain/Repository"
	"github.com/umirode/prom-calendar-russia/src/Domain/Service/DTO"
)

type HolidayService struct {
	holidayRepository Repository.IHolidayRepository
}

func NewHolidayService(holidayRepository Repository.IHolidayRepository) *HolidayService {
	return &HolidayService{
		holidayRepository: holidayRepository,
	}
}

func (s *HolidayService) GetOneByYearMonthAndDay(year uint, month uint, day uint) (*Entity.Holiday, error) {
	holiday, err := s.holidayRepository.FindOneByDayMonthAndYear(day, month, year)
	if err != nil {
		return nil, err
	}

	return holiday, nil
}

func (s *HolidayService) GetAllByYearAndMonth(year uint, month uint, withShortened bool) ([]*Entity.Holiday, error) {
	holidays, err := s.holidayRepository.FindAllByYearAndMonth(month, year, withShortened)
	if err != nil {
		return nil, err
	}

	return holidays, nil
}

func (s *HolidayService) GetAllByYear(year uint, withShortened bool) ([]*Entity.Holiday, error) {
	holidays, err := s.holidayRepository.FindAllByYear(year, withShortened)
	if err != nil {
		return nil, err
	}

	return holidays, nil
}

func (s *HolidayService) GetAll(withShortened bool) ([]*Entity.Holiday, error) {
	holidays, err := s.holidayRepository.FindAll(withShortened)
	if err != nil {
		return nil, err
	}

	return holidays, nil
}

func (s *HolidayService) CreateIfNotExists(holidayDTO *DTO.HolidayDTO) error {
	holiday, err := s.GetOneByYearMonthAndDay(holidayDTO.Year, holidayDTO.Month, holidayDTO.Day)
	if err != nil {
		return err
	}

	if holiday != nil {
		return nil
	}

	return s.holidayRepository.Save(&Entity.Holiday{
		Day:       holidayDTO.Day,
		Month:     holidayDTO.Month,
		Year:      holidayDTO.Year,
		Shortened: holidayDTO.Shortened,
	})
}
