package sigs

import (
	"log"
	"math/big"

	btcutils "github.com/bithyve/research/utils"
)

var Curve = btcutils.Curve

func Create21AOSSig() (*big.Int, *big.Int, *big.Int, btcutils.Point, btcutils.Point, []byte) {
	// Lets assume two parties - Brian and Dom.
	// Lets assume Brian and Dom's P are Pb and Pd. Lets assume their private keys
	// are b and d.
	// Then lets assume Dom wants to create a ring signature over C1 and C2 where
	// C1 = xG + 1H and C2 = xG where x is the blinding factor

	b, err := btcutils.NewPrivateKey()
	if err != nil {
		log.Fatal(err)
	}

	d, err := btcutils.NewPrivateKey()
	if err != nil {
		log.Fatal(err)
	}

	Pb := btcutils.PointFromPrivkey(b)
	Pd := btcutils.PointFromPrivkey(d)

	x, err := btcutils.NewPrivateKey()
	if err != nil {
		log.Fatal(err)
	}

	xG := btcutils.PointFromPrivkey(x)

	shaBytes := btcutils.Sha256(Curve.Params().Gx.Bytes(), Curve.Params().Gy.Bytes()) // btcutils.Sha256(G)

	var H btcutils.Point
	H.Set(Curve.ScalarBaseMult(shaBytes)) // H = btcutils.Point(btcutils.Sha256(G))

	one := []byte{1}
	oneH := btcutils.ScalarMult(H, one)

	C1 := btcutils.Add(xG, oneH) // xG + 1H

	var C2 btcutils.Point
	C2.Set(Curve.ScalarBaseMult(x.Bytes())) // xG

	var m []byte
	m = append(m, append(C1.Bytes(), C2.Bytes()...)...)

	kd, err := btcutils.NewPrivateKey() // random nonce for ring sig
	if err != nil {
		log.Fatal(err)
	}

	var Kd btcutils.Point
	Kd.Set(Curve.ScalarBaseMult(kd.Bytes())) // K = kd*G

	BrianNodeNumber := []byte{2} // assume brian has node number 2

	eb := btcutils.Sha256(Kd.Bytes(), m, BrianNodeNumber)

	sb, err := btcutils.NewPrivateKey() // choose a signature sb at random fro brian
	if err != nil {
		log.Fatal(err)
	}

	var sbG btcutils.Point
	sbG.Set(Curve.ScalarBaseMult(sb.Bytes())) // sb*G

	ebPb := btcutils.ScalarMult(Pb, eb) // eb * Pb

	Kb := btcutils.Sub(sbG, ebPb) // Kb = sb*G - eb*Pb

	DomNodeNumber := []byte{1}

	ed := btcutils.Sha256(Kb.Bytes(), m, DomNodeNumber) // ed = H(Kb || m || D)

	edd := new(big.Int).Mul(new(big.Int).SetBytes(ed), d) // ed * d

	sd := new(big.Int).Add(edd, kd) // ed*d + kd

	return new(big.Int).SetBytes(eb), sb, sd, Pb, Pd, m
}

func Verify21AOSSig() {
	eb, sb, sd, Pb, Pd, m := Create21AOSSig()

	BrianNodeNumber := []byte{2} // assume brian has node number 2
	DomNodeNumber := []byte{1}

	var sbG btcutils.Point
	sbG.Set(Curve.ScalarBaseMult(sb.Bytes())) // sb*G

	ebPb := btcutils.ScalarMult(Pb, eb.Bytes()) // eb*Pb

	Kb := btcutils.Sub(sbG, ebPb) // Kb = sb*G - eb*Pb

	ed := btcutils.Sha256(Kb.Bytes(), m, DomNodeNumber)
	// log.Println("ed: ", ed)

	var sdG btcutils.Point
	sdG.Set(Curve.ScalarBaseMult(sd.Bytes()))

	edPd := btcutils.ScalarMult(Pd, ed)

	Kd := btcutils.Sub(sdG, edPd)

	ebCheck := btcutils.Sha256(Kd.Bytes(), m, BrianNodeNumber)
	ebCheckInt := new(big.Int).SetBytes(ebCheck)

	if ebCheckInt.Cmp(eb) != 0 {
		log.Fatal("Signatures don't match")
	} else {
		log.Println("Ring signatures validated")
	}
}

func main() {
	Verify21AOSSig()
}
