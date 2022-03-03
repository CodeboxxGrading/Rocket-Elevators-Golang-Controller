package main

import (
	"fmt"
)

type ElevatorDetails struct {
	floor         int
	direction     string
	status        string
	floorRequests []int
}

var battery *Battery = NewBattery(1, 4, 60, 6, 5)

func moveAllElevators(c *Column) {
	for i := 0; i < len(c.elevatorsList); i++ {
		for len(c.elevatorsList[i].floorRequestsList) != 0 {
			c.elevatorsList[i].move()
		}
	}
}

func (c *Column) setupElevators(elevatorDetails []ElevatorDetails) {
	for i := 0; i < len(c.elevatorsList); i++ {
		c.elevatorsList[i].currentFloor = elevatorDetails[i].floor
		c.elevatorsList[i].direction = elevatorDetails[i].direction
		c.elevatorsList[i].status = elevatorDetails[i].status
		c.elevatorsList[i].floorRequestsList = elevatorDetails[i].floorRequests
	}
}

func scenario1() (*Column, *Elevator) {
	column := battery.columnsList[1]

	elevator1 := ElevatorDetails{20, "down", "moving", []int{5}}
	elevator2 := ElevatorDetails{3, "up", "moving", []int{15}}
	elevator3 := ElevatorDetails{13, "down", "moving", []int{1}}
	elevator4 := ElevatorDetails{15, "down", "moving", []int{2}}
	elevator5 := ElevatorDetails{6, "down", "moving", []int{2}}

	column.setupElevators([]ElevatorDetails{elevator1, elevator2, elevator3, elevator4, elevator5})
	chosenColumn, chosenElevator := battery.assignElevator(20, "up")
	moveAllElevators(chosenColumn)
	return chosenColumn, chosenElevator

}

func scenario2() (*Column, *Elevator) {
	column := battery.columnsList[2]

	elevator1 := ElevatorDetails{1, "up", "stopped", []int{21}}
	elevator2 := ElevatorDetails{23, "up", "moving", []int{28}}
	elevator3 := ElevatorDetails{33, "down", "moving", []int{1}}
	elevator4 := ElevatorDetails{40, "down", "moving", []int{24}}
	elevator5 := ElevatorDetails{39, "down", "moving", []int{1}}

	column.setupElevators([]ElevatorDetails{elevator1, elevator2, elevator3, elevator4, elevator5})

	chosenColumn, chosenElevator := battery.assignElevator(36, "up")
	moveAllElevators(chosenColumn)
	return chosenColumn, chosenElevator
}

func scenario3() *Elevator {
	column := battery.columnsList[3]

	elevator1 := ElevatorDetails{58, "down", "stopped", []int{1}}
	elevator2 := ElevatorDetails{50, "up", "moving", []int{60}}
	elevator3 := ElevatorDetails{46, "up", "moving", []int{58}}
	elevator4 := ElevatorDetails{1, "up", "moving", []int{54}}
	elevator5 := ElevatorDetails{60, "down", "moving", []int{1}}

	column.setupElevators([]ElevatorDetails{elevator1, elevator2, elevator3, elevator4, elevator5})

	chosenElevator := column.requestElevator(54, "down")
	moveAllElevators(&column)
	return chosenElevator
}

func scenario4() *Elevator {
	column := battery.columnsList[0]

	elevator1 := ElevatorDetails{-4, "", "idle", []int{}}
	elevator2 := ElevatorDetails{1, "", "idle", []int{}}
	elevator3 := ElevatorDetails{-3, "down", "moving", []int{-5}}
	elevator4 := ElevatorDetails{-6, "up", "moving", []int{1}}
	elevator5 := ElevatorDetails{-1, "down", "moving", []int{-6}}

	column.setupElevators([]ElevatorDetails{elevator1, elevator2, elevator3, elevator4, elevator5})

	chosenElevator := column.requestElevator(-3, "up")
	moveAllElevators(&column)
	return chosenElevator
}

func runScenario(scenarioNumber int) {
	switch scenarioNumber {
	case 1:
		scenario1()
	case 2:
		scenario2()
	case 3:
		scenario3()
	case 4:
		scenario4()
	default:
		fmt.Println("Invalid scenario number")
	}

}
