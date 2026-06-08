package criptography

import (
	"errors"
	"fmt"
	"log"
	"math/big"
)

type FieldElement struct {
	Number *big.Int `json:"Number"`
	Prime  *big.Int `json:"Prime"`
}

func NewFieldElement(number, prime *big.Int) (FieldElement, error) {
	if number.Cmp(prime) >= 0 || number.Sign() < 0 {
		return FieldElement{}, errors.New("number not in field range")
	}
	return FieldElement{
		Number: new(big.Int).Set(number),
		Prime:  new(big.Int).Set(prime),
	}, nil
}

func (f FieldElement) String() string {
	if f.Number == nil {
		return "FieldElement(nil)"
	}
	return fmt.Sprintf("FieldElement_%v(%v)", f.Prime, f.Number)
}

func (f FieldElement) Equal(other FieldElement) bool {
	if f.Number == nil || other.Number == nil {
		return f.Number == nil && other.Number == nil
	}
	return f.Number.Cmp(other.Number) == 0 && f.Prime.Cmp(other.Prime) == 0
}

func (f FieldElement) Pow(pow *big.Int) FieldElement {
	expMod := new(big.Int).Sub(f.Prime, big.NewInt(1))
	n := new(big.Int).Mod(pow, expMod)
	result := new(big.Int).Exp(f.Number, n, f.Prime)
	fe, _ := NewFieldElement(result, f.Prime)
	return fe
}

func (f FieldElement) Mult(f2 FieldElement) FieldElement {
	if f.Prime.Cmp(f2.Prime) != 0 {
		log.Println("cannot mult two numbers in different field")
		return FieldElement{}
	}
	result := new(big.Int).Mul(f.Number, f2.Number)
	result.Mod(result, f.Prime)
	fe, _ := NewFieldElement(result, f.Prime)
	return fe
}

func (f FieldElement) Sub(f2 FieldElement) FieldElement {
	if f.Prime.Cmp(f2.Prime) != 0 {
		log.Println("cannot sub two numbers in different field")
		return FieldElement{}
	}
	result := new(big.Int).Sub(f.Number, f2.Number)
	result.Mod(result, f.Prime)
	if result.Sign() < 0 {
		result.Add(result, f.Prime)
	}
	fe, _ := NewFieldElement(result, f.Prime)
	return fe
}

func (f FieldElement) Add(f2 FieldElement) FieldElement {
	if f.Prime.Cmp(f2.Prime) != 0 {
		log.Println("cannot add two numbers in different field")
		return FieldElement{}
	}
	result := new(big.Int).Add(f.Number, f2.Number)
	result.Mod(result, f.Prime)
	fe, _ := NewFieldElement(result, f.Prime)
	return fe
}

func (f FieldElement) Div(f2 FieldElement) FieldElement {
	if f.Prime.Cmp(f2.Prime) != 0 {
		log.Println("cannot div two numbers in different field")
		return FieldElement{}
	}
	inverse := new(big.Int).Exp(f2.Number, new(big.Int).Sub(f.Prime, big.NewInt(2)), f.Prime)
	result := new(big.Int).Mod(new(big.Int).Mul(f.Number, inverse), f.Prime)
	fe, _ := NewFieldElement(result, f.Prime)
	return fe
}
