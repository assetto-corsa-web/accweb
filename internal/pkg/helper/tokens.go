package helper

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/assetto-corsa-web/accweb/internal/pkg/cfg"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GenerateTokenKeysIfNotPresent(config *cfg.Config) {
	// Based on https://stackoverflow.com/questions/64104586/use-golang-to-get-rsa-key-the-same-way-openssl-genrsa
	bitSize := 4096

	logrus.Info("accweb: checking for secrets...")

	// make sure files don't already exist
	if Exists(config.Auth.PublicKeyPath) && Exists(config.Auth.PrivateKeyPath) {
		logrus.WithField("file", config.Auth.PublicKeyPath).Info("Public/private keys already exists, not attempting regeneration")
		return
	} else if Exists(config.Auth.PublicKeyPath) {
		logrus.WithField("publicKeyFile", config.Auth.PublicKeyPath).
			WithField("privateKeyFile", config.Auth.PrivateKeyPath).
			Fatal("Only public key file is present (private key file missing) - remove the public key file to regenerate keys on startup.")
	} else if Exists(config.Auth.PrivateKeyPath) {
		logrus.WithField("publicKeyFile", config.Auth.PublicKeyPath).
			WithField("privateKeyFile", config.Auth.PrivateKeyPath).
			Fatal("Only private key file is present (public key file missing) - remove the private key file to regenerate keys on startup.")
	}

	// Neither key exists, let's generate them
	// create path if required ("secrets" directory by default)
	privateKeyPath := filepath.Join(".", filepath.Dir(config.Auth.PrivateKeyPath))
	if err := os.MkdirAll(privateKeyPath, os.ModePerm); err != nil {
		logrus.WithField("err", err).Fatal("Failed creating public key path directory!")
	}

	publicKeyPath := filepath.Join(".", filepath.Dir(config.Auth.PublicKeyPath))
	if err := os.MkdirAll(publicKeyPath, os.ModePerm); err != nil {
		logrus.WithField("err", err).Fatal("Failed creating public key path directory!")
	}

	key, err := rsa.GenerateKey(rand.Reader, bitSize)

	if err != nil {
		logrus.WithField("err", err).
			WithField("publicKeyPath", publicKeyPath).
			WithField("privateKeyPath", privateKeyPath).
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
	if err := ioutil.WriteFile(config.Auth.PrivateKeyPath, keyPEM, 0700); err != nil {
		logrus.WithField("file", config.Auth.PrivateKeyPath).WithField("err", err).Fatal("Failed writing private key to file.")
	}

	// Write public key to file.
	if err := ioutil.WriteFile(config.Auth.PublicKeyPath, pubPEM, 0755); err != nil {
		logrus.WithField("file", config.Auth.PublicKeyPath).WithField("err", err).Fatal("Failed writing public key to file.")
	}

	logrus.WithField("publicKeyPath", config.Auth.PublicKeyPath).
		WithField("privateKeyPath", config.Auth.PrivateKeyPath).
		Info("accweb: new secrets were generated.")
}
