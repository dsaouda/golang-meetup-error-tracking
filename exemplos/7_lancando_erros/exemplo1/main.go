package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	usuario, err := getUsuarioById(6);

	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(usuario)
}

var (
	ErrorUsuarioNaoEncontrado = errors.New("Usuário não encontrado")
	ErrorEmailDuplicado = errors.New("Email duplicado")
)

type User struct {
	Id int
	Nome string
}

func (user User) String() string {
	return fmt.Sprintf("(Id=%d, Nome=%v)", user.Id, user.Nome)
}

var usuarios = []User{
	{1, "José Barbosa Sousa"},
	{2, "Samuel Sousa Souza"},
	{3, "Sophia Cardoso Cavalcanti"},
	{4, "Isabela Carvalho Ferreira"},
	{5, "Livia Dias Araujo"},
}

func getUsuarioById(id int) (*User, error) {

	for _, user := range usuarios {
		if user.Id == id {
			return &user, nil
		}
	}

	return nil, ErrorUsuarioNaoEncontrado
}

