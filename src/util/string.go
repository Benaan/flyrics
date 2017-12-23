package util

import (
	"regexp"
	"strings"
)

func StripExtra(input string) string {
	var pattern = regexp.MustCompile(`[(\["'{<][^)\]>"}']+[)\]>"}']`)
	return strings.TrimSpace(pattern.ReplaceAllString(input, ""))
}

func RemovePunctuation(input string) string {
	var pattern = regexp.MustCompile(`([ !"#$%&'()*+,\-./0-9:;<=>?@[\\\]^_\{|}~\n` + "`]+)")
	return strings.TrimSpace(pattern.ReplaceAllString(input, ""))
}

func ToMatchable(input string) string {
	cleaned := RemovePunctuation(StripExtra(input))
	if cleaned == "" {
		cleaned = input
	}
	return strings.ToLower(cleaned)
}
