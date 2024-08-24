package libutils

import (
	"regexp"
	"strings"
)

var matchSnake = regexp.MustCompile("(^[A-Za-z])|_([A-Za-z])")

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

type camel struct {
}

var Camel = &camel{}

func (c *camel) ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func (c *camel) ToCamelCase(str string) string {
	return matchSnake.ReplaceAllStringFunc(str, func(s string) string {
		return strings.ToUpper(strings.Replace(s, "_", "", -1))
	})
}
