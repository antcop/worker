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
	Handler *gin.Engine
}

func (router Router) GetHandler() *gin.Engine {
	return router.Handler
}

func (router Router) Register() {
	api := router.GetHandler()

	// Create new resource
	api.POST("/job", createJob)
	// Retrieve all resources
	api.GET("/job", getJobs)
	// Retrieve one resource
	api.GET("/job/:id", getJob)
	// Update all fields for single resource
	api.PUT("job/:id", updateJob)
	// Update one or some fields for single resource
	api.PATCH("job/:id", partlyUpdateJob)
	// Delete resource
	api.DELETE("job/:id", deleteJob)
}