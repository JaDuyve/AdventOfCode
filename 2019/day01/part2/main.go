package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func calculateModuleFuel(moduleMass int) int {
	fuel := moduleMass/3 - 2

	if fuel <= 0 {
		return 0
	}
	return calculateModuleFuel(fuel) + fuel
}

func main() {
	file, err := os.Open("input")
	check(err)

	var result int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var moduleMass int
		_, err := fmt.Sscanf(scanner.Text(), "%d", &moduleMass)
		check(err)
		result += calculateModuleFuel(moduleMass)
	}

	file.Close()

	fmt.Printf("Total fuel needed: %d\n", result)

}
