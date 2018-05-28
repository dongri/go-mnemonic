package mnemonic

import (
	"encoding/hex"
	"strings"
	"testing"

	"golang.org/x/text/unicode/norm"
)

type Vector struct {
	entropy  string
	mnemonic string
	seed     string
}

func TestMnemonicEnglish(t *testing.T) {
	for _, vector := range englishVectors() {
		entropy, err := hex.DecodeString(vector.entropy)
		if err != nil {
			t.Error(err)
		}
		wordList, _ := GetWordList(LanguageEnglish)
		mnemonic, err := entropyToMnemonic(entropy, wordList)
		if err != nil {
			t.Error(err)
		}
		actual := strings.Join(mnemonic, " ")
		if vector.mnemonic != actual {
			t.Error(vector.mnemonic + "!=" + actual)
		}
		seed := ToSeedHex(actual, "TREZOR")
		if vector.seed != seed {
			t.Error(vector.seed + "!=" + seed)
		}
	}
}
func TestMnemonicJapanese(t *testing.T) {
	for _, vector := range japaneseVectors() {
		entropy, err := hex.DecodeString(vector.entropy)
		if err != nil {
			t.Error(err)
		}
		wordList, _ := GetWordList(LanguageJapanese)
		mnemonic, err := entropyToMnemonic(entropy, wordList)
		if err != nil {
			t.Error(err)
		}

		mnemonicArray := []string{}
		for _, word := range mnemonic {
			mnemonicArray = append(mnemonicArray, norm.NFKC.String(word))
		}
		actual := strings.Join(mnemonicArray, "\u3000")
		if vector.mnemonic != actual {
			t.Error(vector.mnemonic + "!=" + actual)
		}
		seed := ToSeedHex(vector.mnemonic, "TREZOR")
		if vector.seed != seed {
			t.Error(vector.seed + "!=" + seed)
		}
	}
}

