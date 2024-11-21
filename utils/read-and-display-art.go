package utils

import (
	"bufio"
	"log"
	"os"
)

// Opens the specified file to read its contents.
func readLogo(name string) (logo string) {
	file, err := os.Open(name)
	if err != nil {
		log.Printf("Error reading %q: %v\n", name, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		logo += scanner.Text() + "\n"
	}
	return
}
