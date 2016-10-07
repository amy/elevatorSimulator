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