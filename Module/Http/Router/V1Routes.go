package Router

import (
	"github.com/umirode/prom-calendar-russia/Module/Http/Controller/v1"
	"github.com/umirode/prom-calendar-russia/src/Application/Service"
	"github.com/umirode/prom-calendar-russia/src/Infrastructure/Persistence/Gorm/Repository"
)

func (r *Router) setV1Routes() {

	holidayRepository := Repository.NewHolidayRepository()

	holidayController := v1.NewHolidayController(
		Service.NewHolidayService(
			holidayRepository,
		),
	)

	v1Routes := r.Router.Group("/v1")

	holidayRoutes := v1Routes.Group("/holidays")
	holidayRoutes.GET("/:year/:month/:day", holidayController.FindByYearMonthAndDay)
	holidayRoutes.GET("/:year/:month", holidayController.FindByYearAndMonth)
	holidayRoutes.GET("/:year", holidayController.FindByYear)
	holidayRoutes.GET("", holidayController.Find)
}
