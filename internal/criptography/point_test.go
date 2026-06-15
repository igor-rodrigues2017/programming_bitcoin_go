package criptography_test

import (
	"testing"

	"github.com/igor-rodrigues2017/programming_bitcoin_go/internal/criptography"
)

func modPrime(n, prime int64) int64 {
	return ((n % prime) + prime) % prime
}

func pt(a, b, x, y, prime int64) criptography.Point {
	p, _ := criptography.NewPoint(modPrime(a, prime), modPrime(b, prime), modPrime(x, prime), modPrime(y, prime), prime)
	return p
}

func ptInf(a, b, prime int64) criptography.Point {
	return criptography.NewPointAtInfinity(modPrime(a, prime), modPrime(b, prime), prime)
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
			pt(5, 7, -1, -1, 223),
			false,
		},
		{
			"Should create another valid point without errors",
			5, 7, 18, 77,
			pt(5, 7, 18, 77, 223),
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
			got, gotErr := criptography.NewPoint(modPrime(tt.a, 223), modPrime(tt.b, 223), modPrime(tt.x, 223), modPrime(tt.y, 223), 223)
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
	p := criptography.NewPointAtInfinity(modPrime(5, 223), modPrime(7, 223), 223)
	inf := ptInf(5, 7, 223)
	if !p.Equal(inf) {
		t.Errorf("NewPointAtInfinity() = %v, want %v", p, inf)
	}
}

func TestAdditionPoint(t *testing.T) {
	tests := []struct {
		name   string
		pointA criptography.Point
		pointB criptography.Point
		want   criptography.Point
	}{
		{
			"Should add without error one Point and one Point at infinity",
			pt(5, 7, 18, 77, 223),
			ptInf(5, 7, 223),
			pt(5, 7, 18, 77, 223),
		},
		{
			"Should add without error Points are additive inverse",
			pt(5, 7, 18, 77, 223),
			pt(5, 7, 18, -77, 223),
			ptInf(5, 7, 223),
		},
		{
			"Should add without error Points for when x are differents",
			pt(5, 7, 2, 5, 223),
			pt(5, 7, -1, -1, 223),
			pt(5, 7, 3, -7, 223),
		},
		{
			"Should add without error the same point",
			pt(5, 7, -1, -1, 223),
			pt(5, 7, -1, -1, 223),
			pt(5, 7, 18, 77, 223),
		},
		{
			"Should add without error the same point, and y is zero, return point at infinity",
			pt(-1, 0, 1, 0, 223),
			pt(-1, 0, 1, 0, 223),
			ptInf(-1, 0, 223),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.pointA.Add(tt.pointB)
			if !got.Equal(tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScalarMultPoint(t *testing.T) {
	tests := []struct {
		name        string
		pointA      criptography.Point
		coefficient int64
		want        criptography.Point
	}{
		{
			"Should multiply scalar without error one Point",
			pt(0, 7, 15, 86, 223),
			7,
			ptInf(0, 7, 223),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.pointA.Mult(tt.coefficient)
			if !got.Equal(tt.want) {
				t.Errorf("Mult() = %v, want %v", got, tt.want)
			}
		})
	}
}
