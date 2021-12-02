package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("day02/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(b), "\n")

	ans, err := multiplyHorizontalByDepth(lines)
	if err != nil {
		panic(err)
	}
	fmt.Println("the answer to part 1 is", ans)

	ans, err = multiplyUsingAim(lines)
	if err != nil {
		panic(err)
	}
	fmt.Println("the answer to part 2 is", ans)
}

func multiplyHorizontalByDepth(lines []string) (int, error) {
	depth := 0
	horizontal := 0
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		arr := strings.Split(l, " ")
		command := arr[0]
		count, err := strconv.Atoi(arr[1])
		if err != nil {
			return 0, err
		}

		switch command {
		case "down":
			depth += count
		case "up":
			depth -= count
		case "forward":
			horizontal += count
		}
	}
	return horizontal * depth, nil
}

func multiplyUsingAim(lines []string) (int, error) {
	aim := 0
	horizontal := 0
	depth := 0
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		arr := strings.Split(l, " ")
		command := arr[0]
		count, err := strconv.Atoi(arr[1])
		if err != nil {
			return 0, err
		}

		switch command {
		case "down":
			aim += count
		case "up":
			aim -= count
		case "forward":
			horizontal += count
			depth += aim * count
		}
	}
	return horizontal * depth, nil
}