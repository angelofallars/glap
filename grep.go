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
    match_pattern := os.Args[1]

    for _, line := range input_lines {
        if (strings.Contains(line, match_pattern)) {
            fmt.Println(line)
        }
    }
}
