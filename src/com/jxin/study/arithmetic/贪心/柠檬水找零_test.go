package 贪心

import "testing"

func TestService(t *testing.T) {
	lemonadeChange([]int{5, 5, 5, 10, 20})
}
func lemonadeChange(bills []int) bool {
	burse := map[int]int{5: 0, 10: 0}
	for _, v := range bills {
		switch v {
		case 5:
			burse[5]++
			break
		case 10:
			if burse[5] == 0 {
				return false
			}
			burse[5]--
			burse[10]++
			break
		case 20:
			if burse[5] == 0 {
				return false
			}
			if burse[10] != 0 {
				burse[5]--
				burse[10]--
				break
			}
			if burse[5] < 3 {
				return false
			}
			burse[5] = burse[5] - 3
			break
		}
	}
	return true
}
