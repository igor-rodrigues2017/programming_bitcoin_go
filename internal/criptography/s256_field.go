package criptography

import (
	"fmt"
	"math/big"
)

var (
	P  = bigFromHex("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F")
	A  = big.NewInt(0)
	B  = big.NewInt(7)
	N  = bigFromHex("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141")
	Gx = bigFromHex("79BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798")
	Gy = bigFromHex("483ADA7726A3C4655DA4FBFC0E1108A8FD17B448A68554199C47D08FFB10D4B8")
)

var G = func() S256Point {
	p, err := NewS256Point(Gx, Gy)
	if err != nil {
		panic(err)
	}
	return p
}()

func bigFromHex(s string) *big.Int {
	n, ok := new(big.Int).SetString(s, 16)
	if !ok {
		panic("invalid hex " + s)
	}
	return n
}

type S256Field struct {
	FieldElement
}

func NewS256Field(num *big.Int) (S256Field, error) {
	fe, err := NewFieldElementFromBig(num, P)
	if err != nil {
		return S256Field{}, err
	}
	return S256Field{fe}, nil
}

func (f S256Field) String() string {
	return fmt.Sprintf("%064x", f.number)
}

type S256Point struct {
	Point
}

func NewS256Point(x, y *big.Int) (S256Point, error) {
	a, _ := NewFieldElementFromBig(A, P)
	b, _ := NewFieldElementFromBig(B, P)

	if x == nil && y == nil {
		return S256Point{
			NewPointAtInfinityFromFieldElements(a, b),
		}, nil
	}

	xFe, err := NewFieldElementFromBig(x, P)
	if err != nil {
		return S256Point{}, err
	}
	yFe, err := NewFieldElementFromBig(y, P)
	if err != nil {
		return S256Point{}, err
	}

	p, err := NewPointFromFieldElements(a, b, xFe, yFe)
	if err != nil {
		return S256Point{}, err
	}

	return S256Point{p}, nil
}

func (p S256Point) MultBig(coef *big.Int) S256Point {
	ncoef := new(big.Int).Mod(coef, N)
	r := p.Point.MultBig(ncoef)
	return S256Point{r}
}
