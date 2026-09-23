func arrangeCoins(n int) int {
	left := 1
	right := n
	answer := 0

	for left <= right {
		mid := left + (right-left)/2

		needed := mid * (mid + 1) / 2

		if needed == n {
			return mid
		} else if needed < n {
			answer = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return answer
}