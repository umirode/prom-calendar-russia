package Http

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/umirode/prom-calendar-russia/Config"
	"github.com/umirode/prom-calendar-russia/Module/Http/Router"
)

type Module struct {
}

func NewModule() *Module {
	return &Module{}
}

func (m *Module) Init() {
	go func() {
		/**
		Get server address
		*/
		serverConfig := Config.GetServerConfig()
		serverAddress := fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port)

		/**
		Start server
		*/
		logrus.Fatal(Router.NewRouter(serverConfig.Debug).Router.Start(serverAddress))
	}()
}
