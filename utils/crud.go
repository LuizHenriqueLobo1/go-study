package utils

import (
	"fmt"
	"errors"
)

type People struct {
	ID int
	Name string
	Age int
}

func Main() {
	var peoples []People

	people := makePeople(peoples, "Luiz Henrique", 22)
	peoples = append(peoples, people)
	peoples = append(peoples, People { ID: createNewId(peoples), Name: "Fabiana", Age: 23 })	

	// listAllPeoples(peoples)

	targetPeople, err := findPeople(peoples, 2)
	if err != nil {
		fmt.Println(err)
		return
	}

	printPeople(targetPeople.ID, targetPeople.Name, targetPeople.Age)
}

func makePeople(peoples []People, name string, age int) People {
	return People {
		ID: createNewId(peoples),
		Name: name,
		Age: age,
	}
}

func listAllPeoples(peoples []People) {
	fmt.Println("---------------------")
	for _, people := range peoples {
		printPeople(people.ID, people.Name, people.Age)
		fmt.Println("---------------------")
	}
	fmt.Printf("Total de pessoas: %d\n", len(peoples))
	fmt.Println("---------------------")
}

func printPeople(id int, name string, age int) {
	fmt.Printf("ID: %d\n", id)
	fmt.Printf("Nome: %s\n", name)
	fmt.Printf("Idade: %d\n", age)
}

func findPeople(peoples []People, id int) (People, error) {
	for _, people := range peoples {
		if people.ID == id {
			return people, nil
		}
	}
	return People{}, errors.New("People not found!")
}

func createNewId(peoples []People) int {
	return len(peoples) + 1
}
