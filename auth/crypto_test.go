package auth

import "testing"

func TestSHA256Base64(t *testing.T) {
	if len(sha256Base64("string")) != 44 {
		t.Fatal("SHA256 string must be 44 characters long")
	}
}

func TestHashPassword(t *testing.T) {
	password := "asdf1234"
	out := hashPassword(password)

	if out == "" {
		t.Fatalf("Password must have been hashed: %v", out)
	}

	if !comparePassword(password, out) {
		t.Fatal("Password must match hash")
	}
}

func TestComparePassword(t *testing.T) {
	if !comparePassword("test", "$2a$10$t84eU4c/J1gGfsDeL..vv.BoHpK0Go/8EpG4.hoZuhx7ulVNvV1iC") {
		t.Fatal("Passwords must match")
	}

	if comparePassword("foo", "$2a$10$t84eU4c/J1gGfsDeL..vv.BoHpK0Go/8EpG4.hoZuhx7ulVNvV1iC") {
		t.Fatal("Passwords must not match")
	}
}
