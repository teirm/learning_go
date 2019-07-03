package main

import "fmt"

func reverse_array(a *[5]int) {

	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

}

func main() {

	a := [5]int{0, 1, 2, 3, 4}

	reverse_array(&a)

	fmt.Println(a)
}
