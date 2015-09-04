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

package manager

import (
	"github.com/epinion-online-research/ant-worker/entity"
)

type JobManager struct {
	Observer chan string
	Config *entity.Config
	JobProcessors []JobProcessor
}

func (manager *JobManager ) Test() {
}

func (manager *JobManager ) ExampleAction( ex string ){
	manager.Observer <- "Example action executed. This is data from rest server: "  + ex
}

func( manager *JobManager ) NewJob( job *entity.Job ){
	//job := entity.JobInterface( job )
	manager.ProcessJob( job )
}

func (manager *JobManager ) ProcessJob( job *entity.Job ){
	processor := JobProcessor {
		Job: job,
		Config: manager.Config,
	}
	//append( manager.JobProcessors, processor )
	go processor.Process()
}

func (manager *JobManager ) Monitor() {
	go func() {
		for {
			select {
			case msg := <- manager.Observer :
				go func() {
					println( msg )
				}()
			}
		}
	}()
}



/*

func (manager *JobManager) Init(){

}

func( manager *JobManager ) UpdateJob(){

}

func( manager *JobManager ) PauseJob(){

}

func( manager *JobManager ) DeleteJob(){

}
*/

