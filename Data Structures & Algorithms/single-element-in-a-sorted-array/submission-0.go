func singleNonDuplicate(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right - left) / 2

		if mid%2 == 1 {
			mid--
		}

		if nums[mid] == nums[mid+1] {
			left = mid + 2
		} else {
			right = mid
		}
	}
	return nums[left]
}
