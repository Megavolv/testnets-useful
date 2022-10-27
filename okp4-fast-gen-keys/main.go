package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"

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
var from int
var to int
var cpus int

func init() {
	flag.StringVar(&prefix, "prefix", "prefix", "Set prefix for generated keys")
	flag.IntVar(&from, "from", 0, "Lower bound of the range")
	flag.IntVar(&to, "to", 100, "Upper limit of the range")
	flag.IntVar(&cpus, "cpus", runtime.NumCPU(), "Number of cores to use. By default - all cores")

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
	// Текущий диапазон ключей
	chunk := (to - from) / cpus

	var wg sync.WaitGroup

	for n := 0; n < cpus; n++ {
		wg.Add(1)
		go func(n int) {
			start := from + n*chunk
			end := from + (n * chunk) + chunk

			// Особое условие для последнего потока в связи с погрешностью деления
			if n == cpus-1 {
				end = to
			}

			for i := start; i < end; i++ {
				fmt.Println(CreateNextKey(prefix, i))
			}
			wg.Done()
		}(n)
	}
	wg.Wait()
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
