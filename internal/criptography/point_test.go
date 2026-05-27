package criptography_test

import (
	"math/big"
	"testing"

	"github.com/igor-rodrigues2017/programming_bitcoin_go/internal/criptography"
)

func pt(a, b, x, y int64) criptography.Point {
	p, _ := criptography.NewPoint(big.NewInt(a), big.NewInt(b), big.NewInt(x), big.NewInt(y))
	return p
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
			got, gotErr := criptography.NewPoint(
				big.NewInt(tt.a), big.NewInt(tt.b),
				big.NewInt(tt.x), big.NewInt(tt.y),
			)
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

