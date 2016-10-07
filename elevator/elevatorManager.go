package elevator

import (
	"fmt"

	"github.com/amy/elevatorSimulator/simulator"
)

type ElevatorManager struct {
	Topfloor  int
	Elevators []*Elevator
}

// Snapshot shows what floor all the elevators in the system are on
func (em *ElevatorManager) Snapshot() {

	fmt.Println("--- CURRENT STEP ---")

	for _, e := range em.Elevators {
		fmt.Printf("Elevator %v on floor %v\n", e.Id, e.Floor)
	}

	fmt.Println("--------------------")
}

// send request to most suitable elevator
func (em *ElevatorManager) Schedule(r simulator.Request) {

	var suitable *Elevator = nil
	elevators := em.Elevators

	switch r.(type) {
	case Request:
		// @TODO pull out this logic
		for _, e := range elevators {
			if suitable == nil {
				suitable = e
			}

			if e.score(r.(Request), em.Topfloor) > suitable.score(r.(Request), em.Topfloor) {
				suitable = e
			}
		}

		fmt.Printf("SCHEDULED request id %v to elevator %v. Destination floor %v\n", r.(Request).Id, suitable.Id, r.(Request).Destination)
		suitable.add(r.(Request))

		// -- add more Request Type handling here -- //
	}

}

// Step advances elevators in the system for one step. One step involves moving the elevator
// up or down, as well as picking up and dropping off requests.
func (em *ElevatorManager) Step() {
	for _, e := range em.Elevators {
		if e.Idle == false {
			e.pickup()
			e.dropoff()
			e.move()
		}
	}
}
