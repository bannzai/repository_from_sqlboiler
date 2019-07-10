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
			uppers[i] = lower
		}
		return uppers
	}()
)

func specializeIfMatched(
	str string,
	specliazedKeywords []string,
	specialize func(string) string,
) string {
	for _, keyword := range specliazedKeywords {
		if keyword == str {
			return specialize(str)
		}
	}
	return str
}

func SpecializeLowerCamelCase(str string) string {
	return specializeIfMatched(str, lowerSpecializeKeywords, LowerCamelCase)
}

func SpecializeUpperCamelCase(str string) string {
	return specializeIfMatched(str, upperSpecializeKeywords, UpperCamelCase)
}