func englishVectors() []Vector {
	return []Vector{
		Vector{
			entropy:  "00000000000000000000000000000000",
			mnemonic: "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about",
			seed:     "c55257c360c07c72029aebc1b53c05ed0362ada38ead3e3e9efa3708e53495531f09a6987599d18264c1e1c92f2cf141630c7a3c4ab7c81b2f001698e7463b04",
		},
		Vector{
			entropy:  "7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f",
			mnemonic: "legal winner thank year wave sausage worth useful legal winner thank yellow",
			seed:     "2e8905819b8723fe2c1d161860e5ee1830318dbf49a83bd451cfb8440c28bd6fa457fe1296106559a3c80937a1c1069be3a3a5bd381ee6260e8d9739fce1f607",
		},
		Vector{
			entropy:  "80808080808080808080808080808080",
			mnemonic: "letter advice cage absurd amount doctor acoustic avoid letter advice cage above",
			seed:     "d71de856f81a8acc65e6fc851a38d4d7ec216fd0796d0a6827a3ad6ed5511a30fa280f12eb2e47ed2ac03b5c462a0358d18d69fe4f985ec81778c1b370b652a8",
		},
		Vector{
			entropy:  "ffffffffffffffffffffffffffffffff",
			mnemonic: "zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo wrong",
			seed:     "ac27495480225222079d7be181583751e86f571027b0497b5b5d11218e0a8a13332572917f0f8e5a589620c6f15b11c61dee327651a14c34e18231052e48c069",
		},
		Vector{
			entropy:  "000000000000000000000000000000000000000000000000",
			mnemonic: "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon agent",
			seed:     "035895f2f481b1b0f01fcf8c289c794660b289981a78f8106447707fdd9666ca06da5a9a565181599b79f53b844d8a71dd9f439c52a3d7b3e8a79c906ac845fa",
		},
		Vector{
			entropy:  "7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f",
			mnemonic: "legal winner thank year wave sausage worth useful legal winner thank year wave sausage worth useful legal will",
			seed:     "f2b94508732bcbacbcc020faefecfc89feafa6649a5491b8c952cede496c214a0c7b3c392d168748f2d4a612bada0753b52a1c7ac53c1e93abd5c6320b9e95dd",
		},
		Vector{
			entropy:  "808080808080808080808080808080808080808080808080",
			mnemonic: "letter advice cage absurd amount doctor acoustic avoid letter advice cage absurd amount doctor acoustic avoid letter always",
			seed:     "107d7c02a5aa6f38c58083ff74f04c607c2d2c0ecc55501dadd72d025b751bc27fe913ffb796f841c49b1d33b610cf0e91d3aa239027f5e99fe4ce9e5088cd65",
		},
		Vector{
			entropy:  "ffffffffffffffffffffffffffffffffffffffffffffffff",
			mnemonic: "zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo when",
			seed:     "0cd6e5d827bb62eb8fc1e262254223817fd068a74b5b449cc2f667c3f1f985a76379b43348d952e2265b4cd129090758b3e3c2c49103b5051aac2eaeb890a528",
		},
		Vector{
			entropy:  "0000000000000000000000000000000000000000000000000000000000000000",
			mnemonic: "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon art",
			seed:     "bda85446c68413707090a52022edd26a1c9462295029f2e60cd7c4f2bbd3097170af7a4d73245cafa9c3cca8d561a7c3de6f5d4a10be8ed2a5e608d68f92fcc8",
		},
		Vector{
			entropy:  "7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f",
			mnemonic: "legal winner thank year wave sausage worth useful legal winner thank year wave sausage worth useful legal winner thank year wave sausage worth title",
			seed:     "bc09fca1804f7e69da93c2f2028eb238c227f2e9dda30cd63699232578480a4021b146ad717fbb7e451ce9eb835f43620bf5c514db0f8add49f5d121449d3e87",
		},
		Vector{
			entropy:  "8080808080808080808080808080808080808080808080808080808080808080",
			mnemonic: "letter advice cage absurd amount doctor acoustic avoid letter advice cage absurd amount doctor acoustic avoid letter advice cage absurd amount doctor acoustic bless",
			seed:     "c0c519bd0e91a2ed54357d9d1ebef6f5af218a153624cf4f2da911a0ed8f7a09e2ef61af0aca007096df430022f7a2b6fb91661a9589097069720d015e4e982f",
		},
		Vector{
			entropy:  "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
			mnemonic: "zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo vote",
			seed:     "dd48c104698c30cfe2b6142103248622fb7bb0ff692eebb00089b32d22484e1613912f0a5b694407be899ffd31ed3992c456cdf60f5d4564b8ba3f05a69890ad",
		},
		Vector{
			entropy:  "77c2b00716cec7213839159e404db50d",
			mnemonic: "jelly better achieve collect unaware mountain thought cargo oxygen act hood bridge",
			seed:     "b5b6d0127db1a9d2226af0c3346031d77af31e918dba64287a1b44b8ebf63cdd52676f672a290aae502472cf2d602c051f3e6f18055e84e4c43897fc4e51a6ff",
		},
		Vector{
			entropy:  "b63a9c59a6e641f288ebc103017f1da9f8290b3da6bdef7b",
			mnemonic: "renew stay biology evidence goat welcome casual join adapt armor shuffle fault little machine walk stumble urge swap",
			seed:     "9248d83e06f4cd98debf5b6f010542760df925ce46cf38a1bdb4e4de7d21f5c39366941c69e1bdbf2966e0f6e6dbece898a0e2f0a4c2b3e640953dfe8b7bbdc5",
		},
		Vector{
			entropy:  "3e141609b97933b66a060dcddc71fad1d91677db872031e85f4c015c5e7e8982",
			mnemonic: "dignity pass list indicate nasty swamp pool script soccer toe leaf photo multiply desk host tomato cradle drill spread actor shine dismiss champion exotic",
			seed:     "ff7f3184df8696d8bef94b6c03114dbee0ef89ff938712301d27ed8336ca89ef9635da20af07d4175f2bf5f3de130f39c9d9e8dd0472489c19b1a020a940da67",
		},
		Vector{
			entropy:  "0460ef47585604c5660618db2e6a7e7f",
			mnemonic: "afford alter spike radar gate glance object seek swamp infant panel yellow",
			seed:     "65f93a9f36b6c85cbe634ffc1f99f2b82cbb10b31edc7f087b4f6cb9e976e9faf76ff41f8f27c99afdf38f7a303ba1136ee48a4c1e7fcd3dba7aa876113a36e4",
		},
		Vector{
			entropy:  "72f60ebac5dd8add8d2a25a797102c3ce21bc029c200076f",
			mnemonic: "indicate race push merry suffer human cruise dwarf pole review arch keep canvas theme poem divorce alter left",
			seed:     "3bbf9daa0dfad8229786ace5ddb4e00fa98a044ae4c4975ffd5e094dba9e0bb289349dbe2091761f30f382d4e35c4a670ee8ab50758d2c55881be69e327117ba",
		},
		Vector{
			entropy:  "2c85efc7f24ee4573d2b81a6ec66cee209b2dcbd09d8eddc51e0215b0b68e416",
			mnemonic: "clutch control vehicle tonight unusual clog visa ice plunge glimpse recipe series open hour vintage deposit universe tip job dress radar refuse motion taste",
			seed:     "fe908f96f46668b2d5b37d82f558c77ed0d69dd0e7e043a5b0511c48c2f1064694a956f86360c93dd04052a8899497ce9e985ebe0c8c52b955e6ae86d4ff4449",
		},
		Vector{
			entropy:  "eaebabb2383351fd31d703840b32e9e2",
			mnemonic: "turtle front uncle idea crush write shrug there lottery flower risk shell",
			seed:     "bdfb76a0759f301b0b899a1e3985227e53b3f51e67e3f2a65363caedf3e32fde42a66c404f18d7b05818c95ef3ca1e5146646856c461c073169467511680876c",
		},
		Vector{
			entropy:  "7ac45cfe7722ee6c7ba84fbc2d5bd61b45cb2fe5eb65aa78",
			mnemonic: "kiss carry display unusual confirm curtain upgrade antique rotate hello void custom frequent obey nut hole price segment",
			seed:     "ed56ff6c833c07982eb7119a8f48fd363c4a9b1601cd2de736b01045c5eb8ab4f57b079403485d1c4924f0790dc10a971763337cb9f9c62226f64fff26397c79",
		},
		Vector{
			entropy:  "4fa1a8bc3e6d80ee1316050e862c1812031493212b7ec3f3bb1b08f168cabeef",
			mnemonic: "exile ask congress lamp submit jacket era scheme attend cousin alcohol catch course end lucky hurt sentence oven short ball bird grab wing top",
			seed:     "095ee6f817b4c2cb30a5a797360a81a40ab0f9a4e25ecd672a3f58a0b5ba0687c096a6b14d2c0deb3bdefce4f61d01ae07417d502429352e27695163f7447a8c",
		},
		Vector{
			entropy:  "18ab19a9f54a9274f03e5209a2ac8a91",
			mnemonic: "board flee heavy tunnel powder denial science ski answer betray cargo cat",
			seed:     "6eff1bb21562918509c73cb990260db07c0ce34ff0e3cc4a8cb3276129fbcb300bddfe005831350efd633909f476c45c88253276d9fd0df6ef48609e8bb7dca8",
		},
		Vector{
			entropy:  "18a2e1d81b8ecfb2a333adcb0c17a5b9eb76cc5d05db91a4",
			mnemonic: "board blade invite damage undo sun mimic interest slam gaze truly inherit resist great inject rocket museum chief",
			seed:     "f84521c777a13b61564234bf8f8b62b3afce27fc4062b51bb5e62bdfecb23864ee6ecf07c1d5a97c0834307c5c852d8ceb88e7c97923c0a3b496bedd4e5f88a9",
		},
		Vector{
			entropy:  "15da872c95a13dd738fbf50e427583ad61f18fd99f628c417a61cf8343c90419",
			mnemonic: "beyond stage sleep clip because twist token leaf atom beauty genius food business side grid unable middle armed observe pair crouch tonight away coconut",
			seed:     "b15509eaa2d09d3efd3e006ef42151b30367dc6e3aa5e44caba3fe4d3e352e65101fbdb86a96776b91946ff06f8eac594dc6ee1d3e82a42dfe1b40fef6bcc3fd",
		},
	}
}

