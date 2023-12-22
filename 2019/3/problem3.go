package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

type vector string
type vectorWire []vector
type point [2]int
type pointWire []point

func main() {
	//wires := read_wires_from_file("input.txt")

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func read_wires_from_file(filename string) [2]vectorWire {
	var wires [2]vectorWire

	f, err := os.Open(filename)
	check(err)
	defer f.Close()

	csv_reader := csv.NewReader(f)
	for i := 0; i < 2; i++ {
		record, err := csv_reader.Read()
		check(err)

		for _, entry := range record {
			wires[i] = append(wires[i], vector(entry[0:1]))
		}
	}

	return wires
}

func (w vectorWire) to_points() pointWire {
	var current_location = point{0, 0}
	var solution pointWire
	for _, vector := range w {
		current_location := apply_vector_to_point(vector, current_location)
		solution = append(solution, current_location)
	}
	return solution
}

func apply_vector_to_point(v vector, p point) point {
	direction := string(v[0])
	size, _ := strconv.Atoi(string(v[1:]))
	switch direction {
	case "U":
		p[1] += size
	case "D":
		p[1] -= size
	case "L":
		p[0] -= size
	case "R":
		p[0] += size
	}
	return p
}
