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

    if ignore_case {
        match_pattern = strings.ToUpper(match_pattern)
    }

    // Read from STDIN
    input_scanner := bufio.NewScanner(os.Stdin)
    input_lines := []string{}

	for input_scanner.Scan() {
        input_lines = append(input_lines, input_scanner.Text())
	}

	if error := input_scanner.Err(); error != nil {
        log.Fatal(error)
	}
    
    // Check each line, and if it has the string sequence, print it
    for _, line := range input_lines {
        processed_line := line

        if ignore_case {
            processed_line = strings.ToUpper(processed_line)
        }

        if (strings.Contains(processed_line, match_pattern)) {
            fmt.Println(line)
        }
    }
}
