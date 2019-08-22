package main

import (
	"encoding/hex"
	"log"

	// "github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

func main1() {
	// https://blockstream.info/tx/baa591f2f4505f189b84f7e518eb390c5049b4fd0441a4fb798cc21b5a7f91f1?expand
	tx := wire.NewMsgTx(2)
	tx.LockTime = 0
	// TxIn     []*TxIn
	// TxOut    []*TxOut
	// LockTime uint32

	op, err := chainhash.NewHashFromStr("525fc8fea737c827d600b09fcbb002e518d5b9498230c0f596ea1ba60358cc20")
	if err != nil {
		log.Fatal(err)
	}

	prevOut := wire.NewOutPoint(op, 0)

	scriptSig, err := hex.DecodeString("1600143e33e3c857f5e4374eb16c65149eb9eeeb5bd7ce")
	if err != nil {
		log.Fatal(err)
	}

	witness1, err := hex.DecodeString("3044022009ee2956dfe779c8120db092c86d8fc8d95741a193231151a203ed246444fbcd02203bec6d837a7d11e69375a4130bd56251fccbbdd262ed6ba4bb693310209bc33401")
	if err != nil {
		log.Fatal(err)
	}

	witness2, err := hex.DecodeString("039ff0c4745bf9b5b3b85ddacff408b7d6288f832a53bc016267d5f5b6a52f92c4")
	if err != nil {
		log.Fatal(err)
	}

	witness := [][]byte{witness1, witness2}

	txin := wire.NewTxIn(prevOut, scriptSig, witness)
	txin.Sequence = 0xffffffff
	tx.AddTxIn(txin)

	scriptPubkey1, err := hex.DecodeString("a91402edb870b533709fc15643eb24e94b7d31bea22787")
	if err != nil {
		log.Fatal(err)
	}

	txout1 := wire.NewTxOut(int64(25025), scriptPubkey1)

	scriptPubkey2, err := hex.DecodeString("76a914d268b87b79af4c5ab430464df058ece832b98a4988ac")
	if err != nil {
		log.Println("scriptsig error")
		log.Fatal(err)
	}

	txout2 := wire.NewTxOut(int64(275), scriptPubkey2)

	tx.AddTxOut(txout1)
	tx.AddTxOut(txout2)

	log.Println(tx.WitnessHash())
	log.Println(tx.TxHash())
	log.Println(tx.SerializeSize())
}

func main() {
	// https://blockstream.info/tx/a753cb59cdfc769d067fc0d7853ce5317be99e06a56acf2bc3419cd042eb549a
	tx := wire.NewMsgTx(2)
	tx.LockTime = 0
	// TxIn     []*TxIn
	// TxOut    []*TxOut
	// LockTime uint32

	op := chainhash.Hash{}
	prevOut := wire.NewOutPoint(&op, 0xffffffff)

	scriptSig, err := hex.DecodeString("0396050941d757b09af6bbe141d757b09ac809be2f4254432e544f502ffabe6d6d1f8f5ea2928a9ee1b5db50404a89bfba26339b514550af5f2626456a47f2c5798000000000000000db0084143cc4000000000000")
	if err != nil {
		log.Fatal(err)
	}

	witness1, err := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000000")
	if err != nil {
		log.Fatal(err)
	}

	witness := [][]byte{witness1}

	txin := wire.NewTxIn(prevOut, scriptSig, witness)
	txin.Sequence = 0xffffffff
	tx.AddTxIn(txin)

	scriptPubkey1, err := hex.DecodeString("76a914ba507bae8f1643d2556000ca26b9301b9069dc6b88ac")
	if err != nil {
		log.Fatal(err)
	}

	txout1 := wire.NewTxOut(int64(1289673184), scriptPubkey1)

	scriptPubkey2, err := hex.DecodeString("6a24aa21a9edbab6d0ff211c4cb047fbf9d9ddb6b6ee1bef8e8483c4650254f2008017648785")
	if err != nil {
		log.Println("scriptsig error")
		log.Fatal(err)
	}

	txout2 := wire.NewTxOut(int64(0), scriptPubkey2)

	scriptPubkey3, err := hex.DecodeString("6a24b9e11b6d6c8a82f163071ef21c19fd7dbd584f6f4bd2e7aa5836cde657c9c80f43797fc1")
	if err != nil {
		log.Println("scriptsig error")
		log.Fatal(err)
	}

	txout3 := wire.NewTxOut(int64(0), scriptPubkey3)

	tx.AddTxOut(txout1)
	tx.AddTxOut(txout2)
	tx.AddTxOut(txout3)

	log.Println(tx.WitnessHash())
	log.Println(tx.TxHash())
	log.Println(tx.SerializeSize())
}

func Test11TxHash() {
	// https://blockstream.info/tx/f051e59b5e2503ac626d03aaeac8ab7be2d72ba4b7e97119c5852d70d52dcb86
	tx := wire.NewMsgTx(1)
	targetHash := "f051e59b5e2503ac626d03aaeac8ab7be2d72ba4b7e97119c5852d70d52dcb86"
	targetSize := 134

	scriptSig, err := hex.DecodeString("0431dc001b0162")
	if err != nil {
		log.Fatal(err)
	}

	op := chainhash.Hash{}
	prevOut := wire.NewOutPoint(&op, 0xffffffff)

	var x [][]byte
	txin := wire.NewTxIn(prevOut, scriptSig, x)
	txin.Sequence = 0xffffffff
	tx.AddTxIn(txin)

	scriptPubkey, err := hex.DecodeString("4104d64bdfd09eb1c5fe295abdeb1dca4281be988e2da0b6c1c6a59dc226c28624e18175e851c96b973d81b01cc31f047834bc06d6d6edf620d184241a6aed8b63a6ac")
	if err != nil {
		log.Println("scriptsig error")
		log.Fatal(err)
	}

	txout2 := wire.NewTxOut(int64(5000000000), scriptPubkey)
	tx.AddTxOut(txout2)

	if tx.TxHash().String() != targetHash && tx.SerializeSize() != targetSize {
		log.Fatal("target hash doesn't match")
	}
}
