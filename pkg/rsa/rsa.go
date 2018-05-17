package rsa

import (
	"math/big"
	"log"
	"fmt"
	"crypto/rand"
)

// Quick way to generate throw-away big ints.
func B(v int64) (bigInt *big.Int) {
	return big.NewInt(v)
}

// Finds the x in "(a * x) % m = 1" (Can be written as "a * x ≡ 1 (mod m)")
func ModularMultiplicativeInverse(a big.Int, m big.Int) (modularMultiplicativeInverse *big.Int) {
	if a.Cmp(&m) >= 0 {
		log.Fatalf("a (%v) must be less than m (%v) for mod mult inv function to work!", a, m)
		return B(1)
	}

	for i := B(1); i.Cmp(&m) == -1; i.Add(i, B(1)) {
		var ami big.Int
		ami.Mul(&a, i)
		var amiMod big.Int
		amiMod.Mod(&ami, &m)
		if amiMod.Cmp(B(1)) == 0 {
			return i
		}
	}

	fmt.Println("Default branch of mod mult inv function was taken.")
	return B(1)
}

func Totient(prime1 big.Int, prime2 big.Int) (totient big.Int) {
	var prime1Minus1 big.Int
	prime1Minus1.Sub(&prime1, B(1))
	var prime2Minus1 big.Int
	prime2Minus1.Sub(&prime2, B(1))
	totient.Mul(&prime1Minus1, &prime2Minus1)
	return totient
}

func PublicExponent(totient big.Int) (publicExponent *big.Int) {
	for {
		publicExponent, _ = rand.Prime(rand.Reader, totient.BitLen()-1)

		copyPublicExponent := new(big.Int).Set(publicExponent)
		copyTotient := new(big.Int).Set(&totient)
		if Coprime(*copyPublicExponent, *copyTotient) {
			break
		}
	}

	return publicExponent
}

func Coprime(a big.Int, b big.Int) (areCoprime bool) {
	for a.Cmp(&b) != 0 {
		if a.Cmp(&b) == 1 {
			a.Sub(&a, &b)
		} else {
			b.Sub(&b, &a)
		}
	}

	return a.Cmp(B(1)) == 0
}
