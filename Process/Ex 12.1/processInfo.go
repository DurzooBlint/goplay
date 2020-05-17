package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	files, err := ioutil.ReadDir("/proc")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(os.Stat(file.Name()))
	}

}
