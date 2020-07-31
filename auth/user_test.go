package auth

import (
	"os"
	"testing"
)

func TestUserList(t *testing.T) {
	if err := os.RemoveAll(userListFile); err != nil {
		t.Fatal(err)
	}

	LoadUser()
	SetUser("admin", "pwd1", RoleAdmin)
	SetUser("mod", "pwd2", RoleMod)
	SetUser("ro", "pwd3", RoleReadOnly)
	LoadUser()

	if len(user.User) != 3 {
		t.Fatalf("User list must have three entries, but was: %v", len(user.User))
	}

	if user.User[0].Username != "admin" || user.User[0].Role != RoleAdmin ||
		user.User[1].Username != "mod" || user.User[1].Role != RoleMod ||
		user.User[2].Username != "ro" || user.User[2].Role != RoleReadOnly {
		t.Fatal("User entries not as expected")
	}

	RemoveUser("mod")

	if len(user.User) != 2 {
		t.Fatalf("Entry must have been removed, but was: %v", len(user.User))
	}

	if GetUser("admin", "pwd1") == nil {
		t.Fatal("admin must have been found for correct password")
	}

	if GetUser("admin", "incorrect") != nil {
		t.Fatal("admin must not have been found for incorrect password")
	}
}
