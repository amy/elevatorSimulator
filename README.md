# Elevator Simulator

## Installation
```
$ go get -u github.com/amy/elevatorSimulator
```
## Run
```
$ cd elevatorSimulator
$ go build
$ ./elevatorSimulator
// When the elevators all stop moving hit ctrl-c
// see Improvements section that addresses this
```
## Sample Output

![Alt Text](https://raw.githubusercontent.com/amy/elevatorSimulator/master/demo.gif)

## Project Walkthrough


### Overview 
In breaking down the elevator simulator, there are two main conceptual components:

1. **SIMULATOR:** The simulator is responsible for stepping the system forward. It is also responsible for sending requests to the system to schedule. 
2. **ELEVATOR SYSTEM:** The elevator system is the meat of this project. You can create any number of elevators and send any number of requests.

Below, we dive into the setup of each component in more detail. 

## Simulator 
Our simulator has two primary functions defined:

1. **StartStepper:** This will start the time-stepping simulation. It steps and snapshots the system after every tick. There is infrastructure to stop the simulation by sending a signal to the done channel, but that has yet to be implemented. See the improvement section for more details.
2. **SendRequests:** This sends requests at 3 second intervals to the Scheduler to schedule. 

Both functions should be run on separate go routines. 

## Elevator System
After initializing elevators, requests, and a elevator manager, you interact with the system through the elevator manager.

**Scheduling:** Each elevator will be given a "suitability" score based on proximity to the request and direction. The source of this calculation can be found [here.](http://www.columbia.edu/~cs2035/courses/ieor4405.S13/p14.pdf) Once a suitable elevator has been chosen, it is sent to the elevator's PickupRequests field.

**Moving:** At each tick, if the pick up floor has been reached, the elevator will pick up requests from PickupRequests and send them to the appropriate floor in DropoffRequests. The elevator will also dropoff requests from DropoffRequests once the drop off floor is reached. If the elevator's direction is up, it will travel up until the highest request and back down until the lowest request. It will also remain on the same floor when there are no requests to service. 

## Improvements

1. **Add Tests**
2. **Elevator Idle State:** When elevators are idle, I should add logic in the elevator's move() function that positions them close to floors most likely to receive requests. For instance, in the beginning of the workday, elevators should be idle on floor 0. During lunch, elevators should stagger themselves through the building.
3. **Stopping Mechanism**: There's currently infrastructure in the simulator to send a done signal to stop the simulation. However, I never actually send a signal and need to manually quit the simulation with ctrl-c. I would like to avoid involving the elevator component to send the done logic, as it needs to be implemented on the simulator level.
4. **Abstracting Request Handling Logic**: I would like to make a Schedulable interface to abstract out request logic instead of having a blank interface. Currently, there's infrastructure in elevator manager Schedule to accept any struct and specifically handles the elevator.Request type. I would like to change it to a Scheduable interface, and move the scoring logic over to the request so that the request can schedule itself. 