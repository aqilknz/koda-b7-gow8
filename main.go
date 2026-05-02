package main

import (
	"fmt"
	"go-weekly/internals/task1"
	"go-weekly/internals/task2"
	"go-weekly/internals/task3"
	"go-weekly/internals/task4"
	"sync"
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

	// task 2
	fmt.Println("Task-2")
	var wg sync.WaitGroup
	hasilChannel := make(chan string)
	listUrls := []string{
		"https://jsonplaceholder.typicode.com/photos",
		"https://jsonplaceholder.typicode.com/posts",
		"https://jsonplaceholder.typicode.com/comments",
		"https://jsonplaceholder.typicode.com/albums",
		"https://jsonplaceholder.typicode.com/todos",
		"https://jsonplaceholder.typicode.com/users",
	}
	go func() {
		for hasil := range hasilChannel {
			fmt.Println(hasil)
		}
		fmt.Println("Semua URL selesai di-fetch!")
	}()
	for _, url := range listUrls {
		wg.Add(1)
		go task2.WebFetcher(url, hasilChannel, &wg)
	}

	wg.Wait()
	close(hasilChannel)
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
	fmt.Println("")

	// task 4
	fmt.Println("Task-4")
	persegi := task4.Rectangle{Width: 12, Height: 8}
	lingkaran := task4.Circle{JariJari: 14}
	fmt.Println(task4.Calculator(persegi))
	fmt.Println(task4.Calculator(lingkaran))

}
