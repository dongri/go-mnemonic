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
    // さべつ　うけつけ　いほう　ぜんぶ　うやまう　はっきり　せっけん　にうけ　いだい　のっく　ときおり　ちそう

    seed := mnemonic.ToSeedHex(words, "password")
    fmt.Println(seed)
    // 9e745d0d134f30f9caa420d71be96578e246b6f1d9a03df2c1d72c6ef432b53edd31eeb318ba0f95cedd6d261dba7f20abe6f71aaa4514f2ab8c6192d17ce12c

    words, _ = mnemonic.GenerateMnemonic(128, mnemonic.LanguageEnglish)
    fmt.Println(words)
    // adult cloud wink ecology neglect route mom message table matter rapid crime

    words, _ = mnemonic.GenerateMnemonic(128, mnemonic.LanguageKorean)
    fmt.Println(words)
    // 소질 업종 총장 통화 이동 아픔 방식 일본 냄비 무용 그림 메일

    words, _ = mnemonic.GenerateMnemonic(128, mnemonic.LanguageChineseSimplified)
    fmt.Println(words)
    // 央 壮 楚 套 颗 烯 初 饼 挺 道 贸 竟
}
```
