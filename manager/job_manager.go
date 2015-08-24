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
	//"time"
)

type JobManager struct {
	observer chan entity.Job
	//ExampleChannel chan string

}

func (manager *JobManager) Init(){
	//Setup params
	//manager.ExampleChannel = make(chan string )



}


func (manager *JobManager ) Monitor(){

}
func (manager *JobManager ) Example(){

}

func( manager *JobManager ) NewJob( job entity.Job ){

}

func( manager *JobManager ) UpdateJob(){

}

func( manager *JobManager ) PauseJob(){

}

func( manager *JobManager ) DeleteJob(){

}

