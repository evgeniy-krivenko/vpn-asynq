package service

import "testing"

func TestCryptoService_Encrypt(t *testing.T) {
	c := new(CryptoService)
	t.Run("work with 16 symbols", func(t *testing.T) {

		var text, key = "simple_pass", "someKeyjaksfamks"

		_, err := c.Encrypt(text, key)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("work with 24 symbols", func(t *testing.T) {
		var text, key = "simple_pass", "someKeyjaksfamkskaldmscb"

		_, err := c.Encrypt(text, key)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("encypted text correct decrypt", func(t *testing.T) {
		var text, key = "simple_pass", "someKeyjaksfamkskaldmscb"

		e, err := c.Encrypt(text, key)
		if err != nil {
			t.Fatal(err)
		}

		d, _ := c.Decrypt(e, key)

		if d != text {
			t.Errorf("result: %v not equal: %v", d, text)
		}
	})
}

func TestCryptoService_GeneratePassword(t *testing.T) {
	c := new(CryptoService)
	t.Run("correct_len", func(t *testing.T) {
		pLen := 12
		p := c.GeneratePassword(pLen)

		if len(p) != pLen {
			t.Errorf("error pass len: %d not equal %d", len(p), pLen)
		}
	})
}
