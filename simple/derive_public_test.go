package ecdsakeyderivation

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDerivePublic(t *testing.T) {
	randReader := rand.New(rand.NewSource(0))
	derivationKey := []byte{1, 2, 3}

	privKey, err := ecdsa.GenerateKey(elliptic.P521(), randReader)
	require.NoError(t, err)

	derivedPubKey := DerivePublic(privKey.Public().(*ecdsa.PublicKey), derivationKey)

	derivedPrivKey := DerivePrivate(privKey, derivationKey)
	require.NotEqual(t, privKey, derivedPrivKey)

	require.Equal(t,
		derivedPrivKey.Public().(*ecdsa.PublicKey),
		derivedPubKey,
	)
}
