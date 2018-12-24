package Parser

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/umirode/prom-calendar-russia/src/Domain/Model/Entity"
	"io"
	"os"
	"strconv"
	"strings"
)

type CSVCalendarParser struct{}

func NewCSVCalendarParser() *CSVCalendarParser {
	return &CSVCalendarParser{}
}

func (p *CSVCalendarParser) Parse(file *os.File, startYear uint) ([]*Entity.Holiday, error) {
	holidays := make([]*Entity.Holiday, 0)

	reader := csv.NewReader(bufio.NewReader(file))

	firstLine := true
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		if !firstLine {
			year, err := strconv.Atoi(line[0])
			if err != nil {
				return nil, err
			}

			logrus.Println(fmt.Sprintf("Year: %d", year))

			if uint(year) < startYear {
				continue
			}

			for month := 1; month <= 12; month++ {
				logrus.Println(fmt.Sprintf("Year: %d, Month: %d", year, month))
				daysStringArray := strings.Split(line[month], ",")
				for _, dayString := range daysStringArray {
					logrus.Println(fmt.Sprintf("Year: %d, Month: %d, Day: %s", year, month, dayString))

					if strings.Contains(dayString, "*") ||strings.Contains(dayString, "+") {
						continue
					}

					day, err := strconv.Atoi(dayString)
					if err != nil {
						return nil, err
					}

					holidays = append(holidays, &Entity.Holiday{
						Day:   uint(day),
						Month: uint(month),
						Year:  uint(year),
					})
				}
			}
		}

		firstLine = false
	}

	return holidays, nil
}
