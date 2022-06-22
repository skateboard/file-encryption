package EncryptionWriter

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"os"
)

type EncryptionFileReader struct {
	file *os.File

	gcm cipher.AEAD
}

func NewReader(key string, file *os.File) EncryptionFileReader {
	encryptionKey, _ := hex.DecodeString(key)

	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	return EncryptionFileReader{
		file: file,
		gcm:  aesGCM,
	}
}

func (e EncryptionFileReader) ReadFile() ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(e.file)

	for scanner.Scan() {
		lines = append(lines, e.decrypt(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func (e EncryptionFileReader) decrypt(encryptedString string) (decryptedString string) {
	enc, _ := hex.DecodeString(encryptedString)

	//Get the nonce size
	nonceSize := e.gcm.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plaintext, err := e.gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", plaintext)
}
