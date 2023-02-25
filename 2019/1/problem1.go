package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const INPUT_FILE = "input.txt"

func main() {
	module_weights := get_input(INPUT_FILE)
	fmt.Println(convert_modules_to_fuel(module_weights))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func get_input(input_file string) []int {
	var lines []int

	f, err := os.Open(input_file)
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		input_line, err := strconv.Atoi(scanner.Text())
		check(err)

		lines = append(lines, input_line)
	}

	return lines
}

func module_weight_to_fuel(weight int) int {
	return int(math.Floor(float64(weight)/3.0) - 2)
}

func convert_modules_to_fuel(module_weights []int) int {
	fuel_weight := 0
	for _, this_weight := range module_weights {
		fuel_weight = fuel_weight + module_weight_to_fuel(this_weight)
	}
	return fuel_weight
}
