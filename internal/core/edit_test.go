package core

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// mock editor function that replaces file content with mock edited content
func mockEditor(filePath string) error {
	// mock edited content
	editedContent := []byte("This is the edited content.")
	err := ioutil.WriteFile(filePath, editedContent, 0600)
	return err
}

func TestEditEncryptedFile(t *testing.T) {
	key := DeriveKey("testkey1234567890")
	originalContent := "This is the original content."

	tmpDir, err := ioutil.TempDir("", "ecat_edit_test")
	if err != nil {
		t.Fatalf("can't create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	filePath := filepath.Join(tmpDir, "test_edit.ecat")

	ioutil.WriteFile(filePath, []byte(originalContent), 0600)

	encryption := NewEncryption()

	// encrypt the initial file
	err = encryption.Encrypt(key, filePath)
	if err != nil {
		t.Fatalf("encrypt the initial file failed: %v", err)
	}

	err = EditEncryptedFile(key, filePath, mockEditor)
	if err != nil {
		t.Errorf("EditEncryptedFile fail: %v", err)
	}

	decryptedContent, err := encryption.Show(key, filePath)
	if err != nil {
		t.Fatalf("decrypted fail: %v", err)
	}

	expectedContent := "This is the edited content."
	if decryptedContent != expectedContent {
		t.Errorf("the edited file content doesn't match, got: %s, expected: %s", decryptedContent, expectedContent)
	}
}
