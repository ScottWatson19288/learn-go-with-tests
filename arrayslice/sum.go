package main

// Sum adds numbers and returns the sum of all of them
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll takes in a variable number of ints and adds them together
func SumAll(numbersToSum ...[]int) (sums []int) {
	lengthOfNumbers := len(numbersToSum)
	sums = make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}
