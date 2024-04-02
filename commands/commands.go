package commands

import (
	"bufio"
	crypt "encryptor/pckg/crypt"
	"fmt"
	"os"
	"strings"
)

func Commands() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: encryptor <operation> <file>")
		os.Exit(1)
	}

	operation := os.Args[1]
	filename := os.Args[2]

	switch operation {
	case "encrypt":
		_, err := crypt.Encrypt(filename)
		if err != nil {
			fmt.Println("Error during encryption:", err)
			os.Exit(1)
		}
	case "decrypt":
		fmt.Print("Enter decryption key: ")
		reader := bufio.NewReader(os.Stdin)
		decryptionKey, _ := reader.ReadString('\n')
		decryptionKey = strings.TrimSpace(decryptionKey)
		_, err := crypt.Decrypt(decryptionKey, filename)
		if err != nil {
			fmt.Println("Error during decryption:", err)
			os.Exit(1)
		}
	default:
		fmt.Println("Error: Invalid operation. Please enter 'encrypt' or 'decrypt'.")
		os.Exit(1)
	}
}
