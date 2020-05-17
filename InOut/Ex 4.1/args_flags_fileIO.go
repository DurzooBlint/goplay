package main

import (
	"flag"
	"log"
	"os"
)

func main() {

	appendFlg := flag.Bool("a", false, "Opens existing file and appends")

	flag.Parse()
	print(*appendFlg)

	if len(os.Args) < 3 {
		println("Missing argument(s). Exiting")
		os.Exit(1)
	}

	if *appendFlg == true {
		var fname string = os.Args[2]
		var text string = os.Args[3]

		f, err := os.OpenFile(fname, os.O_APPEND, 0666)
		if err != nil {
			println("File does not exist. Exiting.")
			log.Fatal(err)
			os.Exit(1)
		}
		f.WriteString(text)
		defer f.Close()
	} else {
		var fname string = os.Args[1]
		var text string = os.Args[2]
		f, err := os.Create(fname)
		if err != nil {
			log.Fatal(err)
		}
		f.WriteString(text)
		defer f.Close()

	}

}
