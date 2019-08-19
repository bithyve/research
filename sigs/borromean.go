package sigs

import (
	"log"
	"math/big"

	utils "github.com/Varunram/essentials/utils"
	btcutils "github.com/bithyve/research/utils"
)

func SubtractOnCurve(e []byte, P btcutils.Point) (btcutils.Point, *big.Int) {
	s, err := btcutils.NewPrivateKey()
	if err != nil {
		log.Fatal(err)
	}

	var sG btcutils.Point
	sG.Set(Curve.ScalarBaseMult(s.Bytes()))

	eP := btcutils.ScalarMult(P, e)

	return btcutils.Sub(sG, eP), s
}

func SubtractOnCurveS(e []byte, P btcutils.Point, s *big.Int) btcutils.Point {
	var sG btcutils.Point
	sG.Set(Curve.ScalarBaseMult(s.Bytes()))

	eP := btcutils.ScalarMult(P, e)

	return btcutils.Sub(sG, eP)
}

func testBorromean() {
	P := make(map[int]map[int]*btcutils.Point)
	x := make(map[int]*big.Int)

	P[0] = make(map[int]*btcutils.Point, 3)
	P[1] = make(map[int]*btcutils.Point, 3)

	for i := 0; i < 3; i++ {
		key, err := btcutils.NewPrivateKey()
		if err != nil {
			log.Fatal(err)
		}

		x[i] = key
		P[0][i] = new(btcutils.Point)
		P[0][i].Set(btcutils.PubkeyPointsFromPrivkey(key))
	}

	for i := 0; i < 3; i++ {
		key, err := btcutils.NewPrivateKey()
		if err != nil {
			log.Fatal(err)
		}

		x[i] = key
		P[1][i] = new(btcutils.Point)
		P[1][i].Set(btcutils.PubkeyPointsFromPrivkey(key))
	}

	jistar := []int{1, 2, 3, 4, 5, 6} // indices of signer in each ring

	M := btcutils.Sha256([]byte("cool"))

	e := make(map[int]map[int][]byte)
	s := make(map[int]map[int]*big.Int)
	k := make(map[int]*big.Int)

	for i := 0; i < 6; i++ {
		ktemp, err := btcutils.NewPrivateKey()
		if err != nil {
			log.Fatal(err)
		}
		k[i] = ktemp
	}

	// start signing
	for i, loop := range P {
		iByte, err := utils.ToByte(i)
		if err != nil {
			log.Fatal(err)
		}

		e[i] = make(map[int][]byte, len(loop))
		s[i] = make(map[int]*big.Int, len(loop))

		kiGx, kiGy := btcutils.PubkeyPointsFromPrivkey(k[i])
		kiG := append(kiGx.Bytes(), kiGy.Bytes()...)

		jstari := jistar[i]
		jstariByte, err := utils.ToByte(jstari)
		if err != nil {
			log.Fatal(err)
		}

		e[i][jstari+1] = btcutils.Sha256(M, kiG, iByte, jstariByte)

		mi := len(loop)
		for j := jstari + 1; j < mi-1; j++ {
			jByte, err := utils.ToByte(j)
			if err != nil {
				log.Fatal(err)
			}

			var temp btcutils.Point
			temp, s[i][j] = SubtractOnCurve(e[i][j], *P[i][j])
			e[i][j+1] = btcutils.Sha256(M, temp.Bytes(), iByte, jByte)
			log.Println(e[i][j+1])
		}
	}

	toBeHashed := []byte("")
	for i := 0; i <= len(P)-1; i++ {
		miMinusOne := 2
		var temp btcutils.Point
		temp, s[i][miMinusOne] = SubtractOnCurve(e[i][miMinusOne], *P[i][miMinusOne])
		toBeHashed = append(toBeHashed, temp.Bytes()...)
	}

	e0 := btcutils.Sha256(toBeHashed)

	for i := 0; i <= 1; i++ {
		iByte, err := utils.ToByte(i)
		if err != nil {
			log.Fatal(err)
		}

		e[i][0] = e0

		for j := 0; j < jistar[i]; j++ {
			var temp btcutils.Point
			jByte, err := utils.ToByte(j)
			if err != nil {
				log.Fatal(err)
			}

			temp, s[i][j] = SubtractOnCurve(e[i][j], *P[i][j])

			e[i][j+1] = btcutils.Sha256(M, temp.Bytes(), iByte, jByte)

			eijstari := new(big.Int).SetBytes(e[i][jistar[i]])

			xieijstari := new(big.Int).Mul(x[i], eijstari)

			s[i][jistar[i]] = new(big.Int).Add(k[i], xieijstari)
		}
	}

	// log.Println("e0: ", e0)
	// log.Println("sigs: ", s)
	log.Println("e: ", e)

	ex := make(map[int]map[int][]byte)
	r := make(map[int]map[int][]byte)

	for i := 0; i <= 1; i++ {
		ex[i] = make(map[int][]byte)
		ex[i][0] = e0
		iByte, err := utils.ToByte(i)
		if err != nil {
			log.Fatal(err)
		}
		r[i] = make(map[int][]byte)

		for j := 0; j <= 2; j++ {
			jplusoneByte, err := utils.ToByte(j + 1)
			if err != nil {
				log.Fatal(err)
			}

			temp := SubtractOnCurveS(ex[i][j], *P[i][j], s[i][j])

			r[i][j+1] = temp.Bytes()
			e[i][j+1] = btcutils.Sha256(M, temp.Bytes(), iByte, jplusoneByte)
			log.Println(e[i][j+1])
		}
	}

	e0prime := btcutils.Sha256(r[0][2], r[1][2], M)
	log.Println(e0prime, e0)
}
