package core

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/google/uuid"
)

type EditorFunc func(filePath string) error

func EditEncryptedFile(userKey []byte, filePath string, editorFunc EditorFunc) error {
	// 1: read file and decrypt original content
	encryption := NewEncryption()
	plaintext, err := encryption.Show(userKey, filePath)
	if err != nil {
		return err
	}

	// 2: create temp file and write decrypted content
	tempDir := os.TempDir()
	tempFileName := uuid.New().String() + filepath.Ext(filePath)
	tempFilePath := filepath.Join(tempDir, tempFileName)

	err = ioutil.WriteFile(tempFilePath, []byte(plaintext), 0600)
	if err != nil {
		return err
	}
	defer os.Remove(tempFilePath) // ensure temp file is deleted

	// 3: open temp file with default editor
	// editor := os.Getenv("EDITOR")
	// // fmt.Println("editor:", editor)
	// // fmt.Println("tempFilePath:", tempFilePath)

	// if editor == "" {
	// 	editor = "vi"
	// }
	// cmd := exec.Command(editor, tempFilePath)
	// cmd.Stdin = os.Stdin
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	err = editorFunc(tempFilePath)
	if err != nil {
		return err
	}

	// encrypt temp file
	err = encryption.Encrypt(userKey, tempFilePath)
	if err != nil {
		return err
	}
	// read temp file content and write back to original file
	editedContent, err := ioutil.ReadFile(tempFilePath)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filePath, editedContent, 0600)

	return err
}

func DefaultEditor(filePath string) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vi"
	}
	cmd := exec.Command(editor, filePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
