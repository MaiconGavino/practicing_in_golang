/*package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(40)
	contador := 0
	for i := 0; i < 40; i++ {
		go func() {
			v := contador
			v++
			contador = v
			wg.Done()
		}()
		fmt.Println(contador)
	}
}
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var contador int = 0
var input int = 50

func main() {
	proc(input)
	wg.Wait()
}

func proc(ent int) {
	wg.Add(ent)
	for i := 0; i < ent; i++ {
		go func() {
			aux := contador
			runtime.Gosched()
			aux++
			contador = aux
			wg.Done()
		}()
		fmt.Println(contador)
	}
}
