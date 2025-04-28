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

	// Create the function that returns the x and y values in the input file
	// input_vals := CreateInputVals

	// Create all of the variables from the input []bool arguments
	vars := make([]Code, 0, len(str_vars))
	for _, var_str := range str_vars {
		variable, err := ParseInputAssignment(var_str)
		if err != nil {
			log.Fatalf("Error parsing variable: %v", err)
		}
		vars = append(vars, variable)
	}

	// Generate the gates in valid syntax. Used to do that in this function, but then
	// it got kinda long.
	gates := CreateSortedGates(str_gates)

	// Create the return statement of `[]bool` from all of the variables
	// starting with `z__`
	return_stmt := CreateReturnSlice(str_gates)

	// Concatenate the vars and gates into a single slice
	function_statements := slices.Concat(vars, gates)
	function_statements = append(function_statements, return_stmt)

	// Create the default input values
	default_vals_func := CreateDefaultValuesFunc(str_vars)

	f := NewFile("circuits")

	// The default input values
	f.Add(default_vals_func)

	// The circuit function
	f.Func().
		Id("Circuit").
		Params(Id("x"), Id("y").Index().Bool()).
		Index().
		Bool().
		Block(function_statements...).
		Line()

	f.Save("circuits/circuit.go")
	// fmt.Printf("%#v", f)
}

// CreateReturnSlice will count how many gates there are that start with `z`, and then
// create a `[]bool` directly from the z values
func CreateReturnSlice(str_gates []string) Code {
	// Create a slice of all the z gates we see in `str_gates`
	z_str_gates := make([]string, 0)
	for _, gate_assignment := range str_gates {
		parts := strings.Fields(gate_assignment)
		if strings.HasPrefix(parts[len(parts)-1], "z") {
			z_str_gates = append(z_str_gates, parts[len(parts)-1])
		}
	}

	// Make sure the gates are in the right order
	slices.Sort(z_str_gates)

	// Make a `[]bool{z00, z01, ...}` from the strings in `z_str_gates`
	z_slice := make([]Code, len(z_str_gates))
	for i, z := range z_str_gates {
		z_slice[i] = Id(z)
	}

	// Put together the final Code item
	return Return(Index().Bool().Values(z_slice...))
}

// CreateSortedGates takes in the raw, unsorted gate operations, converts each gate
// to a `Code`, and then sorts them so they are syntactically valid.
func CreateSortedGates(str_gates []string) []Code {
	// For use in storing the variable assigned to, alongside the Code item
	type gate struct {
		var_assigned_to string
		code            Code
	}
	// Create all of the logic gate statements
	gates := make([]gate, 0, len(str_gates))

	// Create a DiGraph for dependencies
	graph := utils.NewDiGraph[string]()

	// Create the gates and digraph
	for _, gate_str := range str_gates {
		assignment, deps, err := ParseGate(gate_str)
		if err != nil {
			log.Fatalf("Error parsing gate: %v", err)
		}
		gates = append(gates, gate{deps.val, assignment})
		graph.AddEdge(deps.dep1, deps.val)
		graph.AddEdge(deps.dep2, deps.val)
	}

	// Topologically sort the gates
	sorted_str_gates, err := graph.TopoSort()
	if err != nil {
		log.Fatalf("Error sorting gates: %v", err)
	}

	// Remove all of the gates that start with "x.." or "y.."
	sorted_str_gates = slices.DeleteFunc(sorted_str_gates, func(gate string) bool {
		// If the gate starts with x or y, return true, meaning we want to remove it
		return strings.HasPrefix(gate, "x") ||
			strings.HasPrefix(gate, "y")
	})

	// Create a slice to be returned
	returned_gates := make([]Code, len(gates))

	// Sort `gates` by looking up its index in `sorted_gates`
	for _, gate := range gates {
		// What is the index of this gate in sorted_gates?
		// Note that the `Code` is actually a number of syntax items, and we want
		// the first one from that slice.
		idx := slices.Index(sorted_str_gates, gate.var_assigned_to)
		if idx == -1 {
			log.Fatalf("Gate %s not found in sorted gates", gate)
		}
		returned_gates[idx] = gate.code
	}
	return returned_gates
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

type gate_op struct {
	val, dep1, dep2 string
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
func ParseGate(line string) (Code, gate_op, error) {
	parts := strings.Fields(line)
	if len(parts) < 5 {
		return nil, gate_op{}, errors.New("invalid gate line format")
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
		return nil, gate_op{}, errors.New("unknown gate operation")
	}
	gate := Id(result).Op(":=").Id(left).Op(op).Id(right)
	deps := gate_op{result, left, right}
	return gate, deps, nil
}

// CreateDefaultValuesFunc generates a function that returns the default input values
// from the input text.
// The lines of input values look like
//
// x00: 1
// x01: 1
// x02: 0
// y00: 1
// y01: 0
// y02: 1
func CreateDefaultValuesFunc(str_values []string) Code {
	// Split the values into two slices, one for x and one for y
	x_str_vals := make([]string, 0)
	y_str_vals := make([]string, 0)
	for _, val := range str_values {
		if strings.HasPrefix(val, "x") {
			x_str_vals = append(x_str_vals, val)
		} else if strings.HasPrefix(val, "y") {
			y_str_vals = append(y_str_vals, val)
		}
	}

	// Sort them
	slices.Sort(x_str_vals)
	slices.Sort(y_str_vals)

	// Create the boolean Code values
	x_vals := make([]Code, len(x_str_vals))
	y_vals := make([]Code, len(y_str_vals))
	for i, val := range x_str_vals {
		// Split the string to get the value
		str_val := strings.TrimSpace(strings.Split(val, ":")[1])
		if str_val == "1" {
			x_vals[i] = True()
		} else if str_val == "0" {
			x_vals[i] = False()
		} else {
			log.Fatalf("unknown value %s", str_val)
		}
	}
	for i, val := range y_str_vals {
		// Split the string to get the value
		str_val := strings.TrimSpace(strings.Split(val, ":")[1])
		if str_val == "1" {
			y_vals[i] = True()
		} else if str_val == "0" {
			y_vals[i] = False()
		} else {
			log.Fatalf("unknown value %s", str_val)
		}
	}

	return Func().Id("InputValues").
		Params().
		Parens(List(Index().Bool(), Index().Bool())).
		Block(
			Return(
				Index().Bool().Values(x_vals...),
				Index().Bool().Values(y_vals...),
			),
		).
		Line()
}

func main() {
	// Read the input text from ../day24/input.txt
	raw_text := utils.ReadFile("day24/input.txt")
	GenerateCircuit(raw_text)
}
