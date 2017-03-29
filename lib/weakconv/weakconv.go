package weakconv

import "fmt"

func String(v interface{}) string {
	if v != nil {
		return fmt.Sprint(v)
	}
	return ""
}

// TODO func JSONString(v interface{}) string
