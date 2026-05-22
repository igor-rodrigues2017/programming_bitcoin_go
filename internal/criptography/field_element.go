package criptography

import (
	"errors"
	"math"
)

type FieldElement struct {
	Number int64 `json:"Number"`
	Prime  int64 `json:"Prime"`
}

func (f FieldElement) Pow(pow int64) (FieldElement, error) {
	result := math.Pow(float64(f.Number), float64(pow))
	mod := Mod(int64(result), f.Prime)
	return NewFieldElement(mod, f.Prime)
}

func (f FieldElement) Mult(element2 FieldElement) (FieldElement, error) {
	if f.Prime != element2.Prime {
		return FieldElement{}, errors.New("cannot mult two numbers in different field")
	}
	number := Mod(f.Number*element2.Number, f.Prime)
	return NewFieldElement(number, f.Prime)
}

func (f FieldElement) Sub(f2 FieldElement) (FieldElement, error) {
	if f.Prime != f2.Prime {
		return FieldElement{}, errors.New("cannot sub two numbers in different field")
	}
	num := Mod(f.Number-f2.Number, f.Prime)

	return NewFieldElement(num, f.Prime)
}

func (f *FieldElement) Add(f2 FieldElement) (FieldElement, error) {
	if f.Prime != f2.Prime {
		return FieldElement{}, errors.New("cannot add two numbers in different field")
	}
	num := Mod((f.Number + f2.Number), f.Prime)
	return NewFieldElement(num, f.Prime)
}

func NewFieldElement(number int64, prime int64) (FieldElement, error) {
	if number >= prime || number < 0 {
		return FieldElement{}, errors.New("number not in field range")
	}
	return FieldElement{number, prime}, nil
}
