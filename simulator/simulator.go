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

// delete this later
/*type Request struct {
	id          int
	start       int
	destination int
	direction   string
	state       string // starts out as pick up --> moves to dropoff
} */

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

// @TODO inject time + interval
func SendRequests(shdler Scheduler, requests <-chan Request) {
	// sanitize requests to be 0 < x < top floor

	for r := range requests {
		shdler.Schedule(r)
		time.Sleep(time.Duration(3) * time.Second)
	}

	// Needs to be placed somewhere else
	//simulationDone <- true
}
