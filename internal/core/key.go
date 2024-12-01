package core

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var KEY_FILE_NAME = "key.txt"

func InitKey(data string) {
	var inputKey string
	if data == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Please enter the key: ")
		input, err := reader.ReadString('\n')
		if err != nil || strings.TrimSpace(input) == "" {
			fmt.Printf("Failed to read key: %v", err)
			os.Exit(1)
		}
		inputKey = strings.TrimSpace(input)
	} else {
		inputKey = data
	}

	inputKey = string(DeriveKey(inputKey))

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Failed to get user home directory: %v", err)
		os.Exit(1)
	}
	ecatDir := filepath.Join(homeDir, ".ecat")
	if _, err := os.Stat(ecatDir); os.IsNotExist(err) {
		os.Mkdir(ecatDir, 0700)
	}
	keyFile := filepath.Join(ecatDir, KEY_FILE_NAME)
	err = ioutil.WriteFile(keyFile, []byte(inputKey), 0600)
	if err != nil {
		fmt.Printf("Failed to save key:%v", err)
		os.Exit(1)
	}
	fmt.Println("Key saved!")
}

func GetStoredKey() ([]byte, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("Failed to get user home directory: %v", err)
	}
	keyFile := filepath.Join(homeDir, ".ecat", KEY_FILE_NAME)
	key, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return nil, fmt.Errorf("Failed to read key file: %v, You can run ecat init command to generate a key file.", err)
	}
	return key, nil
}

func DeriveKey(password string) []byte {
	hash := sha256.Sum256([]byte(password))
	return hash[:]
}
