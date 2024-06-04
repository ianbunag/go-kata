package permutation_string

// CheckInclusion
// Worst time complexity:   O(n)
// Average time complexity: O(n)
// Space complexity:        O(1)
func CheckInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Counts, s2Counts, windowEnd, matches := [26]int{}, [26]int{}, 0, 0

	for windowEnd < len(s1) {
		s1Counts[s1[windowEnd]-'a'] += 1
		s2Counts[s2[windowEnd]-'a'] += 1
		windowEnd += 1
	}
	for index := 0; index < 26; index += 1 {
		if s1Counts[index] == s2Counts[index] {
			matches += 1
		}
	}
	for matches != 26 && windowEnd < len(s2) {
		left, right := s2[windowEnd-len(s1)]-'a', s2[windowEnd]-'a'

		s2Counts[left] -= 1
		if s1Counts[left] == s2Counts[left] {
			matches += 1
		} else if s1Counts[left] == s2Counts[left]+1 {
			matches -= 1
		}

		s2Counts[right] += 1
		if s1Counts[right] == s2Counts[right] {
			matches += 1
		} else if s1Counts[right] == s2Counts[right]-1 {
			matches -= 1
		}

		windowEnd += 1
	}

	return matches == 26
}
