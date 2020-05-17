package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"syscall"
)

func main() {

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileinfo, err := os.Stat(file.Name())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(fileinfo.Sys().(*syscall.Stat_t).Uid)
	}

}
