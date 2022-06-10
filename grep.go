package main

import (
    "fmt"
    "os"
    "bufio"
    "log"
    "strings"
)

func main() {
    arg_count := len(os.Args[1:])

    if arg_count <= 0 {
        fmt.Println("usage: grep [PATTERN]")
        os.Exit(1)
    }

    ignore_case := false

    var match_pattern string
    match_pattern_found := false
    
    // Walk through command-line arguments
    // Treat the first non-option as the filter pattern
    for _, arg := range os.Args[1:] {
        switch arg {
            case "-i":
            ignore_case = true
            default:
            match_pattern = arg
            match_pattern_found = true
            break
        }
    }

    if !match_pattern_found {
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
    
    if ignore_case {
        match_pattern = strings.ToUpper(match_pattern)

        for i := range processed_lines {
            processed_lines[i] = strings.ToUpper(processed_lines[i])
        }
    }

    for i, processed_line := range processed_lines {

        if (strings.Contains(processed_line, match_pattern)) {
            fmt.Println(lines[i])
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

