package main

import "fmt"

type Person struct {
	name string
}

func (p Person) handleEvent(vacancies []string) {
	fmt.Println("Hello,", p.name)
	fmt.Println("Choose a vacancy:")
	for _, vacancy := range vacancies {
		fmt.Println(vacancy)
	}

}
