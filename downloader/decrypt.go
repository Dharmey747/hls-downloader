package downloader

import (
	"HLS/common"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"strings"
)

func decrypt(file string, key []byte, iv string) bool {
	payload := common.ServeFile(file)

	if len(payload) == 0 {
		return false
	}

	unHexedIV, err := hex.DecodeString(iv)
	if err != nil {
		return false
	}

	decryptionKey, err := aes.NewCipher(key)
	if err != nil {
		return false
	}

	if len(payload)%aes.BlockSize != 0 {
		return false
	}

	if len(unHexedIV) != aes.BlockSize {
		return false
	}

	decryptionMode := cipher.NewCBCDecrypter(decryptionKey, unHexedIV)
	decryptionMode.CryptBlocks(payload, payload)

	common.CreateFile(strings.Replace(file, "temp_downloads", "temp_decrypt", 1), payload)
	common.RemoveFile(file)

	return true
}
