package main

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
		result += string(value)
	}

	return result
}
