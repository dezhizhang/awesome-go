package scheduler

import "crawler/engine"

type ConcurrentEngine struct {
	Scheduler engine.Scheduler
	WorkCount int
}

type SimpleScheduler struct {
	WorkerChan chan engine.Request
}

func (s *SimpleScheduler) Run() {
	s.WorkerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) WorkReady(chan engine.Request) {
	return
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() {
		s.WorkerChan <- request
	}()
}

func (s *SimpleScheduler) WorkChan() chan engine.Request {
	return s.WorkerChan
}

func (s *SimpleScheduler) ConfigWorkChan(c chan engine.Request) {
	s.WorkerChan = c
}
