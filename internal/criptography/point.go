package criptography

import (
	"errors"
	"fmt"
	"log"
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

func (p Point) String() string {
	if p.x == nil {
		return fmt.Sprintf("Point(infinity)_(%v,%v)", p.a, p.b)
	}
	return fmt.Sprintf("Point(%v,%v)_(%v,%v)", p.x, p.y, p.a, p.b)
}

func NewPoint(a, b, x, y, prime int64) (Point, error) {
	aFE, err := NewFieldElement(a, prime)
	if err != nil {
		return Point{}, err
	}
	bFE, err := NewFieldElement(b, prime)
	if err != nil {
		return Point{}, err
	}
	xFE, err := NewFieldElement(x, prime)
	if err != nil {
		return Point{}, err
	}
	yFE, err := NewFieldElement(y, prime)
	if err != nil {
		return Point{}, err
	}
	return NewPointFromFieldElements(aFE, bFE, xFE, yFE)
}

func NewPointFromString(a, b, x, y, prime string, base int) (Point, error) {
	aFE, err := NewFieldElementFromString(a, prime, base)
	if err != nil {
		return Point{}, err
	}
	bFE, err := NewFieldElementFromString(b, prime, base)
	if err != nil {
		return Point{}, err
	}
	xFE, err := NewFieldElementFromString(x, prime, base)
	if err != nil {
		return Point{}, err
	}
	yFE, err := NewFieldElementFromString(y, prime, base)
	if err != nil {
		return Point{}, err
	}
	return NewPointFromFieldElements(aFE, bFE, xFE, yFE)
}

func NewPointFromFieldElements(a, b, x, y FieldElement) (Point, error) {
	xFE, _ := newFieldElement(x.number, x.prime)
	yFE, _ := newFieldElement(y.number, y.prime)
	p := Point{a: a, b: b, x: &xFE, y: &yFE}
	if !p.IsOnCurve() {
		return Point{}, ErrNotOnCurve
	}
	return p, nil
}

func NewPointAtInfinity(a, b, prime int64) Point {
	aFE, err := NewFieldElement(a, prime)
	if err != nil {
		return Point{}
	}
	bFE, err := NewFieldElement(b, prime)
	if err != nil {
		return Point{}
	}
	return NewPointAtInfinityFromFieldElements(aFE, bFE)
}

func NewPointAtInfinityFromFieldElements(a, b FieldElement) Point {
	return Point{a: a, b: b}
}

func (p *Point) Add(other Point) Point {
	if !p.a.Equal(other.a) || !p.b.Equal(other.b) {
		log.Println(ErrDifferentCurves)
		return Point{}
	}
	if p.x == nil {
		return other
	}
	if other.x == nil {
		return *p
	}

	zero, _ := newFieldElement(big.NewInt(0), p.x.prime)
	if (p.x.Equal(*other.x) && !p.y.Equal(*other.y)) || (p.Equal(other) && p.y.Equal(zero)) {
		return NewPointAtInfinityFromFieldElements(p.a, p.b)
	}

	var s FieldElement
	if p.Equal(other) {
		s = p.tangentSlope()
	} else {
		s = p.slope(other)
	}

	// x3 = s^2 - x1 - x2
	s2 := s.Pow(2)
	x3 := s2.Sub(*p.x)
	x3 = x3.Sub(*other.x)

	// y3 = s*(x1 - x3) - y1
	x1SubX3 := p.x.Sub(x3)
	y3 := s.Mult(x1SubX3)
	y3 = y3.Sub(*p.y)

	result, err := NewPointFromFieldElements(p.a, p.b, x3, y3)
	if err != nil {
		log.Println(err)
		return Point{}
	}
	return result
}

func (p *Point) Mult(coefficient int64) Point {
	return p.multBig(big.NewInt(coefficient))
}

func (p *Point) MultBig(coefficient *big.Int) Point {
	return p.multBig(coefficient)
}

func (p *Point) multBig(coefficient *big.Int) Point {
	coef := new(big.Int).Set(coefficient)
	current := *p
	result := NewPointAtInfinityFromFieldElements(p.a, p.b)
	for coef.Cmp(big.NewInt(0)) > 0 {
		if coef.Bit(0) == 1 {
			result = result.Add(current)
		}
		current = current.Add(current)
		coef.Rsh(coef, 1)
	}
	return result
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
	lhs := p.y.Pow(2)
	x3 := p.x.Pow(3)
	ax := p.a.Mult(*p.x)
	rhs := x3.Add(ax)
	rhs = rhs.Add(p.b)
	return lhs.Equal(rhs)
}

func (p *Point) tangentSlope() FieldElement {
	// s = (3*x^2 + a) / (2*y)
	x2 := p.x.Pow(2)
	three, _ := newFieldElement(big.NewInt(3), p.x.prime)
	num := three.Mult(x2)
	num = num.Add(p.a)
	two, _ := newFieldElement(big.NewInt(2), p.x.prime)
	den := two.Mult(*p.y)
	return num.Div(den)
}

func (p *Point) slope(other Point) FieldElement {
	// s = (y2 - y1) / (x2 - x1)
	diffY := other.y.Sub(*p.y)
	diffX := other.x.Sub(*p.x)
	return diffY.Div(diffX)
}
