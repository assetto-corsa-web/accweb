package helper

import (
	"bytes"
	"unicode/utf16"
	"unicode/utf8"

	"golang.org/x/text/encoding/charmap"
)

// NormalizeEncoding attempts to convert ambiguous byte sequences to proper UTF-8.
// Handles cases:
// 1. Already valid UTF-8 -> returned as-is
// 2. Escaped sequences like Faria\xe7\xe3o -> decode to bytes then to UTF-8 (cp1252/latin1)
// 3. Raw CP1252/Latin1 bytes -> converted to UTF-8
// 4. UTF-16LE/BE without BOM (heuristic: many zero bytes) -> decode
func NormalizeEncoding(src []byte) []byte {
	if len(src) == 0 {
		return src
	}

	// Fast path: valid UTF-8
	if utf8.Valid(src) {
		// However it might contain escaped sequences \xNN literally; try to unescape if present
		if bytes.Contains(src, []byte("\\x")) {
			if unesc, ok := tryUnescapeHexSequences(src); ok && utf8.Valid(unesc) {
				return unesc
			}
		}
		return src
	}

	// Heuristic: UTF-16 (LE/BE) without BOM (look for zero bytes pattern)
	if looksLikeUTF16(src) {
		if decoded, ok := tryDecodeUTF16WithoutBOM(src); ok && utf8.Valid(decoded) {
			return decoded
		}
	}

	// Try Windows-1252 (superset of Latin-1 with common accented chars)
	if dec, err := charmap.Windows1252.NewDecoder().Bytes(src); err == nil && utf8.Valid(dec) {
		return dec
	}

	// Fallback: return original (avoid data loss)
	return src
}

// tryUnescapeHexSequences converts sequences like ...Faria\xe7\xe3o... into proper bytes then tries cp1252 decode
func tryUnescapeHexSequences(src []byte) ([]byte, bool) {
	var out bytes.Buffer
	for i := 0; i < len(src); i++ {
		if i+3 < len(src) && src[i] == '\\' && src[i+1] == 'x' && isHex(src[i+2]) && isHex(src[i+3]) {
			b := fromHex(src[i+2])<<4 | fromHex(src[i+3])
			out.WriteByte(b)
			i += 3
			continue
		}
		out.WriteByte(src[i])
	}

	candidate := out.Bytes()
	// Attempt decode as Windows-1252
	if dec, err := charmap.Windows1252.NewDecoder().Bytes(candidate); err == nil && utf8.Valid(dec) {
		return dec, true
	}
	return candidate, utf8.Valid(candidate)
}

func isHex(b byte) bool {
	return (b >= '0' && b <= '9') || (b >= 'a' && b <= 'f') || (b >= 'A' && b <= 'F')
}

func fromHex(b byte) byte {
	if b >= '0' && b <= '9' {
		return b - '0'
	}
	if b >= 'a' && b <= 'f' {
		return 10 + b - 'a'
	}
	return 10 + b - 'A'
}

// looksLikeUTF16 detects high proportion of zero bytes at even or odd positions
func looksLikeUTF16(b []byte) bool {
	if len(b) < 4 {
		return false
	}
	zerosEven := 0
	zerosOdd := 0
	for i := 0; i < len(b); i++ {
		if b[i] == 0 {
			if i%2 == 0 {
				zerosEven++
			} else {
				zerosOdd++
			}
		}
	}
	propEven := float64(zerosEven) / float64(len(b)/2)
	propOdd := float64(zerosOdd) / float64(len(b)/2)
	return propEven > 0.3 || propOdd > 0.3
}

// tryDecodeUTF16WithoutBOM tries both LE and BE
func tryDecodeUTF16WithoutBOM(b []byte) ([]byte, bool) {
	if len(b)%2 != 0 {
		return nil, false
	}
	// LE first
	le := make([]uint16, 0, len(b)/2)
	for i := 0; i < len(b); i += 2 {
		le = append(le, uint16(b[i])|uint16(b[i+1])<<8)
	}
	if r := utf16.Decode(le); validateRunes(r) {
		utf8b := []byte(string(r))
		if utf8.Valid(utf8b) {
			return utf8b, true
		}
	}
	// BE
	be := make([]uint16, 0, len(b)/2)
	for i := 0; i < len(b); i += 2 {
		be = append(be, uint16(b[i+1])|uint16(b[i])<<8)
	}
	if r := utf16.Decode(be); validateRunes(r) {
		utf8b := []byte(string(r))
		if utf8.Valid(utf8b) {
			return utf8b, true
		}
	}
	return nil, false
}

func validateRunes(r []rune) bool {
	if len(r) == 0 {
		return false
	}
	invalid := 0
	for _, ru := range r {
		if ru == 0xFFFD {
			invalid++
		}
	}
	return float64(invalid) < 0.1*float64(len(r))
}
