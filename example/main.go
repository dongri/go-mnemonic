package main

import (
	"fmt"

	"github.com/dongri/go-mnemonic"
)

func main() {
	words, _ := mnemonic.GenerateMnemonic(128, mnemonic.LanguageJapanese)
	fmt.Println(words)

	seed := mnemonic.ToSeedHex(words, "")
	fmt.Println(seed)

	words, _ = mnemonic.GenerateMnemonic(128, mnemonic.LanguageEnglish)
	fmt.Println(words)

	words, _ = mnemonic.GenerateMnemonic(128, mnemonic.LanguageKorean)
	fmt.Println(words)

	words, _ = mnemonic.GenerateMnemonic(128, mnemonic.LanguageChineseSimplified)
	fmt.Println(words)
}
