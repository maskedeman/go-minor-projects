package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

func findFurthest(nums []int) (int, int, int) {
	highest := max(nums[0], max(nums[1], nums[2]))
	lowest := min(nums[0], min(nums[1], nums[2]))

	var remaining int
	for _, num := range nums {
		if num != highest && num != lowest {
			remaining = num
			break
		}
	}

	if math.Abs(float64(highest-remaining)) > math.Abs(float64(lowest-remaining)) {
		return highest, remaining, lowest
	} else {
		return lowest, remaining, highest
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	rand.Seed(time.Now().UnixNano())
	scanner := bufio.NewScanner(os.Stdin)

	index := 1
	correctAnswers := 0
	totalResponseTime := time.Duration(0)

	for {
		base := rand.Intn(100)
		nums := []int{base, base + rand.Intn(13) - 1, base + rand.Intn(13) - 1}
		// Ensure distinct numbers
		for nums[0] == nums[1] || nums[0] == nums[2] || nums[1] == nums[2] {
			nums = []int{base, base + rand.Intn(3) - 1, base + rand.Intn(3) - 1}
		}
		fmt.Printf(Yellow+"Question %d: %v\n"+Reset, index, nums)

		start := time.Now()
		scanner.Scan()
		userInput, _ := strconv.Atoi(scanner.Text())
		end := time.Now()

		responseTime := end.Sub(start)
		totalResponseTime += responseTime

		furthest, middle, _ := findFurthest(nums)
		if userInput == furthest {
			correctAnswers++
			fmt.Printf(Green+"Correct!"+Reset+Blue+" Response time: %v\n"+Reset, responseTime)
			fmt.Printf(Green+"The number %d is furthest from %d with a difference of %d.\n"+Reset, furthest, middle, int(math.Abs(float64(furthest-middle))))
		} else {
			fmt.Printf(Red+"Incorrect. The correct answer was "+Green+"%d"+Reset+Red+"."+Reset+Blue+" Response time: %v\n"+Reset, furthest, responseTime)
			fmt.Printf(Red+"The number %d is furthest from %d with a difference of %d, while your answer %d has a difference of %d from %d.\n"+Reset, furthest, middle, int(math.Abs(float64(furthest-middle))), userInput, int(math.Abs(float64(userInput-middle))), middle)
		}

		if index%20 == 0 {
			fmt.Printf("\nAfter %d questions, you have "+Green+"%d correct answers."+Reset+Blue+" Average response time: %v\n\n"+Reset, index, correctAnswers, totalResponseTime/time.Duration(index))
			correctAnswers = 0
			totalResponseTime = time.Duration(0)
		}

		index++
	}
}
