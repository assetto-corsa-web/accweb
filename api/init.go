package api

import (
	"crypto/rsa"
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

func init() {
	generateKeyFilesIfRequired()
	publicKey, err := ioutil.ReadFile(os.Getenv("ACCWEB_TOKEN_PUBLIC_KEY"))

	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err, "file": os.Getenv("ACCWEB_TOKEN_PUBLIC_KEY")}).Fatal("Token public key file not found")
	}

	privateKey, err := ioutil.ReadFile(os.Getenv("ACCWEB_TOKEN_PRIVATE_KEY"))

	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err, "file": os.Getenv("ACCWEB_TOKEN_PRIVATE_KEY")}).Fatal("Token private key file not found")
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
	adminPassword = os.Getenv("ACCWEB_ADMIN_PASSWORD")
	modPassword = os.Getenv("ACCWEB_MOD_PASSWORD")
	roPassword = os.Getenv("ACCWEB_RO_PASSWORD")

	if adminPassword == "" {
		logrus.Fatal("ACCWEB_ADMIN_PASSWORD must be set")
	}
}

func generateKeyFilesIfRequired() {
	_, err := os.Stat(os.Getenv("ACCWEB_TOKEN_PUBLIC_KEY"))

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
