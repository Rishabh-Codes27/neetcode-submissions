func findJudge(n int, trust [][]int) int {
	// each person has a trust score
	score := make([]int, n+1)

	for _, t := range trust {

		a := t[0]
		b := t[1]

		score[a]--
		score[b]++

	}

	for person := 1; person <= n; person++ {
		if score[person] == n-1 {
			return person
		}
	}
	return -1
}


// people(n) = 1  2  3  4
// score =     0  0  0  0

// [[1,3], [4,3], [2,3]]

// 1 trusts 3
//  1  2  3  4
// -1  0 +1  0

// 4 trusts 3
//  1  2  3  4
// -1  0 +2 -1

// 2 trusts 3
//  1  2  3  4
// -1 -1 +3 -1