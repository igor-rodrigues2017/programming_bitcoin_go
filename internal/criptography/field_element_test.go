package criptography_test

import (
	"testing"

	"github.com/igor-rodrigues2017/programming_bitcoin_go/internal/criptography"
)

func TestNewFieldElement(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		number  int64
		prime   int64
		want    criptography.FieldElement
		wantErr bool
	}{
		{
			"Should Return a FieldElement without errors",
			7,
			13,
			criptography.FieldElement{7, 13},
			false,
		},
		{
			"Should Return an error when number is bigger than prime",
			32,
			13,
			criptography.FieldElement{},
			true,
		},
		{
			"Should Return an error when number is shorter than zero",
			-1,
			13,
			criptography.FieldElement{7, 13},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := criptography.NewFieldElement(tt.number, tt.prime)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("NewFieldElement() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("NewFieldElement() succeeded unexpectedly")
			}
			if got != tt.want {
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
		wantErr  bool
	}{
		{
			"Should sum two fields without errors",
			criptography.FieldElement{7, 19},
			criptography.FieldElement{8, 19},
			criptography.FieldElement{15, 19},
			false,
		},
		{
			"Should return an error when are Fields different",

			criptography.FieldElement{7, 19},
			criptography.FieldElement{8, 23},
			criptography.FieldElement{},
			true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := tt.element1.Add(tt.element2)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Add() elements failed: %v,", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Add() succeeded unexpectedly")
			}
			if got != tt.want {
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
		wantErr  bool
	}{
		{
			"Should sub two fields without errors",
			criptography.FieldElement{7, 19},
			criptography.FieldElement{8, 19},
			criptography.FieldElement{18, 19},
			false,
		},
		{
			"Should return an error when are Fields different",

			criptography.FieldElement{7, 19},
			criptography.FieldElement{8, 23},
			criptography.FieldElement{},
			true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := tt.element1.Sub(tt.element2)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Sub() elements failed: %v,", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Sub() succeeded unexpectedly")
			}
			if got != tt.want {
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
		wantErr  bool
	}{
		{
			"Should sub two fields without errors",
			criptography.FieldElement{3, 13},
			criptography.FieldElement{12, 13},
			criptography.FieldElement{10, 13},
			false,
		},
		{
			"Should return an error when are Fields different",

			criptography.FieldElement{7, 19},
			criptography.FieldElement{8, 23},
			criptography.FieldElement{},
			true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := tt.element1.Mult(tt.element2)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Mult() elements failed: %v,", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Mult() succeeded unexpectedly")
			}
			if got != tt.want {
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
		wantErr  bool
	}{
		{
			"Should sub two fields without errors",
			criptography.FieldElement{3, 13},
			3,
			criptography.FieldElement{1, 13},
			false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := tt.element1.Pow(tt.pow)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("pow() elements failed: %v,", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("pow() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("pow() = %v, want %v", got, tt.want)
			}
		})
	}
}
