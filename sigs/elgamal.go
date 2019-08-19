package sigs

import (
	"log"

	btcutils "github.com/bithyve/research/utils"
)

func testElgamal() {
	// an elgamal commitment is a small upgrade from Pedersen commitments
	// xG + rH - Pedersen
	// (xG + rH, rG) - Elgamal

	// there is a small nuanace to how this work. lets assume we have xG + rH in Pedersen
	// if a person can enumerate over the entire input space, they can find r and as a result
	// binding is broken (the commitment binds to r but we now have a fake r). If binding
	// is broken, in some applications like CT, one can print infinite money. So pedersen
	// commitments have perfect hiding (no one can predict what value exists) but computational
	// binding (a person with finitely infinite resources can find another r)

	// In Elgamal, we commit to anohter point - rG. We can compute this since we know r (used anyway for Pedersen)
	// but where's the difference?
	// Lets assume an attacker finds r like in the above case. since a Pedersen commitment is xG + rH, they can
	// commit to another value r' where rH = r'H. In Elgamal, the mapping from r to rG is one-one, so we
	// get perfect binding (no one can commit to another value even if they have resources to find r). But
	// an attacker with resources can find r and then compute C - rH to find xG (the hidden commitment value).
	// Hence, Pedersen offers perfect binding and computational hiding

	// note that we're only implementing the commitment scheme here not the signature scheme. Signature scheme could be
	// AOS or something similar

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

	var elgamal btcutils.Elgamal
	elgamal.Set(pedersen, rG)

	log.Println("elgamal: ", elgamal)
}
