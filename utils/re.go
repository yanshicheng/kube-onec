package utils

import (
	"errors"
	"regexp"
	"strings"
)

// 去掉结尾 ADN
func RemoveQueryADN(queryStr strings.Builder) string {
	re := regexp.MustCompile(`(?i)\bAND\b\s*$`)
	return re.ReplaceAllString(queryStr.String(), " ")
}

// ValidateString 验证字符串是否符合规则
func ValidateString(input string) error {
	// 检查长度是否超过 20 位
	if len(input) > 20 {
		return errors.New("长度超过 20 个字符")
	}

	// 检查是否符合规则：开头为数字，包含小写字母和中横线，中横线不能结尾
	matched, err := regexp.MatchString(`^[0-9][a-z0-9-]*[a-z0-9]$`, input)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("字符串不符合规则：开头必须是数字，结尾不能是中横线，只能包含小写字母、数字和中横线")
	}

	return nil
}
