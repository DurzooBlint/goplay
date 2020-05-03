package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	if len(os.Args) < 1 {
		println("Missing argument(s). Exiting")
		os.Exit(1)
	}

	var fname string = os.Args[1]

	// Seek won't work if O_APPEND is used.
	f, err := os.OpenFile(fname, os.O_WRONLY, 0666)
	if err != nil {
		println("File does not exist. Exiting.")
		log.Fatal(err)
		os.Exit(1)
	}

	var ofst int64 = 0
	var newOfst int64

	// Seek won't work if O_APPEND is used on OpenFile.
	newOfst, err = f.Seek(ofst, 0)
	if err != nil {
		println("File does not exist. Exiting.")
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Printf("Current position: %d ", newOfst)
	f.WriteString("XxX")

	defer f.Close()
}
