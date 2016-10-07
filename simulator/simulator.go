package simulator

import (
	"sync"
	"time"
)

// Stepper is an interface that steps the Stepper system forward
type Stepper interface {
	Step()
	Snapshot()
}

// Scheduler schedules requests to the Scheduler system
type Scheduler interface {
	Schedule(r Request)
}

type Request interface{}

// StartStepper starts steps simulation at the interval provided by ticker
func StartStepper(stpr Stepper, ticker *time.Ticker, wg *sync.WaitGroup, done chan bool) {
	defer wg.Done()

	for {
		select {
		case <-ticker.C:
			stpr.Snapshot()
			stpr.Step()
		case <-done:
			ticker.Stop()
			return
		}
	}
}

// SendRequests sends schedule requests at 3 second intervals
func SendRequests(shdler Scheduler, requests <-chan Request) {
	// @TODO inject time + interval

	for r := range requests {
		shdler.Schedule(r)
		time.Sleep(time.Duration(3) * time.Second)
	}

	// Needs to be placed somewhere else
	//simulationDone <- true
}
