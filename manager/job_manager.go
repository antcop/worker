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
	"github.com/epinion-online-research/ant-worker/entity"
	. "github.com/epinion-online-research/ant-worker/module"
	"fmt"
)

type JobManager struct {
	Observer chan string
	Module Module
	JobProcessors [] JobProcessor
}

const STATUS_CREATED  int = 1;
const STATUS_PAUSED   int = 2;
const STATUS_STOPPED  int = 3;
const STATUS_PENDING  int = 4;
const STATUS_PROGRESS int = 5;
const STATUS_FINISHED int = 6;

const PROGRESS_START  int = 1;
const PROGRESS_DONE   int = 2;

func (manager *JobManager) Create(data entity.JobApi) (*entity.Job, error) {
	db := manager.Module.Sql.Db
	job := entity.Job {
		Name: data.Name,
		Description: data.Description,
		Type: data.Type,
		Callback: data.Callback,
		Status: STATUS_CREATED,
	}
	err := db.Create(&job)
	if err != nil {
		panic(err)
	}
	fmt.Println("CREATED ", job.Id)
	return &job, nil
}

func( manager *JobManager ) GetAll() ([]entity.Job, error) {
	db := manager.Module.Sql.Db
	var jobs []entity.Job
	db.Find(&jobs)
	return jobs, nil
}

func (manager *JobManager) Get(jobId int) (entity.Job, error) {
	db := manager.Module.Sql.Db
	var job entity.Job
	db.Select("*").Where("id = ?", jobId).Find(&job)
	return job, nil
}

func( manager *JobManager ) Update(jobId int, data entity.JobApi) (entity.Job, error) {
	var job entity.Job
	db := manager.Module.Sql.Db
	db.Select("*").Where("id = ?", jobId).First(&job)

	job.Name = data.Name
	job.Description = data.Description
	job.Type = data.Type
	job.Callback = data.Callback
	
	db.Save(&job)
	return job, nil
}

/*
func( manager *JobManager ) ParlyUpdate(data entity.JobApi) (entity.Job, error) {	
	return job, nil
}*/

func (manager *JobManager) Delete(jobId int) error {
	var job entity.Job
	db := manager.Module.Sql.Db
	db.Where("id = ?", jobId).Delete(&job)
	return nil
}

func (manager *JobManager) Start(jobId int)  (error) {
	job, err := manager.Get(jobId)
	if err != nil {
		return err
	}
	job.Status = STATUS_PROGRESS
	job.Progress = 0
	db := manager.Module.Sql.Db
	db.Save(&job)
	return nil
}

func (manager *JobManager) Pause(jobId int)  (error) {
	job, err := manager.Get(jobId)
	if err != nil {
		return err
	}
	job.Status = STATUS_PAUSED
	// job.Progress is not changed
	db := manager.Module.Sql.Db
	db.Save(&job)
	return nil
}

func (manager *JobManager) Resume(jobId int) (error) {
	job, err := manager.Get(jobId)
	if err != nil {
		return err
	}
	job.Status = STATUS_PROGRESS
	// job.Progress is not changed
	db := manager.Module.Sql.Db
	db.Save(&job)
	return nil
}

func (manager *JobManager) Stop(jobId int) (error) {
	job, err := manager.Get(jobId)
	if err != nil {
		return err
	}
	job.Status = STATUS_STOPPED
	// job.Progress is not changed
	db := manager.Module.Sql.Db
	db.Save(&job)
	return nil
}

func (manager *JobManager) Finish(jobId int) (error) {
	job, err := manager.Get(jobId)
	if err != nil {
		return err
	}
	job.Status   = STATUS_FINISHED
	job.Progress = PROGRESS_DONE
	db := manager.Module.Sql.Db
	db.Save(&job)
	return nil
}

func (manager *JobManager) Process(job entity.Job) {
	processor := JobProcessor {
		Job: job,
		Config: manager.Module.Config,
	}
	//append( manager.JobProcessors, processor)
	go processor.Process()
}

func (manager *JobManager ) Monitor() {
	go func() {
		for {
			select {
			case msg := <- manager.Observer :
				go func() {
					println( msg )
				}()
			}
		}
	}()
}



/*

func (manager *JobManager) Init(){

}

func( manager *JobManager ) UpdateJob(){

}

func( manager *JobManager ) PauseJob(){

}

func( manager *JobManager ) DeleteJob(){

}
*/

