package main

import "testing"

type vector_point_test struct {
	v vector
	p point
}

func TestVectorToPoint(t *testing.T) {
	test_cases := []vector_point_test{
		{"U2", point{0, 2}},
		{"R4", point{4, 0}},
		{"D20", point{0, -20}},
		{"L12", point{-12, 0}},
	}

	for _, test_case := range test_cases {
		initial := point{0, 0}
		actual := apply_vector_to_point(test_case.v, initial)
		if actual != test_case.p {
			t.Errorf("For vector %v expected %v but got %v", test_case.v, test_case.p, actual)
		}
	}

}
