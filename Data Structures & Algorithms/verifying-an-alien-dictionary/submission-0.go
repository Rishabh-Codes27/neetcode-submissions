func isAlienSorted(words []string, order string) bool {
	rank := make(map[byte]int)
	for i := 0; i < len(order); i++ {
		rank[order[i]] = i
	}

	for i := 0; i < len(words)-1; i++ {
		word1 := words[i]
		word2 := words[i+1]

		minLen := len(word1)
		if len(word2) < minLen {
			minLen = len(word2)
		}

		for j := 0; j < minLen; j++ {
			if word1[j] != word2[j] {
				if rank[word1[j]] > rank[word2[j]] {
					return false
				}
				break
			}
		}
		if len(word1) > len(word2) && word1[:minLen] == word2[:minLen] {
			return false
		}
	}

	return true
}
