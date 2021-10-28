package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	scenarioNumber, err := strconv.Atoi(os.Args[1])
	if err == nil {
		runScenario(scenarioNumber)
	} else {
		fmt.Println(err)
	}
}
