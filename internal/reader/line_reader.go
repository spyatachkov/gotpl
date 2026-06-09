package reader

import (
	"bufio"
	"os"
	"strings"
)

func ReadFileByLine(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	ar := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		ar = append(ar, line)
	}

	return ar, scanner.Err()
}