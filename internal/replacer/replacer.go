package replacer

import (
	"strings"
)

func Process(keyMap map[string]string, source string, fs string, es string) string {
	res := source

	for key, value := range keyMap {
		// fmt.Printf("s=%s", value)
		placeholder := fs + key + es
		// fmt.Println(placeholder)
		res = strings.ReplaceAll(res, placeholder, value)
	}

	return res
}