package main

import (
	"fmt"
)

func main() {

	var i int
	var plaintext string
	var key int
	var ciphertext string

	fmt.Print("Press 1 for encryption or 2 for decryption: ")
	fmt.Scan(&i)
	switch i {
	case 1:
		fmt.Print("Enter your plaintext: ")
		fmt.Scan(&plaintext)
		fmt.Print("Enter the key: ")
		fmt.Scan(&key)
		encryptedString := enryption(plaintext, key)
		fmt.Println("Plaintext:", plaintext)
		fmt.Println("Encrypted:", encryptedString)

	case 2:
		fmt.Print("Enter your ciphertext: ")
		fmt.Scan(&ciphertext)
		fmt.Print("Enter the key: ")
		fmt.Scan(&key)
		decryptedString := decryption(ciphertext, key)
		fmt.Println("Ciphertext:", ciphertext)
		fmt.Println("Plaintext:", decryptedString)
	}

}

func enryption(plaintext string, key int) string {
	ciphertext := ""
	for _, char := range plaintext {
		if char >= 'A' && char <= 'Z' {
			ciphertext += string((char-'A'+rune(key))%26 + 'A')
		} else if char >= 'a' && char <= 'z' {
			ciphertext += string((char-'a'+rune(key))%26 + 'a')
		} else {
			ciphertext += string(char)
		}
	}
	return ciphertext
}

func decryption(ciphertext string, key int) string {
	encrypted := enryption(ciphertext, 26-key)
	return encrypted
}
