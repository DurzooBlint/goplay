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
		fileinfo, _ := os.Stat(file.Name())
		fmt.Println(fileinfo.Sys())
		fmt.Println(fileinfo)
	}

}
