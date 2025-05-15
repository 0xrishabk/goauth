package keys

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

var (
	// Private Key which is loaded.
	PrivateKey *ecdsa.PrivateKey
	// Public Key which is loaded
	PublicKey *ecdsa.PublicKey
)

func ConfigureKeys() error {
	if _, err := os.Stat("keys/private_key.pem"); os.IsNotExist(err) {
		var err error
		curve := elliptic.P256()
		PrivateKey, err = ecdsa.GenerateKey(curve, rand.Reader)
		if err != nil {
			return err
		}
		PublicKey = &PrivateKey.PublicKey
		os.Mkdir("keys", 0755)
		err = savePrivateKey(PrivateKey)
		if err != nil {
			return err
		}
		err = savePublicKey(PublicKey)
		if err != nil {
			return err
		}
	} else {
		err := loadPrivateKey()
		if err != nil {
			return err
		}
		err = loadPublicKey()
		if err != nil {
			return err
		}
	}
	return nil
}

func loadPrivateKey() error {
	pemData, err := os.ReadFile("keys/private_key.pem")
	if err != nil {
		return err
	}
	block, _ := pem.Decode(pemData)
	if block == nil || block.Type != "EC PRIVATE KEY" {
		return errors.New("failed to decode PEM block containing private key")
	}

	PrivateKey, err = x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return err
	}
	fmt.Println("Loaded Private Key from the PEM block")
	return nil
}

func loadPublicKey() error {
	pemData, err := os.ReadFile("keys/public_key.pem")
	if err != nil {
		return err
	}
	block, _ := pem.Decode(pemData)
	if block == nil || block.Type != "PUBLIC KEY" {
		return errors.New("failed to decode PEM block containing public key")
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	var ok bool
	PublicKey, ok = pubKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("not ECDSA public key")
	}
	fmt.Println("Loaded Public Key from the PEM block")
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
