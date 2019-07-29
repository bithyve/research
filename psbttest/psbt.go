package main

import (
	"encoding/json"
	// "encoding/hex"
	// "github.com/bithyve/research/bech32"
	// btcutils "github.com/bithyve/research/utils"
	// "github.com/bithyve/research/hdwallet"
	//"github.com/bithyve/research/bip39"
	utils "github.com/Varunram/essentials/utils"
	rpc "github.com/bithyve/research/rpc"
	//"math/big"
	"log"
)

type GetNewAddressReturn struct {
	Result string `json:"result"`
	Error  string `json:"error"`
}

type GetAddressesInfoReturn struct {
	Result struct {
		Address      string `json:"address"`
		ScriptPubKey string `json:"scriptPubKey"`
		Ismine       bool   `json:"ismine"`
		Solvable     bool   `json:"solvable"`
		Desc         string `json:"desc"`
		Iswatchonly  bool   `json:"iswatchonly"`
		Isscript     bool   `json:"isscript"`
		Iswitness    bool   `json:"iswitness"`
		Script       string `json:"script"`
		Hex          string `json:"hex"`
		Pubkey       string `json:"pubkey"`
		Embedded     struct {
			Isscript        bool   `json:"isscript"`
			Iswitness       bool   `json:"iswitness"`
			Witness_version int    `json:"witness_version"`
			Witness_program string `json:"witness_program"`
			Pubkey          string `json:"pubkey"`
			Address         string `json:"address"`
			ScriptPubKey    string `json:"scriptPubKey"`
		} `json:"embedded"`
		Label               string `json:"label"`
		Ischange            bool   `json:"ischange"`
		Timestamp           int64  `json:"timestamp"`
		Hdkeypath           string `json:"hdkeypath"`
		Hdseedid            string `json:"hdseedid"`
		Hdmasterfingerprint string `json:"hdmasterfingerprint"`
		Labels              []struct {
			Name    string `json:"name"`
			Purpose string `json:"purpose"`
		} `json:"labels"`
	} `json:"result"`
	Error string `json:"error"`
	Id    string `json:"id"`
}

type AddMultisigAddressReturn struct {
	Result struct {
		Address      string `json:"address"`
		RedeemScript string `json:"redeemScript"`
	} `json:"result"`
	Error string `json:"error"`
	Id    string `json:"id"`
}

type PsbtReturn struct {
	Result struct {
		Psbt      string  `json:"psbt"`
		Fee       float64 `json:"fee"`
		ChangePos int
	} `json:"result"`
	Error string `json:"error"`
	Id    string `json:"id"`
}

type FinalizePSBTReturn struct {
	Result struct {
		Hex      string `json:"hex"`
		Complete bool
	} `json:"result"`
	Error string `json:"error"`
	Id    string `json:"id"`
}

type SendRawTransactionReturn struct {
	Result string `json:"result"`
	Error  string `json:"error"`
	Id     string `json:"id"`
}

func main() {
	log.Println(utils.ToBigInt("100"))

	// Aalice, Abob, Acarol
	Aalicedata, err := rpc.GetNewAddress("", "")
	if err != nil {
		log.Fatal("could not generate Aalice")
	}
	Abobdata, err := rpc.GetNewAddress("", "")
	if err != nil {
		log.Fatal("could not generate Abob")
	}
	Acaroldata, err := rpc.GetNewAddress("", "")
	if err != nil {
		log.Fatal("could not generate Acarol")
	}

	var Aalice, Abob, Acarol GetNewAddressReturn

	err = json.Unmarshal(Aalicedata, &Aalice)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(Abobdata, &Abob)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(Acaroldata, &Acarol)
	if err != nil {
		log.Fatal(err)
	}

	// log.Println("AALICE: ", Aalice.Result)
	// getaddressinfo
	Kalicedata, err := rpc.GetAddressesInfo(Aalice.Result)
	if err != nil {
		log.Fatal("getaddressinfo failed for Kalice")
	}
	Kbobdata, err := rpc.GetAddressesInfo(Abob.Result)
	if err != nil {
		log.Fatal("getaddressinfo failed for Kbob")
	}
	Kcaroldata, err := rpc.GetAddressesInfo(Acarol.Result)
	if err != nil {
		log.Fatal("getaddressinfo failed for Kcarol")
	}

	var Kalice, Kbob, Kcarol GetAddressesInfoReturn

	err = json.Unmarshal(Kalicedata, &Kalice)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(Kbobdata, &Kbob)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(Kcaroldata, &Kcarol)
	if err != nil {
		log.Fatal(err)
	}

	// log.Println("KALICE: ", Kalice.Result.Address)
	// generate multisig address
	Amultidata, err := rpc.AddMultisigAddress("2", Kalice.Result.Address, Kbob.Result.Address, Kcarol.Result.Address)
	if err != nil {
		log.Fatal("could not generate new multisig address")
	}

	var Amulti AddMultisigAddressReturn

	err = json.Unmarshal(Amultidata, &Amulti)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("MULTISIG ADDRESS: ", Amulti.Result.Address)

	_, err = rpc.GenerateToAddress("101", Amulti.Result.Address)
	if err != nil {
		log.Fatal(err)
	}

	var Asend GetNewAddressReturn
	Asenddata, err := rpc.GetNewAddress("", "")
	if err != nil {
		log.Fatal("could not generate Aalice")
	}

	err = json.Unmarshal(Asenddata, &Asend)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("ASEND: ", Asend.Result)
	var inputs []interface{}
	outputs := make(map[string]int)
	outputs[Asend.Result] = 1

	var locktime int

	options := make(map[string]interface{})
	var temp []int
	temp = append(temp, 0)
	options["subtractFeeFromOutputs"] = temp
	options["includeWatching"] = true

	var bip32Derivs bool

	// construct a raw payload since we can't seem to parse the json (first param is [] and can't be parsed)
	rawpayload := `{"jsonrpc":"1.0","id":"curltext","method":"walletcreatefundedpsbt","params":[`
	rawpayload += `[],{"` + Asend.Result + `":50},0]}`

	psbtData, err := rpc.WalletCreateFundedPSBT(inputs, outputs, locktime, options, bip32Derivs, rawpayload)
	if err != nil {
		log.Fatal(err)
	}

	var psbt PsbtReturn
	err = json.Unmarshal(psbtData, &psbt)
	if err != nil {
		log.Fatal(err)
	}

	P2data, err := rpc.WalletProcessPSBT(psbt.Result.Psbt, true, "", false)
	if err != nil {
		log.Fatal(err)
	}

	var P2 PsbtReturn
	err = json.Unmarshal(P2data, &P2)
	if err != nil {
		log.Fatal(err)
	}

	// Bob validates P here
	P3data, err := rpc.WalletProcessPSBT(P2.Result.Psbt, true, "", false)
	if err != nil {
		log.Fatal(err)
	}

	var P3 PsbtReturn
	err = json.Unmarshal(P3data, &P3)
	if err != nil {
		log.Fatal(err)
	}

	Tdata, err := rpc.FinalizePSBT(P3.Result.Psbt, false)
	if err != nil {
		log.Fatal(err)
	}

	var T FinalizePSBTReturn
	err = json.Unmarshal(Tdata, &T)
	if err != nil {
		log.Fatal(err)
	}

	resultData, err := rpc.SendRawTransaction(T.Result.Hex, false)
	if err != nil {
		log.Fatal(err)
	}

	var result SendRawTransactionReturn

	err = json.Unmarshal(resultData, &result)
	if err != nil {
		log.Fatal(err)
	}

	txhash := result.Result
	log.Println("txhash: ", txhash)
}
