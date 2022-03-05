package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	// for valor, arg := range os.Args[:] {
	// 	fmt.Printf("%d : %s\n", valor, arg)
	// }
	tempo := time.Since(start)
	log.Printf("%s", tempo)

}
