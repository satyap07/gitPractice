package main

import "fmt"

func main() {

	//2 Dimensional Arrays
	arr2D := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for _, row := range arr2D {
		for _, val := range row {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}

	//1 Dimensional slice
	slice1d := make([]int, 5)
	fmt.Println(slice1d)

	//2 Dimensional slices
	slice2D := make([][]int, 6)
	for i, _ := range slice2D {
		slice2D[i] = make([]int, 3)
	}
	fmt.Println(slice2D)

	//Map
	mp := make(map[string]int)
	mp["apple"] = 5
	mp["banana"] = 6
	mp["pears"] = 8
	mp["grapes"] = 10
	mp["strawberry"] = 23
	mp["papaya"] = 69
	fmt.Println(mp)

	//Map with integer key
	mp2 := make(map[int]string)
	mp2[1] = "apple"
	mp2[2] = "banana"
	mp2[3] = "pears"
	mp2[4] = "papaya"
	mp2[5] = "strawberry"
	fmt.Println(mp2)

}
