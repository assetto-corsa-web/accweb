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
	userListFile        = "data/user.yml"
	RoleAdmin           = "admin"
	RoleMod             = "moderator"
	RoleReadOnly        = "read-only"
	defaultUser         = "admin"
	defaultUserPassword = "admin"
)

// User is used for authentication and role management.
type User struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"` // SHA256 + salt
	Role     string `yaml:"role"`     // admin, moderator, read-only
}

// UserList is a list of users that can login to accweb.
type UserList struct {
	user []User
	m    sync.Mutex
}

// NewUserList creates and loads the list of users for authentication.
func NewUserList() *UserList {
	list := &UserList{
		user: make([]User, 0),
	}
	list.load()
	return list
}

// Get returns the user for given username and password (clear text) or nil if not found.
func (list *UserList) Get(username, password string) *User {
	list.m.Lock()
	defer list.m.Unlock()
	username = strings.ToLower(username)
	password = sha256base64(password)
	logbuch.Debug("Getting user", logbuch.Fields{"username": username, "password": password})

	for _, u := range list.user {
		if strings.ToLower(u.Username) == username && u.Password == password {
			return &u
		}
	}

	return nil
}

// GetAll returns all users in the list.
func (list *UserList) GetAll() []User {
	list.m.Lock()
	defer list.m.Unlock()
	user := make([]User, len(list.user))
	copy(user, list.user)
	return user
}

// Set adds/updates the user for given username.
func (list *UserList) Set(username, password, role string) {
	list.m.Lock()
	defer list.m.Unlock()
	username = strings.ToLower(username)
	password = sha256base64(password)
	index := -1

	for i, u := range list.user {
		if strings.ToLower(u.Username) == username {
			index = i
			break
		}
	}

	list.setUser(index, username, password, role)
}

func (list *UserList) setUser(index int, username, password, role string) {
	if index > -1 {
		list.user[index].Password = password
		list.user[index].Role = role
	} else {
		list.user = append(list.user, User{
			username,
			password,
			role,
		})
	}

	list.save()
}

// Remove removes a user by its username.
func (list *UserList) Remove(username string) {
	list.m.Lock()
	defer list.m.Unlock()
	username = strings.ToLower(username)

	for i, u := range list.user {
		if strings.ToLower(u.Username) == username {
			list.user = append(list.user[:i], list.user[i+1:]...)
			list.save()
			return
		}
	}
}

func (list *UserList) load() {
	logbuch.Info("Loading user file...")
	data, err := ioutil.ReadFile(userListFile)

	if os.IsNotExist(err) {
		logbuch.Info("User list file not found, creating default")
		list.Set(defaultUser, defaultUserPassword, RoleAdmin)
		logbuch.Info("Default user file created")
	} else if err != nil {
		logbuch.Fatal("Error loading user list file", logbuch.Fields{"err": err})
	}

	list.m.Lock()
	defer list.m.Unlock()

	if err := yaml.Unmarshal(data, &list.user); err != nil {
		logbuch.Fatal("Error parsing user list file", logbuch.Fields{"err": err})
	}

	logbuch.Info("Users loaded", logbuch.Fields{"users": len(list.user)})
}

func (list *UserList) save() {
	logbuch.Info("Saving user file")
	out, err := yaml.Marshal(list.user)

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
