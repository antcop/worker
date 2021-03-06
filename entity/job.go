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

package entity

type JobApi struct {
	Name string
	Type string
	Description string
	Callback string
	Endpoint map[string] interface {}
	Params map[string] interface {}
	Workload []map[string] interface {}
}

type EndPoint struct {
	Id int
	Url string
	Method string
	Data string
}

type WorkloadUnit map[string] interface {}

type Workload struct {
	Id int
	Data [] WorkloadUnit
}

type Job struct {
	Id uint `gorm:"primary_key"' sql:"AUTO_INCREMENT"`
	Name string
	Type string
	Description string
	Callback string
	Status int
	Progress int
	//EndPoint EndPoint `gorm:"many2many:job_endpoint;"`
	//Params string
	//Workload []Workload `gorm:"many2many:job_workloads;"`
}
