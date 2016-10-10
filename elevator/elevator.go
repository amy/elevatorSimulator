package elevator

import (
	"fmt"
	"math"
)

type Elevator struct {
	Id              int
	Floor           int
	Direction       string
	PickupRequests  [10][]Request
	DropoffRequests [10][]Request
	Idle            bool
}

// pickup moves pickup requests to respective dropoff floor
func (e *Elevator) pickup() {
	requests := e.PickupRequests[e.Floor]

	if len(requests) > 0 {
		for _, r := range requests {
			e.DropoffRequests[r.Destination] = append(e.DropoffRequests[r.Destination], r)
		}
	}

	if len(requests) > 0 {
		fmt.Printf("PICKED UP request(s) with id(s) ")
		for _, r := range requests {
			fmt.Printf("%v ", r.Id)
		}
		fmt.Printf("on elevator %v\n", e.Id)
	}

	e.PickupRequests[e.Floor] = []Request{}
}

// dropoff removes dropoff requests at assigned floor
func (e *Elevator) dropoff() {
	requests := e.DropoffRequests[e.Floor]

	if !e.hasRequests() {
		e.Idle = true
	}

	if len(requests) > 0 {

		fmt.Print("DROPPED OFF request id(s)")

		for _, r := range requests {
			fmt.Printf(" %v ", r.Id)
		}

		fmt.Printf("on floor %v\n", e.Floor)
	}

	e.DropoffRequests[e.Floor] = []Request{}
}

// hasRequests tells you if elevator has any assigned requests
func (e *Elevator) hasRequests() bool {
	for _, r := range e.PickupRequests {
		if len(r) > 0 {
			return true
		}
	}

	for _, r := range e.DropoffRequests {
		if len(r) > 0 {
			return true
		}
	}

	return false
}

// move steps elevator up or down
func (e *Elevator) move() {

	if e.Idle {
		return
	}

	if e.Direction == "up" {
		e.Floor++
		if e.Floor >= e.max() {
			e.Direction = "down"
		}
		return
	}

	e.Floor--
	if e.Floor <= e.min() {
		e.Direction = "up"
	}
}

// min returns lowest floor with either a pick up or drop off request
func (e *Elevator) min() int {
	pmin := 0
	for i := 0; i < len(e.PickupRequests); i++ {
		if len(e.PickupRequests[i]) > 0 {
			pmin = i
			break
		}
	}

	dmin := 0
	for i := 0; i < len(e.DropoffRequests); i++ {
		if len(e.DropoffRequests[i]) > 0 {
			dmin = i
			break
		}
	}

	if pmin < dmin {
		return pmin
	}

	return dmin
}

// max returns highest floor with either a pick up or drop off request
func (e *Elevator) max() int {
	pmax := 0
	for i := len(e.PickupRequests) - 1; i >= 0; i-- {
		if len(e.PickupRequests[i]) > 0 {
			pmax = i
			break
		}
	}

	dmax := 0
	for i := len(e.DropoffRequests) - 1; i >= 0; i-- {
		if len(e.DropoffRequests[i]) > 0 {
			dmax = i
			break
		}
	}

	if pmax > dmax {
		return pmax
	}

	return dmax
}

// add request to Elevator's pick up requests
func (e *Elevator) add(r Request) {
	e.Idle = false

	e.PickupRequests[r.Start] = append(e.PickupRequests[r.Start], r)
}

// calculate "suitability score"
// http://www.columbia.edu/~cs2035/courses/ieor4405.S13/p14.pdf
/**
(1) Towards a call, same direction
	FS = (N + 2) - d
(2) Towards the call, opposite direction
	FS = (N + 1) - d
(3) Away from call
	FS = 1

	N = # Floors – 1;
	d = distance between elevator and call
*/
func (e *Elevator) score(r Request, topfloor int) float64 {
	n := float64(topfloor - 1)
	d := math.Abs(float64(e.Floor - r.Start))

	s := float64(1)

	// towards call
	towards := (r.Start >= e.Floor && e.Direction == "up") || (r.Start <= e.Floor && e.Direction == "down")
	if towards {
		if e.Direction == r.Direction {
			s = n + 2 - d
		} else {
			s = n + 1 - d
		}
	}

	return s
}
