package background_job

import (
	"crawler_go_facebook/pkg/logger"
	"fmt"
	"sync"
	"time"
)

type JobExecute func() error

type backgroundJob struct {
	chQuit chan bool
}

func newBackgroundJob() *backgroundJob {
	return &backgroundJob{
		chQuit: make(chan bool),
	}
}

func (m *backgroundJob) execute(name string, execute JobExecute, wg *sync.WaitGroup, interval time.Duration) {
	defer wg.Done()
	for {
		select {
		default:
			logger.InfoF("job=[%s] execute", name)

			if err := execute(); err != nil {
				logger.ErrorF("job=[%s]:execute error=[%s]", name, err.Error())
			}
			time.Sleep(interval)
		case <-m.chQuit:
			logger.InfoF(fmt.Sprintf("job=[%s]: stopping", name))
			return
		}
	}
}

func (m *backgroundJob) quit() {
	m.chQuit <- true
}

type BackgroundJobManager struct {
	jobs []*backgroundJob
}

func NewBackgroundJobManager() *BackgroundJobManager {
	return &BackgroundJobManager{jobs: make([]*backgroundJob, 0)}
}

func (m *BackgroundJobManager) AddJob(name string, execute JobExecute, wg *sync.WaitGroup, interval time.Duration) {
	job := newBackgroundJob()
	m.jobs = append(m.jobs, job)

	wg.Add(1)
	go job.execute(name, execute, wg, interval)
}

func (m *BackgroundJobManager) Shutdown() {
	logger.InfoF("background jobs shutdown !!!")
	var wg sync.WaitGroup

	for _, job := range m.jobs {
		wg.Add(1)
		go func(job *backgroundJob) {
			defer wg.Done()
			job.quit()
		}(job)
	}

	wg.Wait()
}
