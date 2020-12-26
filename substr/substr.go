package substr

import (
	"sort"
	"strings"
)

func Substr(input string) (string, int) {

	var seen []string
	substrings := make(map[int]string)

	for _, ch := range input {
		str := string(ch)

		if contains(seen, str) {
			s := strings.Join(seen, "")
			substrings[len(s)] = s
			seen = []string{str}
			continue
		}
		seen = append(seen, str)

	}

	keys := make([]int, len(substrings))
	i := 0
	for key := range substrings {
		keys[i] = key
		i++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	return substrings[keys[0]], keys[0]
}

func contains(s []string, str string) bool {
	for _, ch := range s {
		if ch == str {
			return true
		}
	}
	return false
}
