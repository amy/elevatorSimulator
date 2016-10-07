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
## Elevator System

### Improvements

1. **Add Tests**
2. **Elevator Idle State:** When elevators are idle, I should add logic in the elevator's move() function that positions them close to floors most likely to receive requests. For instance, in the beginning of the workday, elevators should be idle on floor 0. During lunch, elevators should stagger themselves through the building.
3. **Stopping Mechanism**: There's currently infrastructure in the simulator to send a done signal to stop the simulation. However, I never actually send a signal and need to manually quit the simulation with ctrl-c. I would like to avoid involving the elevator component to send the done logic, as it needs to be implemented on the simulator level.