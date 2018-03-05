package main

import (
	"fmt"
	"time"
)

type animal struct {
	classe    string
	especie   string
	estimacao bool
	nome      string
	nasc      time.Time
}

func (a animal) imprimir() {
	if a.nome == "" {
		fmt.Println("Nome: Indefinido")
	} else {
		fmt.Printf("Nome: %s\n", a.nome)
	}

	fmt.Printf("Classe: %s\n", a.classe)
	fmt.Printf("Espécie: %s\n", a.especie)

	// Não tem operador ternário :-(
	if a.estimacao {
		fmt.Println("Animal de estimação")
	} else {
		fmt.Println("Animal selvagem")
	}

	if a.nasc.IsZero() {
		fmt.Println("Nascimento: Indefinido")
	} else {
		fmt.Printf("Nascimento: %s\n", a.nasc.Format("02/01/2006"))
	}
}

func main() {
	//Mon Jan 2 15:04:05 MST 2006
	// -   1  2  3  4  5 -7:00  6
	nasc, err := time.Parse("02/01/2006", "23/05/2006")
	if err != nil {
		fmt.Println(err.Error())
	}

	meuPet := animal{"Mamífero", "Cachorro", true, "Marley", nasc}
	aveSelvagem := animal{classe: "Ave", especie: "Falcão", estimacao: false}

	var cobraSelvagem animal
	cobraSelvagem.classe = "Réptil"
	cobraSelvagem.especie = "Cobra"
	cobraSelvagem.estimacao = false

	meuPet.imprimir()
	fmt.Println()
	aveSelvagem.imprimir()
	fmt.Println()
	cobraSelvagem.imprimir()
}
