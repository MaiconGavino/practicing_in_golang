/*package main

import (
	"fmt"
	"sort"
)

func main() {
	ss := []string{"maicon", "artur", "vinicio", "gabi", "fernando", "nociam", "yago"}
	fmt.Println(ss)
	sort.Strings(ss)
	fmt.Println(ss)

	is := []int{21234, 123, 4245, 12453, 1221, 236, 84}
	fmt.Println(is)
	sort.Ints(is)
	fmt.Println(is)

}

Aplicação de um Sort*/

package main

import (
	"fmt"
	"sort"
)

type person struct {
	name    string
	surname string
	age     int
}

type pessoa []person

func (p pessoa) Len() int {
	return len(p)
}
func (p pessoa) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p pessoa) Less(i, j int) bool {
	return p[i].age < p[j].age
}
func main() {
	pess := []person{
		{"maicon", "serrão", 25},
		{"bob", "rafale", 17},
		{"jenny", "baiany", 28},
	}
	sort.Sort(pessoa(pess))
	fmt.Println(pess)
}
