package auth

import (
	"crypto/rsa"
	"github.com/assetto-corsa-web/accweb/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/emvi/logbuch"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

// LoadConfig loads the JWT private/public key and generates them if required.
func LoadConfig() {
	cfg := config.Get()
	generateKeyFilesIfRequired()
	publicKey, err := ioutil.ReadFile(cfg.Auth.PublicKeyPath)

	if err != nil {
		logbuch.Fatal("Token public key file not found", logbuch.Fields{"err": err, "file": cfg.Auth.PublicKeyPath})
	}

	privateKey, err := ioutil.ReadFile(cfg.Auth.PrivateKeyPath)

	if err != nil {
		logbuch.Fatal("Token private key file not found", logbuch.Fields{"err": err, "file": cfg.Auth.PrivateKeyPath})
	}

	verify, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)

	if err != nil {
		logbuch.Fatal("Error parsing token public key", logbuch.Fields{"err": err})
	}

	sign, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)

	if err != nil {
		logbuch.Fatal("Error parsing token private key", logbuch.Fields{"err": err})
	}

	verifyKey = verify
	signKey = sign

	if cfg.Auth.AdminPassword == "" {
		logbuch.Fatal("Admin password must be set")
	}
}

func generateKeyFilesIfRequired() {
	_, err := os.Stat(config.Get().Auth.PublicKeyPath)

	if os.IsNotExist(err) {
		logbuch.Info("Private/public token key files not found, generating new ones")
		keyFileScript := "./scripts/gen_rsa_keys.cmd"

		if runtime.GOOS != "windows" {
			keyFileScript = "./scripts/gen_rsa_keys.sh"
		}

		_, err = os.Stat(keyFileScript)

		if err == nil {
			cmd := exec.Command(keyFileScript)

			if err := cmd.Run(); err != nil {
				logbuch.Fatal("Error generating key files through OpenSSL. Make sure the script gen_rsa_keys is in place", logbuch.Fields{"err": err})
			}

			logbuch.Info("Private/public token key files generated")
		}
	}
}
