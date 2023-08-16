package ssh

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestGenKeyPair(t *testing.T) {
	t.Run("Generate and Verify Key Pair", func(t *testing.T) {
		dir, err := os.MkdirTemp("", "x8sync")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(dir)
		bitSize := 4096
		if err := GenKeyPair(bitSize, dir); err != nil {
			t.Errorf("Error generating key pair: %v", err)
			return
		}

		privateKeyPath := filepath.Join(dir, "id_rsa_x8sync")
		publicKeyPath := filepath.Join(dir, "id_rsa_x8sync.pub")

		// Read private and public key files
		privateKeyBytes, err := ioutil.ReadFile(privateKeyPath)
		if err != nil {
			t.Errorf("Error reading private key file: %v", err)
			return
		}

		publicKeyBytes, err := ioutil.ReadFile(publicKeyPath)
		if err != nil {
			t.Errorf("Error reading public key file: %v", err)
			return
		}

		// Example: Verify that private key starts with "-----BEGIN RSA PRIVATE KEY-----"
		expectedPrivatePrefix := "-----BEGIN RSA PRIVATE KEY-----"
		if string(privateKeyBytes[:len(expectedPrivatePrefix)]) != expectedPrivatePrefix {
			t.Errorf("Private key format is incorrect")
		}

		// Example: Verify that public key starts with "ssh-rsa"
		expectedPublicPrefix := "ssh-rsa"
		if string(publicKeyBytes[:len(expectedPublicPrefix)]) != expectedPublicPrefix {
			t.Errorf("Public key format is incorrect")
		}
	})
}
