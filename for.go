package main

import "fmt"

func main() {
	var i int
	for {
		i++
		if i == 5 {
			continue
		}
		if i == 10 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("최종 i 값은:", i)
}
