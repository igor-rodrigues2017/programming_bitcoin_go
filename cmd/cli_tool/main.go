package main

import (
	"fmt"
	"math"

	"github.com/igor-rodrigues2017/programming_bitcoin_go/internal/criptography"
)

func main() {
	multipleFineteField()

	elements := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	for _, e := range elements {
		var p int64 = 19
		criptography.Mod(int64(math.Pow(float64(e), float64(p-1))), p)

	}
}

func multipleFineteField() {
	fmt.Println("init multipleFineteField")

	elements := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	for _, k := range []int{1, 3, 7, 13, 18} {
		newElements := []int64{}
		for _, element := range elements {
			number := k * element
			newMod := criptography.Mod(int64(number), 19)
			newElements = append(newElements, newMod)

		}
		fmt.Println(newElements)
	}
	fmt.Println("end multipleFineteField")
}
