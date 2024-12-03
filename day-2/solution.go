package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func check(direction string, a int, b int) (bool, string) {

	diff := a - b
	diff_abs := absInt(diff)

	if diff_abs < 1 || diff_abs > 3 {
		return false, direction
	}

	if direction == "unknown" {
		if diff > 0 {
			direction = "increasing"
		} else {
			direction = "decreasing"
		}
	} else {
		if direction == "increasing" && diff < 0 {
			return false, direction
		} else if direction == "decreasing" && diff > 0 {
			return false, direction
		}
	}

	return true, direction
}

func validLevelDiff(a int, b int) bool {
	diff := a - b
	diff_abs := absInt(diff)

	return diff_abs >= 1 && diff_abs <= 3
}

func direction(a int, b int) string {
	diff := a - b

	if diff < 0 {
		return "decreasing"
	} else {
		return "increasing"
	}
}

func validate(report *[]int, removed int) {
	for i := 0; i < len(*report)-1;i++ {
		
	}
}

func main() {
	// INPUTS

	// input_path := "./input.txt"
	// input_path := "./input-sample.txt"
	input_path := "./input-sample-2.txt"

	_, source_file_path, _, _ := runtime.Caller(0)
	source_dir := filepath.Dir(source_file_path)

	file_content := string(lo.Must(os.ReadFile(filepath.Join(source_dir, input_path))))

	fmt.Println(file_content)

	file_lines := strings.Split(file_content, "\n")
	file_lines = file_lines[:len(file_lines)-1] // Last element is blank

	fmt.Println(file_lines)
	fmt.Println()

	reports := make([][]int, len(file_lines))

	for i, line := range file_lines {
		line_split := strings.Split(line, " ")

		fmt.Println(line_split)

		reports[i] = lo.Map(line_split, func(s string, _ int) int { return lo.Must(strconv.Atoi(s)) })

	}
	fmt.Println()

	fmt.Println(reports)
	fmt.Println()

	// PROBLEM 1

	safe_count := 0

ReportsLoop:
	for _, report := range reports {
		direction := "unknown"

		fmt.Println(report)

		for i := 0; i < len(report)-1; i++ {
			safe, new_direction := check(direction, report[i], report[i+1])

			if !safe {
				fmt.Println("Unsafe")
				continue ReportsLoop
			}

			direction = new_direction
		}

		fmt.Println("Safe")
		safe_count++
	}
	fmt.Println()

	fmt.Println("Problem 1 Result:", safe_count) // 624
	fmt.Println()

	// PROBLEM 2

	damped_safe_count := 0

DampedReportsLoop:
	for _, report := range reports {
		direction := "unknown"
		dampener_used := false

		fmt.Println(report)

		for i := 0; i < len(report)-1; i++ {
			safe, new_direction := check(direction, report[i], report[i+1])

			if !safe {
				if dampener_used {
					fmt.Println("Unsafe, i:", i)
					continue DampedReportsLoop
				}

				if (i + 2) >= len(report) {
					break
				}

				damped_safe, new_direction := check(direction, report[i], report[i+2])
				direction = new_direction

				if !damped_safe {
					fmt.Println("Unsafe, i:", i)
					continue DampedReportsLoop
				} else {
					fmt.Println("Dampener used, i:", i)
					dampener_used = true
					i++
				}
			} else {
				direction = new_direction
			}
		}

		fmt.Println("Safe")
		damped_safe_count++
	}
	fmt.Println()

	fmt.Println("Problem 2 Result:", damped_safe_count) // 647
	fmt.Println()
}