func japaneseVectors() []Vector {
	return []Vector{
		Vector{
			entropy:  "00000000000000000000000000000000",
			mnemonic: "あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あおぞら",
			seed:     "5a6c23b5abdd5c3e1f7d77ad25ecd715647bdafb44dab324c730a76a45d7421daccee1a4ff0739715a2c56a8a9f1e527a5e3496224d91293bfcd9b5393bfff83",
		},
		Vector{
			entropy:  "7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f",
			mnemonic: "そつう　れきだい　ほんやく　わかす　りくつ　ばいか　ろせん　やちん　そつう　れきだい　ほんやく　わかめ",
			seed:     "9d269b22155b3c915b09abfefd4e1104573c528f6977cde89c6a68152c3c714dc6c7e0e62f221c322f3f76e4d0bcca66c06e3d2f6a8d70d612c87dd6dee63976",
		},

		Vector{
			entropy:  "80808080808080808080808080808080",
			mnemonic: "そとづら　あまど　おおう　あこがれる　いくぶん　けいけん　あたえる　いよく　そとづら　あまど　おおう　あかちゃん",
			seed:     "17914bd3fe4b9e1224c968ec6b967fc6144a5795adbb2636a17f77da9b6b118200ad788672fd06096ca62683940523f5178f6ce3845c967cbd4ad2b3643cc660",
		},

		Vector{
			entropy:  "ffffffffffffffffffffffffffffffff",
			mnemonic: "われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　ろんぶん",
			seed:     "4bd21b75de4f262b0771a97d6fc877ee19329236ced6e974c4c81a094a5f896758033f7eae270216d727539eee3bc9ba5cad21132a1c6e41a50820e0ac928e83",
		},

		Vector{
			entropy:  "000000000000000000000000000000000000000000000000",
			mnemonic: "あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あらいぐま",
			seed:     "a59401a14bb821cce86ec32add8f273a3e07e9c8b1ed430d5d1a06dbf3c083ff2ffb4bb26a384b8faecb58f6cb4c07cfbf2c91108385f6773f2fefd1581926b5",
		},

		Vector{
			entropy:  "7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f",
			mnemonic: "そつう　れきだい　ほんやく　わかす　りくつ　ばいか　ろせん　やちん　そつう　れきだい　ほんやく　わかす　りくつ　ばいか　ろせん　やちん　そつう　れいぎ",
			seed:     "809861f80877e3adc842b0204e401d5aeac1d16d24072f387107f9cf95b639d0a76141ab25d3dc90752472787307a7d8b1a534bea237c2bb348faac973e17488",
		},

		Vector{
			entropy:  "808080808080808080808080808080808080808080808080",
			mnemonic: "そとづら　あまど　おおう　あこがれる　いくぶん　けいけん　あたえる　いよく　そとづら　あまど　おおう　あこがれる　いくぶん　けいけん　あたえる　いよく　そとづら　いきなり",
			seed:     "01187da93480d0369fff3fc5331284ad6a60cd3ce1f60dbec60899191afa2a2b807cd030038a93ddaf14d4f75d6de4a0e049ee58c92197eb9ca995770b558486",
		},

		Vector{
			entropy:  "ffffffffffffffffffffffffffffffffffffffffffffffff",
			mnemonic: "われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　りんご",
			seed:     "a1385ef66f20a905bbfc70f8be6ecfec341ff76d208e89e1a400ccea34313c99e93f4fba9c6f0729397b9002972af93179dc9dd8af7704fa3d28e656248274dc",
		},

		Vector{
			entropy:  "0000000000000000000000000000000000000000000000000000000000000000",
			mnemonic: "あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　いってい",
			seed:     "c91afc204a8b098524c5e2134bf4955b9a9ddd5d4bb78c2184bb4378a306e851b60f3e4032fc910ecb48acfb9e441dd3ceaaab9e14700b11396b94e27e8ac2da",
		},

		Vector{
			entropy:  "7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f",
			mnemonic: "そつう　れきだい　ほんやく　わかす　りくつ　ばいか　ろせん　やちん　そつう　れきだい　ほんやく　わかす　りくつ　ばいか　ろせん　やちん　そつう　れきだい　ほんやく　わかす　りくつ　ばいか　ろせん　まんきつ",
			seed:     "79aff5bc7868b9054f6c35bb3fa286c72a6931d5999c6c45a029ad31da550b71c8db72e594875e1d61788371b31a03b70fe1d9484840d403e56a1a2783bf9d7e",
		},

		Vector{
			entropy:  "8080808080808080808080808080808080808080808080808080808080808080",
			mnemonic: "そとづら　あまど　おおう　あこがれる　いくぶん　けいけん　あたえる　いよく　そとづら　あまど　おおう　あこがれる　いくぶん　けいけん　あたえる　いよく　そとづら　あまど　おおう　あこがれる　いくぶん　けいけん　あたえる　うめる",
			seed:     "0f46c02350b3f1227c3566dea2ff0f2caf716495a95725b320a31a3058d5d62596fdb816be75909d2c5f7094beb171dc504ea8ea60f5e2e40bd8aa0d9339aab0",
		},

		Vector{
			entropy:  "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
			mnemonic: "われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　らいう",
			seed:     "a0705c2feebefb61509dcc49c57586c35379c1981c688fc1d452da44443d9a651a374f1ad2ee3d7847b50655cf9241d7e607be436c0df7c8bac42f2a82985a79",
		},

		Vector{
			entropy:  "77c2b00716cec7213839159e404db50d",
			mnemonic: "せまい　うちがわ　あずき　かろう　めずらしい　だんち　ますく　おさめる　ていぼう　あたる　すあな　えしゃく",
			seed:     "b7f5478674839a3487a271014f066059490161a381ec57e9a00de0a3c7311ab51f20b53989c7bcbc923f956b5a16556bc6a4c143265e280769f12792d0e0913e",
		},

		Vector{
			entropy:  "b63a9c59a6e641f288ebc103017f1da9f8290b3da6bdef7b",
			mnemonic: "ぬすむ　ふっかつ　うどん　こうりつ　しつじ　りょうり　おたがい　せもたれ　あつめる　いちりゅう　はんしゃ　ごますり　そんけい　たいちょう　らしんばん　ぶんせき　やすみ　ほいく",
			seed:     "a5fe510d0485f7d74dec53fbc1aeb7bf3d527075dcc5ef657e0b3a8ff613554228099faa1cc9332f9a1dde264cefa6493f70ca3828c514781e78dd7c5e39877d",
		},

		Vector{
			entropy:  "3e141609b97933b66a060dcddc71fad1d91677db872031e85f4c015c5e7e8982",
			mnemonic: "くのう　てぬぐい　そんかい　すろっと　ちきゅう　ほあん　とさか　はくしゅ　ひびく　みえる　そざい　てんすう　たんぴん　くしょう　すいようび　みけん　きさらぎ　げざん　ふくざつ　あつかう　はやい　くろう　おやゆび　こすう",
			seed:     "3ca539f28db49e01d56b8dca1b513131dcd57833e961caabad88b7bbf2347ce5ece844c025bc88bd7a90fe4069a5ce2115f5571da9021af64e782539267fc687",
		},

		Vector{
			entropy:  "0460ef47585604c5660618db2e6a7e7f",
			mnemonic: "あみもの　いきおい　ふいうち　にげる　ざんしょ　じかん　ついか　はたん　ほあん　すんぽう　てちがい　わかめ",
			seed:     "1bd33e347a219ff2ff2dbacc0c6149a97d09e20f7dd4951552e1516eb865710387dc011c22b256270661094ff9bfb080b939eb6dd1cb8705afabe0f38cf3b74d",
		},

		Vector{
			entropy:  "72f60ebac5dd8add8d2a25a797102c3ce21bc029c200076f",
			mnemonic: "すろっと　にくしみ　なやむ　たとえる　へいこう　すくう　きない　けってい　とくべつ　ねっしん　いたみ　せんせい　おくりがな　まかい　とくい　けあな　いきおい　そそぐ",
			seed:     "37a76adf17a8330e495ea6e8b41cbb590ae7672a48bbcae709483b4a0b1b5104cacc5c5df6595a9de22c0116a33138233d15ede90c4fc7ba7cb97488d168c137",
		},

		Vector{
			entropy:  "2c85efc7f24ee4573d2b81a6ec66cee209b2dcbd09d8eddc51e0215b0b68e416",
			mnemonic: "かほご　きうい　ゆたか　みすえる　もらう　がっこう　よそう　ずっと　ときどき　したうけ　にんか　はっこう　つみき　すうじつ　よけい　くげん　もくてき　まわり　せめる　げざい　にげる　にんたい　たんそく　ほそく",
			seed:     "ba369b6718743db50a501ca4bc452763b9230370e923063cd7be7fafaf537c7fadd677cfd2066f78c752f5d5830fb3794983b7e896d58722d559e26060b44309",
		},

		Vector{
			entropy:  "eaebabb2383351fd31d703840b32e9e2",
			mnemonic: "めいえん　さのう　めだつ　すてる　きぬごし　ろんぱ　はんこ　まける　たいおう　さかいし　ねんいり　はぶらし",
			seed:     "065cfeac3b160a68307b6a4d5879b6c8f7ed6c9de396abb8bbd26f4dde61c4b45f5977187bd69a228cd521fd0d901a80df90df07a8115c3de05831e549b14b4a",
		},

		Vector{
			entropy:  "7ac45cfe7722ee6c7ba84fbc2d5bd61b45cb2fe5eb65aa78",
			mnemonic: "せんぱい　おしえる　ぐんかん　もらう　きあい　きぼう　やおや　いせえび　のいず　じゅしん　よゆう　きみつ　さといも　ちんもく　ちわわ　しんせいじ　とめる　はちみつ",
			seed:     "a3e06b761cd1ddde4f652856c495b53c67f84e23a545f0a97b79f94e84ebcab5999439124275e2e118cb03d34772f5b03bb2d3d048a532e019aa6e7121b39b9c",
		},

		Vector{
			entropy:  "4fa1a8bc3e6d80ee1316050e862c1812031493212b7ec3f3bb1b08f168cabeef",
			mnemonic: "こころ　いどう　きあつ　そうがんきょう　へいあん　せつりつ　ごうせい　はいち　いびき　きこく　あんい　おちつく　きこえる　けんとう　たいこ　すすめる　はっけん　ていど　はんおん　いんさつ　うなぎ　しねま　れいぼう　みつかる",
			seed:     "37ed8facbb2fcad238893671e9e12fe25f612f1ec5c39c38f3c0b332d6e5b9fb38902dfc9b3e664029a13adab9e8a1ed5869ed9d0a5854974dd5f608676064b7",
		},

		Vector{
			entropy:  "18ab19a9f54a9274f03e5209a2ac8a91",
			mnemonic: "うりきれ　さいせい　じゆう　むろん　とどける　ぐうたら　はいれつ　ひけつ　いずれ　うちあわせ　おさめる　おたく",
			seed:     "db0b8914d12023ea9c2ffacca9e98cde2afd22aa636811c1043ec5df842c8f8f71a5425b7c2d579d88e214f5c27f4a24b940666c6c8542b5b46414ad8e023930",
		},

		Vector{
			entropy:  "18a2e1d81b8ecfb2a333adcb0c17a5b9eb76cc5d05db91a4",
			mnemonic: "うりきれ　うねる　せっさたくま　きもち　めんきょ　へいたく　たまご　ぜっく　びじゅつかん　さんそ　むせる　せいじ　ねくたい　しはらい　せおう　ねんど　たんまつ　がいけん",
			seed:     "6a6436f5a2353a9fc8f091d49bedc6f51ca23987dc32ea9798786a2d94191146f36604aecffd8494db8c5eac7e858e7e17e1e2eeae8b7dead483e02ea9c939a6",
		},

		Vector{
			entropy:  "15da872c95a13dd738fbf50e427583ad61f18fd99f628c417a61cf8343c90419",
			mnemonic: "うちゅう　ふそく　ひしょ　がちょう　うけもつ　めいそう　みかん　そざい　いばる　うけとる　さんま　さこつ　おうさま　ぱんつ　しひょう　めした　たはつ　いちぶ　つうじょう　てさぎょう　きつね　みすえる　いりぐち　かめれおん",
			seed:     "37ff351d26601c20cab59aed72ba7cdff4bd485fdb70fc2bb25c96d6815ce6c506468cc3fc4bd233cd67affa04bd759c29d61ac3e18db0a4301ef28ef230e792",
		},
	}
}
