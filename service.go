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
	"github.com/epinion-online-research/ant-worker/entity"

)

func main() {

	var wg sync.WaitGroup
	wg.Add( 1 )

	//Configurations //TODO Read config from file
	config := entity.Config{
		MaxWorker: 10,
		MaxWorkerPerJob: 2,
	}

	//JobManager
	manager := manager.JobManager{
		Observer: make( chan string ),
		Config: &config,
	}

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
	manager.Monitor()


	wg.Wait()

}
