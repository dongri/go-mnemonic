package mnemonic

import (
	"github.com/dongri/go-mnemonic/wordlists"
)

// Language ...
type Language string

// Language const
const (
	LanguageEnglish            Language = "english"
	LanguageFrench             Language = "french"
	LanguageItalian            Language = "italian"
	LanguageJapanese           Language = "japanese"
	LanguageKorean             Language = "korean"
	LanguageSpanish            Language = "spanish"
	LanguageChineseSimplified  Language = "chinese_simplified"
	LanguageChineseTraditional Language = "chinese_traditional"
)

// GetWordList ...
func GetWordList(language Language) ([]string, error) {
	var wordlist = []string{}
	switch language {
	case LanguageEnglish:
		wordlist = wordlists.WordlistEnglish
	case LanguageFrench:
		wordlist = wordlists.WordlistFrench
	case LanguageItalian:
		wordlist = wordlists.WordlistItalian
	case LanguageJapanese:
		wordlist = wordlists.WordlistJapanese
	case LanguageKorean:
		wordlist = wordlists.WordlistKorean
	case LanguageSpanish:
		wordlist = wordlists.WordlistSpanish
	case LanguageChineseSimplified:
		wordlist = wordlists.WordlistChineseSimplified
	case LanguageChineseTraditional:
		wordlist = wordlists.WordlistChineseTraditional
	default:
		wordlist = wordlists.WordlistEnglish
	}
	return wordlist, nil
}
