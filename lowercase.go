package pggoquery

import "strings"

// takes in a sql string and returns a postgres syntax tree

func Lowercase(s string) string {
	return strings.ToLower(s)

}
