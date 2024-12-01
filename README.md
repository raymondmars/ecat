# ecat

ecat is an enhanced version of the Unix `cat` command with encryption and decryption capabilities. It's designed to seamlessly handle both encrypted and plain text files.

## Features

- **Smart File Detection**: Automatically detects if a file is encrypted by ecat
- **Transparent Reading**: Works like `cat` for non-encrypted files
- **File Encryption**: Encrypt files with AES encryption
- **Base64 Encoding**: Encrypted content is stored in Base64 format
- **Key Management**: Store a default encryption key or provide one at runtime

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
ecat decrypt /path/to/file

# Show encrypted file content
ecat show /path/to/file
```

### Using Custom Keys
```bash
# Encrypt with custom key
ecat encrypt -k "yourkey" /path/to/file

# Decrypt with custom key
ecat decrypt -k "yourkey" /path/to/file

# Show with custom key
ecat show -k "yourkey" /path/to/file
ecat -k "yourkey" /path/to/file
```
## Development    
Requirements:

  - Go 1.18 or higher   
  - Make   

## License    
MIT License

## Contributing    
Contributions are welcome! Please feel free to submit a Pull Request.
