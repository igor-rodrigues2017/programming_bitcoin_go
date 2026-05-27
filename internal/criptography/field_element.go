package criptography

import (
	"errors"
	"math/big"
)

type FieldElement struct {
	Number big.Int `json:"Number"`
	Prime  big.Int `json:"Prime"`
}

func NewFieldElement(number, prime *big.Int) (FieldElement, error) {
	if number.Cmp(prime) >= 0 || number.Sign() < 0 {
		return FieldElement{}, errors.New("number not in field range")
	}
	var fe FieldElement
	fe.Number.Set(number)
	fe.Prime.Set(prime)
	return fe, nil
}

func (f FieldElement) Equal(other FieldElement) bool {
	return f.Number.Cmp(&other.Number) == 0 && f.Prime.Cmp(&other.Prime) == 0
}

func (f FieldElement) Pow(pow *big.Int) (FieldElement, error) {
	expMod := new(big.Int).Sub(&f.Prime, big.NewInt(1))
	n := new(big.Int).Mod(pow, expMod)
	result := new(big.Int).Exp(&f.Number, n, &f.Prime)
	return NewFieldElement(result, &f.Prime)
}

func (f FieldElement) Mult(f2 FieldElement) (FieldElement, error) {
	if f.Prime.Cmp(&f2.Prime) != 0 {
		return FieldElement{}, errors.New("cannot mult two numbers in different field")
	}
	result := new(big.Int).Mul(&f.Number, &f2.Number)
	result.Mod(result, &f.Prime)
	return NewFieldElement(result, &f.Prime)
}

func (f FieldElement) Sub(f2 FieldElement) (FieldElement, error) {
	if f.Prime.Cmp(&f2.Prime) != 0 {
		return FieldElement{}, errors.New("cannot sub two numbers in different field")
	}
	result := new(big.Int).Sub(&f.Number, &f2.Number)
	result.Mod(result, &f.Prime)
	if result.Sign() < 0 {
		result.Add(result, &f.Prime)
	}
	return NewFieldElement(result, &f.Prime)
}

func (f FieldElement) Add(f2 FieldElement) (FieldElement, error) {
	if f.Prime.Cmp(&f2.Prime) != 0 {
		return FieldElement{}, errors.New("cannot add two numbers in different field")
	}
	result := new(big.Int).Add(&f.Number, &f2.Number)
	result.Mod(result, &f.Prime)
	return NewFieldElement(result, &f.Prime)
}

func (f FieldElement) Div(f2 FieldElement) (FieldElement, error) {
	if f.Prime.Cmp(&f2.Prime) != 0 {
		return FieldElement{}, errors.New("cannot add two numbers in different field")
	}
	inverse := new(big.Int).Exp(&f2.Number, new(big.Int).Sub(&f.Prime, big.NewInt(2)), &f.Prime)
	result := new(big.Int).Mod(new(big.Int).Mul(&f.Number, inverse), &f.Prime)
	return NewFieldElement(result, &f.Prime)
}
