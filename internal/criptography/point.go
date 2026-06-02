package criptography

import (
	"errors"
	"math/big"
)

var (
	ErrNilCoordinates  = errors.New("use NewPointAtInfinity for points at infinity")
	ErrNotOnCurve      = errors.New("point is not on the curve")
	ErrNilCurveParams  = errors.New("a and b must not be nil")
	ErrDifferentCurves = errors.New("points are not on the same curve")
)

type Point struct {
	a *big.Int
	b *big.Int
	x *big.Int
	y *big.Int
}

func NewPoint(a, b, x, y *big.Int) (Point, error) {
	if x == nil || y == nil {
		return Point{}, ErrNilCoordinates
	}
	p := Point{
		a: new(big.Int).Set(a),
		b: new(big.Int).Set(b),
		x: new(big.Int).Set(x),
		y: new(big.Int).Set(y),
	}
	if !p.IsOnCurve() {
		return Point{}, ErrNotOnCurve
	}
	return p, nil
}

func NewPointAtInfinity(a, b *big.Int) (Point, error) {
	if a == nil || b == nil {
		return Point{}, ErrNilCurveParams
	}
	return Point{
		a: new(big.Int).Set(a),
		b: new(big.Int).Set(b),
	}, nil
}

func (p *Point) Add(other Point) (Point, error) {
	if p.a.Cmp(other.a) != 0 || p.b.Cmp(other.b) != 0 {
		return Point{}, ErrDifferentCurves
	}
	if p.x == nil {
		return other, nil
	}
	if other.x == nil {
		return *p, nil
	}
	if p.x.Cmp(other.x) == 0 && p.y.Cmp(other.y) != 0 {
		return NewPointAtInfinity(p.a, p.b)
	}

	diffY := new(big.Int).Sub(other.y, p.y)
	diffX := new(big.Int).Sub(other.x, p.x)
	s := diffY.Div(diffY, diffX)

	x3 := new(big.Int).Mul(s, s)
	x3.Sub(x3, p.x).Sub(x3, other.x)

	y3 := new(big.Int).Mul(s, new(big.Int).Sub(p.x, x3))
	y3.Sub(y3, p.y)

	return NewPoint(new(big.Int).Set(p.a), new(big.Int).Set(p.b), x3, y3)
}

func (p *Point) Equal(other Point) bool {
	if p.a.Cmp(other.a) != 0 || p.b.Cmp(other.b) != 0 {
		return false
	}
	if p.x == nil && other.x == nil {
		return true
	}
	if p.x == nil || other.x == nil {
		return false
	}
	return p.x.Cmp(other.x) == 0 && p.y.Cmp(other.y) == 0
}

func (p *Point) IsOnCurve() bool {
	lhs := new(big.Int).Mul(p.y, p.y)

	x3 := new(big.Int).Mul(p.x, p.x)
	x3.Mul(x3, p.x)

	ax := new(big.Int).Mul(p.a, p.x)

	rhs := new(big.Int).Add(x3, ax)
	rhs.Add(rhs, p.b)
	return lhs.Cmp(rhs) == 0
}
