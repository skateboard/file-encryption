package EncryptionWriter

import (
	"os"
	"testing"
)

func TestReading(t *testing.T) {
	logsFile, err := os.OpenFile("log_encrypted.txt", os.O_RDONLY, 0644)
	if err != nil {
		t.Error(err)
	}

	r := NewReader(os.Getenv("ENCRYPTION_KEY"), logsFile)

	contents, err := r.ReadFile()
	if err != nil {
		return
	}

	logsFile, err = os.OpenFile("log_decrypted.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		t.Error(err)
	}

	for _, content := range contents {
		_, err := logsFile.WriteString(content)
		if err != nil {
			t.Error(err)
			return
		}
	}
}
