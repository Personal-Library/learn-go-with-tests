package arrays

// Sum will take an array of 5 ints
func Sum(numbers []int) int {
	sum := 0

	// _ is the blank identifier, and allows us to ignore the index
	// The second argument is the value of the array at that index
	// You can also ignore the value and use the index by swapping _
	// The size of the array is encoded in its type. If you need a
	// collection of any size, then use a slice! The syntax is very
	// similar... just omit the size when declaring
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// Go lets you write variadic functions that can take a variable num of args
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

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
