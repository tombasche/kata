package goal

import "bytes"

func GoalParser(input string) string {

	var out bytes.Buffer
	for i, ch := range input {
		str := string(ch)
		switch str {
		case "G":
			out.WriteString(str)
		case "(":
			if string(input[i+1]) == ")" {
				out.WriteString("o")
				continue
			}
			out.WriteString("al")
		}
	}

	return out.String()
}
