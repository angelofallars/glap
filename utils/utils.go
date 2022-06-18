package utils

import (
	"bufio"
	"log"
	"os"
)

func CloneSlice[T interface{}](original_slice []T) []T {
	copied_slice := make([]T, len(original_slice))
	copy(copied_slice, original_slice)

	return copied_slice
}

func ReadInputLines() []string {
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
