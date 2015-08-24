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
	"github.com/epinion-online-research/ant-worker/manager"
	"sync"
	"github.com/epinion-online-research/ant-worker/api"
	"time"

)

func main() {

	var wg sync.WaitGroup
	wg.Add( 1 )

	//JobManager
	manager := manager.JobManager{
		Observer: make( chan string ),
	}

	//manager.Init()

	//observer := make( chan string )

	//Rest Service
	rs := api.RestServer  {
		Port: "2345",
		JobManager: &manager,
		//Observer: observer,
	}

	rs.Start()


	//Redis Service
	//TODO


	//Socket Service
	//TODO

	//...


	//Start monitoring input from services
	//go manager.Monitor()

	go func() {
		for {
			select {
			case msg := <- manager.Observer :
				go func() {
					println("Wowwwwwww")
					time.Sleep( 1 * time.Second )
					println( msg )
				}()

			}
		}
	}()


	wg.Wait()




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
