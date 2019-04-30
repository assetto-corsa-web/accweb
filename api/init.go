package api

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

var (
	verifyKey     *rsa.PublicKey
	signKey       *rsa.PrivateKey
	adminPassword string
	modPassword   string
	roPassword    string // read only
)

func init() {
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
