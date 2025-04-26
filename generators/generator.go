package main

import (
	"errors"
	"log"
	"slices"
	"strings"

	. "github.com/dave/jennifer/jen"
	"github.com/natemcintosh/aoc_2024/utils"
)

func GenerateCircuit(raw_text string) {
	// Break up the raw text by double newlines
	text := strings.SplitN(strings.TrimSpace(raw_text), "\n\n", 2)
	if len(text) < 2 {
		log.Println(text)
		log.Fatalf("Invalid input format")
	}
	str_vars := strings.Split(text[0], "\n")
	str_gates := strings.Split(text[1], "\n")

	// Create all of the variables from the input []bool arguments
	vars := make([]Code, 0, len(str_vars))
	for _, var_str := range str_vars {
		variable, err := ParseInputAssignment(var_str)
		if err != nil {
			log.Fatalf("Error parsing variable: %v", err)
		}
		vars = append(vars, variable)
	}

	// Create all of the logic gate statements
	gates := make([]Code, 0, len(str_gates))
	for _, gate_str := range str_gates {
		assignment, err := ParseGate(gate_str)
		if err != nil {
			log.Fatalf("Error parsing gate: %v", err)
		}
		gates = append(gates, assignment)
	}

	// Concatenate the vars and gates into a single slice
	all_circuit_statements := slices.Concat(vars, gates)

	f := NewFile("circuits")
	f.
		Func().
		Id("circuit").Params(Id("x"), Id("y").Index().Bool()).
		Index().
		Bool().
		Block(all_circuit_statements...)

	f.Save("circuits/circuit.go")

}

// ParseInputAssignment parses a single input assignment line and returns the
// corresponding assignment. Inputs look like:
//
//	x00: 1
//	x01: 0
//	y00: 1
//	y01: 1
//
// Note that this function does not actually assign the values seen in the input
// lines. It only creates the assignment statements. The values come from []bool
// argument to the function.
func ParseInputAssignment(line string) (Code, error) {
	parts := strings.SplitN(strings.TrimSpace(line), ":", 2)
	if len(parts) != 2 {
		return nil, errors.New("invalid input assignment format")
	}
	var_name := strings.TrimSpace(parts[0])

	// Get the index from the variable name
	idx := utils.ParseInt(var_name[1:])

	// Get the first letter of the variable name to determine if it's x or y
	prefix := string(var_name[0])

	// Create the assignment
	return Id(var_name).Op(":=").Id(prefix).Index(Lit(idx)), nil

}

// ParseGate parses a single gate line and returns the corresponding assignment
// Input line examples look like:
//
//	y33 AND x33 -> bfn
//	y32 XOR x32 -> rck
//	x30 AND y30 -> gns
//	y36 XOR x36 -> hbh
//	cng XOR mwt -> z42
//	bsw OR bfp -> pwp
func ParseGate(line string) (Code, error) {
	parts := strings.Fields(line)
	if len(parts) < 5 {
		return nil, errors.New("invalid gate line format")
	}
	result := strings.TrimSpace(parts[4])
	left := parts[0]
	right := parts[2]
	var op string
	switch strings.ToUpper(parts[1]) {
	case "AND":
		op = "&&"
	case "OR":
		op = "||"
	case "XOR":
		op = "!="
	default:
		return nil, errors.New("unknown gate operation")
	}
	return Id(result).Op(":=").Id(left).Op(op).Id(right), nil
}

func main() {
	// Read the input text from ../day24/input.txt
	raw_text := utils.ReadFile("day24/input.txt")
	GenerateCircuit(raw_text)
}
