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
	"github.com/epinion-online-research/ant-worker/entity"
	//. "github.com/epinion-online-research/ant-worker/util"
	"strconv"
	"net/http"
)

const MaxListJob int = 20

type JobHandler struct {
	manager manager.JobManager
}

// GET /test
func (handler JobHandler) Test(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H {"status": true})
}

// POST api/v1/job
func (handler JobHandler) Create(context *gin.Context) {
	var jobData entity.JobApi
	if context.BindJSON(&jobData) == nil {
		job, err := handler.manager.Create(jobData)
		if err != nil {
			handler.Error(context, http.StatusOK, 1200, "Internal Server Error")
			return
		}
		context.JSON(http.StatusOK, gin.H {
			"status": true,
			"id": job.Id,
			"location": "http://localhost/api/job/v1/" + strconv.Itoa(job.Id),
		})
		return
	}
	handler.Error(context, http.StatusOK, 1200, "Invalid json format")
}

// GET /api/v1/job
func (handler JobHandler) GetAll(context *gin.Context ) {
	var jobList = make([]gin.H, 20)
	for i:=0; i<=100; i++ {
		jobList = append(jobList, gin.H {
			"key" : "12355",
			"name" : " Do some thing",
			"status" : 1,
			"progress" : 90,
			"estimate" : 25,
		})
	}
	context.JSON(200, gin.H {
		"status": true,
		"jobs": jobList,
	})
}

// GET /api/v1/job/:id
func (handler JobHandler) Get(context *gin.Context) {
	context.JSON(200, gin.H {
		"status": false,
		"job": gin.H {
			"key" : context.Param("id"),
			"name" : " Do some thing",
			"status": 1,
			"progress": 90,
			"estimate": 25,
		},
	})
}

func (handler JobHandler) Update(context *gin.Context) {
	context.JSON(200, gin.H {
		"status": false,
		"job": gin.H {
			"key" : context.Param("id"),
			"name" : " Do some thing",
			"status": 1,
			"progress": 90,
			"estimate": 25,
		},
	})
}

func (handler JobHandler) PartlyUpdate(context *gin.Context) {
	context.JSON(200, gin.H {
		"status": false,
		"job": gin.H {
			"key" : context.Param("id"),
			"name" : " Do some thing",
			"status": 1,
			"progress": 90,
			"estimate": 25,
		},
	})
}

func (handler JobHandler) Delete(context *gin.Context) {
	context.JSON(200, gin.H {
		"status": false,
		"job": gin.H {
			"key" : context.Param("id"),
			"name" : " Do some thing",
			"status": 1,
			"progress": 90,
			"estimate": 25,
		},
	})
}

func (handler JobHandler) Error(context *gin.Context, status int, code int, message string) {
	context.JSON(http.StatusOK, gin.H {
		"status": false,
		"error": gin.H {
			"code": code,
			"message": message,
		},
	})
}