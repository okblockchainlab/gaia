package account

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/cosmos/go-bip39"
	"runtime"
	"strconv"
)

const (
	TuringName     = "Turing"
	TuringMnemonic = "hover tumble cool toilet drop rhythm rude history provide man memory family strategy language tuna pool extend reform obey circle found usage scrub iron"

	DefaultEncryptPasswd = "12345678"
	defaultBip39Passwd   = ""
	defaultAccountNumber = uint32(0)
)

var (
	goRoutineNum = runtime.NumCPU()
)

func GenerateAccount(kb keys.Keybase, addrNum uint32) []keys.Info {
	indexChan := make(chan uint32, 5000)
	infoChan := make(chan keys.Info, 5000)

	for i := 0; i < goRoutineNum; i++ {
		go genGoRoutine(kb, indexChan, infoChan)
	}

	var addrIndex uint32
	for ; addrIndex < addrNum; addrIndex++ {
		indexChan <- addrIndex
	}

	accounts := make([]keys.Info, 0)

	for i := uint32(0); i < addrNum; i++ {
		info := <-infoChan
		accounts = append(accounts, info)
	}

	close(indexChan)

	return accounts
}

func genGoRoutine(kb keys.Keybase, indexChan chan uint32, accChan chan keys.Info) error {
	for i := range indexChan {
		seed, err := bip39.NewEntropy(entropySize)
		if err != nil {
			return err
		}

		mnemonic, err := bip39.NewMnemonic(seed)
		if err != nil {
			return err
		}

		name := strconv.Itoa(int(i))

		info, err := kb.CreateAccount(name, mnemonic, defaultBip39Passwd, DefaultEncryptPasswd, defaultAccountNumber, i)
		if err != nil {
			return err
		}
		fmt.Println(fmt.Sprintf("%s:%s", info.GetName(), info.GetAddress().String()))
		accChan <- info

	}
	return nil
}
