package ssh

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"path/filepath"

	"golang.org/x/crypto/ssh"
)

const (
	KeyBitSize = 4096
)

// GenKeyPair generates a pair of ssh keys (private and public) with the specified bit size.
func GenKeyPair(bitSize int, dir string) error {
	/*
		1. Generate private key
		2. Encode private key to PEM format bytes
		3. Write encoded private key bytes to file
		4. Encode public key bytes from private key
		5. Write public key bytes to file
	*/
	privateKey, err := genPrivateKey(bitSize)
	if err != nil {
		return err
	}
	privateKeyBytes := encodePrivateKeyToPEM(privateKey)
	privateKeyFile := filepath.Join(dir, "id_rsa_x8sync")
	if err := writeKeyToFile(privateKeyBytes, privateKeyFile); err != nil {
		return err
	}
	publicKeyBytes, err := encodePublicKey(&privateKey.PublicKey)
	if err != nil {
		return err
	}
	publicKeyFile := filepath.Join(dir, "id_rsa_x8sync.pub")
	if err := writeKeyToFile(publicKeyBytes, publicKeyFile); err != nil {
		return err
	}
	return nil
}

func genPrivateKey(bitSize int) (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, err
	}

	err = privateKey.Validate()
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func encodePrivateKeyToPEM(privateKey *rsa.PrivateKey) []byte {
	// Get ASN.1 DER format
	privDER := x509.MarshalPKCS1PrivateKey(privateKey)

	// pem.Block
	privBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privDER,
	}

	// Private key in PEM format
	privatePEM := pem.EncodeToMemory(&privBlock)
	return privatePEM
}

func encodePublicKey(privatekey *rsa.PublicKey) ([]byte, error) {
	publicRsaKey, err := ssh.NewPublicKey(privatekey)
	if err != nil {
		return nil, err
	}

	publicKeyBytes := ssh.MarshalAuthorizedKey(publicRsaKey)
	return publicKeyBytes, nil
}

func writeKeyToFile(keyBytes []byte, saveFileTo string) error {
	err := ioutil.WriteFile(saveFileTo, keyBytes, 0600)
	if err != nil {
		return err
	}
	return nil
}
