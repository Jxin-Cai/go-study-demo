package main

func main() {

}
func plusOne(digits []int) []int {
	if 0 == len(digits) {
		return []int{1}
	}
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i] += 1
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}
