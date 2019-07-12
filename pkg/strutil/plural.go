package strutil

import "regexp"

func Plural(str string) string {
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

	return str + "s"
}
