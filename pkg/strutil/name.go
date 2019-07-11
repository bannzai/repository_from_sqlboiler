package strutil

import (
	"regexp"
	"strings"
)

func SnakeCase(str string) string {
	matchFirstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// To Lower Camel case
func camelCase(str string) string {
	firstString := strings.ToLower(str[:1])
	droppedFirstString := str[1:]
	matchRegex := regexp.MustCompile("(_[a-zA-Z])")
	tailContent := matchRegex.ReplaceAllStringFunc(droppedFirstString, func(matched string) string {
		underscore := "_"
		return strings.ToUpper(matched[len(underscore):])
	})
	return firstString + tailContent
}

func LowerCamelCase(str string) string {
	return camelCase(str)
}

func UpperCamelCase(target string) string {
	firstString := strings.ToUpper(target[:1])
	tailContent := camelCase(target[1:])
	return firstString + tailContent
}

var (
	lowerSpecializeKeywords = []string{"url", "os", "id", "uid"}
	upperSpecializeKeywords = func() []string {
		uppers := make([]string, len(lowerSpecializeKeywords))
		for i, lower := range lowerSpecializeKeywords {
			uppers[i] = strings.ToUpper(lower)
		}
		return uppers
	}()
	reservedWords = []string{"type"}
)

func EscapedReservedWord(str string) string {
	for _, word := range reservedWords {
		if word == str {
			str = "_" + str
		}
	}
	return str
}

func SpecializeLowerCamelCase(str string) string {
	for _, keyword := range upperSpecializeKeywords {
		if keyword == str {
			return strings.ToLower(str)
		}
	}
	return LowerCamelCase(str)
}

func SpecializeUpperCamelCase(str string) string {
	for _, keyword := range lowerSpecializeKeywords {
		if keyword == str {
			return strings.ToUpper(str)
		}
	}
	return UpperCamelCase(str)
}

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
