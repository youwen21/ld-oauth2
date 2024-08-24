package libutils

import "regexp"

// 是否是合法手机号
func IsValidPhone(phone string) bool {

	pattern := "\\d{11}" // 反斜杠要转义

	isMatched, _ := regexp.MatchString(pattern, phone)

	return isMatched
}
