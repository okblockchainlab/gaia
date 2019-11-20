package account

import "github.com/cosmos/cosmos-sdk/crypto/keys"

const (
	RicherName     = "Turing"
	RicherMnemonic = "hover tumble cool toilet drop rhythm rude history provide man memory family strategy language tuna pool extend reform obey circle found usage scrub iron"

	defaultEncryptPasswd = "12345678"
	workerNum            = 100
)

func GenerateAccount(kb keys.Keybase, addrNum uint) []keys.Info {
	index := make(chan uint, 5000)
	infoChan := make(chan keys.Info, 5000)

	for i := 0; i < workerNum; i++ {
		
	}
}
