package v1

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/umirode/prom-calendar-russia/Module/Http/Controller"
	"github.com/umirode/prom-calendar-russia/src/Application/Hydrator"
	"github.com/umirode/prom-calendar-russia/src/Domain/Service"
)

type HolidayController struct {
	Controller.BaseController

	HolidayService  Service.IHolidayService
	HolidayHydrator *Hydrator.HolidayHydrator
}

func NewHolidayController(holidayService Service.IHolidayService) *HolidayController {
	controller := &HolidayController{
		HolidayService:  holidayService,
		HolidayHydrator: &Hydrator.HolidayHydrator{},
	}

	return controller
}

func (c *HolidayController) FindByYear(context echo.Context) error {
	year, err := c.GetParam(context, "year", "uint")
	if err != nil {
		return err
	}

	shortened := false
	shortenedString := context.QueryParam("shortened")
	if shortenedString != "" {
		shortened, err = strconv.ParseBool(shortenedString)
		if err != nil {
			return err
		}
	}

	holidays, err := c.HolidayService.GetAllByYear(year.(uint), shortened)
	if err != nil {
		return err
	}

	holidaysMapArray := make([]map[string]interface{}, 0)
	for _, holiday := range holidays {
		holidayMap, _ := c.HolidayHydrator.Extract(holiday)
		holidaysMapArray = append(holidaysMapArray, holidayMap)
	}

	return c.Response(context, http.StatusOK, holidaysMapArray, "")
}

func (c *HolidayController) Find(context echo.Context) error {
	shortened := false
	shortenedString := context.QueryParam("shortened")
	if shortenedString != "" {
		var err error
		shortened, err = strconv.ParseBool(shortenedString)
		if err != nil {
			return err
		}
	}

	holidays, err := c.HolidayService.GetAll(shortened)
	if err != nil {
		return err
	}

	holidaysMapArray := make([]map[string]interface{}, 0)
	for _, holiday := range holidays {
		holidayMap, _ := c.HolidayHydrator.Extract(holiday)
		holidaysMapArray = append(holidaysMapArray, holidayMap)
	}

	return c.Response(context, http.StatusOK, holidaysMapArray, "")
}

func (c *HolidayController) FindByYearAndMonth(context echo.Context) error {
	year, err := c.GetParam(context, "year", "uint")
	if err != nil {
		return err
	}
	month, err := c.GetParam(context, "month", "uint")
	if err != nil {
		return err
	}

	shortened := false
	shortenedString := context.QueryParam("shortened")
	if shortenedString != "" {
		shortened, err = strconv.ParseBool(shortenedString)
		if err != nil {
			return err
		}
	}

	holidays, err := c.HolidayService.GetAllByYearAndMonth(year.(uint), month.(uint), shortened)
	if err != nil {
		return err
	}

	holidaysMapArray := make([]map[string]interface{}, 0)
	for _, holiday := range holidays {
		holidayMap, _ := c.HolidayHydrator.Extract(holiday)
		holidaysMapArray = append(holidaysMapArray, holidayMap)
	}

	return c.Response(context, http.StatusOK, holidaysMapArray, "")
}

func (c *HolidayController) FindByYearMonthAndDay(context echo.Context) error {
	year, err := c.GetParam(context, "year", "uint")
	if err != nil {
		return err
	}
	month, err := c.GetParam(context, "month", "uint")
	if err != nil {
		return err
	}
	day, err := c.GetParam(context, "day", "uint")
	if err != nil {
		return err
	}

	holiday, err := c.HolidayService.GetOneByYearMonthAndDay(year.(uint), month.(uint), day.(uint))
	if err != nil {
		return err
	}

	holidayMap, _ := c.HolidayHydrator.Extract(holiday)

	return c.Response(context, http.StatusOK, holidayMap, "")
}
