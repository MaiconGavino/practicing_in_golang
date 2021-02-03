/*
Partindo do código abaixo, utilize marshal para transformar []user em JSON.
package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "maicon",
		Age:   23,
	}
	u2 := user{
		First: "thiago",
		Age:   32,
	}
	u3 := user{
		First: "elisagela",
		Age:   22,
	}
	users := []user{u1, u2, u3}
	b, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}


package main

import (
	"encoding/json"
	"fmt"
)

type Data []struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	input := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	//fmt.Println(input)

	var out Data
	err := json.Unmarshal([]byte(input), &out)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(out)

}

### Na prática: exercício #3

- Partindo do código abaixo, utilize NewEncoder() e Encode() para enviar as informações como JSON para Stdout.
    - https://play.golang.org/p/BVRZTdlUZ_
- Desafio: descubra o que é, e utilize, method chaining para conectar os dois métodos acima.

package main

import (
	"encoding/json"
	"os"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}
	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}
	users := []user{u1, u2, u3}
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(users)
}

### Na prática: exercício #4

- Partindo do código abaixo, ordene a []int e a []string.
	- https://play.golang.org/p/H_q75mpmHW

package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}

### Na prática: exercício #5

- Partindo do código abaixo, ordene os []user por idade e sobrenome.
    - https://play.golang.org/p/BVRZTdlUZ_
- Os valores no campo Sayings devem ser ordenados tambem, e demonstrados de maneira esteticamente harmoniosa.
*/
package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type Age []user

func (u Age) Len() int { return len(u) }

func (u Age) Swap(i, j int) { u[i], u[j] = u[j], u[i] }

func (u Age) Less(i, j int) bool { return u[i].Age < u[j].Age }

type Last []user

func (l Last) Len() int { return len(l) }

func (l Last) Swap(i, j int) { l[i], l[j] = l[j], l[i] }

func (l Last) Less(i, j int) bool { return l[i].Last < l[j].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	//ordenando as frases sayings
	for _, value := range users {
		sort.Strings(value.Sayings)
	}

	//implementando a interface idade
	sort.Sort(Age(users))
	fmt.Println("Ordenar por idade")
	estetica(users)

	//implementando a interface Last
	sort.Sort(Last(users))
	fmt.Println("Ordenar por Sobrenome")
	estetica(users)
}

func estetica(input []user) {
	for i, value := range input {
		fmt.Printf("%d° User - Nome completo: %s %s  Idade: %d  \n", i+1, value.First, value.Last, value.Age)
		for i, value := range value.Sayings {
			fmt.Printf("\t Citação: %d° - %s\n", i+1, value)
		}
	}
}
