package day1

import "fmt"
import "os"
import "strings"
import "strconv"
import "math"
import "sort"
import ("aoc2024/lib")

func Part1() {
	// input_buff, err := os.ReadFile("input/day1/example.txt");
	input_buff, err := os.ReadFile("input/day1/input.txt");
	if err != nil {
		fmt.Println("Error reading input file!")
		fmt.Println(err.Error)
		os.Exit(1)
	}
	
	input := string(input_buff)

	var listA []int
	var listB []int

	for _, line := range strings.Split(input, "\n") {
		// fmt.Println(line);
		sections := common.Filter(
			strings.Split(line, " "),
			func(s string) bool { return !common.IsWhitespace(s) },
		)
		if len(sections) == 2 {
			a, err := strconv.Atoi(sections[0]);
			if err != nil {
				fmt.Println("Error: ", err);
				os.Exit(1);
			}
			b, err := strconv.Atoi(sections[1]);
			if err != nil {
				fmt.Println("Error: ", err);
				os.Exit(1);
			}
			listA = append(listA, a)
			listB = append(listB, b)
		}
	}

	sort.Ints(listA);
	sort.Ints(listB);

	var total float64 = 0
	for i, a := range listA {
		min := math.Min(float64(a), float64(listB[i]))
		max := math.Max(float64(a), float64(listB[i]))
		total += max - min
	}

	fmt.Println(int(total))
}


func Part2() {
	// input_buff, err := os.ReadFile("input/day1/example.txt");
	input_buff, err := os.ReadFile("input/day1/input.txt");
	if err != nil {
		fmt.Println("Error reading input file!")
		fmt.Println(err.Error)
		os.Exit(1)
	}
	
	input := string(input_buff)

	var listA []int
	var listB []int

	for _, line := range strings.Split(input, "\n") {
		// fmt.Println(line);
		sections := common.Filter(
			strings.Split(line, " "),
			func(s string) bool { return !common.IsWhitespace(s) },
		)
		if len(sections) == 2 {
			a, err := strconv.Atoi(sections[0]);
			if err != nil {
				fmt.Println("Error: ", err);
				os.Exit(1);
			}
			b, err := strconv.Atoi(sections[1]);
			if err != nil {
				fmt.Println("Error: ", err);
				os.Exit(1);
			}
			listA = append(listA, a)
			listB = append(listB, b)
		}
	}

	var total int = 0

	for _, item := range listA {
		items := common.Filter(listB, func(s int) bool {
			return s == item
		},)
		total += item * len(items)
	}

	fmt.Println(total)
}
