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
	. "github.com/epinion-online-research/ant-worker/daemon"
	. "github.com/epinion-online-research/ant-worker/manager"
	. "github.com/epinion-online-research/ant-worker/api"

)

func start(daemon Daemon) {
	manager := Manager {
		server : Server {
			port : daemon.GetPort(),
		}
	}
	daemon.Print("")
}

func stop(daemon Daemon) {
	
}


func main() {

	//Web server

	ws := api.RestApiServer{}
	ws.Start("1234");



	/*
	daemon := Daemon {
		Name: "antworker",
		Description: "Job server",
		Port : 1234,
		OnStart: start,
		OnStop:  stop,
	}
	daemon.Run(true)
	*/
}
