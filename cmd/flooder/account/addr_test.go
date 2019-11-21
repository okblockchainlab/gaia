package account

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"runtime"
	"testing"
)

func TestGenerateAccount(t *testing.T) {
	fmt.Println(runtime.GOMAXPROCS(4))
	kb := keys.NewInMemory()
	accs := GenerateAccount(kb, 100)
	fmt.Println(accs)
}
