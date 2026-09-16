func mySqrt(x int) int {
    l := 1
    r := x

    for l <= r {
        mid := l + (r - l) / 2

        if mid*mid == x {
            return mid
        } else if mid*mid < x {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }

    return r
}
