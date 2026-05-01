package task1

import "fmt"

func ProcessNumber(listAngka []int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	alllist := []int{}

	if listAngka == nil {
		panic("Panic: No Data Provided")
	}
	if len(listAngka) == 0 {
		panic("Panic: Empty List Provided")
	}
	for _, value := range listAngka {
		alllist = append(alllist, value*2)
	}
	fmt.Println("Task-1")
	fmt.Println("Hasil: ", alllist)

}
