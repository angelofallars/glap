package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Options struct {
	ignore_case bool
}

func main() {
	arg_count := len(os.Args[1:])

	if arg_count <= 0 {
		fmt.Println("usage: grep [PATTERN]")
		os.Exit(1)
	}

	options := Options{}

	var pattern string
	pattern_arg_found := false

	for _, arg := range os.Args[1:] {
		switch arg {
        
        case "--ignore-case":
            fallthrough
		case "-i":
			options.ignore_case = true

		default:
			// Treat the first non-option argument as the filter pattern
			pattern = arg
			pattern_arg_found = true
			break
		}
	}

	if !pattern_arg_found {
		fmt.Println("usage: grep [PATTERN]")
		os.Exit(1)
	}

	lines := read_input_lines()

	// processed_lines is needed because the internal representation of
	// things to filter won't always be the same as the original input,
	// e.g. when ignoring letter case, where everything turns into uppercase.
	var processed_lines []string

	// Clone the `lines` slice into the `processed_lines` slice
	processed_lines = append(processed_lines, lines...)

	if options.ignore_case {
		pattern = strings.ToUpper(pattern)

		for i := range processed_lines {
			processed_lines[i] = strings.ToUpper(processed_lines[i])
		}
	}

	print_matching_lines(lines, processed_lines, pattern)
}

func print_matching_lines(orig_lines []string, lines []string, pattern string) {
	for i, line := range lines {

		if strings.Contains(line, pattern) {
			fmt.Println(orig_lines[i])
		}
	}
}

func read_input_lines() []string {
	input_scanner := bufio.NewScanner(os.Stdin)
	input_lines := []string{}

	for input_scanner.Scan() {
		input_lines = append(input_lines, input_scanner.Text())
	}

	if error := input_scanner.Err(); error != nil {
		log.Fatal(error)
	}

	return input_lines
}
