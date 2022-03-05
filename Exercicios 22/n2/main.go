package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()
	str, sep := "", ""
	for _, arg := range os.Args[1:] {
		str += sep + arg
		sep = " "
	}
	fmt.Println(str)
	tempo := time.Since(start)
	log.Printf("%s", tempo)
}
