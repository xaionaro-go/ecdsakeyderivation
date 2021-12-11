package ecdsakeyderivation

import (
	"crypto/sha512"
	"math/big"
)

func calcDerivationKey(derivationKeySeed []byte) *big.Int {
	b := sha512.Sum512(derivationKeySeed)
	return new(big.Int).SetBytes(b[:])
}

func bigIntToScalar(i *big.Int) []byte {
	return i.Bytes()
}
