package mnemonic

import (
	"bufio"
	"fmt"

	_ "github.com/dongri/go-mnemonic/statik" // statik dir
	"github.com/rakyll/statik/fs"
)

// Language ...
type Language string

// Language const
const (
	LanguageEnglish            Language = "english"
	LanguageFrench             Language = "frech"
	LanguageItalian            Language = "italian"
	LanguageJapanese           Language = "japanese"
	LanguageKorean             Language = "korean"
	LanguageSpanish            Language = "spanish"
	LanguageChineseSimplified  Language = "chinese_simplified"
	LanguageChineseTraditional Language = "chinese_traditional"
)

// GetWordList ...
func GetWordList(language Language) ([]string, error) {
	var wordlist []string
	filename := fmt.Sprintf("/%s.txt", language)
	FS, err := fs.New()
	if err != nil {
		return wordlist, err
	}
	file, err := FS.Open(filename)
	if err != nil {
		return wordlist, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wordlist = append(wordlist, scanner.Text())
	}
	return wordlist, nil
}
