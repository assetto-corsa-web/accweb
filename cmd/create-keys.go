package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/assetto-corsa-web/accweb/internal/pkg/cfg"
	"github.com/assetto-corsa-web/accweb/internal/pkg/helper"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
)

const configFile = "config.yml"

// Based on https://stackoverflow.com/questions/64104586/use-golang-to-get-rsa-key-the-same-way-openssl-genrsa
func main() {
	bitSize := 4096

	c := cfg.Load(configFile)

	logrus.Info("accweb: checking for secrets...")

	// make sure files don't already exist
	if helper.Exists(c.Auth.PublicKeyPath) {
		logrus.WithField("file", c.Auth.PublicKeyPath).Fatal("Public key file already exists!")
	}

	if helper.Exists(c.Auth.PrivateKeyPath) {
		logrus.WithField("file", c.Auth.PrivateKeyPath).Fatal("Private key file already exists!")
	}

	// create path if required ("secrets" directory by default)
	privateKeyPath := filepath.Join(".", filepath.Dir(c.Auth.PrivateKeyPath))
	if err := os.MkdirAll(privateKeyPath, os.ModePerm); err != nil {
		logrus.WithField("err", err).Fatal("Failed creating public key path directory!")
	}

	publicKeyPath := filepath.Join(".", filepath.Dir(c.Auth.PublicKeyPath))
	if err := os.MkdirAll(publicKeyPath, os.ModePerm); err != nil {
		logrus.WithField("err", err).Fatal("Failed creating public key path directory!")
	}

	key, err := rsa.GenerateKey(rand.Reader, bitSize)

	if err != nil {
		logrus.WithField("err", err).Fatal("Key generation failed!")
	}

	privateKeyBytes, err := x509.MarshalPKCS8PrivateKey(key)

	if err != nil {
		logrus.WithField("err", err).Fatal("Failed converting private token to a proper PKCS8 block.")
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
		logrus.WithField("err", err).Fatal("Failed converting public token to a proper PEM block.")
	}

	pubPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: publicKeyBytes,
		},
	)

	// Write private key to file.
	if err := ioutil.WriteFile(c.Auth.PrivateKeyPath, keyPEM, 0700); err != nil {
		logrus.WithField("file", c.Auth.PrivateKeyPath).WithField("err", err).Fatal("Failed writing private key to file.")
	}

	// Write public key to file.
	if err := ioutil.WriteFile(c.Auth.PublicKeyPath, pubPEM, 0755); err != nil {
		logrus.WithField("file", c.Auth.PublicKeyPath).WithField("err", err).Fatal("Failed writing public key to file.")
	}

	logrus.Info("accweb: new secrets have been generated.")
}
