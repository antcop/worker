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
	"testing"
	"github.com/stretchr/testify/assert"
	. "github.com/epinion-online-research/ant-worker/util"
	. "github.com/epinion-online-research/ant-worker/manager"
	. "github.com/epinion-online-research/ant-worker/module"
	"github.com/epinion-online-research/ant-worker/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	//"fmt"
	//"log"
	"bytes"
	"strconv"
	"os"
)

// Create mockup HTTP Request
func makeMockupRequest(router *gin.Engine, method, url string, data Json) *httptest.ResponseRecorder {
	query, _ := json.Marshal(data)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(query))
	request.Header.Set("Content-Type", "application/json")
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	return writer
}

func GetJobManager(assert *assert.Assertions) JobManager {
	module := Module {}
	pwd := os.Getenv("PWD")
	configFile := pwd + "/../test.conf"
	_, err := os.Stat(configFile)
	assert.Nil(err)
	module.Load(configFile)
	return JobManager {
		Module: module,
	}
}

func SetupJob(assert *assert.Assertions, manager JobManager) {
	db := manager.Module.Sql.Db
	db.CreateTable(&entity.Job{})
	count := 1
	db.Model(entity.Job{}).Count(&count)
	assert.Equal(0, count)
}

func TeardownJob(assert *assert.Assertions, manager JobManager) {
	db := manager.Module.Sql.Db
	db.DropTable(&entity.Job{})
	db.Close()
}

func GetJobHandler(manager JobManager) JobHandler {
	return JobHandler {
		Job: manager,
	}
}

func TestJobCrudApi(t *testing.T) {
	assert := assert.New(t)
	// Job Manager
	manager := GetJobManager(assert)
	// Job Handler
	handler := GetJobHandler(manager)
	// Setup job
	SetupJob(assert, manager)
	// Test Api
	ApiTestJob(assert, handler)
	// Create three jobs
	jobId1 := ApiCreateJob(assert, handler)
	jobId2 := ApiCreateJob(assert, handler)
	jobId3 := ApiCreateJob(assert, handler)
	assert.Equal(1, jobId1)
	assert.Equal(2, jobId2)
	assert.Equal(3, jobId3)
	// Get job
	ApiGetJob(assert, handler, jobId1)
	ApiGetJob(assert, handler, jobId2)
	ApiGetJob(assert, handler, jobId3)
	// Get all jobs
	ApiGetAllJobs(assert, handler)
	// Update job
	ApiUpdateJob(assert, handler, jobId1)
	ApiUpdateJob(assert, handler, jobId2)
	ApiUpdateJob(assert, handler, jobId3)
	// Delete 2 jobs
	ApiDeleteJob(assert, handler, jobId1)
	ApiDeleteJob(assert, handler, jobId2)
	ApiDeleteJob(assert, handler, jobId3)
	// Teardown
	TeardownJob(assert, manager)
}

// TEST api/v1/test
func ApiTestJob(assert *assert.Assertions, handler JobHandler) {
	var router = gin.Default()
	router.GET("/api/v1/test", handler.Test)
	response := makeMockupRequest(router, "GET", "/api/v1/test", Json {})
	assert.NotNil(response)
	data := ToJsonObject(response.Body)
	assert.NotNil(data)
	status := data["status"].(bool)
	assert.Equal(true, status)
}

// POST api/v1/job
func ApiCreateJob(assert *assert.Assertions, handler JobHandler) int {
	var router = gin.Default()
	router.POST("/api/v1/job", handler.Create)
	response := makeMockupRequest(router, "POST", "/api/v1/job", Json {
		"name": "sendsms",
		"description": "Send SMS Message",
		"type": "api_call",
		"callback": "http://google.com",
	})
	assert.NotNil(response)
	data := ToJsonObject(response.Body)
	assert.NotNil(data)
	id := data["id"].(float64)
	assert.Equal(true, id > 0)
	return int(id)
}

// GET api/v1/job
func ApiGetAllJobs(assert *assert.Assertions, handler JobHandler) {
	var router = gin.Default()
	router.GET("/api/v1/job", handler.GetAll)
	response := makeMockupRequest(router, "GET", "/api/v1/job", Json {})
	assert.NotNil(response)
	data := ToJsonArray(response.Body)
	assert.NotNil(data)
	assert.Equal(3, len(data))
	for i:=0; i<len(data); i++ {
		row := data[i]
		assert.Equal("sendsms", row["name"].(string))
		assert.Equal("http://localhost/api/job/v1/" + strconv.Itoa(int(row["id"].(float64))), row["location"].(string))
	}
}

// GET api/v1/job/:id
func ApiGetJob(assert *assert.Assertions, handler JobHandler, jobId int) Json {
	var router = gin.Default()
	router.GET("/api/v1/job/:id", handler.Get)
	response := makeMockupRequest(router, "GET", "/api/v1/job/" + strconv.Itoa(jobId), Json {})
	assert.NotNil(response)
	data := ToJsonObject(response.Body)
	assert.NotNil(data)
	return data
}

// PUT api/v1/job/:id
func ApiUpdateJob(assert *assert.Assertions, handler JobHandler, jobId int) {
	var router = gin.Default()
	router.PUT("/api/v1/job/:id", handler.Update)
	response := makeMockupRequest(router, "PUT", "/api/v1/job/" + strconv.Itoa(jobId),
	Json {
		"name" : "testname",
	})
	assert.NotNil(response)
	data := ToJsonObject(response.Body)
	assert.NotNil(data)
	assert.Equal(true, data["status"].(bool))
	// Verify resource
	data = ApiGetJob(assert, handler, jobId)
	assert.Equal("testname", data["name"].(string))
}

// DELETE api/v1/job/:id
func ApiDeleteJob(assert *assert.Assertions, handler JobHandler, jobId int) {
	var router = gin.Default()
	router.DELETE("/api/v1/job/:id", handler.Delete)
	response := makeMockupRequest(router, "DELETE", "/api/v1/job/" + strconv.Itoa(jobId), Json {})
	assert.NotNil(response)
	data := ToJsonObject(response.Body)
	assert.NotNil(data)
}