package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	orbits := readAndPrepareData()

	orbitTree := buildOrbitTree(orbits)
	fmt.Print(orbitTree)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readAndPrepareData() [][]string {
	dat, err := ioutil.ReadFile("input")
	check(err)

	data := string(dat)
	orbits := strings.Split(data, "\n")
	orbitArray := make([][]string, len(orbits))

	for i, orbit := range orbits {
		bodies := strings.Split(orbit, ")")
		orbitArray[i] = bodies
	}

	return orbitArray
}

func buildOrbitTree(orbits [][]string) Object {

	var orbitTree Object

	for i := range orbits {

		addNode(i, orbits, &orbitTree)

	}

	return orbitTree
}

func getParentIndex(objectName string, orbits [][]string) int {

	for i, orbit := range orbits {
		if orbit[1] == objectName {
			return i
		}
	}

	return -1
}

func addNode(index int, orbits [][]string, orbitTree *Object) {
	bodies := orbits[index]

	parentIndex := getParentIndex(bodies[0], orbits)
	if parentIndex != -1 {
		if orbitTree.findNode(orbits[parentIndex][0]) == nil {
			addNode(parentIndex, orbits, orbitTree)
		}

	} else {
		orbitTree = &Object{name: bodies[0], next: nil}
	}

	child := Object{name: bodies[1], next: nil}
	orbitTree.next = append(orbitTree.next, &child)
}
