package core

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	key := DeriveKey("testkey1234567890")
	originalContent := []byte("This is a test content.")

	tmpDir, err := ioutil.TempDir("", "ecat_test")
	if err != nil {
		t.Fatalf("Unable to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	originalFilePath := filepath.Join(tmpDir, "test.txt")

	err = ioutil.WriteFile(originalFilePath, originalContent, 0644)
	if err != nil {
		t.Fatalf("Unable to write raw file: %v", err)
	}

	encryption := NewEncryption()

	err = encryption.Encrypt(key, originalFilePath)
	if err != nil {
		t.Errorf("Encryption failed: %v", err)
	}

	encryptedContent, err := ioutil.ReadFile(originalFilePath)
	if err != nil {
		t.Fatalf("Unable to read encrypted file: %v", err)
	}

	if bytes.Equal(encryptedContent, originalContent) {
		t.Errorf("The encrypted content should not be the same as the original content.")
	}

	err = encryption.Decrypt(key, originalFilePath)
	if err != nil {
		t.Errorf("Decryption failed: %v", err)
	}

	decryptedContent, err := ioutil.ReadFile(originalFilePath)
	if err != nil {
		t.Fatalf("Unable to read decrypted file: %v", err)
	}

	if !bytes.Equal(decryptedContent, originalContent) {
		t.Errorf("The decrypted content does not match the original content.")
	}
}

func TestShow(t *testing.T) {
	key := DeriveKey("testkey1234567890")
	originalContent := []byte("This is a test content for show method.")

	tmpDir, err := ioutil.TempDir("", "ecat_test_show")
	if err != nil {
		t.Fatalf("Unable to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	filePath := filepath.Join(tmpDir, "test_show.txt")

	err = ioutil.WriteFile(filePath, originalContent, 0644)
	if err != nil {
		t.Fatalf("Unable to write raw file: %v", err)
	}

	encryption := NewEncryption()

	err = encryption.Encrypt(key, filePath)
	if err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}

	content, err := encryption.Show(key, filePath)
	if err != nil {
		t.Errorf("Show method failed: %v", err)
	}

	if content != string(originalContent) {
		t.Errorf("The content returned by the Show method does not match the original content.")
	}
}
