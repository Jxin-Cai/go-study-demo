package 字符串

import "testing"

func TestReverseStr(t *testing.T) {
	reverseStr("abcdefg", 2)
}
func reverseStr1(s string, k int) string {
	if len(s) < 0 {
		return ""
	}
	for offset := 0; len(s)-1 > offset; offset += 2 * k {
		if len(s)-1 < offset+k {
			s1 := s[:offset]
			s2 := s[offset:]
			s3 := ""
			for _, v := range s2 {
				s3 = string(v) + s3
			}
			return s1 + s3
		}
		s1 := s[:offset]
		s2 := s[offset : offset+k]
		s3 := s[offset+k:]
		for _, v := range s2 {
			s3 = string(v) + s3
		}
		s = s1 + s3
	}
	return s
}

func reverseStr(s string, k int) string {
	if len(s) < 0 {
		return ""
	}
	arr := []byte(s)

	for i := 0; i < len(arr); i += 2 * k {
		if len(arr)-1 < i+k {
			for j := 0; j < (len(arr)-i)/2; j++ {
				arr[i+j], arr[len(arr)-j-1] = arr[len(arr)-j-1], arr[i+j]
			}
			return string(arr)
		}

		for j := 0; j < k/2; j++ {
			arr[i+j], arr[i+k-j-1] = arr[i+k-j-1], arr[i+j]
		}
	}
	return string(arr)
}
