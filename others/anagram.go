//Package others contains implementation for algorithms
package others

import (
	"sort"
	"strings"
)

//IsAnagram verify, case insensitively, if two strings are anagrams
func IsAnagram(a, b string) bool {
	a, b = strings.ReplaceAll(a, " ", ""), strings.ReplaceAll(b, " ", "")

	if len(a) != len(b) {
		return false
	}

	x, y := strings.Split(strings.ToLower(a), ""), strings.Split(strings.ToLower(b), "")
	sort.Strings(x)
	sort.Strings(y)

	return strings.Join(x, "") == strings.Join(y, "")
}
