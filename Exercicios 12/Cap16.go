/*
TRABALHANDO COM JSON
package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name        string
	Surname     string
	Age         int
	Profession  string
	BankAccount float64
}

func main() {
	tiago := person{
		Name:        "Tiago",
		Surname:     "Hipolito",
		Age:         29,
		Profession:  "Develop",
		BankAccount: 265.421,
	}
	victor := person{
		Name:        "Victo",
		Surname:     "Viscoothynne",
		Age:         25,
		Profession:  "Professor",
		BankAccount: 123.895,
	}
	//utilizar a func json.Marshal
	t, err := json.Marshal(tiago)
	if err != nil {
		fmt.Println("erro", err)
	}
	fmt.Println(string(t))
	v, err := json.Marshal(victor)
	if err != nil {
		fmt.Println("erro", err)
	}
	fmt.Println(string(v))
}

package main

import (
	"encoding/json"
	"os"
)

type person struct {
	Name    string
	Surname string
	Age     int
}

func main() {
	pessoa := person{
		Name:    "Tales",
		Surname: "Aivilin",
		Age:     17,
	}
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(pessoa)

}


unMashal
package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name        string  `json:"Name"`
	Surname     string  `json:"Surname"`
	Age         int     `json:"Age"`
	Profession  string  `json:"Profession"`
	BankAccount float64 `json:"BankAccount"`
}

func main() {
	var input = []byte(`[{"Name":"Victo","Surname":"Viscoothynne","Age":25,"Profession":"Professor","BankAccount":123.895}]`)
	var pessoa []person
	err := json.Unmarshal(input, &pessoa)
	if err != nil {
		fmt.Println("Erro", err)
	}
	fmt.Printf("%+v", pessoa)
}
*/
package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type person struct {
	Name    string `json:"Name"`
	Surname string `json:"Surname"`
}

func main() {
	var input = (`[{"Name":"Victo","Surname":"Viscoothynne"}]`)

	decoder := json.NewDecoder(strings.NewReader(input))
	t, _ := decoder.Token()
	fmt.Printf("%T: %v\n", t, t)
	for decoder.More() {
		var P person
		err := decoder.Decode(&P)
		if err != nil {
			fmt.Printf("erro", err)
		}
		fmt.Printf("%+v", P)
	}
	t, _ = decoder.Token()
	fmt.Printf("%T: %v\n", t, t)

}
