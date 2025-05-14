package generator

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func ConfigureKeys() error {
	if _, err := os.Stat("keys/private_key.pem"); os.IsNotExist(err) {
		curve := elliptic.P256()
		privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
		if err != nil {
			return err
		}
		publicKey := &privateKey.PublicKey
		os.Mkdir("keys", 0755)
		err = savePrivateKey(privateKey)
		if err != nil {
			return err
		}
		err = savePublicKey(publicKey)
		if err != nil {
			return err
		}
	}
	return nil
}

func savePrivateKey(privateKey *ecdsa.PrivateKey) error {
	privateKeyFile, err := os.Create("keys/private_key.pem")
	if err != nil {
		return err
	}
	defer privateKeyFile.Close()

	privateKeyBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return err
	}
	err = pem.Encode(privateKeyFile, &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: privateKeyBytes,
	})
	if err != nil {
		return err
	}
	fmt.Println("Private key saved to", privateKeyFile.Name())
	return nil
}

func savePublicKey(publicKey *ecdsa.PublicKey) error {
	publicKeyFile, err := os.Create("keys/public_key.pem")
	if err != nil {
		return err
	}
	defer publicKeyFile.Close()

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	err = pem.Encode(publicKeyFile, &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})
	if err != nil {
		return err
	}
	fmt.Println("Public key saved to", publicKeyFile.Name())
	return nil
}
