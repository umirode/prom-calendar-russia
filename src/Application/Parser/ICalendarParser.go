package Parser

import (
	"github.com/umirode/prom-calendar-russia/src/Domain/Model/Entity"
	"os"
)

type ICalendarParser interface {
	Parse(file *os.File, startYear uint) ([]*Entity.Holiday, error)
}
