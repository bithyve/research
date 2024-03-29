package utils

import (
	"math/big"
)

type Point struct {
	X *big.Int
	Y *big.Int
}

func (p *Point) Set(x, y *big.Int) {
	p.X = x
	p.Y = y
}

func (p *Point) AddCoords(x1, y1, x2, y2 *big.Int) {
	p.X, p.Y = Curve.Add(x1, y1, x2, y2)
}

func (p *Point) Add(x1, x2 Point) {
	p.X, p.Y = Curve.Add(x1.X, x1.Y, x2.X, x2.Y)
}

func (p *Point) Cmp (a Point) bool {
	return p.X.Cmp(a.X) == 0 && p.Y.Cmp(a.Y) == 0
}

func Add(x1, x2 Point) Point {
	var p Point
	p.X, p.Y = Curve.Add(x1.X, x1.Y, x2.X, x2.Y)
	return p
}

func (p *Point) ScalarMult(a []byte) {
	p.X, p.Y = Curve.ScalarMult(p.X, p.Y, a)
}

func ScalarMult(P Point, a []byte) Point {
	var p Point
	p.X, p.Y = Curve.ScalarMult(P.X, P.Y, a)
	return p
}

func ScalarBaseMult(a []byte) Point {
	var p Point
	p.X, p.Y = Curve.ScalarBaseMult(a)
	return p
}

func Sub(x1, x2 Point) Point {
	var p Point

	negx2 := new(big.Int).Neg(x2.Y)

	p.Set(Curve.Add(x1.X, x1.Y, x2.X, new(big.Int).Mod(negx2, Curve.P)))

	return p
}

func (p *Point) Bytes() []byte {
	return append(p.X.Bytes(), p.Y.Bytes()...)
}

type Elgamal struct {
	X Point // since each one of these is a curve point
	Y Point // since each one of these is a curve point
}

func (e *Elgamal) Set(x, y Point) {
	e.X = x
	e.Y = y
}

func PubkeyPointsFromPrivkey(privkey *big.Int) (*big.Int, *big.Int) {
	x, y := Curve.ScalarBaseMult(privkey.Bytes())
	return x, y
}

func PointFromPrivkey(privkey *big.Int) Point {
	var x Point
	x.X, x.Y = Curve.ScalarBaseMult(privkey.Bytes())
	return x
}
