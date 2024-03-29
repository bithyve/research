package main

import (
	"encoding/hex"
	"log"

	// "github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

func addWitness(tx *wire.TxIn, witnesses ...string) error {
	if len(witnesses) == 0 {
		var witness [][]byte
		tx.Witness = witness
		return nil
	}
	witness := make([][]byte, len(witnesses))
	for i, val := range witnesses {
		witnessBytes, err := hex.DecodeString(val)
		if err != nil {
			return err
		}
		witness[i] = witnessBytes
	}
	tx.Witness = witness
	return nil
}

func addScriptSig(input *wire.TxIn, ss string) error {
	scriptPubKey, err := hex.DecodeString(ss)
	if err != nil {
		return err
	}

	input.SignatureScript = scriptPubKey
	return nil
}

func addPrevOut(input *wire.TxIn, prevoutString string, index uint32) error {
	if prevoutString == "" {
		op := chainhash.Hash{}
		prevOut := wire.OutPoint{
			Hash:  op,
			Index: 0xffffffff,
		}

		input.PreviousOutPoint = prevOut
		return nil
	}

	op, err := chainhash.NewHashFromStr(prevoutString)
	if err != nil {
		log.Fatal(err)
	}

	prevOut := wire.NewOutPoint(op, index)
	input.PreviousOutPoint = *prevOut

	return nil
}

func addSequence(input *wire.TxIn, sequence uint32) {
	input.Sequence = sequence
}

func addScriptPubkey(tx *wire.MsgTx, pubkey string, value int64) error {
	scriptPubKey, err := hex.DecodeString(pubkey)
	if err != nil {
		return err
	}

	txOut := wire.NewTxOut(value, scriptPubKey)
	tx.AddTxOut(txOut)
	return nil
}

func setLockTime(tx *wire.MsgTx, locktime uint32) {
	tx.LockTime = locktime
}

func test1() {
	// https://blockstream.info/tx/f051e59b5e2503ac626d03aaeac8ab7be2d72ba4b7e97119c5852d70d52dcb86
	targetHash := "f051e59b5e2503ac626d03aaeac8ab7be2d72ba4b7e97119c5852d70d52dcb86"
	targetSize := 134

	tx := wire.NewMsgTx(1)

	txin := new(wire.TxIn)

	err := addPrevOut(txin, "", 0xffffffff)
	if err != nil {
		log.Fatal(err)
	}

	err = addScriptSig(txin, "0431dc001b0162")
	if err != nil {
		log.Fatal(err)
	}

	err = addWitness(txin)
	if err != nil {
		log.Fatal(err)
	}

	addSequence(txin, 0xffffffff)

	tx.AddTxIn(txin)

	err = addScriptPubkey(tx, "4104d64bdfd09eb1c5fe295abdeb1dca4281be988e2da0b6c1c6a59dc226c28624e18175e851c96b973d81b01cc31f047834bc06d6d6edf620d184241a6aed8b63a6ac", 5000000000)
	if err != nil {
		log.Fatal(err)
	}

	if tx.TxHash().String() != targetHash || tx.SerializeSize() != targetSize {
		log.Fatal("test 1 target hash doesn't match")
	}
}

func test2() {
	targetHash := "baa591f2f4505f189b84f7e518eb390c5049b4fd0441a4fb798cc21b5a7f91f1"
	targetSize := 249

	tx := wire.NewMsgTx(2)

	txin := new(wire.TxIn)
	err := addScriptSig(txin, "1600143e33e3c857f5e4374eb16c65149eb9eeeb5bd7ce")
	if err != nil {
		log.Fatal(err)
	}

	err = addPrevOut(txin, "525fc8fea737c827d600b09fcbb002e518d5b9498230c0f596ea1ba60358cc20", 0)
	if err != nil {
		log.Fatal(err)
	}

	err = addWitness(txin, "3044022009ee2956dfe779c8120db092c86d8fc8d95741a193231151a203ed246444fbcd02203bec6d837a7d11e69375a4130bd56251fccbbdd262ed6ba4bb693310209bc33401",
		"039ff0c4745bf9b5b3b85ddacff408b7d6288f832a53bc016267d5f5b6a52f92c4")
	if err != nil {
		log.Fatal(err)
	}

	txin.Sequence = 0xffffffff
	tx.AddTxIn(txin)

	err = addScriptPubkey(tx, "a91402edb870b533709fc15643eb24e94b7d31bea22787", 250250)
	if err != nil {
		log.Fatal(err)
	}

	err = addScriptPubkey(tx, "76a914d268b87b79af4c5ab430464df058ece832b98a4988ac", 2750)
	if err != nil {
		log.Fatal(err)
	}

	if tx.TxHash().String() != targetHash || tx.SerializeSize() != targetSize {
		log.Fatal("test 2 target hash doesn't match")
	}
}

func test3() {
	// https://blockstream.info/tx/f051e59b5e2503ac626d03aaeac8ab7be2d72ba4b7e97119c5852d70d52dcb86
	tx := wire.NewMsgTx(1)
	targetHash := "a753cb59cdfc769d067fc0d7853ce5317be99e06a56acf2bc3419cd042eb549a"
	targetSize := 300

	txin := new(wire.TxIn)

	err := addScriptSig(txin, "0396050941d757b09af6bbe141d757b09ac809be2f4254432e544f502ffabe6d6d1f8f5ea2928a9ee1b5db50404a89bfba26339b514550af5f2626456a47f2c5798000000000000000db0084143cc4000000000000")
	if err != nil {
		log.Fatal(err)
	}

	err = addPrevOut(txin, "", 0xffffffff)
	if err != nil {
		log.Fatal(err)
	}

	err = addWitness(txin, "0000000000000000000000000000000000000000000000000000000000000000")
	if err != nil {
		log.Fatal(err)
	}

	addSequence(txin, 0xffffffff)

	tx.AddTxIn(txin)

	err = addScriptPubkey(tx, "76a914ba507bae8f1643d2556000ca26b9301b9069dc6b88ac", 1289673184)
	if err != nil {
		log.Fatal(err)
	}

	err = addScriptPubkey(tx, "6a24aa21a9edbab6d0ff211c4cb047fbf9d9ddb6b6ee1bef8e8483c4650254f2008017648785", 0)
	if err != nil {
		log.Fatal(err)
	}

	err = addScriptPubkey(tx, "6a24b9e11b6d6c8a82f163071ef21c19fd7dbd584f6f4bd2e7aa5836cde657c9c80f43797fc1", 0)
	if err != nil {
		log.Fatal(err)
	}

	if tx.TxHash().String() != targetHash || tx.SerializeSize() != targetSize {
		log.Fatal("test 3 target hash doesn't match")
	}
}

func test4() {
	// https://blockstream.info/tx/cd141d5cbbd081f2d5807d835d49d49f15a9728c2fa674e7ec79182c1315d207?expand
	tx := wire.NewMsgTx(1)
	targetHash := "cd141d5cbbd081f2d5807d835d49d49f15a9728c2fa674e7ec79182c1315d207"
	targetSize := 370

	txin1 := new(wire.TxIn)
	txin2 := new(wire.TxIn)

	err := addScriptSig(txin1, "483045022100c477f46a7bc670e11ee5466d312c9fd587e01a0328752fadd82107a40c28fe420220728acd291534aa8a2e2e2970ded7c39dfdd83f6847a1c208701e560461654246012103b23d3a7097833e05193307b104c7bbceb482a5e69d9c754120fc3f82c7d1475c")
	if err != nil {
		log.Fatal(err)
	}

	err = addPrevOut(txin1, "fa65a71767d5be55620b3d55214fe67eaf108860427f2b368a2b17ba301e77cc", 1)
	if err != nil {
		log.Fatal(err)
	}

	err = addWitness(txin1)
	if err != nil {
		log.Fatal(err)
	}

	addSequence(txin1, 0xffffffff)

	tx.AddTxIn(txin1)

	err = addScriptSig(txin2, "47304402203459b842186a374015e409ffe4d2d67101c7cc6bc68f62aeed6f5499c2086c68022034471874999d4e2452c1493bf002562a7b279bb58af6ee414a50fe04892fcce2012103b16ac8c3821b24d5071d26b2a123031515c65e026bbf53b3d7a2ddd5e4270950")
	if err != nil {
		log.Fatal(err)
	}

	err = addPrevOut(txin2, "b67e3daf478b28cf68828f75632fb5e281c9475d3204242c2c6e46e031dd2191", 0)
	if err != nil {
		log.Fatal(err)
	}

	err = addWitness(txin2)
	if err != nil {
		log.Fatal(err)
	}

	addSequence(txin2, 0xfffffffe)

	tx.AddTxIn(txin2)

	err = addScriptPubkey(tx, "76a914081fbdf61f3e1fc1d7cb252109870a56f64fa80c88ac", 25777652)
	if err != nil {
		log.Fatal(err)
	}

	err = addScriptPubkey(tx, "6a146f6d6e69000000000000001f0000002115d1ed00", 0)
	if err != nil {
		log.Fatal(err)
	}

	if tx.TxHash().String() != targetHash || tx.SerializeSize() != targetSize {
		log.Fatal("test 4 target hash doesn't match")
	}
}

func main() {
	test1()
	test2()
	test3()
	test4()
}
