package ssm

import (
	"fmt"
	"testing"
)

func TestCreateECDSAKeyPair(t *testing.T) {
	privKey, pubKey, cpubKey, _ := CreateECDSAKeyPair()
	fmt.Println("privKey=", privKey)
	fmt.Println("pubKey=", pubKey)
	fmt.Println("cpubKey=", cpubKey)
}

func TestSignMessage(t *testing.T) {
	// 0x35096AD62E57e86032a3Bb35aDaCF2240d55421D
	privKey := "fe13c8e55444107c32f50cca04965f9772fe4fd720ffedd30b347e541fe7a97c"
	message := "0x3e4f9a460233ec33862da1ac3dabf5b32db01400fba166cdec40ad6dc735b4ab"
	signature, err := SignECDSAMessage(privKey, message)
	if err != nil {
		fmt.Println("sign tx fail")
	}
	fmt.Println("Signature: ", signature)
}

func TestVerifySign(t *testing.T) {
	msgHash := "0000000000000000000000000000000000000000000000000000000000000000"
	// 04fd13e3e6895a2d9d38821e47f60f6c701da7bf243cc1f0a21a15ff8b09322ceda16bbc6d07df0e408dc19ffb756a75bf0c7cbf6836ef89594e0d97cbfeff319c
	publicKey := "04fd13e3e6895a2d9d38821e47f60f6c701da7bf243cc1f0a21a15ff8b09322ceda16bbc6d07df0e408dc19ffb756a75bf0c7cbf6836ef89594e0d97cbfeff319c"
	signature := "9e922e8deea041d06461b1b1d228593d6ce56427cd5df6b10d177bc2dbf138117bc5543bdc44790f8734f8171903fd324547bad54c403d5f7f7652af115571be01"
	ok := VerifyEdDSASign(publicKey, msgHash, signature)
	fmt.Println("ok==", ok)
}
