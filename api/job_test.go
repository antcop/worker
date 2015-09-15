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
	"bytes"
	"os"
)

// END POINT TESTING

var router = gin.Default()

// Create mockup HTTP Request
func makeMockupRequest(method, url string, data Json) *httptest.ResponseRecorder {
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
	GetJobHandler(manager)
	//handler := GetJobHandler(manager)
	// Setup job
	SetupJob(assert, manager)
	// Test Api
	//ApiTestJob(assert, handler)
	// Create first job
	//ApiCreateJob(assert, handler)
	// Create second job
	//job2 := CreateJobApi(assert, handler)
	// Get all jobs
	//GetAllJobsApi(assert, handler)
	// Get job1
	//GetJobApi(assert, handler, int(job1.Id))
	//GetJobApi(assert, handler, int(job2.Id))
	// Partly update job 1
	//UpdateJobApi(assert, handler, int(job1.Id))
	// Delete 2 jobs
	//DeleteJobApi(assert, handler, int(job1.Id), int(job2.Id))
	// Teardown job
	TeardownJob(assert, manager)
}

func ApiTestJob(assert *assert.Assertions, handler JobHandler) {
	router.GET("/api/v1/test", handler.Test)
	response := makeMockupRequest("GET", "/api/v1/test", Json {})
	assert.NotNil(response)
	data := ToJson(response.Body)
	assert.NotNil(data)
	status := data["status"].(bool)
	assert.Equal(true, status)
}

/*
// GET /test
func TestJobTest( t *testing.T ) {
	assert := assert.New(t)
	manager := getJobManager(assert)

	handler := JobHandler {
		Job: manager,
	}
	defer manager.Module.Sql.Db.Close()
	
	
}

// POST api/v1/job
func TestCreateJob(t *testing.T) {
	/*assert := assert.New(t)
	handler := JobHandler {
		Job: getJobManager(assert),
	}
	
	router.POST("/api/v1/job", handler.Create)
	response := makeMockupRequest("POST", "/api/v1/job", Json {
		"name": "sendsms",
		"description": "Send SMS Message",
		"type": "api_call",
		"callback": "http://google.com",
	})
	assert.NotNil(response)
	data := ToJson(response.Body)
	assert.NotNil(data)
	fmt.Println(data)
	id := data["id"].(int)
	assert.Equal(true, id > 0)
}

// GET api/v1/job
func TestGetAllJob(t *testing.T) {
	assert := assert.New(t)
	handler := JobHandler {
		Job: getJobManager(assert),
	}
	router.GET("/api/v1/job", handler.GetAll)
	response := makeMockupRequest("GET", "/api/v1/job", Json {})
	assert.NotNil(response)
	//data := ToJson(response.Body)
	//assert.Equal(true, data != nil)
	//status := data["status"].(bool)
	//assert.Equal(true, status)
}

// GET api/v1/job/:id
func TestGetJob(t *testing.T) {
	assert := assert.New(t)
	handler := JobHandler {
		Job: getJobManager(assert),
	}
	router.GET("/api/v1/job/:id", handler.Get)
	response := makeMockupRequest("GET", "/api/v1/job/1234", Json {})
	assert.NotNil(response)
	data := ToJson(response.Body)
	assert.NotNil(data)
	key := data["key"]
	assert.Equal("1234", key)
}

// PUT api/v1/job/:id
func TestUpdateJob(t *testing.T) {
	assert := assert.New(t)
	handler := JobHandler {
		Job: getJobManager(assert),
	}
	router.PUT("/api/v1/job/:id", handler.Update)
	response := makeMockupRequest("PUT", "/api/v1/job/1234", Json {})
	assert.NotNil(response)
	data := ToJson(response.Body)
	assert.NotNil(data)
	key := data["key"]
	assert.Equal("1234", key)
}

// PATCH api/v1/job/:id
func TestParlyUpdateJob(t *testing.T) {
	assert := assert.New(t)
	handler := JobHandler {
		Job: getJobManager(assert),
	}
	router.PATCH("/api/v1/job/:id", handler.PartlyUpdate)
	response := makeMockupRequest("PATCH", "/api/v1/job/1234", Json {})
	assert.NotNil(response)
	data := ToJson(response.Body)
	assert.NotNil(data)
	key := data["key"]
	assert.Equal("1234", key)
}

// DELETE api/v1/job/:id
func TestDeleteJob(t *testing.T) {
	assert := assert.New(t)
	handler := JobHandler {
		Job: getJobManager(assert),
	}
	router.DELETE("/api/v1/job/:id", handler.PartlyUpdate)
	response := makeMockupRequest("DELETE", "/api/v1/job/1234", Json {})
	assert.NotNil(response)
	data := ToJson(response.Body)
	assert.NotNil(data)
	key := data["key"]
	assert.Equal("1234", key)
}*/