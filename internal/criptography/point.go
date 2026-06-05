package criptography

import (
	"errors"
	"math/big"
)

var (
	ErrNotOnCurve      = errors.New("point is not on the curve")
	ErrDifferentCurves = errors.New("points are not on the same curve")
)

type Point struct {
	a FieldElement
	b FieldElement
	x *FieldElement
	y *FieldElement
}

func NewPoint(a, b, x, y FieldElement) (Point, error) {
	xFE, _ := NewFieldElement(x.Number, x.Prime)
	yFE, _ := NewFieldElement(y.Number, y.Prime)
	p := Point{a: a, b: b, x: &xFE, y: &yFE}
	if !p.IsOnCurve() {
		return Point{}, ErrNotOnCurve
	}
	return p, nil
}

func NewPointAtInfinity(a, b FieldElement) Point {
	return Point{a: a, b: b}
}

func (p *Point) tangentSlope() (FieldElement, error) {
	// s = (3*x^2 + a) / (2*y)
	x2, err := p.x.Pow(big.NewInt(2))
	if err != nil {
		return FieldElement{}, err
	}
	three, _ := NewFieldElement(big.NewInt(3), p.x.Prime)
	num, err := three.Mult(x2)
	if err != nil {
		return FieldElement{}, err
	}
	num, err = num.Add(p.a)
	if err != nil {
		return FieldElement{}, err
	}
	two, _ := NewFieldElement(big.NewInt(2), p.x.Prime)
	den, err := two.Mult(*p.y)
	if err != nil {
		return FieldElement{}, err
	}
	return num.Div(den)
}

func (p *Point) slope(other Point) (FieldElement, error) {
	// s = (y2 - y1) / (x2 - x1)
	diffY, err := other.y.Sub(*p.y)
	if err != nil {
		return FieldElement{}, err
	}
	diffX, err := other.x.Sub(*p.x)
	if err != nil {
		return FieldElement{}, err
	}
	return diffY.Div(diffX)
}

func (p *Point) Add(other Point) (Point, error) {
	if !p.a.Equal(other.a) || !p.b.Equal(other.b) {
		return Point{}, ErrDifferentCurves
	}
	if p.x == nil {
		return other, nil
	}
	if other.x == nil {
		return *p, nil
	}

	zero, _ := NewFieldElement(big.NewInt(0), p.x.Prime)
	if (p.x.Equal(*other.x) && !p.y.Equal(*other.y)) || (p.Equal(other) && p.y.Equal(zero)) {
		return NewPointAtInfinity(p.a, p.b), nil
	}

	var s FieldElement
	var err error
	if p.Equal(other) {
		s, err = p.tangentSlope()
	} else {
		s, err = p.slope(other)
	}
	if err != nil {
		return Point{}, err
	}

	// x3 = s^2 - x1 - x2
	s2, err := s.Pow(big.NewInt(2))
	if err != nil {
		return Point{}, err
	}
	x3, err := s2.Sub(*p.x)
	if err != nil {
		return Point{}, err
	}
	x3, err = x3.Sub(*other.x)
	if err != nil {
		return Point{}, err
	}

	// y3 = s*(x1 - x3) - y1
	x1SubX3, err := p.x.Sub(x3)
	if err != nil {
		return Point{}, err
	}
	y3, err := s.Mult(x1SubX3)
	if err != nil {
		return Point{}, err
	}
	y3, err = y3.Sub(*p.y)
	if err != nil {
		return Point{}, err
	}

	return NewPoint(p.a, p.b, x3, y3)
}

func (p *Point) Equal(other Point) bool {
	if !p.a.Equal(other.a) || !p.b.Equal(other.b) {
		return false
	}
	if p.x == nil && other.x == nil {
		return true
	}
	if p.x == nil || other.x == nil {
		return false
	}
	return p.x.Equal(*other.x) && p.y.Equal(*other.y)
}

func (p *Point) IsOnCurve() bool {
	// y^2 = x^3 + a*x + b
	lhs, err := p.y.Pow(big.NewInt(2))
	if err != nil {
		return false
	}
	x3, err := p.x.Pow(big.NewInt(3))
	if err != nil {
		return false
	}
	ax, err := p.a.Mult(*p.x)
	if err != nil {
		return false
	}
	rhs, err := x3.Add(ax)
	if err != nil {
		return false
	}
	rhs, err = rhs.Add(p.b)
	if err != nil {
		return false
	}
	return lhs.Equal(rhs)
}
