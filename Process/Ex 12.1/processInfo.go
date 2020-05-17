package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"syscall"
)

func main() {

	files, err := ioutil.ReadDir("/proc")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
		fileNfo, err := os.Stat(file.Name())
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(fileNfo.Sys().(*syscall.Stat_t).Uid)
	}

}
