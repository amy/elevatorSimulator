package main

import (
	"math/rand"
	"sync"
	"time"

	"github.com/amy/elevatorSimulator/elevator"
	"github.com/amy/elevatorSimulator/simulator"
)

var (
	simulationDone = make(chan bool)
)

func main() {

	var wg sync.WaitGroup

	elevatorSystem := InitElevatorManager(3)
	ticker := time.NewTicker(time.Duration(1) * time.Second)
	wg.Add(1)

	// Start elevator system
	go simulator.StartStepper(elevatorSystem, ticker, &wg, simulationDone)

	// Send elevator system requests
	var requests = InitRequestChannel(4)
	go simulator.SendRequests(elevatorSystem, requests)

	wg.Wait()
}

// InitElevatorManager initializes an elevator manager with
// n elevators starting on floor 0
func InitElevatorManager(n int) *elevator.ElevatorManager {

	system := []*elevator.Elevator{}

	for i := 0; i < n; i++ {

		e := &elevator.Elevator{
			Id:              i,
			Floor:           0,
			Direction:       "up",
			PickupRequests:  [10][]elevator.Request{},
			DropoffRequests: [10][]elevator.Request{},
			Idle:            true,
		}

		system = append(system, e)
	}

	return &elevator.ElevatorManager{Topfloor: 10, Elevators: system}

}

// InitRequestChannel initializes the request channel with n requests
// starting from floor 0 and ending on a randomized floor
func InitRequestChannel(n int) chan simulator.Request {

	var requests = make(chan simulator.Request, n)

	for i := 0; i < n; i++ {
		rand.Seed(time.Now().UnixNano())
		dest := rand.Intn(9)

		requests <- elevator.Request{
			Id:          i,
			Start:       0,
			Destination: dest,
			Direction:   "up",
		}
	}
	close(requests)

	return requests

}
