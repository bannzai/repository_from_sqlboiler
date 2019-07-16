package strutil

import (
	"fmt"
	"regexp"
)

// NOTE: reference https://github.com/volatiletech/inflect/blob/e7201282ae8da26cd97aed2e516f75c0bd91bb93/inflect_test.go#L17
var SingularToPlural = map[string]string{
	"search":      "searches",
	"switch":      "switches",
	"fix":         "fixes",
	"box":         "boxes",
	"process":     "processes",
	"address":     "addresses",
	"case":        "cases",
	"stack":       "stacks",
	"wish":        "wishes",
	"fish":        "fish",
	"jeans":       "jeans",
	"funky jeans": "funky jeans",
	"category":    "categories",
	"query":       "queries",
	"ability":     "abilities",
	"agency":      "agencies",
	"movie":       "movies",
	"archive":     "archives",
	"index":       "indices",
	"wife":        "wives",
	"safe":        "saves",
	"half":        "halves",
	"move":        "moves",
	"salesperson": "salespeople",
	"person":      "people",
	"spokesman":   "spokesmen",
	"man":         "men",
	"woman":       "women",
	"basis":       "bases",
	"diagnosis":   "diagnoses",
	"datum":       "data",
	"medium":      "media",
	"stadium":     "stadia",
	"analysis":    "analyses",
	"node_child":  "node_children",
	"child":       "children",
	"experience":  "experiences",
	"day":         "days",
	"comment":     "comments",
	"foobar":      "foobars",
	"newsletter":  "newsletters",
	"old_news":    "old_news",
	"news":        "news",
	"series":      "series",
	"species":     "species",
	"quiz":        "quizzes",
	"perspective": "perspectives",
	"ox":          "oxen",
	"photo":       "photos",
	"buffalo":     "buffaloes",
	"tomato":      "tomatoes",
	"dwarf":       "dwarves",
	"elf":         "elves",
	"information": "information",
	"equipment":   "equipment",
	"bus":         "buses",
	"status":      "statuses",
	"status_code": "status_codes",
	"mouse":       "mice",
	"louse":       "lice",
	"house":       "houses",
	"octopus":     "octopi",
	"virus":       "viri",
	"alias":       "aliases",
	"portfolio":   "portfolios",
	"vertex":      "vertices",
	"matrix":      "matrices",
	"matrix fu":   "matrix fu",
	"matrix fus":  "matrix fus",
	"fus":         "fus",
	"axis":        "axes",
	"testis":      "testes",
	"crisis":      "crises",
	"rice":        "rice",
	"shoe":        "shoes",
	"horse":       "horses",
	"prize":       "prizes",
	"edge":        "edges",
	"database":    "databases",
}

func PluralSuffix(str string) string {
	// specialize Rule
	ruleSets := SingularToPlural
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
		if UpperCamelCase(key) == str[len(str)-len(key):] {
			replaced := str[:len(str)-len(key)] + UpperCamelCase(converted)
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
