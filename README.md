# go-mnemonic

Go implementation of Bitcoin BIP39: Mnemonic code for generating deterministic keys

# Examples

```go
import (
	"fmt"

	"github.com/dongri/go-mnemonic"
)

func main() {
    words, _ := mnemonic.GenerateMnemonic(128, mnemonic.LanguageJapanese)
    fmt.Println(words)
    // あいさつ　ひつじゅひん　うろこ　すふれ　てきとう　こさめ　くどく　こくさい　ようす　くげん　むさぼる　ひさん

    seed := mnemonic.ToSeedHex(words, "password")
    fmt.Println(seed)
    // bdd000956333c5abd39be64128159add10a200dee6e68d207c8bf31275cbac8cebc2c55704080c8973d862d6424c5d9298b9da3d9d7beb4f6b333ff2a6ab729c

    words, _ = mnemonic.GenerateMnemonic(128, mnemonic.LanguageEnglish)
    fmt.Println(words)

    words, _ = mnemonic.GenerateMnemonic(128, mnemonic.LanguageKorean)
    fmt.Println(words)

    words, _ = mnemonic.GenerateMnemonic(128, mnemonic.LanguageChineseSimplified)
    fmt.Println(words)
}
```
