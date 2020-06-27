package main

func main() {
	a := []int{1, 2, 3, 0, 0, 0}
	b := 3
	c := []int{2, 5, 6}
	d := 3
	merge(a, b, c, d)
}
func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
		return
	}

	for m > 0 && n > 0 {
		if nums1[m-1] <= nums2[n-1] {
			nums1[m+n-1] = nums2[n-1]
			n--
			continue
		}
		nums1[m+n-1] = nums1[m-1]
		m--
	}

	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
	}
}
