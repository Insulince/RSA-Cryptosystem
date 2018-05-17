package main

import (
	"rsa-cryptosystem/pkg/models"
	"rsa-cryptosystem/pkg/rsa"
	"fmt"
	"flag"
	"os"
)

func main() () {
	messageValue := flag.Int64("message", 0, "The message to be encrypted.")
	publicKeyValue := flag.String("public-key", "", "Your public key.")
	flag.Parse()

	publicKey := models.Deserialize(*publicKeyValue)

	message := rsa.B(*messageValue)
	if message.Cmp(&publicKey.Modulus) >= 0 {
		panic("Message must be smaller than the public key's modulus (prime product)!")
	}

	encryptedMessage := rsa.B(0).Mod(rsa.B(0).Exp(message, &publicKey.Exponent, nil), &publicKey.Modulus)
	fmt.Fprintf(os.Stderr, "%v\n", encryptedMessage)
}
