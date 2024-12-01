package core

type Encryption interface {
	Encrypt(key []byte, filePath string) error
	Decrypt(key []byte, filePath string) error
	Show(key []byte, filePath string) (string, error)
}

func NewEncryption() Encryption {
	return &encryption{}
}
