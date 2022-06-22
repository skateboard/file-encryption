package EncryptionWriter

import (
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

func TestWriter(t *testing.T) {
	logsFile, err := os.OpenFile("log_encrypted.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		t.Error(err)
	}

	w := NewWriter(os.Getenv("ENCRYPTION_KEY"), logsFile)

	test := io.MultiWriter(os.Stdout, w)
	log.SetOutput(test)

	log.Println("Test1!")
	log.Println("Test2!")
	log.Println("Test3!")
	log.Println("Test4!")
	log.Println("Test5!")

	for i := 0; i < 1_000_000; i++ {
		log.Println(fmt.Sprintf("Test%d!", i))
	}
}
