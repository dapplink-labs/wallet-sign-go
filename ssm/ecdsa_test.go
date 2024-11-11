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
	privKey := "960014c429fc4afe27ce42a026a6f51d39f744b2b589322633f11310942d81b2"
	message := "0xf6b0ccfd4c046b06fdca66537a7374e510cad44ca01c1b6364ec6bf8f879b019"
	signature, err := SignECDSAMessage(privKey, message)
	if err != nil {
		fmt.Println("sign tx fail")
	}
	fmt.Println("Signature: ", signature)
}

func TestVerifySign(t *testing.T) {
	msgHash := "0000000000000000000000000000000000000000000000000000000000000000"
	// 04fd13e3e6895a2d9d38821e47f60f6c701da7bf243cc1f0a21a15ff8b09322ceda16bbc6d07df0e408dc19ffb756a75bf0c7cbf6836ef89594e0d97cbfeff319c
	publicKey := "fd13e3e6895a2d9d38821e47f60f6c701da7bf243cc1f0a21a15ff8b09322ceda16bbc6d07df0e408dc19ffb756a75bf0c7cbf6836ef89594e0d97cbfeff319c"
	signature := "9e922e8deea041d06461b1b1d228593d6ce56427cd5df6b10d177bc2dbf138117bc5543bdc44790f8734f8171903fd324547bad54c403d5f7f7652af115571be01"
	ok := VerifyEdDSASign(publicKey, msgHash, signature)
	fmt.Println("ok==", ok)
}
