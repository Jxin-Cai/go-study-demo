package 回溯

import (
	"fmt"
	"strconv"
	"testing"
)

func TestService1(t *testing.T) {
	ipArr := restoreIpAddresses("25525111135")
	fmt.Println(ipArr)
}

func restoreIpAddresses(s string) []string {
	var ret []string
	for i := 1; i < 4; i++ {
		for j := 1; j < 4; j++ {
			for x := 1; x < 4; x++ {
				for y := 1; y < 4; y++ {
					if i+j+x+y != len(s) {
						continue
					}
					s1 := s[:i]
					s2 := s[i : i+j]
					s3 := s[i+j : i+j+x]
					s4 := s[i+j+x : i+j+x+y]
					if !check(s1) {
						continue
					}
					if !check(s2) {
						continue
					}
					if !check(s3) {
						continue
					}
					if !check(s4) {
						continue
					}
					ret = append(ret, fmt.Sprintf("%v.%v.%v.%v", s1, s2, s3, s4))
				}
			}
		}
	}
	return ret
}

func check(s string) bool {
	if len(s) > 1 && s[:1] == "0" {
		return false
	}
	i, _ := strconv.Atoi(s)
	if i > 255 {
		return false
	}
	return true
}
