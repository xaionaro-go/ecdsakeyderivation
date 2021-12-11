package ecdsakeyderivation

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDerivePrivate(t *testing.T) {
	randReader := rand.New(rand.NewSource(0))
	derivationKey := []byte{1, 2, 3}
	message := []byte{4, 5, 6, 7}
	var noHash crypto.Hash

	privKey, err := ecdsa.GenerateKey(elliptic.P521(), randReader)
	require.NoError(t, err)
	{
		signature, err := privKey.Sign(randReader, message, noHash)
		require.NoError(t, err)
		pubKey := privKey.Public().(*ecdsa.PublicKey)
		require.True(t, ecdsa.VerifyASN1(pubKey, message, signature))
	}

	derivedPrivKey := DerivePrivate(privKey, derivationKey)

	signature, err := ecdsa.SignASN1(randReader, derivedPrivKey, message)
	require.NoError(t, err)

	derivedPubKey := derivedPrivKey.Public().(*ecdsa.PublicKey)
	require.True(t, ecdsa.VerifyASN1(derivedPubKey, message, signature))
}
