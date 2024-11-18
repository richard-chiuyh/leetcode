/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 */

// @lc code=start
func isMatch(s string, p string) bool {
	s += "0"
	p += "0"
	for p != "0" && s != "0" {
		if p[0] >= 'a' && p[0] <= 'z' {
			if len(p) > 1 && p[1] == '*' {
				tmp := p[0]
				p = p[2:]
				if p[0] == tmp {
					p = p[1:]
				}
				for s != "0" && s[0] == tmp {
					s = s[1:]
				}
			} else {
				if s[0] == p[0] {
					s = s[1:]
					p = p[1:]
				} else {
					return false
				}
			}
		} else if p[0] == '.' {
			if len(p) > 1 && p[1] == '*' {
				p = p[2:]
				if p == "0" {
					return true
				}
				for s != "0" && s[0] != p[0] {
					s = s[1:]
				}
			} else {
				s = s[1:]
				p = p[1:]
			}
		} else {
			return false
		}
	}
	if p != s {
		return false
	}
	return true
}

// @lc code=end

