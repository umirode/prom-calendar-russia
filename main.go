package prom_calendar_russia

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/joho/godotenv"
	"github.com/umirode/prom-calendar-russia/Database"
	"github.com/umirode/prom-calendar-russia/Http/Router"
	"github.com/umirode/prom-calendar-russia/configs"
)

func main() {
	/**
	Load .env variables
	*/
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal(err.Error())
	}

	/**
	Create database connection
	*/
	databaseConfig := configs.GetDatabaseConfig()
	db, err := Database.NewDatabase(
		&Database.Config{
			Driver:   databaseConfig.Driver,
			Debug:    databaseConfig.Debug,
			Database: databaseConfig.Database,
			Host:     databaseConfig.Host,
			Port:     databaseConfig.Port,
			Username: databaseConfig.Username,
			Password: databaseConfig.Password,
			Params:   databaseConfig.Params,
		},
	)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	defer db.Close()

	/**
	Run migrations
	*/
	Database.RunMigrations(db)

	/**
	Get server address
	*/
	serverConfig := configs.GetServerConfig()
	serverAddress := fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port)

	/**
	Start server
	*/
	logrus.Fatal(Router.NewRouter(db, serverConfig.Debug).Router.Start(serverAddress))
}
