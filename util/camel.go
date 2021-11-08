package util

import "strings"

// ToCamel Converts a string to CamelCase
func ToCamel(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}

	n := strings.Builder{}
	n.Grow(len(s))
	next := true
	for _, v := range []byte(s) {
		vIsCap := v >= 'A' && v <= 'Z'
		vIsLow := v >= 'a' && v <= 'z'

		// char to upper
		if next && vIsLow {
			v -= 32
		}

		// current char already done, next need not handle
		if vIsCap || vIsLow {
			n.WriteByte(v)
			next = false
		} else if vIsNum := v >= '0' && v <= '9'; vIsNum { // current is num, next need upper
			n.WriteByte(v)
			next = true
		} else { // if current in those, next need upper
			next = v == '_' || v == ' ' || v == '-' || v == '.'
		}
	}
	return n.String()
}
