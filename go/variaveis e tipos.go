package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Car struct {
	Name string `json:"Nomeeee"`
	Year int    `json:"Carrooo"`
}

func (c Car) Drive() {
	fmt.Println(c.Name, "Andou!")
}

func main() {
	//sao tipados
	a := 10
	b := "Rodrigo"
	c := 10.45
	d := true
	e := 'R'
	f := `Rodrigo 
	            Prado`

	fmt.Printf("Format: %v \n", a)
	fmt.Printf("Format: %v \n", b)
	fmt.Printf("Format: %v \n", c)
	fmt.Printf("Format: %v \n", d)
	fmt.Printf("Format: %v \n", e)
	fmt.Printf("Format: %v \n", f)

	fmt.Printf("Format: %T \n", a)
	fmt.Printf("Format: %T \n", b)
	fmt.Printf("Format: %T \n", c)
	fmt.Printf("Format: %T \n", d)
	fmt.Printf("Format: %T \n", e)
	fmt.Printf("Format: %T \n", f)

	car1 := Car{
		Name: "Focus",
		Year: 2013,
	}

	car1.Drive()

	result, _ := json.Marshal(car1)
	fmt.Println(string(result))

	x := 10
	y := &x
	fmt.Println(y)  //endereço da memoria
	fmt.Println(*y) //ponteiro para ver o valor registrado

	vetor := [10]int{4, 6, 8, 1, 2} // array com 10 posicoes
	fmt.Println(vetor)

	slice := make([]int, 5) //ele tem slice no lugar do array
	slice[0] = 1
	fmt.Println(slice)

	slice = append(slice, 1, 2, 3, 4) //ele pode adicionar mais posicoes no vetores
	fmt.Println(slice)

	sliceString := []string{
		"rod",
		"prado",
	}
	fmt.Println(sliceString)

	aa, bb, err := nome_e_idade("rod", 37)
	if err != nil {
		panic("Exception")
	}
	fmt.Println(aa, bb)
}

func nome_e_idade(nome string, idade int) (string, int, error) {
	if idade > 37 {
		return "", 0, errors.New("Deu erro")
	}
	return nome, idade, nil
}
