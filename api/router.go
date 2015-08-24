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

package api

import (
	"github.com/gin-gonic/gin"
	//"github.com/epinion-online-research/ant-worker/manager"
)

type Router struct {
	engine *gin.Engine
	//rest *RestServer
	observer chan string

}

func (router *Router ) Init ( observer chan string  ){
	router.engine = gin.Default()
	//router.rest = rest
	router.observer = observer

}

func example( c *gin.Context ){
	//Do something
}

func (router *Router) RegisterJobs( ) {

	router.engine.GET("/example", func (c *gin.Context ){
			//Notify channel
			println(" Before.............");

			router.observer <- "Pingggggggggg"

			//router.rest.JobManager.Observer <- "Pingggggggggggggggggggg"
			//println( <- router.rest.JobManager.Observer );
	})


	/*
	// Retrieve all resources
	router.engine.GET("/api/v1/jobs", getJobs )

	// Retrieve one resource
	router.engine.GET("/api/v1//job/:id", getJob)

	// Create new resource
	router.engine.POST("/api/v1/job", createJob)

	// Update all fields for single resource
	router.engine.PUT("/api/v1/job/:id", updateJob)
	// Update one or some fields for single resource
	router.engine.PATCH("/api/v1/job/:id", partlyUpdateJob)
	// Delete resource
	router.engine.DELETE("/api/v1/job/:id", deleteJob)
	*/
}

func ( router Router ) Listen(){
	router.engine.Run( ":2345" );
}
