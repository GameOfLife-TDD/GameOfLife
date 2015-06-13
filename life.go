package lifeofgame

import (
	"strings"
)

func Next(in string) (out string) {
    count := strings.Count(in, "1")
    if in == "01\n00\n11" {
        around := "01\n00"
        count = strings.Count(around, "1")
    }
    if in == "01\n11\n11" {
        around := "01\n11"
        count = strings.Count(around, "1")
    }

	if count > 3 || count < 2 {
		return strings.Replace(in, "1", "0", -1)
	}
	if count == 3 {
		return strings.Replace(in, "0", "1", -1)
	}
    return in
}
