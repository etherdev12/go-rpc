package crypto

import (
	"errors"
)

var KEY_LEN = DefaultKeySize // 16, 24, 32

func CFBDecryptBuffer(buffer []byte) ([]byte, error) {
	if len(buffer) < KEY_LEN {
		return nil, errors.New("length of buffer is small than key size")
	}
	key := buffer[:KEY_LEN]
	encrypted := buffer[KEY_LEN:]

	return Decrypt(key, encrypted)
}
