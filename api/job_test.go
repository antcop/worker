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
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"io"
	"encoding/json"
	"bytes"
)

// END POINT TESTING

var job Job

var router = gin.Default()
type Json map[string] interface {}

// Create mockup HTTP Request
func makeMockupRequest(method, url string, data Json) *httptest.ResponseRecorder {
	query, _ := json.Marshal(data)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(query))
	request.Header.Set("Content-Type", "application/json")
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	return writer
}

// Convert string to JSON type
func toJson(data io.Reader) Json {
	decoder := json.NewDecoder(data)
	var json Json
	decoder.Decode(&json)
	return json
}

// GET /test
func TestJobTest( t *testing.T ) {
	assert := assert.New(t)
	router.GET("/test", job.Test)
	response := makeMockupRequest("GET", "/test", Json {})
	assert.NotNil(response)
	data := toJson(response.Body)
	assert.NotNil(data)
	status := data["status"].(bool)
	assert.Equal(true, status)
}

// POST api/v1/job
func TestCreateJob(t *testing.T) {
	assert := assert.New(t)
	router.POST("/api/v1/job", job.Test)
	response := makeMockupRequest("POST", "/api/v1/job", Json {})
	assert.NotNil(response)
	data := toJson(response.Body)
	assert.NotNil(data)
	status := data["status"].(bool)
	assert.Equal(true, status)
}

// GET api/v1/job
func TestGetAllJob(t *testing.T) {
	assert := assert.New(t)
	router.GET("/api/v1/job", job.GetAll)
	response := makeMockupRequest("GET", "/api/v1/job", Json {})
	assert.NotNil(response)
	data := toJson(response.Body)
	assert.NotNil(data)
	status := data["status"].(bool)
	assert.Equal(true, status)
}

// GET api/v1/job/:id
func TestGetJob(t *testing.T) {
	assert := assert.New(t)
	router.GET("/api/v1/job/:id", job.Get)
	response := makeMockupRequest("GET", "/api/v1/job/1234", Json {})
	assert.NotNil(response)
	data := toJson(response.Body)
	assert.NotNil(data)
	key := data["job"].(map[string] interface {})["key"]
	assert.Equal("1234", key)
}

// PUT api/v1/job/:id
func TestUpdateJob(t *testing.T) {
	assert := assert.New(t)
	router.PUT("/api/v1/job/:id", job.Update)
	response := makeMockupRequest("PUT", "/api/v1/job/1234", Json {})
	assert.NotNil(response)
	data := toJson(response.Body)
	assert.NotNil(data)
	key := data["job"].(map[string] interface {})["key"]
	assert.Equal("1234", key)
}

// PATCH api/v1/job/:id
func TestParlyUpdateJob(t *testing.T) {
	assert := assert.New(t)
	router.PATCH("/api/v1/job/:id", job.PartlyUpdate)
	response := makeMockupRequest("PATCH", "/api/v1/job/1234", Json {})
	assert.NotNil(response)
	data := toJson(response.Body)
	assert.NotNil(data)
	key := data["job"].(map[string] interface {})["key"]
	assert.Equal("1234", key)
}

// DELETE api/v1/job/:id
func TestDeleteJob(t *testing.T) {
	assert := assert.New(t)
	router.DELETE("/api/v1/job/:id", job.PartlyUpdate)
	response := makeMockupRequest("DELETE", "/api/v1/job/1234", Json {})
	assert.NotNil(response)
	data := toJson(response.Body)
	assert.NotNil(data)
	key := data["job"].(map[string] interface {})["key"]
	assert.Equal("1234", key)
}