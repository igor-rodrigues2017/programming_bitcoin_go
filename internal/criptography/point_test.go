package criptography_test

import (
	"math/big"
	"testing"

	"github.com/igor-rodrigues2017/programming_bitcoin_go/internal/criptography"
)

func pt(a, b, x, y *big.Int) criptography.Point {
	p, _ := criptography.NewPoint(a, b, x, y)
	return p
}

func ptInf(a, b *big.Int) criptography.Point {
	p, _ := criptography.NewPointAtInfinity(a, b)
	return p
}

func TestNewPoint(t *testing.T) {
	tests := []struct {
		name    string
		a, b    *big.Int
		x, y    *big.Int
		want    criptography.Point
		wantErr bool
	}{
		{
			"Should create a valid point without errors",
			big.NewInt(5), big.NewInt(7), big.NewInt(-1), big.NewInt(-1),
			pt(big.NewInt(5), big.NewInt(7), big.NewInt(-1), big.NewInt(-1)),
			false,
		},
		{
			"Should create another valid point without errors",
			big.NewInt(5), big.NewInt(7), big.NewInt(18), big.NewInt(77),
			pt(big.NewInt(5), big.NewInt(7), big.NewInt(18), big.NewInt(77)),
			false,
		},
		{
			"Should return error when point is not on the curve",
			big.NewInt(5), big.NewInt(7), big.NewInt(-1), big.NewInt(-2),
			criptography.Point{},
			true,
		},
		{
			"Should not create the point at infinity when x and y is nil ",
			big.NewInt(5), big.NewInt(7), nil, nil,
			criptography.Point{},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := criptography.NewPoint(tt.a, tt.b, tt.x, tt.y)
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
	tests := []struct {
		name    string
		a, b    *big.Int
		wantErr bool
	}{
		{
			"Should create point at infinity without errors",
			big.NewInt(5), big.NewInt(7),
			false,
		},
		{
			"Should return error when a and b are nil",
			nil, nil,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := criptography.NewPointAtInfinity(tt.a, tt.b)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("NewPointAtInfinity() failed: %v", err)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("NewPointAtInfinity() succeeded unexpectedly")
			}
		})
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
			pt(big.NewInt(5), big.NewInt(7), big.NewInt(18), big.NewInt(77)),
			ptInf(big.NewInt(5), big.NewInt(7)),
			pt(big.NewInt(5), big.NewInt(7), big.NewInt(18), big.NewInt(77)),
			false,
		},
		{
			"Should add without error Points are additive inverse",
			pt(big.NewInt(5), big.NewInt(7), big.NewInt(18), big.NewInt(77)),
			pt(big.NewInt(5), big.NewInt(7), big.NewInt(18), big.NewInt(-77)),
			ptInf(big.NewInt(5), big.NewInt(7)),
			false,
		},
		{
			"Should add without error Points for when x are differents",
			pt(big.NewInt(5), big.NewInt(7), big.NewInt(2), big.NewInt(5)),
			pt(big.NewInt(5), big.NewInt(7), big.NewInt(-1), big.NewInt(-1)),
			pt(big.NewInt(5), big.NewInt(7), big.NewInt(3), big.NewInt(-7)),
			false,
		},
		{
			"Should add without error the same point",
			pt(big.NewInt(5), big.NewInt(7), big.NewInt(-1), big.NewInt(-1)),
			pt(big.NewInt(5), big.NewInt(7), big.NewInt(-1), big.NewInt(-1)),
			pt(big.NewInt(5), big.NewInt(7), big.NewInt(18), big.NewInt(77)),
			false,
		},
		{
			"Should add without error the same point, and y is zero, return point at infinity",
			pt(big.NewInt(-1), big.NewInt(0), big.NewInt(1), big.NewInt(0)),
			pt(big.NewInt(-1), big.NewInt(0), big.NewInt(1), big.NewInt(0)),
			ptInf(big.NewInt(-1), big.NewInt(0)),
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, error := tt.pointA.Add(tt.pointB)
			if error != nil {
				if !tt.wantErr {
					t.Errorf("Add() = %v, want %v", got, tt.want)
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
