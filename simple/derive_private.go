package ecdsakeyderivation

import (
	"crypto/ecdsa"
	"math/big"
)

// DerivePrivate just adds derivationKeySeed hash to the private key:
// derivedKey = key + hash(derivationKeySeed)
func DerivePrivate(key *ecdsa.PrivateKey, derivationKeySeed []byte) *ecdsa.PrivateKey {
	// derivationKey = hash(derivationKeySeed)
	derivationKey := calcDerivationKey(derivationKeySeed)

	// Derived private key
	var derivedKey ecdsa.PrivateKey
	derivedKey.Curve = key.Curve
	derivedKey.D = new(big.Int).Add(key.D, derivationKey)

	// Derived public key
	derivedKey.X, derivedKey.Y = derivedKey.Curve.ScalarBaseMult(scalarToBytes(derivedKey.D))

	return &derivedKey
}
