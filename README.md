# ecat   
  [![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/ecat)](https://goreportcard.com/report/github.com/yourusername/ecat)
  [![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)   

ecat is a command-line tool that enhances the functionality of the standard Unix `cat` command by adding encryption and decryption capabilities. With ecat, you can securely view, encrypt, decrypt, and edit files directly from the terminal, handling both encrypted and plain text files seamlessly.


## Features

- **Smart File Detection**: ecat automatically detects if a file is encrypted with its encryption scheme. If the file is encrypted, it attempts to decrypt it using your stored key or a provided key. If not, it displays the file content just like the regular `cat` command.

- **Transparent Reading**: For non-encrypted files, ecat behaves exactly like the standard `cat` command, allowing you to use it as a drop-in replacement without changing your workflow.

- **File Encryption**: Easily encrypt your files using AES encryption to protect sensitive information. Encrypted files are stored securely and can only be decrypted with the correct key.

- **File Decryption**: Decrypt encrypted files to retrieve the original content. ecat ensures that only users with the correct key can access the decrypted content.

- **Edit Encrypted Files**: Securely edit encrypted files using your preferred text editor. ecat decrypts the file to a temporary location, opens it for editing, and then re-encrypts it upon saving, ensuring your data remains protected throughout the process.

- **Base64 Encoding**: Encrypted content is stored in Base64 format, making it safe for transmission and storage in systems that may not handle binary data well.

- **Key Management**: You can store a default encryption key for convenience or specify a custom key at runtime for additional security.

## Installation

### From Source

```bash
# Clone the repository
git clone https://github.com/yourusername/ecat.git

# Change to project directory
cd ecat

# Install to /usr/local/bin
make install

# Uninstall
make uninstall
```
## Usage

### Basic Usage

```bash
# View file content (works for both encrypted and plain files)
ecat /path/to/file

# Initialize encryption key
ecat init

# Encrypt a file
ecat encrypt /path/to/file

# Decrypt a file
ecat decrypt /path/to/encrypted_file

# Edit encrypted file
ecat edit /path/to/encrypted_file

# Show encrypted file content
ecat show /path/to/encrypted_file
```

### Using Custom Keys
```bash
# Encrypt with custom key
ecat encrypt -k "yourkey" /path/to/file

# Decrypt with custom key
ecat decrypt -k "yourkey" /path/to/encrypted_file

# Show with custom key
ecat show -k "yourkey" /path/to/encrypted_file
ecat -k "yourkey" /path/to/encrypted_file

# Edit encrypted file with custom key
ecat edit -k "yourkey" /path/to/encrypted_file
```

## Editing Encrypted Files
The edit command allows you to securely edit encrypted files. the ecat handles the decryption and encryption process for you:

  1. Decryption: ecat decrypts the specified encrypted file using your key.   
  2. Temporary Editing: It creates a temporary decrypted file and opens it in your default text editor (specified by the EDITOR environment variable, defaults to vi if not set).    
  3. Encryption: After you save and close the editor, ecat encrypts the modified content and updates the original encrypted file.    
  4. Cleanup: The temporary decrypted file is securely deleted after editing to prevent unauthorized access.   

### Example:   
```bash
ecat edit /path/to/encrypted_file
```
### Note:  
- Ensure you have initialized your encryption key using ecat init, or provide a custom key using the -k option.  
- The edit command requires a text editor to be set in the EDITOR environment variable. If not set, it defaults to vi.  
```bash
export EDITOR=nano
```  

## Development    
Requirements:

  - Go 1.18 or higher   
  - Make   

## License    
MIT License

## Contributing    
Contributions are welcome! Please feel free to submit a Pull Request.
