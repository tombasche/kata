package validip

import (
	"strconv"
	"strings"
)

const Ipv4 = "IPv4"
const Ipv6 = "IPv6"
const Neither = "Neither"

func ValidIP(ip string) string {
	ip = strings.ToLower(ip)
	tokens := strings.Split(ip, ".")
	if isValidIpv4(tokens) {
		return Ipv4
	}

	tokens = strings.Split(ip, ":")
	if isValidIpv6(tokens) {
		return Ipv6
	}

	return Neither
}

func isValidIpv4(toks []string) bool {
	if len(toks) == 1 {
		return false
	}
	for _, t := range toks {
		if len(t) > 1 && string(t[0]) == "0" {
			return false
		}
		num, err := strconv.Atoi(t)
		if err != nil {
			return false
		}
		if num < 0 || num > 255 {
			return false
		}

	}
	return true
}

func validIPv6Chars() []string {
	return []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
}

func isValidIpv6Char(ch string) bool {
	for _, c := range validIPv6Chars() {
		if ch == c {
			return true
		}
	}
	return false
}

func isValidIpv6(toks []string) bool {
	if len(toks) == 1 {
		return false
	}
	for _, t := range toks {
		if len(t) > 4 {
			return false
		}
		for _, ch := range t {
			if !isValidIpv6Char(string(ch)) {
				return false
			}
		}

	}
	return true
}
