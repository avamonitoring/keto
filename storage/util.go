package storage

import (
	"regexp"
	"github.com/gobwas/glob"
)

func exactMatch(target string, source []string) bool {
	for _, i := range source {
		if i == target {
			return true
		}
	}
	return false
}

func globMatch(target string, source []string) bool {
  g := glob.MustCompile(target, ':')
	for _, i := range source {
		if g.Match(i)  {
			return true
		}
	}
	return false
}

func regexMatch(target string, source []string) bool {
	re := regexp.MustCompile(target)
	for _, i := range source {
		if re.MatchString(i)  {
			return true
		}
	}
	return false
}
