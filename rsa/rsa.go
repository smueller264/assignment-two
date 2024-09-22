package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func encrypt(plaintext string, e, n *big.Int) []*big.Int {
	var encrypted []*big.Int
	for _, char := range plaintext {
		m := big.NewInt(int64(char))
		c := new(big.Int).Exp(m, e, n)
		encrypted = append(encrypted, c)
	}
	return encrypted
}

func decrypt(ciphertext []*big.Int, d, n *big.Int) string {
	decrypted := ""
	for _, c := range ciphertext {
		m := new(big.Int).Exp(c, d, n)
		decrypted += string(rune(m.Int64()))
	}
	return decrypted
}

func main() {

	var i int
	var p big.Int
	var q big.Int
	var e big.Int
	var plaintext string

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter p: ")
	fmt.Scan(&p)
	fmt.Print("Enter q: ")
	fmt.Scan(&q)
	fmt.Print("Enter e: ")
	fmt.Scan(&e)

	phi := new(big.Int).Mul(new(big.Int).Sub(&p, big.NewInt(1)), new(big.Int).Sub(&q, big.NewInt(1)))
	n := new(big.Int).Mul(&p, &q)

	kpriv := new(big.Int)
	kpriv.ModInverse(&e, phi)

	fmt.Println("phi:", phi)
	fmt.Println("n:", n)
	fmt.Println("d:", kpriv)

	fmt.Print("Press 1 for encryption or 2 for decryption: ")
	fmt.Scan(&i)
	switch i {
	case 1:
		fmt.Print("Enter your plaintext: ")
		fmt.Scan(&plaintext)
		encrypted := encrypt(plaintext, &e, n)
		fmt.Println("Ciphertext:", encrypted)

	case 2:
		fmt.Print("Enter your ciphertext in 4 space-seperated digits: ")
		encrypted, _ := reader.ReadString('\n')
		encrypted = strings.TrimSpace(encrypted)
		parts := strings.Split(encrypted, " ")
		fmt.Print(parts)
		var message []*big.Int
		for _, part := range parts {
			c := new(big.Int)
			c.SetString(part, 10)
			message = append(message, c)
		}
		decrypted := decrypt(message, kpriv, n)
		fmt.Println("Plaintext:", decrypted)

	}

}
