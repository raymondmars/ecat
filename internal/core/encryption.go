package core

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
)

var ECAT_FILE_HEADER_IDENTITY = []byte("_ECAT_")

type encryption struct {
}

func (e *encryption) Encrypt(key []byte, filePath string) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
		return err
	}
	encryptedData, err := e.internalEncrypt(data, key)
	if err != nil {
		log.Fatalf("Encryption failed: %v", err)
		return err
	}
	finalData := append(ECAT_FILE_HEADER_IDENTITY, encryptedData...)
	encoded := base64.StdEncoding.EncodeToString(finalData)

	// outputFile := filePath + ".ecat"
	return ioutil.WriteFile(filePath, []byte(encoded), 0600)
}

func (e *encryption) Decrypt(key []byte, filePath string) error {
	decryptedData, err := e.getEncryptData(key, filePath, false)
	if err != nil {
		log.Fatalf("Decryption failed: %v", err)
		return err
	}
	return ioutil.WriteFile(filePath, decryptedData, 0600)
}

func (e *encryption) Show(key []byte, filePath string) (string, error) {
	decryptedData, err := e.getEncryptData(key, filePath, true)
	if err != nil {
		log.Fatalf("Decryption failed: %v", err)
		return "", err
	}
	return string(decryptedData), nil
}

func (e *encryption) getEncryptData(key []byte, filePath string, onlyShow bool) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("Failed to read file: %v", err)
	}

	// Base64 decoding
	decodedData, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		if onlyShow {
			return data, nil
		}
		return nil, fmt.Errorf("Base64 decoding failed: %v", err)
	}

	// Check the file header
	if !bytes.HasPrefix(decodedData, ECAT_FILE_HEADER_IDENTITY) {
		if onlyShow {
			return data, nil
		}
		return nil, fmt.Errorf("The file format is incorrect or it is not a file encrypted by ecat.")
	}

	// Remove file header
	encryptedData := decodedData[len(ECAT_FILE_HEADER_IDENTITY):]

	decryptedData, err := e.internalDecrypt(encryptedData, key)
	if err != nil {
		return nil, fmt.Errorf("Decryption failed: %v", err)
	}
	return decryptedData, nil
}

func (e *encryption) internalEncrypt(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)
	return ciphertext, nil
}

func (e *encryption) internalDecrypt(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(data) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(data, data)
	return data, nil
}
