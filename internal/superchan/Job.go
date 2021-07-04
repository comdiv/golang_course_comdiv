package superchan

import (
	"fmt"
	"sync/atomic"
)

type Job struct {
	Id     int32
	Cancel func()
	Wait   func()
	OnFinish func()
}

// Finish stops, wait and optionally cleanup job
func (j *Job) Finish(){
	j.Cancel()
	j.Wait()
	if j.OnFinish!=nil {
		j.OnFinish()
	}
	fmt.Println("job finished")
}

var jobId int32
func NextJobId() int32 {
	return atomic.AddInt32(&jobId, 1)
}