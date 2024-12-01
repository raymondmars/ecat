package core

import (
	"bytes"
	"crypto/sha256"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestDeriveKey(t *testing.T) {
	password := "testpassword"
	expectedHash := sha256.Sum256([]byte(password))

	derivedKey := DeriveKey(password)

	if !bytes.Equal(derivedKey, expectedHash[:]) {
		t.Errorf("DeriveKey() = %x, want %x", derivedKey, expectedHash[:])
	}
}

func TestInitKey(t *testing.T) {
	tempHomeDir, err := ioutil.TempDir("", "testhomedir")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempHomeDir)

	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)
	os.Setenv("HOME", tempHomeDir)

	inputKey := "testinputkey"

	InitKey(inputKey)

	keyFilePath := filepath.Join(tempHomeDir, ".ecat", KEY_FILE_NAME)
	keyData, err := ioutil.ReadFile(keyFilePath)
	if err != nil {
		t.Fatalf("Failed to read key file: %v", err)
	}

	storedKey := DeriveKey(inputKey)
	if string(keyData) != string(storedKey) {
		t.Errorf("Key file content = %s, want %s", string(keyData), storedKey)
	}

	fileInfo, err := os.Stat(keyFilePath)
	if err != nil {
		t.Fatalf("Failed to stat key file: %v", err)
	}

	expectedPerms := os.FileMode(0600)
	if fileInfo.Mode().Perm() != expectedPerms {
		t.Errorf("Key file permissions = %v, want %v", fileInfo.Mode().Perm(), expectedPerms)
	}
}

func TestGetStoredKey(t *testing.T) {
	tempHomeDir, err := ioutil.TempDir("", "testhomedir")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempHomeDir)

	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)
	os.Setenv("HOME", tempHomeDir)

	keyContent := []byte("testkeycontent")
	keyDir := filepath.Join(tempHomeDir, ".ecat")
	err = os.MkdirAll(keyDir, 0700)
	if err != nil {
		t.Fatalf("Failed to create key directory: %v", err)
	}

	keyFilePath := filepath.Join(keyDir, KEY_FILE_NAME)
	err = ioutil.WriteFile(keyFilePath, keyContent, 0600)
	if err != nil {
		t.Fatalf("Failed to write key file: %v", err)
	}

	storedKey, err := GetStoredKey()
	if err != nil {
		t.Fatalf("GetStoredKey() returned an error: %v", err)
	}

	if !bytes.Equal(storedKey, keyContent) {
		t.Errorf("GetStoredKey() = %s, want %s", storedKey, keyContent)
	}
}

func TestGetStoredKeyNoKeyFile(t *testing.T) {
	tempHomeDir, err := ioutil.TempDir("", "testhomedir")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempHomeDir)

	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)
	os.Setenv("HOME", tempHomeDir)

	storedKey, err := GetStoredKey()
	if err == nil {
		t.Fatalf("Expected error when key file does not exist, but got nil")
	}

	if storedKey != nil {
		t.Errorf("Expected nil key when key file does not exist, but got %v", storedKey)
	}
}
