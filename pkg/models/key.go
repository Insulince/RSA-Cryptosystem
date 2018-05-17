package models

import (
	"math/big"
	"encoding/base64"
	"bytes"
	"encoding/gob"
)

// Public and Private key structure.
type Key struct {
	Modulus             big.Int `json:"modulus"`
	Exponent            big.Int `json:"exponent"`
}

func (key *Key) Serialize() (serializedKey string) {
	// Serialize Key into Binary.
	serializationBuffer := bytes.Buffer{}
	err := gob.NewEncoder(&serializationBuffer).Encode(key)
	if err != nil {
		panic(err)
	}

	// Encode Binary into Base64.
	serializedKey = base64.StdEncoding.EncodeToString(serializationBuffer.Bytes())

	return serializedKey
}

func Deserialize(serializedKey string) (deserializedKey Key) {
	// Decode Base64 into Binary.
	decodedKey, err := base64.StdEncoding.DecodeString(serializedKey)
	if err != nil {
		panic(err)
	}

	// Deserialize Binary into Key.
	deserializationBuffer := bytes.NewBuffer(decodedKey)
	err = gob.NewDecoder(deserializationBuffer).Decode(&deserializedKey)
	if err != nil {
		panic(err)
	}

	return deserializedKey
}
