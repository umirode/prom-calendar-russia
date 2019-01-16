package Command

import (
	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"github.com/umirode/prom-calendar-russia/src/Application/Parser"
	Service2 "github.com/umirode/prom-calendar-russia/src/Application/Service"
	"github.com/umirode/prom-calendar-russia/src/Domain/Service"
	"github.com/umirode/prom-calendar-russia/src/Domain/Service/DTO"
	"github.com/umirode/prom-calendar-russia/src/Infrastructure/Persistence/Gorm/Repository"
	"os"
	"strconv"
)

type CalendarCommand struct {
	Database *gorm.DB

	HolidayService Service.IHolidayService
}

func NewCalendarCommand(db *gorm.DB) *CalendarCommand {
	return &CalendarCommand{
		Database:       db,
		HolidayService: Service2.NewHolidayService(Repository.NewHolidayRepository(db)),
	}
}

func (c *CalendarCommand) GetCommand() *cobra.Command {
	mainCommand := c.getMainCommand()

	mainCommand.AddCommand(c.getUpdateCsvCommand())

	return mainCommand
}

func (c *CalendarCommand) getMainCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "calendar",
		Short: "",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
	}

	return command
}

func (c *CalendarCommand) getUpdateCsvCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "update_csv [file] [start year]",
		Short: "Parse calendar",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			csvFile, err := os.Open(args[0])
			if err != nil {
				logrus.Fatal(err)
			}

			startYear, err := strconv.Atoi(args[1])
			if err != nil {
				logrus.Fatal(err)
			}

			calendarParser := Parser.NewCSVCalendarParser()

			holidays, err := calendarParser.Parse(csvFile, uint(startYear))
			if err != nil {
				logrus.Fatal(err)
			}

			for _, holiday := range holidays {
				c.HolidayService.CreateIfNotExists(&DTO.HolidayDTO{
					Year:      holiday.Year,
					Month:     holiday.Month,
					Day:       holiday.Day,
					Shortened: holiday.Shortened,
				})
			}
		},
	}

	return command
}
