package main

import (
	"log"

	"github.com/ecat/internal/core"
	"github.com/spf13/cobra"
)

var action = core.NewAction()

var rootCmd = &cobra.Command{
	Use:   "ecat",
	Short: "ecat is a cat command with encryption and decryption capabilities.",
	Args:  cobra.MaximumNArgs(1),
	Run:   action.RootAction,
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize and store encryption key",
	Run:   action.InitAction,
}

var encryptCmd = &cobra.Command{
	Use:   "encrypt [flags] <file>",
	Short: "Encrypt a file",
	Args:  cobra.ExactArgs(1),
	Run:   action.EncryptAction,
}

var decryptCmd = &cobra.Command{
	Use:   "decrypt [flags] <file>",
	Short: "Decrypt a file",
	Args:  cobra.ExactArgs(1),
	Run:   action.DecryptAction,
}

var showCmd = &cobra.Command{
	Use:   "show [flags] <file>",
	Short: "Show the content of an encrypted file",
	Args:  cobra.ExactArgs(1),
	Run:   action.ShowAction,
}

var editCmd = &cobra.Command{
	Use:   "edit [flags] <file>",
	Short: "Edit the content of an encrypted file",
	Args:  cobra.ExactArgs(1),
	Run:   action.EditAction,
}

func init() {
	rootCmd.PersistentFlags().StringP("key", "k", "", "Specify a custom encryption key")
	rootCmd.PersistentFlags().BoolVarP(&core.UseEncrypt, "encrypt", "e", false, "Encrypt a file")
	rootCmd.PersistentFlags().BoolVarP(&core.UseDecrypt, "decrypt", "d", false, "Decrypt a file")

	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(encryptCmd)
	rootCmd.AddCommand(decryptCmd)
	rootCmd.AddCommand(showCmd)
	rootCmd.AddCommand(editCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
