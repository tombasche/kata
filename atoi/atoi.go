package atoi

import (
	"math"
	"strconv"
	"strings"
)

func digits() []string {
	return []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
}

func isDigit(str string) bool {
	for _, s := range digits() {
		if str == s {
			return true
		}
	}
	return false
}

func Atoi(input string) int {

	var candidate []string
	input = strings.TrimSpace(input)

	if len(input) == 0 {
		return 0
	}
	signed := string(input[0]) == "-"
	for _, ch := range input {
		str := string(ch)
		if isDigit(str) {
			candidate = append(candidate, str)
			continue
		}
		if !signed && len(candidate) == 0 {
			return 0
		}
	}
	if signed {
		candidate = append([]string{"-"}, candidate...)
	}

	i, err := strconv.Atoi(strings.Join(candidate, ""))
	if err != nil {
		return 0
	}
	if i > math.MaxInt32 {
		return math.MaxInt32
	}
	if i < math.MinInt32 {
		return math.MinInt32
	}
	return i
}
