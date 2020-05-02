package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	var sourceName string = os.Args[1]
	var destName string = os.Args[2]

	if len(os.Args) < 2 {
		println("Missing argument(s). Exiting")
		os.Exit(1)
	}

	sourceFile, err := os.OpenFile(sourceName, os.O_RDONLY, 0444)

	//Creating a slice of n bytes to use for file read
	body := make([]byte, 50)

	var nBytesRead int = 0
	var nBytesWrite int = 0
	//Reading file by byte - as max size of slice b1
	nBytesRead, err = sourceFile.Read(body)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	//printing number of bytes read and content
	fmt.Printf("%d bytes: %s\n", nBytesRead, string(body[:nBytesRead]))

	destFile, err := os.OpenFile(destName, os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("Failed to write to file %s. Exiting.", destName)
		log.Fatal(err)
		os.Exit(1)
	} else {
		nBytesWrite, err = destFile.Write(body)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
			//printing number of bytes wrote and content
			fmt.Printf("%d bytes: %s\n", nBytesWrite, string(body[:nBytesWrite]))
		}

		defer sourceFile.Close()
		defer destFile.Close()

	}
}
