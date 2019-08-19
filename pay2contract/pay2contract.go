package main

import (
	"encoding/hex"
	"log"
	"strconv"
	"strings"

	hdwallet "github.com/bithyve/research/hdwallet"
	btcutils "github.com/bithyve/research/utils"
)

func main() {

	seed, err := hdwallet.GenSeed(256)
	if err != nil {
		log.Fatal(err)
	}

	masterprv := hdwallet.MasterKey(seed)
	log.Println("Master priv key: ", masterprv)

	hKey := uint32(0x800003e7) // 3e7 is 999 in hex
	chainPrivm999H, err := masterprv.Child(hKey)
	if err != nil {
		log.Fatal(err)
	}

	hKey = uint32(0x80000000) // m/999'/0
	chainPrivm9990H, err := chainPrivm999H.Child(hKey)
	if err != nil {
		log.Fatal(err)
	}

	doc1 := []byte("blah")
	doc2 := []byte("blah2")

	sha1 := btcutils.Sha256(doc1)
	sha2 := btcutils.Sha256(doc2)

	sha1Hex := hex.EncodeToString(sha1)
	sha2Hex := hex.EncodeToString(sha2)

	var pos []byte

	if strings.Compare(sha1Hex, sha2Hex) == -1 {
		pos = append(pos, append(sha1, sha2...)...)
	} else {
		pos = append(pos, append(sha2, sha1...)...)
	}

	shaHash := btcutils.Sha256(pos)
	combinedHashString := hex.EncodeToString(shaHash)

	var slices string
	remaining := combinedHashString

	for len(remaining) > 4 {
		slices = slices + remaining[0:4] + "/"
		remaining = remaining[4:len(remaining)]
	}

	slices += remaining
	levels := strings.Split(slices, "/")

	finalMasterPrv := hdwallet.MasterKey([]byte(""))

	for i, level := range levels {
		// convert hex to string
		level = "0x" + level
		intConv, _ := strconv.ParseUint(level, 0, 64)

		if i == 0 {
			finalMasterPrv, err = chainPrivm9990H.Child(uint32(intConv))
			if err != nil {
				log.Fatal(err)
			}
		} else {
			finalMasterPrv, err = finalMasterPrv.Child(uint32(intConv))
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	pub := finalMasterPrv.Pub()
	address := pub.Address()
	log.Println("address: ", address)
}
