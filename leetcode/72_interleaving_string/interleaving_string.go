package interleaving_string

// Worst time complexity:   O(m*n)
// Average time complexity: O(m*n)
// Space complexity:        O(m)
func IsInterleave(s1, s2, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	results := make([]bool, len(s2)+1)

	for i := len(s1); i >= 0; i -= 1 {
		for j := len(s2); j >= 0; j -= 1 {
			if i < len(s1) && s1[i] == s3[i+j] && results[j] ||
				j < len(s2) && s2[j] == s3[i+j] && results[j+1] ||
				i+j == len(s3) {
				results[j] = true
				continue
			}

			results[j] = false
		}
	}

	return results[0]
}
