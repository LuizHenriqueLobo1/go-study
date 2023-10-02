package utils

import "fmt"

// Quando o nome de uma função tem a primeira letra maiúscula, a função está sendo exportada
func SayHelloTo(name string) {
	finalPhrase := fmt.Sprintf("Olá, %s!", name)
	fmt.Println(finalPhrase)
}
