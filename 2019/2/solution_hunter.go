package main

import "errors"

func hunt_for_outcomes(desired int, base_program program) (int, error) {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			this_prog := apply_noun_and_verb_to_program(base_program, noun, verb)
			state := system_state{prog: this_prog, pc: 0}
			output_state := run_program(state)
			if output_state.prog[0] == pc_value(desired) {
				return 100*noun + verb, nil
			}
		}
	}
	return 0, errors.New("failed to find solution")
}

func apply_noun_and_verb_to_program(base_program program, noun int, verb int) program {
	new_prog := make(program, len(base_program))
	copy(new_prog, base_program)
	new_prog[1] = pc_value(noun)
	new_prog[2] = pc_value(verb)
	return new_prog
}
