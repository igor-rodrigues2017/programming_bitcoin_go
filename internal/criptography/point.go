package criptography

import (
	"errors"
	"math/big"
)

type Point struct {
	a *big.Int
	b *big.Int
	x *big.Int
	y *big.Int
}

func NewPoint(a, b, x, y *big.Int) (Point, error) {
	p := Point{
		a: new(big.Int).Set(a),
		b: new(big.Int).Set(b),
		x: new(big.Int).Set(x),
		y: new(big.Int).Set(y),
	}
	if !p.IsOnCurve() {
		return Point{}, errors.New("point is not on the curve")
	}
	return p, nil
}

func (p Point) Equal(other Point) bool {
	return p.a.Cmp(other.a) == 0 &&
		p.b.Cmp(other.b) == 0 &&
		p.x.Cmp(other.x) == 0 &&
		p.y.Cmp(other.y) == 0
}

func (p Point) IsOnCurve() bool {
	lhs := new(big.Int).Mul(p.y, p.y)

	x3 := new(big.Int).Mul(p.x, p.x)
	x3.Mul(x3, p.x)

	ax := new(big.Int).Mul(p.a, p.x)

	rhs := new(big.Int).Add(x3, ax)
	rhs.Add(rhs, p.b)
	return lhs.Cmp(rhs) == 0
}
