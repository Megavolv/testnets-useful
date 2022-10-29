package okp4lib

import (
	"testing"
	//	"github.com/davecgh/go-spew/spew"
)

var test_entropy []byte = []byte{
	0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1,
	0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1,
}

var test_mnemonic string = "absurd amount doctor acoustic avoid letter advice cage absurd amount doctor acoustic avoid letter advice cage absurd amount doctor acoustic avoid letter advice comic"

func TestEntropy(t *testing.T) {
	o := NewOkp4()

	seed, err := o.GetEntropySeed()
	if err != nil {
		t.Error(err)
	}
	if l := len(seed); l != 32 {
		t.Error("Expect 32 byte, got ", l)
	}
}

func TestMnemonic(t *testing.T) {
	o := NewOkp4()

	mnemonic, err := o.NewMnemonic(test_entropy)
	if err != nil {
		t.Error(err)
	}

	if mnemonic != test_mnemonic {
		t.Error("Expected diff menmonic")
	}
}

func TestNewKey(t *testing.T) {
	o := NewOkp4()
	key, err := o.createKey("x", test_mnemonic)
	if err != nil {
		t.Error(err)
	}

	if key.Name != "x" {
		t.Error("Expected name x, got ", key.Name)
	}

	if key.Address != "okp41v4d0dhfljhmwnpgtuwene4txqceat2w267drx8" {
		t.Error("Expected address okp41v4d0dhfljhmwnpgtuwene4txqceat2w267drx8, got ", key.Address)
	}

	if key.Mnemonic != test_mnemonic {
		t.Error("Expected ", test_mnemonic, " got ", key.Mnemonic)
	}

}
