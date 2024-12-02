package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func distanceCalculator() int {
	file, err := os.Open("./Day1/Part-2/input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	similarity := 0
	leftFrequency := make(map[int]int)
	rightFrequency := make(map[int]int)

	const maxInt = int(^uint(0) >> 1)
	leftMinValue := maxInt
	leftMaxValue := 0

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

		if leftMaxValue < leftId {
			leftMaxValue = leftId
		}
	}

	left := leftMinValue

	for left <= leftMaxValue {
		leftValue, ok := leftFrequency[left]
		if !ok || leftValue <= 0 {
			left++
			continue
		}

		similarity += leftValue * (left * rightFrequency[left])
		left++
	}

	return similarity
}

func main() {
	sum := distanceCalculator()
	fmt.Println(sum)
}
