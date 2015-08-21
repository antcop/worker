package daemon

import (
	"os"
	"log"
	"github.com/epinion-online-research/service"
)

var logger service.Logger

type Daemon struct {
	Name string
	Description string
	Port int
	OnStart func(daemon Daemon)
	OnStop func(daemon Daemon)
}

func (daemon Daemon) Print(message string) {
	logger.Infof(message)
}

func (daemon Daemon) GetName() string {
	return daemon.Name
}

func (daemon Daemon) GetDescription() string {
	return daemon.Description
}

func (daemon Daemon) GetPort() int {
	return daemon.Port
}

func (daemon Daemon) Start(s service.Service) error {
	logger.Infof("On Start")
	go daemon.OnStart(daemon)
	return nil
}

func (daemon Daemon) Stop(s service.Service) error {
	logger.Infof("On Stop")
	daemon.OnStop(daemon)
	return nil
}

func (daemon Daemon) Run(bgMode bool) {
	config := &service.Config {
		Name: daemon.GetName(),
		DisplayName: daemon.GetName(),
		Description: daemon.GetDescription(),
	}
	prg := &daemon
	s, err := service.New(prg, config)
	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) > 1 {
		err = service.Control(s, os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	if (bgMode) {
		err = s.Run()
		if err != nil {
			logger.Error(err)
		}
	}
}