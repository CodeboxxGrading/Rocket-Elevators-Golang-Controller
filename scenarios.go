package main

import (
	"fmt"
)

var battery *Battery = NewBattery(1, 4, 60, 6, 5)

func moveAllElevators(c *Column) {
	for i := 0; i < len(c.elevatorsList); i++ {
		for len(c.elevatorsList[i].floorRequestsList) != 0 {
			c.elevatorsList[i].move()
		}
	}
}

func scenario1() (*Column, *Elevator) {
	column := battery.columnsList[1]

	column.elevatorsList[0].currentFloor = 20
	column.elevatorsList[0].direction = "down"
	column.elevatorsList[0].status = "moving"
	column.elevatorsList[0].floorRequestsList = append(column.elevatorsList[0].floorRequestsList, 5)

	column.elevatorsList[1].currentFloor = 3
	column.elevatorsList[1].direction = "up"
	column.elevatorsList[1].status = "moving"
	column.elevatorsList[1].floorRequestsList = append(column.elevatorsList[1].floorRequestsList, 15)

	column.elevatorsList[2].currentFloor = 13
	column.elevatorsList[2].direction = "down"
	column.elevatorsList[2].status = "moving"
	column.elevatorsList[2].floorRequestsList = append(column.elevatorsList[2].floorRequestsList, 1)

	column.elevatorsList[3].currentFloor = 15
	column.elevatorsList[3].direction = "down"
	column.elevatorsList[3].status = "moving"
	column.elevatorsList[3].floorRequestsList = append(column.elevatorsList[3].floorRequestsList, 2)

	column.elevatorsList[4].currentFloor = 6
	column.elevatorsList[4].direction = "down"
	column.elevatorsList[4].status = "moving"
	column.elevatorsList[4].floorRequestsList = append(column.elevatorsList[4].floorRequestsList, 2)

	chosenColumn, chosenElevator := battery.assignElevator(20, "up")
	moveAllElevators(chosenColumn)
	return chosenColumn, chosenElevator

}

func scenario2() (*Column, *Elevator) {
	column := battery.columnsList[2]

	column.elevatorsList[0].currentFloor = 1
	column.elevatorsList[0].direction = "up"
	column.elevatorsList[0].status = "stopped"
	column.elevatorsList[0].floorRequestsList = append(column.elevatorsList[0].floorRequestsList, 21)

	column.elevatorsList[1].currentFloor = 23
	column.elevatorsList[1].direction = "up"
	column.elevatorsList[1].status = "moving"
	column.elevatorsList[1].floorRequestsList = append(column.elevatorsList[1].floorRequestsList, 28)

	column.elevatorsList[2].currentFloor = 33
	column.elevatorsList[2].direction = "down"
	column.elevatorsList[2].status = "moving"
	column.elevatorsList[2].floorRequestsList = append(column.elevatorsList[2].floorRequestsList, 1)

	column.elevatorsList[3].currentFloor = 40
	column.elevatorsList[3].direction = "down"
	column.elevatorsList[3].status = "moving"
	column.elevatorsList[3].floorRequestsList = append(column.elevatorsList[3].floorRequestsList, 24)

	column.elevatorsList[4].currentFloor = 39
	column.elevatorsList[4].direction = "down"
	column.elevatorsList[4].status = "moving"
	column.elevatorsList[4].floorRequestsList = append(column.elevatorsList[4].floorRequestsList, 1)

	chosenColumn, chosenElevator := battery.assignElevator(36, "up")
	moveAllElevators(chosenColumn)
	return chosenColumn, chosenElevator
}

func scenario3() *Elevator {
	column := battery.columnsList[3]

	column.elevatorsList[0].currentFloor = 58
	column.elevatorsList[0].direction = "down"
	column.elevatorsList[0].status = "moving"
	column.elevatorsList[0].floorRequestsList = append(column.elevatorsList[0].floorRequestsList, 1)

	column.elevatorsList[1].currentFloor = 50
	column.elevatorsList[1].direction = "up"
	column.elevatorsList[1].status = "moving"
	column.elevatorsList[1].floorRequestsList = append(column.elevatorsList[1].floorRequestsList, 60)

	column.elevatorsList[2].currentFloor = 46
	column.elevatorsList[2].direction = "up"
	column.elevatorsList[2].status = "moving"
	column.elevatorsList[2].floorRequestsList = append(column.elevatorsList[2].floorRequestsList, 58)

	column.elevatorsList[3].currentFloor = 1
	column.elevatorsList[3].direction = "up"
	column.elevatorsList[3].status = "moving"
	column.elevatorsList[3].floorRequestsList = append(column.elevatorsList[3].floorRequestsList, 54)

	column.elevatorsList[4].currentFloor = 60
	column.elevatorsList[4].direction = "down"
	column.elevatorsList[4].status = "moving"
	column.elevatorsList[4].floorRequestsList = append(column.elevatorsList[4].floorRequestsList, 1)

	chosenElevator := column.requestElevator(54, "down")
	moveAllElevators(&column)
	return chosenElevator
}

func scenario4() *Elevator {
	column := battery.columnsList[0]

	column.elevatorsList[0].currentFloor = -4

	column.elevatorsList[1].currentFloor = 1

	column.elevatorsList[2].currentFloor = -3
	column.elevatorsList[2].direction = "down"
	column.elevatorsList[2].status = "moving"
	column.elevatorsList[2].floorRequestsList = append(column.elevatorsList[2].floorRequestsList, -5)

	column.elevatorsList[3].currentFloor = -6
	column.elevatorsList[3].direction = "up"
	column.elevatorsList[3].status = "moving"
	column.elevatorsList[3].floorRequestsList = append(column.elevatorsList[3].floorRequestsList, 1)

	column.elevatorsList[4].currentFloor = -1
	column.elevatorsList[4].direction = "down"
	column.elevatorsList[4].status = "moving"
	column.elevatorsList[4].floorRequestsList = append(column.elevatorsList[4].floorRequestsList, -6)

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
