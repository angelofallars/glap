package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	flag "github.com/spf13/pflag"
)

type Options struct {
	show_help        bool
	ignore_case      bool
	show_line_number bool
	only_show_count  bool
	invert_match     bool
}

var options Options

func init() {
	flag.BoolVarP(&options.show_help,
		"help", "h", false,
		"show help message and exit")

	flag.BoolVarP(&options.ignore_case,
		"ignore-case", "i", false,
		"ignore case when finding matches")

	flag.BoolVarP(&options.show_line_number,
		"line-number", "n", false,
		"print line number before matching lines")

	flag.BoolVarP(&options.only_show_count,
		"count", "c", false,
		"only display the count of matching lines")

	flag.BoolVarP(&options.invert_match,
		"invert-match", "v", false,
		"display non-matching lines instead")
}

func main() {

	flag.Parse()

	if options.show_help {
		print_help_message()
		os.Exit(0)
	}

	remaining_args := flag.Args()

	if len(remaining_args) == 0 {
		print_usage()
		os.Exit(1)
	}

	pattern := remaining_args[0]

	lines := read_input_lines()

	// processed_lines is needed because the internal representation of
	// things to filter won't always be the same as the original input,
	// e.g. when ignoring letter case, where everything turns into uppercase.
	processed_lines := copy_string_slice(lines)

	pattern, processed_lines = prepare_for_matching(pattern, processed_lines, options)

	print_matching_lines(lines, processed_lines, pattern, options)
}

func prepare_for_matching(pattern string, lines []string, options Options) (string, []string) {
	if options.ignore_case {
		pattern = strings.ToUpper(pattern)

		for i := range lines {
			lines[i] = strings.ToUpper(lines[i])
		}
	}

	return pattern, lines
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
	flag.PrintDefaults()
}

func copy_string_slice(original_slice []string) []string {
	copied_slice := make([]string, len(original_slice))
	copy(copied_slice, original_slice)

	return copied_slice
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
