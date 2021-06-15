package main

import (
	"os"
	"strconv"
)

func main() {
    scenarioNumber, err := strconv.Atoi(os.Args[1])
	if err != nil {
		runScenario(scenarioNumber)
	}
}
