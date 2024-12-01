package main

import (
	"fmt"
	"log"

	"github.com/ecat/internal/core"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ecat",
	Short: "ecat is a tool for encrypting and decrypting files",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
		userKey := getKey(cmd)
		filePath := args[0]
		encryption := core.NewEncryption()
		content, err := encryption.Show(userKey, filePath)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(content)
	},
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize and store encryption key",
	Run: func(cmd *cobra.Command, args []string) {
		core.InitKey("")
	},
}

var encryptCmd = &cobra.Command{
	Use:   "encrypt [flags] <file>",
	Short: "Encrypt a file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		userKey := getKey(cmd)
		filePath := args[0]
		encryption := core.NewEncryption()
		err := encryption.Encrypt(userKey, filePath)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Encryption successful!")
	},
}

var decryptCmd = &cobra.Command{
	Use:   "decrypt [flags] <file>",
	Short: "Decrypt a file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		userKey := getKey(cmd)
		filePath := args[0]
		encryption := core.NewEncryption()
		err := encryption.Decrypt(userKey, filePath)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Decryption successful!")
	},
}

var showCmd = &cobra.Command{
	Use:   "show [flags] <file>",
	Short: "Show the content of an encrypted file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		userKey := getKey(cmd)
		filePath := args[0]
		encryption := core.NewEncryption()
		content, err := encryption.Show(userKey, filePath)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(content)
	},
}

var editCmd = &cobra.Command{
	Use:   "edit [flags] <file>",
	Short: "Edit the content of an encrypted file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		userKey := getKey(cmd)
		filePath := args[0]
		err := core.EditEncryptedFile([]byte(userKey), filePath, core.DefaultEditor)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("File edited successfully!")
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("key", "k", "", "Specify a custom encryption key")

	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(encryptCmd)
	rootCmd.AddCommand(decryptCmd)
	rootCmd.AddCommand(showCmd)
	rootCmd.AddCommand(editCmd)
}

func getKey(cmd *cobra.Command) []byte {
	keyValue, err := cmd.Flags().GetString("key")
	if err != nil {
		log.Fatalf("failed to retrieve key: %v", err)
	}

	if keyValue != "" {
		return core.DeriveKey(keyValue)
	}

	storedKey, err := core.GetStoredKey()
	if err != nil {
		log.Fatalf("failed to retrieve stored key: %v", err)
	}

	return storedKey
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
