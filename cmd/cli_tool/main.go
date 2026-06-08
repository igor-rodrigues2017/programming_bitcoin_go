package main

import (
	"fmt"
	"math/big"

	"github.com/igor-rodrigues2017/programming_bitcoin_go/internal/criptography"
)

func main() {
	var prime int64 = 223
	a, _ := criptography.NewFieldElement(big.NewInt(0), big.NewInt(prime))
	b, _ := criptography.NewFieldElement(big.NewInt(7), big.NewInt(prime))
	x, _ := criptography.NewFieldElement(big.NewInt(192), big.NewInt(prime))
	y, _ := criptography.NewFieldElement(big.NewInt(105), big.NewInt(prime))
	point, _ := criptography.NewPoint(a, b, x, y)

	x2, _ := criptography.NewFieldElement(big.NewInt(17), big.NewInt(prime))
	y2, _ := criptography.NewFieldElement(big.NewInt(56), big.NewInt(prime))
	point2, _ := criptography.NewPoint(a, b, x2, y2)

	sum := point.Add(point2)
	fmt.Println(sum)

	x3, _ := criptography.NewFieldElement(big.NewInt(47), big.NewInt(prime))
	y3, _ := criptography.NewFieldElement(big.NewInt(71), big.NewInt(prime))
	point3, _ := criptography.NewPoint(a, b, x3, y3)

	result := point3
	for i := 1; i < 21; i++ {
		result = result.Add(point3)

		fmt.Printf("%v * %v = %v ", i+1, point3, result)
		fmt.Println()
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
