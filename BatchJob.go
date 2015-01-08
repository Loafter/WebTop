package main

import "errors"

type BatchJob struct {
	Job    func()
	runJob bool
	done   chan bool
}

func (batchJob *BatchJob) Start() error {
	if batchJob.runJob {

		return errors.New("error: can't stop not stopped job")
	}
	if batchJob.Job == nil {
		return errors.New("error: empty job function")
	}

	if !batchJob.runJob {
		batchJob.done = make(chan bool, 1)
	}

	go batchJob.execution(batchJob.done)
	batchJob.runJob = true
	return nil
}
func (batchJob *BatchJob) IsRunning() bool {
	return batchJob.runJob

}

func (batchJob *BatchJob) Stop() error {
	if !batchJob.runJob {
		return errors.New("error: can't stop not stopted job")
	}
	batchJob.runJob = false
	isDone := <-batchJob.done
	if isDone {
		close(batchJob.done)
		return nil
	}
	return errors.New("error: failed stop job")
}

func (batchJob *BatchJob) execution(done chan bool) {
	for {
		if batchJob.runJob {
			batchJob.Job()
		} else {
			done <- true
			return
		}

	}
}
