package wallets

import (
	"crypto/rand"
	"testing"
)

func TestByteToString(t *testing.T) {
	{
		expectedBytes := []byte("hello world")
		expected := string(expectedBytes)
		actual := b2s(expectedBytes)
		if expected != actual {
			t.Errorf("expected: %s, actual: %s", expected, actual)
		}
	}
	{
		expectedBytes := make([]byte, 1024*1024)
		if _, err := rand.Read(expectedBytes); err != nil {
			t.Error(err)
			t.FailNow()
		}

		expected := string(expectedBytes)
		actual := b2s(expectedBytes)
		if expected != actual {
			t.Errorf("expected: %s, actual: %s", expected, actual)
		}
	}
	{
		expectedBytes := make([]byte, 1024*1024)
		expected := string(expectedBytes)
		actual := b2s(expectedBytes)
		if expected != actual {
			t.Errorf("expected: %s, actual: %s", expected, actual)
		}
	}

}
