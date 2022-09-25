package wildcard

import "strings"

func Match(pattern string, s string) (bool, error) {
	// Edge cases.
	if pattern == "*" {
		return true, nil
	}

	if pattern == "" {
		if s == "" {
			return true, nil
		}
		return false, nil
	}

	// If pattern does not contain wildcard chars, just compare the strings
	// to avoid extra memory allocation.
	if !strings.Contains(pattern, "*") && !strings.Contains(pattern, ".") {
		return pattern == s, nil
	}

	// Initialize DP.
	lp := len(pattern)
	ls := len(s)
	dp := make([][]bool, lp+1)
	for i := 0; i < lp+1; i++ {
		dp[i] = make([]bool, ls+1)
	}

	dp[0][0] = true

	for i := 0; i < lp; i++ {
		if pattern[i] == '*' {
			dp[i+1][0] = dp[i][0]
		} else {
			dp[i+1][0] = false
		}
	}

	for j := 0; j < ls; j++ {
		dp[0][j+1] = false
	}

	// Start DP.
	for i := 0; i < lp; i++ {
		for j := 0; j < ls; j++ {
			pc := pattern[i]
			sc := s[j]
			switch pattern[i] {
			case '*':
				dp[i+1][j+1] = dp[i][j] || dp[i][j+1] || dp[i+1][j]
			case '.':
				dp[i+1][j+1] = dp[i][j]
			default:
				if pc == sc {
					dp[i+1][j+1] = dp[i][j]
				} else {
					dp[i+1][j+1] = false
				}
			}
		}
	}

	return dp[lp][ls], nil
}
