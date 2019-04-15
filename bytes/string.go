package bytes

import (
	"bytes"
	"strconv"
)

func Join(ss ...string) string {
	var buffer bytes.Buffer
	for _, s := range ss {
		buffer.WriteString(s)
	}
	return buffer.String()
}

func JoinObj(objs ...interface{}) string {
	var buffer bytes.Buffer
	for _, o := range objs {
		switch o.(type) {
		case int:
			buffer.WriteString(strconv.Itoa(o.(int)))
		case string:
			buffer.WriteString(o.(string))
		}
	}
	return buffer.String()
}
