package reader

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadFileByChunk(filePath string) error {
	chunkSize := 1024 * 1 // 1KB

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	buf := make([]byte, chunkSize)

	for {
		n, err := reader.Read(buf)
		if n > 0 {
		 process(buf[:n])
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}

	return nil
	
}

func process(buf []byte) {
	fmt.Println(string(buf))
}