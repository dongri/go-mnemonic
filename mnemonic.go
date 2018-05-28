package mnemonic

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/text/unicode/norm"
)

// const ...
const (
	InvalidStrength = "Invalid strength"
	InvalidEntropy  = "Invalid entropy"
)

var (
	chunksRe = regexp.MustCompile(".{1,11}")
)

// GenerateMnemonic ...
func GenerateMnemonic(strength uint16, language Language) (string, error) {
	if strength%32 != 0 {
		return "", errors.New(InvalidStrength)
	}
	entropy := randomBytes(strength / 8)
	wordlist, err := GetWordList(language)
	if err != nil {
		return "", err
	}
	words, err := entropyToMnemonic(entropy, wordlist)
	if err != nil {
		return "", err
	}
	sep := " "
	if language == LanguageJapanese {
		sep = "\u3000"
	}
	return strings.Join(words, sep), nil
}

// ToSeedHex ...
func ToSeedHex(mnemonic, password string) string {
	seed := pbkdf2.Key([]byte(mnemonic), []byte("mnemonic"+password), 2048, 64, sha512.New)
	return hex.EncodeToString(seed)
}

func randomBytes(length uint16) []byte {
	token := make([]byte, length)
	rand.Read(token)
	return token
}

func entropyToMnemonic(entropy []byte, wordlist []string) ([]string, error) {
	length := len(entropy)
	if length < 16 || length > 32 || length%4 != 0 {
		return nil, errors.New(InvalidEntropy)
	}
	entropyBits := bytesToBinary(entropy)
	checksumBits := deriveChecksumBits(entropy)
	bits := entropyBits + checksumBits
	chunks := chunksRe.FindAllString(bits, -1)
	words := []string{}
	for _, binary := range chunks {
		i, err := binaryToByte(binary)
		if err != nil {
			return words, err
		}
		words = append(words, norm.NFKC.String(wordlist[i]))
	}
	return words, nil
}

func deriveChecksumBits(bytes []byte) string {
	ENT := len(bytes) * 8
	CS := ENT / 32
	s := sha256.New()
	s.Write(bytes)
	hash := s.Sum(nil)
	return bytesToBinary(hash)[:CS]
}

func bytesToBinary(bytes []byte) string {
	bits := ""
	for _, b := range bytes {
		bits += fmt.Sprintf("%08b", b)
	}
	return bits
}

func binaryToByte(bin string) (int64, error) {
	return strconv.ParseInt(bin, 2, 16)
}
