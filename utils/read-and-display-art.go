package utils

import (
	"bufio"
	"log"
	"net"
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

// Uses Write property to display information of the specified file.
func displayLogo(conn net.Conn) {
	conn.Write([]byte(readLogo("linux-art/linux-art.txt")))
}
