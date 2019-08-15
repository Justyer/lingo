package url

import (
	"net/url"
	"strings"
)

func percentReplace(str string) string {
	str = strings.Replace(str, "+", "%20", -1)
	str = strings.Replace(str, "*", "%2A", -1)
	str = strings.Replace(str, "%7E", "~", -1)

	return str
}

func URLEncode(qStr string) string {
	x := url.QueryEscape(qStr)
	return percentReplace(x)
}
