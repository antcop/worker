/**
 * Ant Worker Project
 *
 * Copyright (c) 2015 Epinion Online Research Team
 *
 * --------------------------------------------------------------------
 *
 * This program is free software: you can redistribute it and/or
 * modify it under the terms of the GNU Affero General Public License
 * as published by the Free Software Foundation, either version 3
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public
 * License along with this program.
 * If not, see <http://www.gnu.org/licenses/>.
 *
 * --------------------------------------------------------------------
 *
 * Author:
 *     Jerry Pham       <jerry@andjerry.com>
 *     Loi Nguyen       <loint@penlook.com>
 */

package daemon

import (
	"os"
	"log"
	"github.com/kardianos/service"
	. "github.com/epinion-online-research/ant-worker/manager"
)

var logger service.Logger

type Daemon struct {
	Name string
	Description string
	Bind string
	Port int
	OnStart func(daemon Daemon)
	OnStop func(daemon Daemon)
	Manager Manager
}

func (daemon Daemon) Println(message string) {
	logger.Infof(message)
}

func (daemon Daemon) GetName() string {
	return daemon.Name
}

func (daemon Daemon) GetDescription() string {
	return daemon.Description
}

func (daemon Daemon) GetBind() string {
	return daemon.Bind
}

func (daemon Daemon) GetPort() int {
	return daemon.Port
}

func (daemon Daemon) Start(s service.Service) error {
	go daemon.OnStart(daemon)
	return nil
}

func (daemon Daemon) Stop(s service.Service) error {
	daemon.OnStop(daemon)
	return nil
}

func (daemon Daemon) RunAsService(bgMode bool) {
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