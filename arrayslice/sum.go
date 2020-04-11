package main

// Sum adds numbers and returns the sum of all of them
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAllTails takes in a variable number of ints and adds them together
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {

			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
