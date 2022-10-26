package main

import (
	"fmt"
	"strconv"

	"okp4-fast-gen-keys/pkg/cosmos"

	"github.com/cosmos/cosmos-sdk/codec"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/go-bip39"
	flag "github.com/spf13/pflag"
)

var KeysCdc *codec.LegacyAmino
var prefix string
var num int

func init() {
	flag.StringVar(&prefix, "name", "prefix", "Set prefix for generated wallets")
	flag.IntVar(&num, "num", 100, "Number of wallets generated")
	flag.Parse()

	KeysCdc = codec.NewLegacyAmino()
	cryptocodec.RegisterCrypto(KeysCdc)
	KeysCdc.Seal()

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("okp4", "okp4pub")
	config.SetBech32PrefixForValidator("okp4valoper", "okp4valoperpub")
	config.SetBech32PrefixForConsensusNode("okp4valcons", "okp4valconspub")
	config.SetCoinType(118)
	config.Seal()
}

func main() {
	for i := 0; i < num; i++ {
		fmt.Println(CreateNextKey(prefix, i))
	}
}

func CreateNextKey(name string, n int) string {
	entropySeed, err := bip39.NewEntropy(256)
	if err != nil {
		panic(err)
	}

	mnemonic, err := bip39.NewMnemonic(entropySeed)
	if err != nil {
		panic(err)
	}

	pk, err := cosmos.ParseMnemonic(mnemonic)
	if err != nil {
		panic(err)
	}

	pubKey := pk.PubKey()

	addr := sdk.AccAddress(pubKey.Address())

	out := keyring.KeyOutput{
		Name:     name + strconv.Itoa(n),
		Type:     "local",
		Address:  addr.String(),
		PubKey:   pubKey.String(),
		Mnemonic: mnemonic,
	}

	jsonString, err := KeysCdc.MarshalJSON(out)
	if err != nil {
		panic(err)
	}

	return string(jsonString)
}
