package criptography_test

import (
	"math/big"
	"testing"

	"github.com/igor-rodrigues2017/programming_bitcoin_go/internal/criptography"
)

func fe(number, prime int64) criptography.FieldElement {
	f, _ := criptography.NewFieldElement(big.NewInt(number), big.NewInt(prime))
	return f
}

func TestNewFieldElement(t *testing.T) {
	tests := []struct {
		name    string
		number  int64
		prime   int64
		want    criptography.FieldElement
		wantErr bool
	}{
		{
			"Should Return a FieldElement without errors",
			7, 13,
			fe(7, 13),
			false,
		},
		{
			"Should Return an error when number is bigger than prime",
			32, 13,
			criptography.FieldElement{},
			true,
		},
		{
			"Should Return an error when number is shorter than zero",
			-1, 13,
			criptography.FieldElement{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := criptography.NewFieldElement(big.NewInt(tt.number), big.NewInt(tt.prime))
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("NewFieldElement() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("NewFieldElement() succeeded unexpectedly")
			}
			if !got.Equal(tt.want) {
				t.Errorf("NewFieldElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumElements(t *testing.T) {
	testCases := []struct {
		name     string
		element1 criptography.FieldElement
		element2 criptography.FieldElement
		want     criptography.FieldElement
	}{
		{
			"Should sum two fields without errors",
			fe(7, 19), fe(8, 19), fe(15, 19),
		},
		{
			"Should return empty FieldElement when fields have different primes",
			fe(7, 19), fe(8, 23),
			criptography.FieldElement{},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.element1.Add(tt.element2)
			if !got.Equal(tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubElements(t *testing.T) {
	testCases := []struct {
		name     string
		element1 criptography.FieldElement
		element2 criptography.FieldElement
		want     criptography.FieldElement
	}{
		{
			"Should sub two fields without errors",
			fe(7, 19), fe(8, 19), fe(18, 19),
		},
		{
			"Should return empty FieldElement when fields have different primes",
			fe(7, 19), fe(8, 23),
			criptography.FieldElement{},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.element1.Sub(tt.element2)
			if !got.Equal(tt.want) {
				t.Errorf("Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultElements(t *testing.T) {
	testCases := []struct {
		name     string
		element1 criptography.FieldElement
		element2 criptography.FieldElement
		want     criptography.FieldElement
	}{
		{
			"Should mult two fields without errors",
			fe(3, 13), fe(12, 13), fe(10, 13),
		},
		{
			"Should return empty FieldElement when fields have different primes",
			fe(7, 19), fe(8, 23),
			criptography.FieldElement{},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.element1.Mult(tt.element2)
			if !got.Equal(tt.want) {
				t.Errorf("Mult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPowElements(t *testing.T) {
	testCases := []struct {
		name     string
		element1 criptography.FieldElement
		pow      int64
		want     criptography.FieldElement
	}{
		{
			"Should pow field element without errors",
			fe(3, 13), 3, fe(1, 13),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.element1.Pow(big.NewInt(tt.pow))
			if !got.Equal(tt.want) {
				t.Errorf("Pow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivElements(t *testing.T) {
	testCases := []struct {
		name     string
		element1 criptography.FieldElement
		divisor  criptography.FieldElement
		want     criptography.FieldElement
	}{
		{
			"Should division field element without errors",
			fe(2, 19), fe(7, 19), fe(3, 19),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.element1.Div(tt.divisor)
			if !got.Equal(tt.want) {
				t.Errorf("Div() = %v, want %v", got, tt.want)
			}
		})
	}
}
