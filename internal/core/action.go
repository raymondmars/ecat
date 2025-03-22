package core

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

type Action interface {
	InitAction(cmd *cobra.Command, args []string)
	RootAction(cmd *cobra.Command, args []string)
	EncryptAction(cmd *cobra.Command, args []string)
	DecryptAction(cmd *cobra.Command, args []string)
	EditAction(cmd *cobra.Command, args []string)
	ShowAction(cmd *cobra.Command, args []string)
}

var UseEncrypt bool
var UseDecrypt bool

type action struct {
	encryption Encryption
}

func NewAction() Action {
	return &action{encryption: NewEncryption()}
}

func (a *action) RootAction(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
		return
	}

	if UseEncrypt {
		a.EncryptAction(cmd, args)
		return
	}

	if UseDecrypt {
		a.DecryptAction(cmd, args)
		return
	}

	// default execute show action
	a.ShowAction(cmd, args)
}

func (a *action) InitAction(cmd *cobra.Command, args []string) {
	InitKey("")
}

func (a *action) EncryptAction(cmd *cobra.Command, args []string) {
	userKey := a.getKey(cmd)
	filePath := args[0]
	err := a.encryption.Encrypt(userKey, filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Encryption successful!")
}

func (a *action) DecryptAction(cmd *cobra.Command, args []string) {
	userKey := a.getKey(cmd)
	filePath := args[0]
	err := a.encryption.Decrypt(userKey, filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Decryption successful!")
}

func (a *action) EditAction(cmd *cobra.Command, args []string) {
	userKey := a.getKey(cmd)
	filePath := args[0]
	err := EditEncryptedFile([]byte(userKey), filePath, DefaultEditor)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File edited successfully!")
}

func (a *action) ShowAction(cmd *cobra.Command, args []string) {
	userKey := a.getKey(cmd)
	filePath := args[0]
	content, err := a.encryption.Show(userKey, filePath)
	if err != nil {
		log.Fatal(err)
	}
	trimmedContent := strings.TrimSuffix(content, "\n")
	fmt.Println(trimmedContent)
}

func (a *action) getKey(cmd *cobra.Command) []byte {
	keyValue, err := cmd.Flags().GetString("key")

	if err != nil {
		log.Fatalf("failed to retrieve key: %v", err)
	}

	if keyValue != "" {
		return DeriveKey(keyValue)
	}

	storedKey, err := GetStoredKey()
	if err != nil {
		log.Fatalf("failed to retrieve stored key: %v", err)
	}

	return storedKey
}
