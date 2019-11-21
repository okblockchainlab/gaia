package account

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/cosmos/go-bip39"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"sync"
)

const entropySize = 256

var (
	mnemonicNum uint
	mnemonFilePath string
	addrFilePath string
	number uint
)

func getMnemonicAndAddrToFile() *cobra.Command{
	var maCmd = &cobra.Command{
		Use:   "save_to_file",
		Short: "generate mnemonics and address to file",
		Long:  "generate mnemonics and address to file",
		RunE:  mnemonicAndAddrToFile,
	}

	flags:=maCmd.Flags()
	flags.StringVar(&mnemonFilePath,"mnemonic_file_path","prikey.txt","file path of mnemonics")
	flags.StringVar(&addrFilePath,"address_file_path","addr.txt","file path of addresses")
	flags.UintVar(&number,"number",1000,"the number of mnemonics&&addresses")
	return maCmd
}

func getMnemonicCmd() *cobra.Command {
	var mnemonicCmd = &cobra.Command{
		Use:   "mnemonic",
		Short: "generate mnemonics",
		Long:  "generate the account mnemonics",
		RunE:  generateMnemonics,
	}

	mnemonicCmd.Flags().UintVarP(&mnemonicNum, "number", "n", 1, "the number of mnemonics 2 generate")
	return mnemonicCmd
}

func generateMnemonics(*cobra.Command, []string) error {
	var i uint
	for ; i < mnemonicNum; i++ {
		seed, err := bip39.NewEntropy(entropySize)
		if err != nil {
			return err
		}

		mnemonic, err := bip39.NewMnemonic(seed)
		if err != nil {
			return err
		}
		fmt.Println(mnemonic)
	}
	return nil
}

func mnemonicAndAddrToFile(*cobra.Command, []string) error {
	os.Remove(mnemonFilePath)
	os.Remove(addrFilePath)
	w1,err:=os.OpenFile(mnemonFilePath,os.O_RDWR|os.O_CREATE,0644)
	if err!=nil{
		fmt.Println(err)
		return err
	}
	defer w1.Close()

	w2,err:=os.OpenFile(addrFilePath,os.O_RDWR|os.O_CREATE,0644)
	if err!=nil{
		fmt.Println(err)
		return err
	}
	defer w2.Close()

	kb := keys.NewInMemory()
	fmt.Println("generating ...")
	accs:=GenerateMnemonicAndAccInfo(kb,number)
	fmt.Println("writing 2 files")
	for i,v:=range accs{
		content1,content2:=[]byte(v.Mnemonic+"\n"),[]byte(v.Address+"\n")
		_,err := w1.Write(content1)
		if err != nil {
			fmt.Println("error:",i,err,v)
		}
		_,err = w2.Write(content2)
		if err != nil {
			fmt.Println("error:",i,err,v)
		}
	}
	return nil
}

// mnemonic and address
type Combination struct{
	Mnemonic string
	Address string
}


func GenerateMnemonicAndAccInfo(kb keys.Keybase, genNum uint) []Combination {
	indexChan := make(chan uint, 1000)
	combinationChan := make(chan Combination, 1000)

	for i := 0; i < goRoutineNum; i++ {
		go genCombinations(kb, indexChan, combinationChan)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		var addrIndex uint
		for ; addrIndex < genNum; addrIndex++ {
			indexChan <- addrIndex
		}
		wg.Done()
	}()

	combinations := make([]Combination, 0)
	var i uint
	for ; i < genNum; i++ {
		com := <-combinationChan
		combinations = append(combinations, com)
	}
	wg.Wait()
	return combinations
}

func genCombinations(kb keys.Keybase, indexChan chan uint, comChan chan Combination) error {
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

		info, err := kb.CreateAccount(name, mnemonic, defaultBip39Passwd, DefaultEncryptPasswd, defaultAccountNumber, 0)
		if err != nil {
			return err
		}
		//fmt.Println(fmt.Sprintf("%s:%s,%s len(camChan):%d", info.GetName(), info.GetAddress().String(),mnemonic,len(comChan)))
		comChan <- Combination{mnemonic,info.GetAddress().String()}

	}
	return nil
}
