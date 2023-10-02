package main

import "fmt"

func main() {
	// hello()
	// variables()
	arrays()
}

func hello() {
	fmt.Println("Hello, world!")
}

func variables() {
	// Forma mais utilizada para declaração de variáveis com escopo mais aberto (globais)
	var name = "Luiz Henrique"

	// Forma mais utilizada para declaração de variáveis com escopo mais fechado
	age := 22
	
	fmt.Println(fmt.Sprintf("Meu nome é %s, tenho %d anos.", name, age))

	// Forma utilizada para declarar variáveis sem necessariamente atribuir um valor a elas, é importante que seja declarado o tipo da mesma
	var bestGameInTheWorld string

	bestGameInTheWorld = "Red Dead Redemption 2"

	// Constante
	const pi = 3.14

	fmt.Println(bestGameInTheWorld)
}

func arrays() {
	// Declarando array sem inicializá-lo
	var emptyArray []string

	// Adicionando elemento ao array
	emptyArray = append(emptyArray, "Arthur")

	// Declarando array e incializando o mesmo
	names := []string{"Abigail", "John", "Jack"}

	// Pegando um range específico de elementos de outro array
	names = names[0:2]

	// Tamanho do array
	namesLength := len(names)

	// Percorrendo array
	for index, value := range names {
		fmt.Println(index, value)
	}

	fmt.Println(emptyArray, names, namesLength)
}
