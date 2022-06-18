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
	"fmt"
	"glap/utils"
	"io/ioutil"
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

	files_to_read := remaining_args[1:]

	total_lines_matched := 0

	if len(files_to_read) == 0 {
		lines := utils.ReadInputLines()

		matching_lines := searchPattern(lines, pattern, options)

		printMatches(matching_lines, "", options)

		total_lines_matched += len(matching_lines)

	} else {
		for _, file_name := range files_to_read {
			var file, err = ioutil.ReadFile(file_name)

			if err != nil {
				fmt.Printf("grep: %v: No such file or directory\n", file_name)
				continue
			}

			lines := strings.Split(string(file), "\n")

			matching_lines := searchPattern(lines, pattern, options)

			if len(files_to_read) == 1 {
				printMatches(matching_lines, "", options)
			} else {
				printMatches(matching_lines, file_name, options)
			}

			total_lines_matched += len(matching_lines)
		}
	}

	if total_lines_matched > 0 {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}

func printMatches(lines []string, prefix string, options Options) {
	if prefix != "" {
		prefix = prefix + ":"
	}

	if !options.only_show_count {
		for _, line := range lines {
			fmt.Printf("%v%v\n", prefix, line)
		}
	} else {
		fmt.Printf("%v%v\n", prefix, len(lines))
	}
}

func searchPattern(lines []string, pattern string, options Options) []string {
	// processed_lines is needed because the internal representation of
	// things to filter won't always be the same as the original input,
	// e.g. when ignoring letter case, where everything turns into uppercase.
	processed_lines := utils.CloneSlice(lines)

	pattern, processed_lines = prepareForMatching(pattern, processed_lines, options)

	matching_indexes := findMatchingIndexes(processed_lines, pattern, options)

	matching_lines := []string{}

	for _, index := range matching_indexes {
		var line string

		if !options.show_line_number {
			line = lines[index]
		} else {
			line = fmt.Sprintf("%v:%v", index, lines[index])
		}

		matching_lines = append(matching_lines, line)
	}

	return matching_lines
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

func findMatchingIndexes(lines []string, pattern string, options Options) []uint {
	line_count := 0
	matching_indexes := []uint{}

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
			matching_indexes = append(matching_indexes, uint(i))
		}
	}

	return matching_indexes
}

func printUsage() {
	fmt.Println("usage: grep [OPTION]... PATTERN [FILE]")
	fmt.Println("Try 'grep --help' for more information.")
}

func printHelpMessage() {
	fmt.Println("usage: grep [OPTION]... PATTERN [FILE]")
	fmt.Println("Search for PATTERN matches from standard input, or from one file.")
	fmt.Println("Example: ls | grep -i '.go'")
	fmt.Printf("\n")
	fmt.Println("Available options:")
	flag.PrintDefaults()
}
