package scheduler

import "crawler/engine"

type ConcurrentEngine struct {
	Scheduler engine.Scheduler
	WorkCount int
}

type SimpleScheduler struct {
	WorkerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() {
		s.WorkerChan <- request
	}()
}

func (s *SimpleScheduler) ConfigWorkChan(c chan engine.Request) {
	s.WorkerChan = c
}
