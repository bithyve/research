package utils

import (
	"crypto/rand"
	"io"
	"log"
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
	p.X, p.Y = Curve.Add(x1.X, x2.X, x1.Y, x2.Y)
}

func (p *Point) ScalarMult(a []byte) {
	p.X, p.Y = Curve.ScalarMult(p.X, p.Y, a)
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
