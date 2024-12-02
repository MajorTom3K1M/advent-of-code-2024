package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func distanceCalculator() int {
	file, err := os.Open("./Day1/Part-1/input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	distance := 0
	leftFrequency := make(map[int]int)
	rightFrequency := make(map[int]int)

	const maxInt = int(^uint(0) >> 1)
	leftMinValue := maxInt
	rightMinValue := maxInt
	leftMaxValue := 0
	rightMaxValue := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		ids := strings.Fields(line)

		leftId, err := strconv.Atoi(ids[0])
		if err != nil {
			log.Fatalf("error converting left id to integer: %s", err)
		}

		leftFrequency[leftId]++

		rightId, err := strconv.Atoi(ids[1])
		if err != nil {
			log.Fatalf("error converting right id to integer: %s", err)
		}

		rightFrequency[rightId]++

		if leftId < leftMinValue {
			leftMinValue = leftId
		}

		if rightId < rightMinValue {
			rightMinValue = rightId
		}

		if leftMaxValue < leftId {
			leftMaxValue = leftId
		}

		if rightMaxValue < rightId {
			rightMaxValue = rightId
		}
	}

	left, right := leftMinValue, rightMinValue

	for left <= leftMaxValue && right <= rightMaxValue {
		leftValue, ok := leftFrequency[left]
		if !ok || leftValue <= 0 {
			left++
			continue
		}

		rightValue, ok := rightFrequency[right]
		if !ok || rightValue <= 0 {
			right++
			continue
		}

		distance += int(math.Abs(float64(left - right)))

		leftFrequency[left]--
		rightFrequency[right]--
	}

	return distance
}

func main() {
	sum := distanceCalculator()
	fmt.Println(sum)
}
