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
		fileinfo, err := os.Stat(file.Name())
		if err != nil {
			fmt.Println("file does not exist")
			continue
		}
		fmt.Println(fileinfo.Sys().(*syscall.Stat_t).Uid)
	}

}
