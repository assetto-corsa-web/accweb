package helper

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GenerateTokenKeysIfNotPresent(publicKeyPath string, privateKeyPath string) {
	// Based on https://stackoverflow.com/questions/64104586/use-golang-to-get-rsa-key-the-same-way-openssl-genrsa
	bitSize := 4096

	logrus.Info("accweb: checking for secrets...")

	// make sure files don't already exist
	if Exists(publicKeyPath) && Exists(privateKeyPath) {
		logrus.WithField("privateKeyPath", privateKeyPath).
			WithField("publicKeyPath", publicKeyPath).
			Info("Public/private keys already exists, not attempting regeneration")
		return
	} else if Exists(publicKeyPath) {
		logrus.WithField("publicKeyFile", publicKeyPath).
			WithField("privateKeyFile", privateKeyPath).
			Fatal("Only public key file is present (private key file missing) - remove the public key file to regenerate keys on startup.")
	} else if Exists(privateKeyPath) {
		logrus.WithField("publicKeyFile", publicKeyPath).
			WithField("privateKeyFile", privateKeyPath).
			Fatal("Only private key file is present (public key file missing) - remove the private key file to regenerate keys on startup.")
	}

	// Neither key exists, let's generate them
	// create path if required ("secrets" directory by default)
	privateKeyDirectoryPath := filepath.Join(".", filepath.Dir(privateKeyPath))
	if err := os.MkdirAll(privateKeyDirectoryPath, os.ModePerm); err != nil {
		logrus.WithField("err", err).Fatal("Failed creating public key path directory!")
	}

	publicKeyDirectoryPath := filepath.Join(".", filepath.Dir(publicKeyPath))
	if err := os.MkdirAll(publicKeyDirectoryPath, os.ModePerm); err != nil {
		logrus.WithField("err", err).Fatal("Failed creating public key path directory!")
	}

	key, err := rsa.GenerateKey(rand.Reader, bitSize)

	if err != nil {
		logrus.WithField("err", err).
			WithField("privateKeyPath", privateKeyPath).
			WithField("publicKeyPath", publicKeyPath).
			Fatal("Key generation failed!")
	}

	privateKeyBytes, err := x509.MarshalPKCS8PrivateKey(key)

	if err != nil {
		logrus.WithField("err", err).Fatal("Failed marshalling private token to a PKCS8 block.")
	}

	keyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PRIVATE KEY",
			Bytes: privateKeyBytes,
		},
	)

	// Export public key to public key certificate
	pub := key.Public()
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(pub.(*rsa.PublicKey))

	if err != nil {
		logrus.WithField("err", err).Fatal("Failed marshalling public token to a PEM block.")
	}

	pubPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: publicKeyBytes,
		},
	)

	// Write private key to file.
	if err := ioutil.WriteFile(privateKeyPath, keyPEM, 0700); err != nil {
		logrus.WithField("file", privateKeyPath).WithField("err", err).Fatal("Failed writing private key to file.")
	}

	// Write public key to file.
	if err := ioutil.WriteFile(publicKeyPath, pubPEM, 0755); err != nil {
		logrus.WithField("file", publicKeyPath).WithField("err", err).Fatal("Failed writing public key to file.")
	}

	logrus.WithField("privateKeyPath", privateKeyPath).
		WithField("publicKeyPath", publicKeyPath).
		Info("accweb: new secrets were generated.")
}
