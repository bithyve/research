package main

import (
	"log"
	"encoding/hex"

	// "github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
)

func test() {
	// https://blockstream.info/testnet/tx/ffcdbe40c20a35d9243121e2cc5920d033021218a311dc492ce6dffe7484750e?expand
	tx := wire.NewMsgTx(2) // core v0.18
	tx.LockTime = 0
	// TxIn     []*TxIn
	// TxOut    []*TxOut
	// LockTime uint32

	op, err := chainhash.NewHashFromStr("0d591d27ba813c0bccd6e3963694cac32299f783c7b756420d0261c4297e19f5")
	if err != nil {
		log.Fatal(err)
	}

	prevOut := wire.NewOutPoint(op, 1)
	log.Println(prevOut.String())

	// OP_PUSHBYTES_22 0014fb2b0b81452bc77600667856cb57b76d76d7c409
	scriptSig, err := hex.DecodeString("160014fb2b0b81452bc77600667856cb57b76d76d7c409")
	if err != nil {
		log.Fatal(err)
	}

	witness1, err := hex.DecodeString("304402200126318f5fdd61b041c61e5b311d52cf27c063aa3de5059f18ed1f937a97383c02202a6dd7f5e7aa2d0df34e1d8b75c04cd2c497654c70d3a973607e42c9bcdcfa1701")
	if err != nil {
		log.Fatal(err)
	}

	witness2, err := hex.DecodeString("028414165c66a08425b57e63cb98e898c15d91d1d089cf848c17f208f24d89d2df")
	if err != nil {
		log.Fatal(err)
	}

	witness := make([][]byte, 2)
	witness[0] = witness1
	witness[1] = witness2

	txin := wire.NewTxIn(prevOut, scriptSig, witness)
	txin.Sequence = 0xffffffff

	tx.AddTxIn(txin)

	scriptPubkey1, err := hex.DecodeString("a914eb03de286e847950f63e59b374560e69372654cd87")
	if err != nil {
		log.Fatal(err)
	}

	txout1 := wire.NewTxOut(int64(2611841), scriptPubkey1)

	scriptPubkey2, err := hex.DecodeString("6a23535701365bc3bdaa6c458778768d9230619e26e657ba4a5202abb2b6470fea16feb16c")
	if err != nil {
		log.Println("scriptsig error")
		log.Fatal(err)
	}

	txout2 := wire.NewTxOut(int64(0), scriptPubkey2)

	tx.AddTxOut(txout2)
	tx.AddTxOut(txout1)

	log.Println(tx.WitnessHash())
	log.Println(tx.TxHash())
	log.Println(tx.SerializeSize())
}

func main() {
	// https://blockstream.info/testnet/tx/ffcdbe40c20a35d9243121e2cc5920d033021218a311dc492ce6dffe7484750e?expand
	tx := wire.NewMsgTx(1) // core v0.18
	// TxIn     []*TxIn
	// TxOut    []*TxOut
	// LockTime uint32

	// OP_PUSHBYTES_22 0014fb2b0b81452bc77600667856cb57b76d76d7c409
	scriptSig, err := hex.DecodeString("0431dc001b0162")
	if err != nil {
		log.Fatal(err)
	}

	txin := &wire.TxIn {
		PreviousOutPoint: wire.OutPoint{
			Hash:  chainhash.Hash{},
			Index: 0xffffffff,
		},
		SignatureScript: scriptSig,
		Sequence: 0xffffffff,
	}

	tx.AddTxIn(txin)
	scriptPubkey2, err := hex.DecodeString("4104d64bdfd09eb1c5fe295abdeb1dca4281be988e2da0b6c1c6a59dc226c28624e18175e851c96b973d81b01cc31f047834bc06d6d6edf620d184241a6aed8b63a6ac")
	if err != nil {
		log.Println("scriptsig error")
		log.Fatal(err)
	}

	txout2 := wire.NewTxOut(int64(5000000000), scriptPubkey2)

	tx.AddTxOut(txout2)

	log.Println(tx.WitnessHash())
	log.Println(tx.TxHash())
	log.Println(tx.SerializeSize())
}
