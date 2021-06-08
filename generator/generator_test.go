package generator

import (
	"testing"
)

func TestNewGeneratorWithEmptyPath(t *testing.T) {
	bits := 256
	hdpath := ""
	generator, err := NewWalletGenerator(bits, hdpath)
	if err != nil {
		t.Errorf("Error with %v", err)
	}
	if generator == nil {
		t.Error("generator should not be nil")
	}
}

func TestNewGeneratorWithPath(t *testing.T) {
	bits := 256
	hdpath := "m/44'/60'/0'/0"
	generator, err := NewWalletGenerator(bits, hdpath)
	if err != nil {
		t.Errorf("Error with %v", err)
	}
	if generator == nil {
		t.Error("generator should not be nil")
	}
}

func TestNewGeneratorWithWrongConfig(t *testing.T) {
	bits := 512
	hdpath := "m/'44'/60'/0'/0"

	generator, err := NewWalletGenerator(bits, "")
	if err == nil {
		t.Errorf("error should not be nil")
	}
	if generator != nil {
		t.Error("generator should be nil")
	}

	generator, err = NewWalletGenerator(256, hdpath)
	if err == nil {
		t.Errorf("error should not be nil")
	}
	if generator != nil {
		t.Error("generator should be nil")
	}

}
