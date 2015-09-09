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

package main

import (
	. "github.com/epinion-online-research/ant-worker/manager"
	"github.com/epinion-online-research/ant-worker/api"
	. "github.com/epinion-online-research/ant-worker/module/daemon"
	. "github.com/epinion-online-research/ant-worker/module"
)

func startService(daemon Daemon) {
	
	manager := daemon.Manager
	restServer := api.Rest {
		Port: manager.Job.Module.Config.Port,
		Manager: manager,
	}
	restServer.Start()
}

func stopService(daemon Daemon) {
	daemon.Println("Stop Daemon")
}

func main() {
	manager := Manager {}
	module := Module {}
	// Loading standard modules
	module.Load()
	// Inject modules for managers can access global resources
	manager.Job = JobManager {
		Module: module,
	}
	manager.Worker = WorkerManager {
		Module: module,
	}
	// Daemon process as linux service
	Daemon {
		Name: module.Config.Name,
		Description: module.Config.Description,
		Bind: module.Config.Bind,
		Port: module.Config.Port,
		OnStart: startService,
		OnStop: stopService,
		Manager: manager,
	}.RunAsService(true)
}
