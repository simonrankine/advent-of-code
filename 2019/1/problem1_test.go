package main

import "testing"

type fuelTest struct {
	weight, fuel int
}

var moduleFuelTests = []fuelTest{
	{12, 2},
	{14, 2},
	{1969, 654},
	{100756, 33583},
}

func Test_weight_to_fuel(t *testing.T) {
	for _, test := range moduleFuelTests {
		if output := weight_to_fuel(test.weight); output != test.fuel {
			t.Errorf("Output %d not equal to expected %d", output, test.fuel)
		}
	}
}

var totalFuelTests = []fuelTest{
	{14, 2},
	{1969, 966},
	{100756, 50346},
}

func Test_module_to_fuel(t *testing.T) {
	for _, test := range totalFuelTests {
		if output := module_to_fuel(test.weight); output != test.fuel {
			t.Errorf("Output %d not equal to expected %d", output, test.fuel)
		}
	}
}

func Test_module_weights_to_fuel(t *testing.T) {
	input := []int{14, 1969, 100756}
	expected := 51314
	if output := convert_modules_to_fuel(input); output != expected {
		t.Errorf("Expected total fuel %d, got %d", expected, output)
	}
}
