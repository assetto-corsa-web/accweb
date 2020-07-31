package auth

import (
	"crypto/sha256"
	"encoding/base64"
	"github.com/assetto-corsa-web/accweb/config"
	"github.com/emvi/logbuch"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

const (
	userListFile = "user.yml"
	RoleAdmin    = "admin"
	RoleMod      = "moderator"
	RoleReadOnly = "read-only"
)

var (
	user UserList
	m    sync.Mutex
)

// User is used for authentication and role management.
type User struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"` // SHA256 + salt
	Role     string `yaml:"role"`     // admin, moderator, read-only
}

// UserList is a list of users that can login to accweb.
type UserList struct {
	User []User `yaml:"user"`
}

// LoadUser loads the list of users for authentication.
func LoadUser() {
	logbuch.Info("Loading user file")
	data, err := ioutil.ReadFile(userListFile)

	if os.IsNotExist(err) {
		logbuch.Info("User list file not found, you won't be able to login. Create the user.yml in the root directory and add users.")
	} else if err != nil {
		logbuch.Fatal("Error loading user list file", logbuch.Fields{"err": err})
	}

	m.Lock()
	defer m.Unlock()

	if err := yaml.Unmarshal(data, &user); err != nil {
		logbuch.Fatal("Error parsing user list file", logbuch.Fields{"err": err})
	}
}

// GetUser returns the user for given username and password (clear text) or nil if not found.
func GetUser(username, password string) *User {
	m.Lock()
	defer m.Unlock()
	username = strings.ToLower(username)
	password = sha256base64(password)
	logbuch.Debug("Getting user", logbuch.Fields{"username": username, "password": password})

	for _, u := range user.User {
		if strings.ToLower(u.Username) == username && u.Password == password {
			return &u
		}
	}

	return nil
}

// SetUser adds/updates the user for given username.
func SetUser(username, password, role string) {
	m.Lock()
	defer m.Unlock()
	username = strings.ToLower(username)
	password = sha256base64(password)

	for i, u := range user.User {
		if strings.ToLower(u.Username) == username {
			setUser(i, username, password, role)
			return
		}
	}

	setUser(-1, username, password, role)
}

func setUser(index int, username, password, role string) {
	if index > -1 {
		user.User[index].Password = password
		user.User[index].Role = role
	} else {
		user.User = append(user.User, User{
			username,
			password,
			role,
		})
	}

	saveUser()
}

// RemoveUser removes a user by its username.
func RemoveUser(username string) {
	m.Lock()
	defer m.Unlock()
	username = strings.ToLower(username)

	for i, u := range user.User {
		if strings.ToLower(u.Username) == username {
			user.User = append(user.User[:i], user.User[i+1:]...)
			saveUser()
			return
		}
	}
}

func saveUser() {
	logbuch.Info("Saving user file")
	out, err := yaml.Marshal(&user)

	if err != nil {
		logbuch.Fatal("Error marshalling user list", logbuch.Fields{"err": err})
	}

	if err := ioutil.WriteFile(userListFile, out, 0755); err != nil {
		logbuch.Fatal("Error saving user list", logbuch.Fields{"err": err})
	}
}

func sha256base64(password string) string {
	h := sha256.New()
	h.Write([]byte(password + config.Get().Auth.Salt))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
