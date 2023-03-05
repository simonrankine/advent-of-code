package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type pc_value int
type program []pc_value
type operands []pc_value
type outcome_type int

const (
	Exit  outcome_type = 1
	Store outcome_type = 2
)

type outcome struct {
	action   outcome_type
	result   pc_value
	location pc_value
}
type outcome_handler func(outcome, system_state) system_state

type instruction_type struct {
	name     string
	advance  pc_value
	operands int
	handler  instruction_handler
}
type instruction_handler func(operands, system_state) outcome

type system_state struct {
	prog program
	pc   pc_value
}

const INPUT_FILE = "input.txt"

var instruction_set = map[pc_value]instruction_type{
	1: {
		name:     "add",
		advance:  4,
		operands: 3,
		handler:  add_hander,
	},
	2: {
		name:     "multiply",
		advance:  4,
		operands: 3,
		handler:  multiply_handler,
	},
	99: {
		name:     "exit",
		advance:  0,
		operands: 0,
		handler:  exit_handler,
	},
}

var outcomes = map[outcome_type]outcome_handler{
	Exit:  exit_action_handler,
	Store: store_action_handler,
}

func main() {
	base_program := read_program_from_file(INPUT_FILE)
	outcome, err := hunt_for_outcomes(19690720, base_program)
	if err != nil {
		fmt.Println("Failed to find solution")
	}
	fmt.Println(outcome)
}

func run_program(state system_state) system_state {
	for {
		next_instruction := get_next_instruction(state)
		operands := get_operands(state, next_instruction)
		process_outcomes := run_instruction(next_instruction, operands, state)
		if process_outcomes.action == Exit {
			return state
		}
		state = update_state(state, process_outcomes)
		state.pc = update_pc(next_instruction, state.pc)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func read_program_from_file(filename string) program {
	f, err := os.Open(filename)
	check(err)
	defer f.Close()

	csv_reader := csv.NewReader(f)
	record, err := csv_reader.Read()
	check(err)

	var instructions program
	for _, entry := range record {
		int_instruction, err := strconv.Atoi(entry)
		check(err)
		instructions = append(instructions, pc_value(int_instruction))
	}

	return instructions
}

func get_next_instruction(state system_state) instruction_type {
	return opcode_to_instruction(state.prog[state.pc])
}

func get_operands(state system_state, instruction instruction_type) operands {
	return operands(state.prog[state.pc+1 : state.pc+pc_value(instruction.operands+1)])
}

func opcode_to_instruction(opcode pc_value) instruction_type {
	instruction, ok := instruction_set[opcode]
	if !ok {
		panic("Encountered invalid opcode")
	}
	return instruction
}

func run_instruction(instruction instruction_type, input operands, state system_state) outcome {
	return instruction.handler(input, state)
}

func update_pc(instruction instruction_type, pc pc_value) pc_value {
	return pc + instruction.advance
}

func add_hander(input operands, state system_state) outcome {
	return outcome{
		action:   Store,
		result:   state.prog[input[0]] + state.prog[input[1]],
		location: input[2],
	}
}

func multiply_handler(input operands, state system_state) outcome {
	return outcome{
		action:   Store,
		result:   state.prog[input[0]] * state.prog[input[1]],
		location: input[2],
	}
}

func exit_handler(input operands, state system_state) outcome {
	return outcome{
		action: Exit,
	}
}

func update_state(state system_state, outcome outcome) system_state {
	outcome_func := outcomes[outcome.action]
	state = outcome_func(outcome, state)
	return state
}

func store_action_handler(oc outcome, state system_state) system_state {
	state.prog[int(oc.location)] = oc.result
	return state
}

func exit_action_handler(oc outcome, state system_state) system_state {
	os.Exit(0)
	return state
}
