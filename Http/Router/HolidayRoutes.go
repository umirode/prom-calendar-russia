package Router

import (
	"github.com/umirode/prom-calendar-russia/Http/Controller"
	"github.com/umirode/prom-calendar-russia/src/Application/Service"
	"github.com/umirode/prom-calendar-russia/src/Infrastructure/Persistence/Gorm/Repository"
)

func (r *Router) setHolidayRoutes() {
	holidayRepository := Repository.NewHolidayRepository(r.Database)

	holidayController := Controller.NewHolidayController(Service.NewHolidayService(holidayRepository))

	holidayGroup := r.Router.Group("/holidays")

	holidayGroup.GET("/:year/:month/:day", holidayController.FindByYearMonthAndDay)
	holidayGroup.GET("/:year/:month", holidayController.FindByYearAndMonth)
	holidayGroup.GET("/:year", holidayController.FindByYear)
	holidayGroup.GET("", holidayController.Find)
}
