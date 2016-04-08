package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {

	// Make sure there is two Arguments passed
	if len(os.Args) < 3 {
		printOptions()
	}

	// Extracting Arguments
	hashing := os.Args[1]
	filePath := os.Args[2]

	// Choose right encoding to check
	switch hashing {
	case "-sha1":
		checkSha1(filePath)
	default:
		printOptions()
	}

}

func printOptions() {
	fmt.Println()
	fmt.Println()
	fmt.Println("----OPTIONS")
	fmt.Println()
	fmt.Println("    -sha1 filePath")
	fmt.Println()
	fmt.Println("#More usage coming soon! By kroksys.")
	fmt.Println()
	fmt.Println()
	os.Exit(0)
}

func checkSha1(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Wrong file path: ", filePath)
		os.Exit(0)
	}
	defer f.Close()

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		fmt.Println("Error !!!!!!!!!!!!!")
		os.Exit(0)
	}

	fmt.Println("SHA-1 of ", filePath, " ", hex.EncodeToString(h.Sum(nil)))
	os.Exit(0)
}
