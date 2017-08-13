package main

import (
	"sort"
	"strconv"
)

func main() {
	println("Hi")
}

func ReturnInt() int {
	return 1
}

func ReturnFloat() float32 {
	return float32(1.1)
}

func ReturnIntArray() [3]int {
	return [3]int{1, 3, 4}
}

func ReturnIntSlice() []int {
	return []int{1, 2, 3}
}

func IntSliceToString(numbers []int) string {
	var result string

	for _, value := range numbers {
		result += strconv.Itoa(value)
	}

	return result
}

func MergeSlices(firstSlice []float32, secondSlice []int32) []int {
	result := make([]int, 0, len(firstSlice)+len(secondSlice))

	for _, value := range firstSlice {
		result = append(result, int(value))
	}

	for _, value := range secondSlice {
		result = append(result, int(value))
	}

	return result
}

func GetMapValuesSortedByKey(m map[int]string) []string {
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	result := make([]string, 0, len(keys))

	for _, key := range keys {
		result = append(result, m[key])
	}

	return result
}
