package main

import (
	"crypto/rand"
	"rsa-cryptosystem/pkg/models"
	"rsa-cryptosystem/pkg/rsa"
	"flag"
	"fmt"
	"os"
)

func main() () {
	primeBitLength := flag.Int("prime-bit-length", 10, "The bit length of the random primes used for key-generation. Reccomended length is ~10 for quick performance.")
	flag.Parse()

	prime1, err := rand.Prime(rand.Reader, *primeBitLength)
	if err != nil {
		panic(err)
	}
	prime2, err := rand.Prime(rand.Reader, *primeBitLength)
	if err != nil {
		panic(err)
	}
	for prime1.Cmp(prime2) == 0 {
		prime2, err = rand.Prime(rand.Reader, *primeBitLength)
		if err != nil {
			panic(err)
		}
	}
	primeProduct := rsa.B(0).Mul(prime1, prime2)

	totient := rsa.Totient(*prime1, *prime2)

	publicExponent := rsa.PublicExponent(totient)
	privateExponent := rsa.ModularMultiplicativeInverse(*publicExponent, totient)

	publicKey := models.Key{Modulus: *primeProduct, Exponent: *publicExponent}
	privateKey := models.Key{Modulus: *primeProduct, Exponent: *privateExponent}

	fmt.Fprintf(os.Stderr, "Public Key:\n%v\nPrivate Key:\n%v\n", publicKey.Serialize(), privateKey.Serialize())
}
