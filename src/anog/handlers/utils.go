package handlers

import (
	"sort"
	"strings"
)

func 	normalize(subject string) string {
	var symbols []string
	for _, rune := range subject {
		symbols = append(symbols, strings.ToLower(string(rune)))
	}
	sort.Strings(symbols)
	return strings.Join(symbols, "")
}

