package ssm

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/log"
)

func CreateEdDSAKeyPair() (string, string, error) {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		log.Error("create key pair fail:", "err", err)
		return "0x00", "0x00", nil
	}
	return hex.EncodeToString(privateKey), hex.EncodeToString(publicKey), nil
}

func SignEdDSAMessage(priKey string, txMsg string) (string, error) {
	privateKey, _ := hex.DecodeString(priKey)
	txMsgByte, _ := hex.DecodeString(txMsg)
	signMsg := ed25519.Sign(privateKey, txMsgByte)

	//publicKey := privateKey[32:]
	//if !ed25519.Verify(publicKey, txMsgByte, signMsg) {
	//	return "", fmt.Errorf("signature verification failed")
	//}

	return hex.EncodeToString(signMsg), nil
}

func VerifyEdDSASign(pubKey, msgHash, sig string) bool {
	publicKeyByte, _ := hex.DecodeString(pubKey)
	msgHashByte, _ := hex.DecodeString(msgHash)
	signature, _ := hex.DecodeString(sig)
	return ed25519.Verify(publicKeyByte, msgHashByte, signature)
}
