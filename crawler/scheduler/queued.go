package scheduler

import (
	"richie.com/richie/learn_golang/crawler/engine"
)

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueuedScheduler) WorkerChan(
	) chan engine.Request {
	return make(chan  engine.Request)
}


func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(
	w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) ConfigureMasterWorkerChan(r chan engine.Request) {
	panic("implement me")
}

func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan  engine.Request)
	s.requestChan = make(chan engine.Request)

	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {

			var actiteRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				actiteRequest = requestQ[0]
			}

			select {

			case r := <-s.requestChan:
				// shend r to a ? worker
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				// send ? next_request to w
				workerQ = append(workerQ, w)
			case activeWorker <- actiteRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
