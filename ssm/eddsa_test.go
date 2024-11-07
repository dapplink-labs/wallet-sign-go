package ssm

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestCreateEdDSAKeyPair(t *testing.T) {
	privKey, pubKey, _ := CreateEdDSAKeyPair()
	fmt.Println("privKey=", privKey)
	fmt.Println("pubKey=", pubKey)
}

func TestSignEdDSAMessage(t *testing.T) {
	privKey := "09fa5c99a11f3857dccfede0b9f6ead29bc2f5757b43b336796d64d2cdacf74a39f523de37c1218d28ca467a6e0ea0aa0a603064ab402983829513a0feca0039"
	txMsg, _ := SignEdDSAMessage(privKey, common.Hash{}.String())
	fmt.Println("txMsg=", txMsg)
}

func TestVerifyEdDSASign(t *testing.T) {
	signature := "9594742341865c897b066714f10bc90f4c2afebba986c3fef500d9d69ea8eb6df19a19160c210dcf65becbe68c1885820915565f3a84277efc34027b73c89608"
	pubKey := "39f523de37c1218d28ca467a6e0ea0aa0a603064ab402983829513a0feca0039"
	ok := VerifyECDSASign(pubKey, common.Hash{}.String(), signature)
	fmt.Println("ok=", ok)
}
