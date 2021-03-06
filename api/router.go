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
	//"github.com/epinion-online-research/ant-worker/entity"
	//"net/http"
	"strconv"
	"github.com/epinion-online-research/ant-worker/manager"
)

type Router struct {
	engine  *gin.Engine
	rest    *Rest
	manager manager.Manager
}

func (router *Router ) Init (rest *Rest) {
	router.engine = gin.Default()
	router.rest = rest
	router.manager = rest.Manager
}

/*
type Json struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}*/

func (router *Router) RegisterJobs( ) {
	job := JobHandler {
		Job: router.manager.Job,
	}
	// Example
	router.engine.GET("/api/v1/test", job.Test)
	// Create new resource
	router.engine.POST("/api/v1/job", job.Create)
	// Retrieve all resources
	router.engine.GET("/api/v1/jobs", job.GetAll)
	// Retrieve one resource
	router.engine.GET("/api/v1/job/:id", job.Get)
	// Update all fields for single resource
	router.engine.PUT("/api/v1/job/:id", job.Update)
	// Update one or some fields for single resource
	//router.engine.PATCH("/api/v1/job/:id", job.PartlyUpdate)
	// Delete resource
	router.engine.DELETE("/api/v1/job/:id", job.Delete)
}

func (router Router) Listen() {
	router.engine.Run(router.rest.Bind + ":" + strconv.Itoa(router.rest.Port))
}
