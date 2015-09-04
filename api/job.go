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
	"github.com/epinion-online-research/ant-worker/manager"
	"net/http"
)

const MaxListJob int = 20

type Job struct {
	manager *manager.JobManager
}

// GET /test
func (job Job) Test(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H {"status": true})
}

// POST api/v1/job
func (job Job) Create(context *gin.Context ) {
	context.JSON(200, gin.H {
		"status": "Retrieve one job",
	})
}

// GET /api/v1/job
func (job Job) GetAll(context *gin.Context ) {
	var jobList = make([]gin.H, 20)
	jobList = append(jobList, gin.H {
		"key" : "12355",
		"name" : " Do some thing",
		"status": 1,
		"progress": 90,
		"estimate": 25,
	})
	context.JSON(200, gin.H {
		"status": true,
		"jobs": jobList,
	})
}

// GET /api/v1/job/:id
func (job Job) Get(context *gin.Context) {
	context.JSON(200, gin.H {
		"status": true,
		"job": gin.H {
			"key" : context.Param("key"),
			"name" : " Do some thing",
			"status": 1,
			"progress": 90,
			"estimate": 25,
		},
	})
}

func (job Job) Update(context *gin.Context) {
	context.JSON(200, gin.H {
		"status": "Update one job",
	})
}

func (job Job) PartlyUpdate(context *gin.Context) {
	context.JSON(200, gin.H {
		"status": "Update parly job",
	})
}

func (job Job) Delete(context *gin.Context) {
	context.JSON(200, gin.H {
		"status": "Delete job",
	})
}
