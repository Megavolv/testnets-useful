package okp4lib

import (
	"github.com/Megavolv/testnets-useful/okp4lib/pkg/cosmos"
	"github.com/cosmos/cosmos-sdk/codec"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/go-bip39"
)

var KeysCdc *codec.LegacyAmino

func init() {
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

type Okp4 struct {
}

func NewOkp4() *Okp4 {
	return &Okp4{}
}

func (o *Okp4) GetEntropySeed() ([]byte, error) {
	return bip39.NewEntropy(256)
}

func (o *Okp4) NewMnemonic(entropy []byte) (string, error) {
	return bip39.NewMnemonic(entropy)
}

func (o *Okp4) CreateJsonedKey(name string) (string, error) {
	key, err := o.CreateKey(name)
	if err != nil {
		return "", err
	}

	return key.ToJson()
}

func (o *Okp4) CreateKey(name string) (Key, error) {
	entropy, err := o.GetEntropySeed()
	if err != nil {
		return Key{}, err
	}

	mnemonic, err := o.NewMnemonic(entropy)
	if err != nil {
		return Key{}, err
	}

	return o.createKey(name, mnemonic)
}

func (o *Okp4) createKey(name, mnemonic string) (Key, error) {
	pk, err := cosmos.ParseMnemonic(mnemonic)
	if err != nil {
		return Key{}, err
	}

	pubKey := pk.PubKey()

	addr := sdk.AccAddress(pubKey.Address())

	return Key{
		keyring.KeyOutput{
			Name:     name,
			Type:     "local",
			Address:  addr.String(),
			PubKey:   pubKey.String(),
			Mnemonic: mnemonic,
		},
	}, nil
}
