package txt

import (
	"bufio"
	"log"
	"os"
)

func GetFileData(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
