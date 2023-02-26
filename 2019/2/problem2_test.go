package main

import "testing"

var testprogs = [][2]program{
	{{1, 0, 0, 0, 99}, {2, 0, 0, 0, 99}},
	{{2, 3, 0, 3, 99}, {2, 3, 0, 6, 99}},
	{{2, 4, 4, 5, 99, 0}, {2, 4, 4, 5, 99, 9801}},
	{{1, 1, 1, 4, 99, 5, 6, 0, 99}, {30, 1, 1, 4, 2, 5, 6, 0, 99}},
}

func compare_progs(a program, b program) bool {
	for index, element := range a {
		if element != b[index] {
			return false
		}
	}
	return true
}

func Test_run_example_progs(t *testing.T) {
	for program_no, program_set := range testprogs {
		test_state := system_state{
			prog: program_set[0],
			pc:   0,
		}
		result_state := run_program(test_state)
		if !compare_progs(result_state.prog, program_set[1]) {
			t.Errorf("Programs did not match for problem set %d. Expected %v but got %v", program_no, program_set[1], result_state.prog)
		}
	}
}
