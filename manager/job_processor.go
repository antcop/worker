package  manager

import (
	"github.com/epinion-online-research/ant-worker/entity"
)

type JobProcessor struct {
	Job *entity.Job
	Workers []entity.Worker
	Config *entity.Config
}

func ( processor *JobProcessor ) PickWorker() {

}

//Process a job by workers
func( processor *JobProcessor ) Process() {
	println( "Processing: " + processor.Job.Name )
}

