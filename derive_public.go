package ecdsakeyderivation

import (
	"crypto/ecdsa"
)

// DerivePublic just adds derivationKeySeed hash multiplied by basepoint to the public key:
// derivedKey = key + basepoint*hash(derivationKeySeed)
func DerivePublic(key *ecdsa.PublicKey, derivationKeySeed []byte) *ecdsa.PublicKey {
	// derivationKey = hash(derivationKeySeed)
	derivationKey := calcDerivationKey(derivationKeySeed)

	var addKey, derivedKey ecdsa.PublicKey
	addKey.Curve, derivedKey.Curve = key.Curve, key.Curve
	addKey.X, addKey.Y = derivedKey.Curve.ScalarBaseMult(bigIntToScalar(derivationKey))
	derivedKey.X, derivedKey.Y = derivedKey.Curve.Add(key.X, key.Y, addKey.X, addKey.Y)
	return &derivedKey

}
