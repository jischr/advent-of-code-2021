package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	nums, err := ReadFile("day01/input.txt")
	if err != nil {
		panic(err)
	}
	ans := largerThanPreviousMeasurement(nums)
	fmt.Println("The answer to part one is", ans)

	ans = largerThanPreviousSum(nums)
	fmt.Println("The answer to part two is", ans)

}

func largerThanPreviousMeasurement(nums []int) int {
	count := 0
	for i:=1; i<len(nums); i++ {
		if nums[i] > nums[i-1] {
			count++
		}
	}
	return count
}

func largerThanPreviousSum(nums []int) int {
	count := 0
	prevSum := nums[0] + nums[1] + nums[2]
	for i := 3; i < len(nums); i++ {
		sum := nums[i-2] + nums[i-1] + nums[i]
		if sum > prevSum {
			count++
		}
		prevSum = sum
	}
	return count
}

func ReadFile(filename string) (nums []int, err error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	// Assign cap to avoid resize on every append.
	nums = make([]int, 0, len(lines))

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}
		// Atoi better suits the job when we know exactly what we're dealing
		// with. Scanf is the more general option.
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}