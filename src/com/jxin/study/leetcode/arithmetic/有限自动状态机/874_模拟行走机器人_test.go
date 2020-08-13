package 有限自动状态机

import (
	"sort"
	"testing"
)

func TestServer(t *testing.T) {
	//robotSim([]int{-2,-1,8,9,6}, [][]int{{-1,3},{0,1},{-1,5},{-2,-4},{5,4},{-2,-3},{5,-1},{1,-1}, {5,5},{5,2}})
	robotSim([]int{-2, -1, 8, 9, 6}, [][]int{{-1, 3}, {0, 1}, {-1, 5}, {-2, -4}, {5, 4}, {-2, -3}, {5, -1}, {1, -1}, {5, 5}, {5, 2}})
}
func TestServer1(t *testing.T) {
	arr := []int{1, 2, 3}

	t.Log(len(arr))
	t.Log(arr[3:3])

}
func robotSim(commands []int, obstacles [][]int) int {
	status := "上"
	offset := 0
	for true {
		if offset == len(commands) {
			break
		}
		if commands[offset] == -1 {
			status = rightStatus(status)
			offset++
			continue
		}
		if commands[offset] == -2 {
			status = leftStatus(status)
			offset++
			continue
		}
		break
	}
	x := 0
	y := 0
	// 获取路障散列表
	xMap := make(map[int][]int)
	yMap := make(map[int][]int)
	for _, v := range obstacles {
		if v[0] == 0 && v[1] == 0 {
			continue
		}
		if val, ok := xMap[v[0]]; ok {
			xMap[v[0]] = append(val, v[1])
		} else {
			xMap[v[0]] = append([]int{}, v[1])
		}

		if yMap[v[1]] != nil {
			yMap[v[1]] = append(xMap[v[1]], v[0])
		} else {
			yMap[v[1]] = append([]int{}, v[0])
		}
	}

	for offset < len(commands) {
		switch status {
		// 上
		case "上":
			yTemp := y
			y += commands[offset]
			if xMap[x] != nil {
				sort.Ints(xMap[x])
				for _, v := range xMap[x] {
					if v < yTemp {
						continue
					}
					if v > y {
						continue
					}
					y = v - 1
					break
				}
			}
			offset++

			for true {
				if offset == len(commands) {
					break
				}
				if commands[offset] == -1 {
					status = rightStatus(status)
					offset++
					continue
				}
				if commands[offset] == -2 {
					status = leftStatus(status)
					offset++
					continue
				}
				break
			}
			break
		// 下
		case "下":
			yTemp := y
			y -= commands[offset]
			if xMap[x] != nil {
				sort.Sort(sort.Reverse(sort.IntSlice(xMap[x])))
				for _, v := range xMap[x] {
					if v > yTemp {
						continue
					}
					if v < y {
						continue
					}
					y = v + 1
					break
				}
			}
			offset++

			for true {
				if offset == len(commands) {
					break
				}
				if commands[offset] == -1 {
					status = rightStatus(status)
					offset++
					continue
				}
				if commands[offset] == -2 {
					status = leftStatus(status)
					offset++
					continue
				}
				break
			}
			break
		// 左
		case "左":
			xTemp := x
			x -= commands[offset]
			if yMap[y] != nil {
				sort.Sort(sort.Reverse(sort.IntSlice(yMap[y])))
				for _, v := range yMap[y] {
					if v > xTemp {
						continue
					}
					if v < x {
						continue
					}
					x = v + 1
					break
				}
			}
			offset++

			for true {
				if offset == len(commands) {
					break
				}
				if commands[offset] == -1 {
					status = rightStatus(status)
					offset++
					continue
				}
				if commands[offset] == -2 {
					status = leftStatus(status)
					offset++
					continue
				}
				break
			}
			break
		// 右
		case "右":
			xTemp := x
			x += commands[offset]
			if yMap[y] != nil {
				sort.Ints(yMap[y])
				for _, v := range yMap[y] {
					if v < xTemp {
						continue
					}
					if v > x {
						continue
					}
					x = v - 1
					break
				}
			}
			offset++

			for true {
				if offset == len(commands) {
					break
				}
				if commands[offset] == -1 {
					status = rightStatus(status)
					offset++
					continue
				}
				if commands[offset] == -2 {
					status = leftStatus(status)
					offset++
					continue
				}
				break
			}
			break
		}
	}
	return x*x + y*y
}

func leftStatus(s string) string {
	switch s {
	// 上
	case "上":
		return "左"
	// 下
	case "下":
		return "右"
	// 左
	case "左":
		return "下"
	// 右
	default:
		return "上"
	}
}

func rightStatus(s string) string {
	switch s {
	// 上
	case "上":
		return "右"
	// 下
	case "下":
		return "左"
	// 左
	case "左":
		return "上"
	// 右
	default:
		return "下"
	}
}
