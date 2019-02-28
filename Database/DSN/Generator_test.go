package DSN

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/umirode/prom-calendar-russia/Database"
)

func TestNewGenerator(t *testing.T) {
	assert.NotEmpty(t, NewGenerator(&Database.Config{}))
}

func TestDsnGenerator_GetDSN(t *testing.T) {
	drivers := [...]string{
		"mysql",
		"postgres",
		"sqlite3",
		"mssql",
	}

	config := &Database.Config{
		Database: "test",
	}

	for _, driver := range drivers {
		config.Driver = driver

		dsnGenerator := NewGenerator(config)

		assert.NotEmpty(t, dsnGenerator.GetDSN())
	}
}
