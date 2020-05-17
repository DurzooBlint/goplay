package main

import (
	"fmt"
	"io/ioutil"
	"log"
	//"os"
	//"syscall"
)

/*func fileStatFromInfoOs(fi os.FileInfo, flags *uint32, fileStat *FileStat) {
	if statt, ok := fi.Sys().(*syscall.Stat_t); ok {
		*flags |= ssh_FILEXFER_ATTR_UIDGID
		fileStat.UID = statt.Uid
		fileStat.GID = statt.Gid
	}
}*/

func main() {

	files, err := ioutil.ReadDir("/proc")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
		/*fileNfo, err := os.Stat(file.IsDir())
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(fileNfo.Sys().(*syscall.Stat_t).Uid)*/
	}

}
