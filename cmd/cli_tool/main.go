package main

import (
	"fmt"
	"math/big"

	"github.com/igor-rodrigues2017/programming_bitcoin_go/internal/criptography"
)

func main() {
	point, err := criptography.NewPointFromString(
		"0",
		"7",
		"79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798",
		"483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8",
		"fffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc2f",
		16,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(point)

	n, _ := new(big.Int).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)
	fmt.Println(point.MultBig(n))

	coef := big.NewInt(2)
	p := criptography.G.MultBig(coef)
	fmt.Println(p)

	fmt.Println(criptography.G.MultBig(criptography.N))
}

func sumPoints() {
	var prime int64 = 223
	point, _ := criptography.NewPoint(0, 7, 192, 105, prime)
	point2, _ := criptography.NewPoint(0, 7, 17, 56, prime)

	sum := point.Add(point2)
	fmt.Println(sum)

	point3, _ := criptography.NewPoint(0, 7, 47, 71, prime)

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
