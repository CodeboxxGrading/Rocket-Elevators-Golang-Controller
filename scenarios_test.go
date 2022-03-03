package main

import (
	"testing"
)

func TestScenario1(t *testing.T) {

	expectedColumn := battery.columnsList[1]
	expectedElevator := battery.columnsList[1].elevatorsList[4]
	userPosition := 1
	destination := 20
	expectedFinalPositions := []int{5, 15, 1, 2, 20}

	chosenColumn, chosenElevator := scenario1()

	if chosenColumn.ID == expectedColumn.ID {
		t.Log("Correct column selected")
	} else {
		t.Errorf("Wrong column selected, expected Column %v, got Column %v", expectedColumn.ID, chosenColumn.ID)
	}

	if chosenElevator.ID == expectedElevator.ID {
		t.Log("Best elevator selected")
	} else {
		t.Errorf("Wrong elevator selected, expected Elevator %v, got Elevator %v", expectedElevator.ID, chosenElevator.ID)
	}

	var userPickedUp bool
	if contains(chosenElevator.completedRequestsList, userPosition) {
		t.Log("The user was picked up at its floor")
		userPickedUp = true
	} else {
		t.Error("No elevator was sent to pick up the user")
	}

	if userPickedUp && chosenElevator.currentFloor == destination {
		t.Log("The user was brought to destination")
	} else {
		t.Error("The user didn't reach its destination")
	}

	for i := 0; i < len(expectedColumn.elevatorsList); i++ {
		if expectedColumn.elevatorsList[i].currentFloor == expectedFinalPositions[i] {
			t.Logf("Elevator %v ends up at the correct floor", expectedColumn.elevatorsList[i].ID)
		} else {
			t.Errorf("Elevator %v didn't finish at the correct floor, expected %v, got %v", expectedColumn.elevatorsList[i].ID, expectedFinalPositions[i], chosenColumn.elevatorsList[i].currentFloor)
		}
	}

}

func TestScenario2(t *testing.T) {
	expectedColumn := battery.columnsList[2]
	expectedElevator := expectedColumn.elevatorsList[0]
	userPosition := 1
	destination := 36
	expectedFinalPositions := []int{36, 28, 1, 24, 1}

	chosenColumn, chosenElevator := scenario2()

	if chosenColumn.ID == expectedColumn.ID {
		t.Log("Correct column selected")
	} else {
		t.Errorf("Wrong column selected, expected Column %v, got Column %v", expectedColumn.ID, chosenColumn.ID)
	}

	if chosenElevator.ID == expectedElevator.ID {
		t.Log("Best elevator selected")
	} else {
		t.Errorf("Wrong elevator selected, expected Elevator %v, got Elevator %v", expectedElevator.ID, chosenElevator.ID)
	}

	var userPickedUp bool
	if contains(chosenElevator.completedRequestsList, userPosition) {
		t.Log("The user was picked up at its floor")
		userPickedUp = true
	} else {
		t.Error("No elevator was sent to pick up the user")
	}

	if userPickedUp && chosenElevator.currentFloor == destination {
		t.Log("The user was brought to destination")
	} else {
		t.Error("The user didn't reach its destination")
	}

	for i := 0; i < len(expectedColumn.elevatorsList); i++ {
		if expectedColumn.elevatorsList[i].currentFloor == expectedFinalPositions[i] {
			t.Logf("Elevator %v ends up at the correct floor", expectedColumn.elevatorsList[i].ID)
		} else {
			t.Errorf("Elevator %v didn't finish at the correct floor, expected %v, got %v", expectedColumn.elevatorsList[i].ID, expectedFinalPositions[i], chosenColumn.elevatorsList[i].currentFloor)
		}
	}

}

func TestScenario3(t *testing.T) {
	columnUsed := battery.columnsList[3]
	expectedElevator := columnUsed.elevatorsList[0]
	userPosition := 54
	destination := 1
	expectedFinalPositions := []int{1, 60, 58, 54, 1}

	chosenElevator := scenario3()

	if chosenElevator.ID == expectedElevator.ID {
		t.Log("Best elevator selected")
	} else {
		t.Errorf("Wrong elevator selected, expected Elevator %v, got Elevator %v", expectedElevator.ID, chosenElevator.ID)
	}

	var userPickedUp bool
	if contains(chosenElevator.completedRequestsList, userPosition) {
		t.Log("The user was picked up at its floor")
		userPickedUp = true
	} else {
		t.Error("No elevator was sent to pick up the user")
	}

	if userPickedUp && chosenElevator.currentFloor == destination {
		t.Log("The user was brought to destination")
	} else {
		t.Error("The user didn't reach its destination")
	}

	for i := 0; i < len(columnUsed.elevatorsList); i++ {
		if columnUsed.elevatorsList[i].currentFloor == expectedFinalPositions[i] {
			t.Logf("Elevator %v ends up at the correct floor", columnUsed.elevatorsList[i].ID)
		} else {
			t.Errorf("Elevator %v didn't finish at the correct floor, expected %v, got %v", columnUsed.elevatorsList[i].ID, expectedFinalPositions[i], columnUsed.elevatorsList[i].currentFloor)
		}
	}
}

func TestScenario4(t *testing.T) {
	columnUsed := battery.columnsList[0]
	expectedElevator := columnUsed.elevatorsList[3]
	userPosition := -3
	destination := 1
	expectedFinalPositions := []int{-4, 1, -5, 1, -6}

	chosenElevator := scenario4()

	if chosenElevator.ID == expectedElevator.ID {
		t.Log("Best elevator selected")
	} else {
		t.Errorf("Wrong elevator selected, expected Elevator %v, got Elevator %v", expectedElevator.ID, chosenElevator.ID)
	}

	var userPickedUp bool
	if contains(chosenElevator.completedRequestsList, userPosition) {
		t.Log("The user was picked up at its floor")
		userPickedUp = true
	} else {
		t.Error("No elevator was sent to pick up the user")
	}

	if userPickedUp && chosenElevator.currentFloor == destination {
		t.Log("The user was brought to destination")
	} else {
		t.Error("The user didn't reach its destination")
	}

	for i := 0; i < len(columnUsed.elevatorsList); i++ {
		if columnUsed.elevatorsList[i].currentFloor == expectedFinalPositions[i] {
			t.Logf("Elevator %v ends up at the correct floor", columnUsed.elevatorsList[i].ID)
		} else {
			t.Errorf("Elevator %v didn't finish at the correct floor, expected %v, got %v", columnUsed.elevatorsList[i].ID, expectedFinalPositions[i], columnUsed.elevatorsList[i].currentFloor)
		}
	}
}
