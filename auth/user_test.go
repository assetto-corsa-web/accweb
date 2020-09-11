package auth

import (
	"os"
	"testing"
)

func TestUserList(t *testing.T) {
	if err := os.MkdirAll("data", 0755); err != nil {
		t.Fatal(err)
	}

	if err := os.RemoveAll(userListFile); err != nil {
		t.Fatal(err)
	}

	list := NewUserList()
	list.Set("admin", "pwd1", RoleAdmin)
	list.Set("mod", "pwd2", RoleMod)
	list.Set("ro", "pwd3", RoleReadOnly)
	list.load()

	if len(list.user) != 3 {
		t.Fatalf("User list must have three entries, but was: %v", len(user.user))
	}

	if list.user[0].Username != "admin" || list.user[0].Role != RoleAdmin ||
		list.user[1].Username != "mod" || list.user[1].Role != RoleMod ||
		list.user[2].Username != "ro" || list.user[2].Role != RoleReadOnly {
		t.Fatal("User entries not as expected")
	}

	list.Remove("mod")

	if len(list.user) != 2 {
		t.Fatalf("Entry must have been removed, but was: %v", len(user.user))
	}

	if list.Get("admin", "pwd1") == nil {
		t.Fatal("admin must have been found for correct password")
	}

	if list.Get("admin", "incorrect") != nil {
		t.Fatal("admin must not have been found for incorrect password")
	}
}
