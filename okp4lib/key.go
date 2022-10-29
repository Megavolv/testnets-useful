package okp4lib

import "github.com/cosmos/cosmos-sdk/crypto/keyring"

type Key struct {
	keyring.KeyOutput
}

func (k Key) ToJson() (string, error) {
	out, err := KeysCdc.MarshalJSON(k.KeyOutput)
	return string(out), err
}
