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

	seed := mnemonic.ToSeedHex(words, "password")
	fmt.Println(seed)
}
```
