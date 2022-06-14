/* ## gogrep
 * -- print lines that match a pattern
 *
 * A clone of grep written in Golang.
 * Example: ls | grep -i '.png'
 *
 * Only has some of GNU grep's features, but more features will come soon.
 *
 * License: MIT
 */

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
	max_count        int
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

	flag.IntVarP(&options.max_count,
		"max-count", "m", -1,
		"stop after N selected lines")
}

func main() {

	flag.Parse()

	if options.show_help {
		printHelpMessage()
		os.Exit(0)
	}

	remaining_args := flag.Args()

	if len(remaining_args) == 0 {
		printUsage()
		os.Exit(1)
	}

	pattern := remaining_args[0]

	lines := readInputLines()

	// processed_lines is needed because the internal representation of
	// things to filter won't always be the same as the original input,
	// e.g. when ignoring letter case, where everything turns into uppercase.
	processed_lines := copyStringSlice(lines)

	pattern, processed_lines = prepareForMatching(pattern, processed_lines, options)

	matching_lines := findMatchingLines(lines, processed_lines, pattern, options)

	if !options.only_show_count {
		for _, line := range matching_lines {
			fmt.Println(line)
		}

	} else {
		fmt.Println(len(matching_lines))
	}

	if len(matching_lines) > 0 {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}

func prepareForMatching(pattern string, lines []string, options Options) (string, []string) {
	if options.ignore_case {
		pattern = strings.ToUpper(pattern)

		for i := range lines {
			lines[i] = strings.ToUpper(lines[i])
		}
	}

	return pattern, lines
}

func findMatchingLines(original_lines []string, lines []string, pattern string, options Options) []string {
	line_count := 0
	matching_lines := []string{}

	for i, line := range lines {
		if options.max_count >= 0 && line_count >= options.max_count {
			break
		}

		line_is_match := strings.Contains(line, pattern)

		if options.invert_match {
			line_is_match = !line_is_match
		}

		if line_is_match {
			line_count += 1

			matching_line := ""

			if options.show_line_number {
				matching_line = fmt.Sprintf("%v:", i)
			}

			matching_line += original_lines[i]

			matching_lines = append(matching_lines, matching_line)
		}
	}

	return matching_lines
}

func printUsage() {
	fmt.Println("usage: grep [OPTION]... PATTERN")
	fmt.Println("Try 'grep --help' for more information.")
}

func printHelpMessage() {
	fmt.Println("usage: grep [OPTION]... PATTERN")
	fmt.Println("Search for PATTERN matches from standard input. Reading from file support coming soon.")
	fmt.Println("Example: ls | grep -i '.go'")
	fmt.Printf("\n")
	fmt.Println("Available options:")
	flag.PrintDefaults()
}

func copyStringSlice(original_slice []string) []string {
	copied_slice := make([]string, len(original_slice))
	copy(copied_slice, original_slice)

	return copied_slice
}

func readInputLines() []string {
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
