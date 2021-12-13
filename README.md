# A proof of concept of a public key derivation for ECDSA (without knowledge of the private key)

WARNING! Do not use this code in production. It was never validated by a real specialist in cryptography.

# Simple

It is a demonstration of how to implement a simple key derivation algorithm for public keys without knowledge of private keys. This works for any asymmetric crypto algorithm with distributive property enabled for public key. The idea is very simple:

Let's imagine `s0` is the private key, then public key would be `P0 = s0*B` (where `B` is the basepoint of the curve). So if a new private key is derived as `s1 = s0 + d` then a new public key could be derived as `P1 = s1*B = (s0+d)*B`. Thus new public key is: `P1 = s0*B + d*B = P0 + d*B`. That's it.

# Demonstration

See the unit test:
```go
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
```
result:
```sh
xaionaro@void:~/go/src/github.com/xaionaro-go/ecdsakeyderivation$ go test ./... -count=1
ok  	github.com/xaionaro-go/ecdsakeyderivation	0.101s
```
