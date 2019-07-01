package url

import (
	"sort"
	"strings"

	s2m "github.com/Justyer/lingo/struct2map"
)

func OrderArgsString(obj interface{}, tag string) (string, error) {
	rltMap, err := s2m.KeyByTagToString(obj, tag)
	if err != nil {
		return "", err
	}
	var keys []string
	for k := range rltMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i, k := range keys {
		keys[i] = k + "=" + rltMap[k]
	}
	return strings.Join(keys, "&"), err
}
