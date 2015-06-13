package lifeofgame

import (
	"strings"
)

func Next(in string) (out string) {
	if strings.Count(in, "1") > 3 {
		return strings.Replace(in, "1", "0", -1)
	}
	if strings.Count(in, "1") == 3 {
		return strings.Replace(in, "0", "1", -1)
	}
	if strings.Count(in, "1") >= 2 {
		return in
	}
	return strings.Replace(in, "1", "0", -1)
}
