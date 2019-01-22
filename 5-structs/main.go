package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Active bool
}

// Struct como ponteiro para poder alterar os valores e propagar isso
func changePerson(p *Person){
	p.Age = 26
}

func main() {
	p := &Person{
		Name:   "Wilson Júnior",
		Age:    24,
		Active: true,
	}
	//Ver a questão de zero values
	// p := &Person{	Name:   "Wilson Júnior"}
	changePerson(p)
	fmt.Println("Nome", p.Name)
	fmt.Println("Age", p.Age)
	fmt.Println("Active", p.Active)
}
