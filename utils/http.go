package utils

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type GitHubUser struct {
	Name        string `json:"name"`
	AvatarURL   string `json:"avatar_url"`
	ReposAmount int    `json:"public_repos"`
}

func GetDataOfUser(user string) {
	url := "https://api.github.com/users/" + user

	// Faz o request
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Lê a resposta do request
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var userData GitHubUser
	if err := json.Unmarshal(body, &userData); err != nil {
		panic(err)
	}

	// Imprimir os dados
	fmt.Printf("Nome: %s\n", userData.Name)
	fmt.Printf("Avatar URL: %s\n", userData.AvatarURL)
	fmt.Printf("Quantidade de Repositórios: %d\n", userData.ReposAmount)
}
