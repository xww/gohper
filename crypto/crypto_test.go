package crypto

import (
	"testing"
)

func TestEnCrypt(t *testing.T) {
	password := "abcdefg"
	fixsalt := "123456"
	enc, salt, _ := ShaEncrypt(password, fixsalt)
	newEnc, _ := ShaEncryptWithSalt(password, fixsalt, salt)
	if string(newEnc) != string(enc) {
		t.Fail()
	}
}
