package sigs

import (
	"log"
	"math/big"

	btcutils "github.com/bithyve/research/utils"
)

func testSwitch() {
	x, err := btcutils.NewPrivateKey()
	if err != nil {
		log.Fatal(err)
	}

	r, err := btcutils.NewPrivateKey()
	if err != nil {
		log.Fatal(err)
	}

	xG := btcutils.PointFromPrivkey(x)

	shaBytes := btcutils.Sha256(Curve.Params().Gx.Bytes(), Curve.Params().Gy.Bytes()) // btcutils.Sha256(G)

	var H btcutils.Point
	H.Set(Curve.ScalarBaseMult(shaBytes)) // H = btcutils.Point(btcutils.Sha256(G))

	var rH btcutils.Point
	rH.Set(Curve.ScalarMult(H.X, H.Y, r.Bytes()))

	var pedersen btcutils.Point
	pedersen.Add(xG, rH)

	var rG btcutils.Point
	rG = btcutils.PointFromPrivkey(r)

	// xG+(r+H(xG+rH||rG))H is the switch commitment
	// let the ugly term be p ie the commitment is xG = pH
	insideHash := btcutils.Sha256(pedersen.Bytes(), rG.Bytes())
	insideHashNumber := new(big.Int).SetBytes(insideHash)
	p := new(big.Int).Add(r, insideHashNumber)

	var pH btcutils.Point
	pH.Set(Curve.ScalarMult(H.X, H.Y, p.Bytes()))

	var switchCmt btcutils.Point
	switchCmt.Add(xG, pH)

	log.Println("switch commitment: ", switchCmt)
}
