package criptography_test

import (
	"testing"

	"github.com/igor-rodrigues2017/programming_bitcoin_go/internal/criptography"
)

const testPrime = int64(10007)

// fprime cria um FieldElement no campo testPrime, aceitando negativos via redução modular.
// Ex: fprime(-1) → 10006, porque -1 ≡ p-1 (mod p).
func fprime(n int64) criptography.FieldElement {
	p := int64(testPrime)
	return fe(((n%p)+p)%p, testPrime)
}

func pt(a, b, x, y int64) criptography.Point {
	p, _ := criptography.NewPoint(fprime(a), fprime(b), fprime(x), fprime(y))
	return p
}

func ptInf(a, b int64) criptography.Point {
	return criptography.NewPointAtInfinity(fprime(a), fprime(b))
}

func TestNewPoint(t *testing.T) {
	tests := []struct {
		name    string
		a, b    int64
		x, y    int64
		want    criptography.Point
		wantErr bool
	}{
		{
			"Should create a valid point without errors",
			5, 7, -1, -1,
			pt(5, 7, -1, -1),
			false,
		},
		{
			"Should create another valid point without errors",
			5, 7, 18, 77,
			pt(5, 7, 18, 77),
			false,
		},
		{
			"Should return error when point is not on the curve",
			5, 7, -1, -2,
			criptography.Point{},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := criptography.NewPoint(fprime(tt.a), fprime(tt.b), fprime(tt.x), fprime(tt.y))
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("NewPoint() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("NewPoint() succeeded unexpectedly")
			}
			if !got.Equal(tt.want) {
				t.Errorf("NewPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPointAtInfinity(t *testing.T) {
	p := criptography.NewPointAtInfinity(fprime(5), fprime(7))
	inf := ptInf(5, 7)
	if !p.Equal(inf) {
		t.Errorf("NewPointAtInfinity() = %v, want %v", p, inf)
	}
}

func TestAdditionPoint(t *testing.T) {
	tests := []struct {
		name    string
		pointA  criptography.Point
		pointB  criptography.Point
		want    criptography.Point
		wantErr bool
	}{
		{
			"Should add without error one Point and one Point at infinity",
			pt(5, 7, 18, 77),
			ptInf(5, 7),
			pt(5, 7, 18, 77),
			false,
		},
		{
			"Should add without error Points are additive inverse",
			pt(5, 7, 18, 77),
			pt(5, 7, 18, -77),
			ptInf(5, 7),
			false,
		},
		{
			"Should add without error Points for when x are differents",
			pt(5, 7, 2, 5),
			pt(5, 7, -1, -1),
			pt(5, 7, 3, -7),
			false,
		},
		{
			"Should add without error the same point",
			pt(5, 7, -1, -1),
			pt(5, 7, -1, -1),
			pt(5, 7, 18, 77),
			false,
		},
		{
			"Should add without error the same point, and y is zero, return point at infinity",
			pt(-1, 0, 1, 0),
			pt(-1, 0, 1, 0),
			ptInf(-1, 0),
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.pointA.Add(tt.pointB)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("Add() error = %v", err)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Add() succeeded unexpectedly")
			}
			if !got.Equal(tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
