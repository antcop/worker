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

package manager

import (
	"testing"
	"github.com/stretchr/testify/assert"
	. "github.com/epinion-online-research/ant-worker/module"
	 "github.com/epinion-online-research/ant-worker/entity"
	 "os"
)

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

func TestJobCrudCycle(t *testing.T) {
	return
	assert := assert.New(t)
	// Job Manager
	manager := GetJobManager(assert)
	// Setup job table
	SetupJob(assert, manager)
	// Create first job
	job1 := CreateJob(assert, manager)
	// Create second job
	job2 := CreateJob(assert, manager)
	// Get all jobs
	GetAllJobs(assert, manager)
	// Get job1
	GetJob(assert, manager, int(job1.Id))
	GetJob(assert, manager, int(job2.Id))
	// Partly update job 1
	UpdateJob(assert, manager, int(job1.Id))
	// Delete 2 jobs
	DeleteJob(assert, manager, int(job1.Id), int(job2.Id))
	// Teardown job
	TeardownJob(assert, manager)
}

func CreateJob(assert *assert.Assertions, manager JobManager) (entity.Job) {
	data := entity.JobApi {
		Name: "sendmail",
		Description: "Send Email By Using MailChimp",
		Type: "api_call",
		Callback: "http://example.com/callback",
	}
	job, err := manager.Create(data)
	// There is no error
	assert.Equal(true, err == nil)
	assert.Equal(true, job.Id > 0)
	return job
}

func GetAllJobs(assert *assert.Assertions, manager JobManager) [] entity.Job {
	jobs, err := manager.GetAll()
	assert.Equal(true, err == nil)
	// Count number of job
	assert.Equal(2, len(jobs))
	return jobs
}

func GetJob(assert *assert.Assertions, manager JobManager, jobId int) entity.Job {
	assert.Equal(true, jobId > 0)
	job, err := manager.Get(jobId)
	assert.Equal(true, err == nil)
	assert.Equal("sendmail", job.Name)
	assert.Equal("Send Email By Using MailChimp", job.Description)
	assert.Equal("api_call", job.Type)
	assert.Equal("http://example.com/callback", job.Callback)
	return job
}

/*
func PartlyUpdateJob(assert *assert.Assertions, manager JobManager, jobId int) (entity.Job) {
	var db entity.Job	
	job.Name = ""
	db.Save(&user)
	return job
}*/

func UpdateJob(assert *assert.Assertions, manager JobManager, jobId int) (entity.Job) {
	assert.Equal(true, jobId > 0)
	data := entity.JobApi {
		Name: "send_notification",
		Description: "Send Notification to Facebook",
		Type: "api_call2",
		Callback: "http://example.com/callback2",
	}
	job, err := manager.Update(jobId, data)
	assert.Equal(true, err == nil)
	assert.Equal("send_notification", job.Name)
	assert.Equal("Send Notification to Facebook", job.Description)
	assert.Equal("api_call2", job.Type)
	assert.Equal("http://example.com/callback2", job.Callback)
	return job
}

func DeleteJob(assert *assert.Assertions, manager JobManager, job1Id int, job2Id int) {
	assert.Equal(true, job1Id > 0)
	assert.Equal(true, job2Id > 0)
	err := manager.Delete(job1Id)
	assert.Equal(true, err == nil)
	err = manager.Delete(job2Id)
	assert.Equal(true, err == nil)
	jobs, err := manager.GetAll()
	assert.Equal(true, err == nil)
	// Count number of job
	assert.Equal(0, len(jobs))
}