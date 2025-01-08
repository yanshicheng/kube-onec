package utils

import (
	"regexp"
	"strings"
)

// 去掉结尾 ADN
func RemoveQueryADN(queryStr strings.Builder) string {
	re := regexp.MustCompile(`(?i)\bAND\b\s*$`)
	return re.ReplaceAllString(queryStr.String(), " ")
}
