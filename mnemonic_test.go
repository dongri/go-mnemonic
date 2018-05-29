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
		seed := ToSeedHex(vector.mnemonic, "がちょう")
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
			seed:     "87b8f75a471cd521050933bba59ed2b514b539be850ff317af2eebe6238aa5cdc5cfac06a07bd5e5473c1c5ccb0f6fd676555f865b914315f0b101c998db7d82",
		},
		Vector{
			entropy:  "7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f",
			mnemonic: "そつう　れきだい　ほんやく　わかす　りくつ　ばいか　ろせん　やちん　そつう　れきだい　ほんやく　わかめ",
			seed:     "266bbc369da18b3f68cde2cf9aa6e7bd8af5fc26fab2a62025d545a94e6d2f5ef91f3124699ae1d05a6c54219e590d2521645f050a5840b1dce57773ce7e7b68",
		},

		Vector{
			entropy:  "80808080808080808080808080808080",
			mnemonic: "そとづら　あまど　おおう　あこがれる　いくぶん　けいけん　あたえる　いよく　そとづら　あまど　おおう　あかちゃん",
			seed:     "70bdc20925845a47babc52eef9f98d62a1deb1ba168a81612cde4e44ecdccd2eb4e58cebfed471ea889f97d9aad9bb9db78ee2e61e570d378e9536416b41988d",
		},

		Vector{
			entropy:  "ffffffffffffffffffffffffffffffff",
			mnemonic: "われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　ろんぶん",
			seed:     "a90ae280202a8170cc1a5cdd57dc2a212981c5f2bbf140c4e2455a38f98a19bda2f1637f1f4ad6da7d5d7127187e63ced7b6de8390e30487b70ed42da569f082",
		},

		Vector{
			entropy:  "000000000000000000000000000000000000000000000000",
			mnemonic: "あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あらいぐま",
			seed:     "83b5919b0504831f45b65265528933ef64fc0af7354acb6668d35a8c037e4a53b8b2789363cd6665bf633eddd6206e49e3d5eee4a268e6dcd22a812cfe88d43d",
		},

		Vector{
			entropy:  "7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f",
			mnemonic: "そつう　れきだい　ほんやく　わかす　りくつ　ばいか　ろせん　やちん　そつう　れきだい　ほんやく　わかす　りくつ　ばいか　ろせん　やちん　そつう　れいぎ",
			seed:     "43746c110efa6d97562d61501c9756115ac421c8b74ce4c000fb0dcaef82270ca8c28354d8cb0f22e5c6ea462cf6e048a311466343b1ef17a7fb0220a2c62391",
		},

		Vector{
			entropy:  "808080808080808080808080808080808080808080808080",
			mnemonic: "そとづら　あまど　おおう　あこがれる　いくぶん　けいけん　あたえる　いよく　そとづら　あまど　おおう　あこがれる　いくぶん　けいけん　あたえる　いよく　そとづら　いきなり",
			seed:     "a9842f86a5fe7b72fe9c17e197ef805b04afb19534a8a74414a047ea109f06723e7ddce2222b3ecb05d93276af20345907b52e6e90dbf775c022c168f9a481a8",
		},

		Vector{
			entropy:  "ffffffffffffffffffffffffffffffffffffffffffffffff",
			mnemonic: "われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　りんご",
			seed:     "c752d15cb7cd60169185f3a227ae62b434df0d13996798b48c4905a11ea9d080a7288ff586e9d1f438eaf9440fc9a779b6130eba53f64bded5426ef4cd8e931f",
		},

		Vector{
			entropy:  "0000000000000000000000000000000000000000000000000000000000000000",
			mnemonic: "あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　いってい",
			seed:     "56e219a639fe994f7b00f40cf0abe3487079eb7a131b6b6cd409a7ad3dd2a37df494a90782b646793cf2630991599cf23e87ef9e81193b3a8aa46706de2093e0",
		},

		Vector{
			entropy:  "7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f",
			mnemonic: "そつう　れきだい　ほんやく　わかす　りくつ　ばいか　ろせん　やちん　そつう　れきだい　ほんやく　わかす　りくつ　ばいか　ろせん　やちん　そつう　れきだい　ほんやく　わかす　りくつ　ばいか　ろせん　まんきつ",
			seed:     "47f8730c69f9d80ec95c1bfdaed516cab76981edad3d0677e10667c301f2dc52823b6ada818c912a36d083b7c1b4839b68ec4e4f53776636a903b56f3cd497e3",
		},

		Vector{
			entropy:  "8080808080808080808080808080808080808080808080808080808080808080",
			mnemonic: "そとづら　あまど　おおう　あこがれる　いくぶん　けいけん　あたえる　いよく　そとづら　あまど　おおう　あこがれる　いくぶん　けいけん　あたえる　いよく　そとづら　あまど　おおう　あこがれる　いくぶん　けいけん　あたえる　うめる",
			seed:     "627b71a70c170625016895c0df96dfc5065353c69646b82a893af7367df8eba9bbf7e86e41a50d16a46b643daa3f2fbb65c2072caf0b3fd87880aefc196ebe0c",
		},

		Vector{
			entropy:  "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
			mnemonic: "われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　われる　らいう",
			seed:     "b9d167e30c4462106d425e4128fc73dfbdc2169d4b5eab19a5e1d4db8dd95c701d2993d09185ef1f8902f14cc415553addfa10c41e54c20ae7263fb6e1954ff9",
		},

		Vector{
			entropy:  "77c2b00716cec7213839159e404db50d",
			mnemonic: "せまい　うちがわ　あずき　かろう　めずらしい　だんち　ますく　おさめる　ていぼう　あたる　すあな　えしゃく",
			seed:     "733c95dff13bf833921be380fee0441bfdd173e0ad2896fe0d953eae2e79b33752539f2a036080c87d7bf6d96af749b77ffec9abf2a884a62d3b84b94ca08f66",
		},

		Vector{
			entropy:  "b63a9c59a6e641f288ebc103017f1da9f8290b3da6bdef7b",
			mnemonic: "ぬすむ　ふっかつ　うどん　こうりつ　しつじ　りょうり　おたがい　せもたれ　あつめる　いちりゅう　はんしゃ　ごますり　そんけい　たいちょう　らしんばん　ぶんせき　やすみ　ほいく",
			seed:     "00f1d43653c3dd25f619da0ccae8e8b96e15f8d56dbc64f57044346d278486bdba3984b99199dda9c115982914ab19c591a2ef37aec35141b662bc75a068a94a",
		},

		Vector{
			entropy:  "3e141609b97933b66a060dcddc71fad1d91677db872031e85f4c015c5e7e8982",
			mnemonic: "くのう　てぬぐい　そんかい　すろっと　ちきゅう　ほあん　とさか　はくしゅ　ひびく　みえる　そざい　てんすう　たんぴん　くしょう　すいようび　みけん　きさらぎ　げざん　ふくざつ　あつかう　はやい　くろう　おやゆび　こすう",
			seed:     "06e20222cc51574e49e352a7c1da974f63ab341e8016fa5f8cc7991a5823c398de0d1b56baa91979a9a6b8dcf9715388f2d119139b5785d049e6d74383502975",
		},

		Vector{
			entropy:  "0460ef47585604c5660618db2e6a7e7f",
			mnemonic: "あみもの　いきおい　ふいうち　にげる　ざんしょ　じかん　ついか　はたん　ほあん　すんぽう　てちがい　わかめ",
			seed:     "5960ec7b5db03b3c24d4d6bf99696c3d3a50ca185e3f2ee56a39d8c59ef569f6e3d4c5969e102499c9abece2456aa07009a2094ba33c90c32ca59d0c9ea8b5e8",
		},

		Vector{
			entropy:  "72f60ebac5dd8add8d2a25a797102c3ce21bc029c200076f",
			mnemonic: "すろっと　にくしみ　なやむ　たとえる　へいこう　すくう　きない　けってい　とくべつ　ねっしん　いたみ　せんせい　おくりがな　まかい　とくい　けあな　いきおい　そそぐ",
			seed:     "d19e9c6e6877677b61c5cf9a71bdef6de98bb3892ab33f596c615274f187bb830bcc03d2ef14b8401d86691e95b2870fff35574bc700de6e67da328a33c160ea",
		},

		Vector{
			entropy:  "2c85efc7f24ee4573d2b81a6ec66cee209b2dcbd09d8eddc51e0215b0b68e416",
			mnemonic: "かほご　きうい　ゆたか　みすえる　もらう　がっこう　よそう　ずっと　ときどき　したうけ　にんか　はっこう　つみき　すうじつ　よけい　くげん　もくてき　まわり　せめる　げざい　にげる　にんたい　たんそく　ほそく",
			seed:     "51fa46c7a4fd42222202235285bc5db24e5f72c6cba90d18c13ae10405df03a715b57c06b33c2b67c21e6fdf2c52c339f896f4dc895f0a4e2a57ae2654440040",
		},

		Vector{
			entropy:  "eaebabb2383351fd31d703840b32e9e2",
			mnemonic: "めいえん　さのう　めだつ　すてる　きぬごし　ろんぱ　はんこ　まける　たいおう　さかいし　ねんいり　はぶらし",
			seed:     "19a813ff6d24f42b0aa355ea9146e2dcdf260828a89e4b68d25b6b0db7d846b6d99cabf77f57a28b59ac0772dc00035870c0a4349dcf17a348f46e2101fc27e7",
		},

		Vector{
			entropy:  "7ac45cfe7722ee6c7ba84fbc2d5bd61b45cb2fe5eb65aa78",
			mnemonic: "せんぱい　おしえる　ぐんかん　もらう　きあい　きぼう　やおや　いせえび　のいず　じゅしん　よゆう　きみつ　さといも　ちんもく　ちわわ　しんせいじ　とめる　はちみつ",
			seed:     "2616eed5faefecc3d71015da23b041c84d6f8ba3680975d3affee6989fd7594f9392e0c2bbe98030ebb87b4f2fcc206049dd86d0eafccb121e4821148c9aca09",
		},

		Vector{
			entropy:  "4fa1a8bc3e6d80ee1316050e862c1812031493212b7ec3f3bb1b08f168cabeef",
			mnemonic: "こころ　いどう　きあつ　そうがんきょう　へいあん　せつりつ　ごうせい　はいち　いびき　きこく　あんい　おちつく　きこえる　けんとう　たいこ　すすめる　はっけん　ていど　はんおん　いんさつ　うなぎ　しねま　れいぼう　みつかる",
			seed:     "86113f18e13ce0ea4e3077e660a648bbd4d517cefed8b29ef965194aae1563397380d6193414446357bf41e66d4a8dadf0e17668ac6866e0673767dc7677ddde",
		},

		Vector{
			entropy:  "18ab19a9f54a9274f03e5209a2ac8a91",
			mnemonic: "うりきれ　さいせい　じゆう　むろん　とどける　ぐうたら　はいれつ　ひけつ　いずれ　うちあわせ　おさめる　おたく",
			seed:     "ba8927473ee626306bc1a0b88473a11645c199f9942f6699e0c3e642b82362e4a5a1752af1f0e78eee9f14cb92a908e577c1e951275a9855a46ba06d514c1a59",
		},

		Vector{
			entropy:  "18a2e1d81b8ecfb2a333adcb0c17a5b9eb76cc5d05db91a4",
			mnemonic: "うりきれ　うねる　せっさたくま　きもち　めんきょ　へいたく　たまご　ぜっく　びじゅつかん　さんそ　むせる　せいじ　ねくたい　しはらい　せおう　ねんど　たんまつ　がいけん",
			seed:     "65ee1d9d9b393b25eff8966e2401cddeda67c3d82f4f21643156b8036f739922ae275444c5832866b4ce0aace7ea00f17d3f7ad427aab57435179ab46781d67e",
		},

		Vector{
			entropy:  "15da872c95a13dd738fbf50e427583ad61f18fd99f628c417a61cf8343c90419",
			mnemonic: "うちゅう　ふそく　ひしょ　がちょう　うけもつ　めいそう　みかん　そざい　いばる　うけとる　さんま　さこつ　おうさま　ぱんつ　しひょう　めした　たはつ　いちぶ　つうじょう　てさぎょう　きつね　みすえる　いりぐち　かめれおん",
			seed:     "79897293a5982c04a4c7732e0dfa5aa493eb24567bc84a43481ea804b6c092e7b0695c5ae5c7be618619120e51d0555ad542e902af3362f0b71435c892364d53",
		},
	}
}
