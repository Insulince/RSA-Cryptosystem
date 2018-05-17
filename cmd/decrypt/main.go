package main

import (
	"fmt"
	"os"
	"flag"
	"rsa-cryptosystem/pkg/models"
	"rsa-cryptosystem/pkg/rsa"
)

func main() () {
	encryptedMessageValue := flag.Int64("encrypted-message", 0, "The encrypted-message to be decrypted.")
	privateKeyValue := flag.String("private-key", "", "Your private key.")
	flag.Parse()

	privateKey := models.Deserialize(*privateKeyValue)

	encryptedMessage := rsa.B(*encryptedMessageValue)

	decryptedMessage := rsa.B(0).Mod(rsa.B(0).Exp(encryptedMessage, &privateKey.Exponent, nil), &privateKey.Modulus)
	fmt.Fprintf(os.Stderr, "%v\n", decryptedMessage)
}
