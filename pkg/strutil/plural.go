package strutil

import (
	"fmt"
	"regexp"
)

var ruleSets = map[string]string{
	"": "",
}

func Plural(str string) string {
	// specialize Rule
	for key, converted := range ruleSets {
		if len(key) > len(str) {
			continue
		}
		if len(key) <= 0 {
			continue
		}
		if key == str[len(str)-len(key):] {
			replaced := str[:len(str)-len(key)] + converted
			return replaced
		}
	}

	// normalize convert plural
	if len(str) < 2 {
		return str
	}
	if str[len(str)-1:] == "s" ||
		str[len(str)-2:] == "sh" ||
		str[len(str)-2:] == "ch" ||
		str[len(str)-1:] == "o" ||
		str[len(str)-1:] == "x" {
		return str + "es"
	}
	if str[len(str)-1:] == "f" {
		return str[0:len(str)-1] + "ves"
	}
	if str[len(str)-2:] == "fe" {
		return str[0:len(str)-2] + "ves"
	}
	isExists, _ := regexp.MatchString(".+[^aiueo]y$", str)
	if isExists {
		return str[0:len(str)-1] + "ies"
	}

	fmt.Printf("str = %v\n", str)
	return str + "s"
}
