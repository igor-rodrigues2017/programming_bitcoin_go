package criptography

import (
	"errors"
	"fmt"
	"log"
	"math/big"
)

type FieldElement struct {
	number *big.Int
	prime  *big.Int
}

func newFieldElement(number, prime *big.Int) (FieldElement, error) {
	if number.Cmp(prime) >= 0 || number.Sign() < 0 {
		return FieldElement{}, errors.New("number not in field range")
	}
	return FieldElement{
		number: new(big.Int).Set(number),
		prime:  new(big.Int).Set(prime),
	}, nil
}

func NewFieldElement(number, prime int64) (FieldElement, error) {
	return newFieldElement(big.NewInt(number), big.NewInt(prime))
}

func NewFieldElementFromBig(number, prime *big.Int) (FieldElement, error) {
	return newFieldElement(number, prime)
}

func NewFieldElementFromString(number, prime string, base int) (FieldElement, error) {
	n, ok := new(big.Int).SetString(number, base)
	if !ok {
		return FieldElement{}, errors.New("invalid field element number")
	}
	p, ok := new(big.Int).SetString(prime, base)
	if !ok {
		return FieldElement{}, errors.New("invalid field element prime")
	}
	return newFieldElement(n, p)
}

func (f FieldElement) String() string {
	if f.number == nil {
		return "FieldElement(nil)"
	}
	return fmt.Sprintf("FieldElement_%v(%v)", f.prime, f.number)
}

func (f FieldElement) Equal(other FieldElement) bool {
	if f.number == nil || other.number == nil {
		return f.number == nil && other.number == nil
	}
	return f.number.Cmp(other.number) == 0 && f.prime.Cmp(other.prime) == 0
}

func (f FieldElement) Pow(pow int64) FieldElement {
	return f.powBig(big.NewInt(pow))
}

func (f FieldElement) PowBig(pow *big.Int) FieldElement {
	return f.powBig(pow)
}

func (f FieldElement) powBig(pow *big.Int) FieldElement {
	expMod := new(big.Int).Sub(f.prime, big.NewInt(1))
	n := new(big.Int).Mod(pow, expMod)
	result := new(big.Int).Exp(f.number, n, f.prime)
	fe, _ := newFieldElement(result, f.prime)
	return fe
}

func (f FieldElement) Mult(f2 FieldElement) FieldElement {
	if f.prime.Cmp(f2.prime) != 0 {
		log.Println("cannot mult two numbers in different field")
		return FieldElement{}
	}
	result := new(big.Int).Mul(f.number, f2.number)
	result.Mod(result, f.prime)
	fe, _ := newFieldElement(result, f.prime)
	return fe
}

func (f FieldElement) Sub(f2 FieldElement) FieldElement {
	if f.prime.Cmp(f2.prime) != 0 {
		log.Println("cannot sub two numbers in different field")
		return FieldElement{}
	}
	result := new(big.Int).Sub(f.number, f2.number)
	result.Mod(result, f.prime)
	if result.Sign() < 0 {
		result.Add(result, f.prime)
	}
	fe, _ := newFieldElement(result, f.prime)
	return fe
}

func (f FieldElement) Add(f2 FieldElement) FieldElement {
	if f.prime.Cmp(f2.prime) != 0 {
		log.Println("cannot add two numbers in different field")
		return FieldElement{}
	}
	result := new(big.Int).Add(f.number, f2.number)
	result.Mod(result, f.prime)
	fe, _ := newFieldElement(result, f.prime)
	return fe
}

func (f FieldElement) Div(f2 FieldElement) FieldElement {
	if f.prime.Cmp(f2.prime) != 0 {
		log.Println("cannot div two numbers in different field")
		return FieldElement{}
	}
	inverse := new(big.Int).Exp(f2.number, new(big.Int).Sub(f.prime, big.NewInt(2)), f.prime)
	result := new(big.Int).Mod(new(big.Int).Mul(f.number, inverse), f.prime)
	fe, _ := newFieldElement(result, f.prime)
	return fe
}
