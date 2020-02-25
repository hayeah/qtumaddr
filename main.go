package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/btcsuite/btcutil/base58"
)

const (
	qtumMainPubKeyHashAddrID = 58
	qtumMainScriptHashAddrID = 50

	qtumTestNetPubKeyHashAddrID = 120
	qtumTestNetScriptHashAddrID = 110
)

var (
	flagHexAddress    = kingpin.Flag("hex", "address to convert in hex").Short('h').String()
	flagBase58Address = kingpin.Flag("base58", "address to convert in hex").Short('b').String()
	flagIsTestnet     = kingpin.Flag("test", "generate base58 address for testnet").Short('t').Bool()
)

func base58ToHex(b58addr string) (string, error) {
	addr, _, err := base58.CheckDecode(b58addr)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(addr), nil
}

func hexToBase58(hexaddr string, netid byte) (string, error) {
	addr, err := hex.DecodeString(hexaddr)
	if err != nil {
		return "", err
	}

	b58addr := base58.CheckEncode(addr, netid)

	return b58addr, nil
}

func run() error {
	kingpin.Parse()

	var err error
	var out string

	var version byte = qtumMainPubKeyHashAddrID
	if *flagIsTestnet {
		version = qtumTestNetPubKeyHashAddrID
	}

	hexAddr := *flagHexAddress
	base58Addr := *flagBase58Address

	if hexAddr != "" {
		out, err = hexToBase58(hexAddr, version)
		if err != nil {
			return err
		}
	} else if base58Addr != "" {
		out, err = base58ToHex(base58Addr)
		if err != nil {
			return err
		}
	} else {
		log.Fatalln("Must specify an address to convert using either --hex or --base58")
	}

	fmt.Println(out)

	return nil
}

func main() {
	err := run()
	if err != nil {
		log.Fatalln(err)
	}
}
