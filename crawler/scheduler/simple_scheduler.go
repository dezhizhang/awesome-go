package scheduler

import "crawler/engine"

type Scheduler interface {
	Submit(request engine.Request)
	ConfigWorkChan(chan engine.Request)
}

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
}

type SimpleScheduler struct {
	WorkerChan chan engine.Request
}
