# RSA Cryptosystem

This is a rudimentary implementation of the RSA cryptosystem in Go.

A difference between this system and how most RSA systems work is I choose a prime for the public key exponent (`e`) randomly instead of the usual 65537, and a consequence of that is that the private key exponent (`d`, the modular multiplicative inverse of `e`) is also randomly shifted each run. I do not know is this leads to a stronger cryptosystem, and in fact I doubt it does, but I wanted to implement it that way.

## Building and Running

Note: This project was created in a *Windows* environment

• Build the executables. I reccomend using the `./bin` folder from the project root:
```bash
go build -o ./bin/keygen.exe ./cmd/keygen/main.go
go build -o ./bin/encrypt.exe ./cmd/encrypt/main.go
go build -o ./bin/decrypt.exe ./cmd/decrypt/main.go
```
• Run `keygen` with the `-prime-bit-length` flag. I set it to 10 for quick performance. This will output your public and private keys:
```bash
./bin/keygen.exe -prime-bit-length=10
```
• Run `encrypt` with the `-message` flag set to whatever int64 you want to encrypt and then set `-public-key` to the value output above for the public key. This will output the encrypted version of your message:
```bash
./bin/encrypt.exe -message=12 -public-key=[YOUR_PUBLIC_KEY]
```
• Run `decrypt` with the `-encrypted-message` flag set to the encrypted messagte from above, and then set `-private-key` to the value output above in the first command for the private key. This will output the decrypted message, which should be the original message:
```bash
./bin/decrypt.exe -encrypted-message=[YOUR_ENCRYPTED_MESSAGE] -private-key=[YOUR_PRIVATE_KEY]
```

### Limitations

• You can only encrypt int64s currently.

• The value of your message must be less than the prime product (p * q), which is around 800,000 on average when the `prime-bit-length` is set to 10. Increasing the `prime-bit-length` increases that limit, however it also decreases the speed of the cryptographic operations.
