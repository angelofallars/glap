package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Options struct {
	ignore_case      bool
	show_line_number bool
	only_show_count  bool
	invert_match     bool
}

func main() {
	arg_count := len(os.Args[1:])

	if arg_count <= 0 {
		print_usage()
		os.Exit(1)
	}

	options := Options{}

	var pattern string
	pattern_arg_found := false

	for _, arg := range os.Args[1:] {
		switch arg {

		case "--help":
			fallthrough
		case "--h":
			print_help_message()
			os.Exit(0)

		case "--ignore-case":
			fallthrough
		case "-i":
			options.ignore_case = true

		case "--line-number":
			fallthrough
		case "-n":
			options.show_line_number = true

		case "--count":
			fallthrough
		case "-c":
			options.only_show_count = true

		case "--invert-match":
			fallthrough
		case "-v":
			options.invert_match = true

		default:
			if pattern_arg_found {
				continue
			}

			// Treat the first non-option argument as the filter pattern
			pattern = arg
			pattern_arg_found = true
		}
	}

	if !pattern_arg_found {
		print_usage()
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

	print_matching_lines(lines, processed_lines, pattern, options)
}

func print_matching_lines(original_lines []string, lines []string, pattern string, options Options) {
	line_count := 0

	for i, line := range lines {
		line_is_match := strings.Contains(line, pattern)

		if options.invert_match {
			line_is_match = !line_is_match
		}

		if line_is_match {

			if options.only_show_count {
				line_count += 1
				continue
			}

			if options.show_line_number {
				fmt.Printf("%v:", i)
			}

			fmt.Println(original_lines[i])
		}
	}

	if options.only_show_count {
		fmt.Println(line_count)
	}
}

func print_usage() {
	fmt.Println("usage: grep [OPTION]... PATTERN")
	fmt.Println("Try 'grep --help' for more information.")
}

func print_help_message() {
	fmt.Println("usage: grep [OPTION]... PATTERN")
	fmt.Println("Search for PATTERN matches from standard input. Reading from file support coming soon.")
	fmt.Println("Example: ls | grep -i '.go'")
	fmt.Printf("\n")
	fmt.Println("Available options:")
	fmt.Println(" --h, --help                show help message and exit")
	fmt.Println("  -i, --ignore-case         ignore case when finding matches")
	fmt.Println("  -n, --line-number         print line number before matching lines")
	fmt.Println("  -c, --count               only display the count of matching lines")
	fmt.Println("  -v, --invert-match        display non-matching lines instead")
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
