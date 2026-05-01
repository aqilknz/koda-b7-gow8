package main

import (
	"fmt"
	"go-weekly/internals/task1"
	"go-weekly/internals/task3"
)

func main() {
	// task 1
	fmt.Println("Task-1")
	input := []int{1, 2, 3, 4}
	var inputNil []int
	inputKosong := []int{}

	task1.ProcessNumber(input)
	task1.ProcessNumber(inputNil)
	task1.ProcessNumber(inputKosong)
	fmt.Println("")

	// task 3
	fmt.Println("Task-3")
	um := task3.NewUserManager()
	um.AddUser(1, "Aqil")
	um.AddUser(2, "Khairun")
	um.AddUser(3, "Nadzar")
	// tester add
	um.AddUser(1, "Ahmad")
	// tester get
	fmt.Println(um.GetUser(1))
	fmt.Println(um.GetUser(2))
	fmt.Println(um.GetUser(3))
	// undefined user
	fmt.Println(um.GetUser(4))

}
