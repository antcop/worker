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
)

type Router struct {
	handler *gin.Engine
}


func (router Router) RegisterJobs() {
	handler := router.GetHandler()
	router := router.handler.Default();

	// Create new resource
	router.POST("/api/v1/job", createJob)
	// Retrieve all resources
	router.GET("/api/v1//job", getJobs)
	// Retrieve one resource
	router.GET("/api/v1//job/:id", getJob)
	// Update all fields for single resource
	router.PUT("/api/v1/job/:id", updateJob)
	// Update one or some fields for single resource
	router.PATCH("/api/v1/job/:id", partlyUpdateJob)
	// Delete resource
	router.DELETE("/api/v1/job/:id", deleteJob)
}

func ( router Router ) Listen( port int ){
	router.handler.Default().Run( ":" + port );
}
