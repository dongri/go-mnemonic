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
}
