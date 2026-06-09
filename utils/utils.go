package utils

import (
	"errors"
	"strings"
)

func CreateMapFromLinesBySeparator(lines []string, separator string) (map[string]string, error) {
	m := make(map[string]string)

	for _, line := range lines {
		if strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.Split(line, separator)
		if len(parts) < 2 {
			return nil, errors.New("must be in format key:value, invalid line: " + line)
		}
		
		if len(parts) == 2 {
			m[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}

	return m, nil
}