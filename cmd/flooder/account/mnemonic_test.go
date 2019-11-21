package account

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"os"
	"testing"
)

// generate Mnemonics and addresses 2 files
func TestGenerateMnemonicAndAccInfo(t *testing.T) {
	filePath1 := "prikey.txt"
	filePath2 := "addr.txt"
	os.Remove(filePath1)
	os.Remove(filePath2)
	w1, err := os.OpenFile(filePath1, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	w2, err := os.OpenFile(filePath2, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	kb := keys.NewInMemory()
	fmt.Println("generating ...")
	accs := GenerateMnemonicAndAccInfo(kb, 300000)
	fmt.Println("writing 2 files")
	for i, v := range accs {
		if i%10000 == 0 {
			fmt.Println(fmt.Sprintf("%d already ...", i))
		}
		content1, content2 := []byte(v.Mnemonic), []byte(v.Address)
		content1 = append(content1, 0x0D)
		content2 = append(content2, 0x0D)
		w1.Write(content1)
		w2.Write(content2)

	}

	w1.Close()
	w2.Close()

}

func TestMultiSend(t *testing.T) {
	kb := keys.NewInMemory()
	fmt.Println("generating ...")
	accs := GenerateMnemonicAndAccInfo(kb, 3000)
	var bz []byte
	for _, v := range accs {
		bz=append(bz,[]byte(v.Address)...)
		bz=append(bz,',')
	}

	w1, err := os.OpenFile("para.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer w1.Close()
	w1.Write(bz)
}
