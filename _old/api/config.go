package api

import (
	"crypto/rsa"
	"github.com/assetto-corsa-web/accweb/cfg"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
)

var (
	verifyKey     *rsa.PublicKey
	signKey       *rsa.PrivateKey
	adminPassword string
	modPassword   string
	roPassword    string // read only
)

func LoadConfig() {
	config := cfg.Get()
	generateKeyFilesIfRequired()
	publicKey, err := ioutil.ReadFile(config.Auth.PublicKeyPath)

	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err, "file": config.Auth.PublicKeyPath}).Fatal("Token public key file not found")
	}

	privateKey, err := ioutil.ReadFile(config.Auth.PrivateKeyPath)

	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err, "file": config.Auth.PrivateKeyPath}).Fatal("Token private key file not found")
	}

	verify, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)

	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Fatal("Error parsing token public key")
	}

	sign, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)

	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Fatal("Error parsing token private key")
	}

	verifyKey = verify
	signKey = sign
	adminPassword = config.Auth.AdminPassword
	modPassword = config.Auth.ModeratorPassword
	roPassword = config.Auth.ReadOnlyPassword

	if adminPassword == "" {
		logrus.Fatal("Admin password must be set")
	}
}

func generateKeyFilesIfRequired() {
	_, err := os.Stat(cfg.Get().Auth.PublicKeyPath)

	if os.IsNotExist(err) {
		logrus.Info("Private/public token key files not found, generating new ones")
		keyFileScript := "./gen_rsa_keys.cmd"

		if runtime.GOOS != "windows" {
			keyFileScript = "./gen_rsa_keys.sh"
		}

		_, err = os.Stat(keyFileScript)

		if err == nil {
			cmd := exec.Command(keyFileScript)

			if err := cmd.Run(); err != nil {
				logrus.WithFields(logrus.Fields{"err": err}).Fatal("Error generating key files through OpenSSL. Make sure the script gen_rsa_keys is in place")
			}

			logrus.Info("Private/public token key files generated")
		}
	}
}
