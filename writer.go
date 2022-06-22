package EncryptionWriter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

type EncryptionFileWriter struct {
	file *os.File

	gcm cipher.AEAD
}

func NewWriter(key string, file *os.File) EncryptionFileWriter {
	encryptionKey, _ := hex.DecodeString(key)

	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	return EncryptionFileWriter{
		file: file,
		gcm:  aesGCM,
	}
}

func (e EncryptionFileWriter) Write(p []byte) (n int, err error) {
	e.file.WriteString(e.encrypt(string(p)))

	return 0, nil
}

func (e EncryptionFileWriter) encrypt(stringToEncrypt string) (encryptedString string) {
	plaintext := []byte(stringToEncrypt)

	nonce := make([]byte, e.gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	ciphertext := e.gcm.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x\n", ciphertext)
}
