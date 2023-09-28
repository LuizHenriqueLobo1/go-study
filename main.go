package main

import "fmt"

func main() {
	// hello()
	variables()
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

	fmt.Println(bestGameInTheWorld)
}
