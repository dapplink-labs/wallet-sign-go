package ssm

import (
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

func CreateECDSAKeyPair() (string, string, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Error("generate key fail", "err", err)
		return "0x00", "0x00", "0x00", err
	}
	priKeyStr := hex.EncodeToString(crypto.FromECDSA(privateKey))
	pubKeyStr := hex.EncodeToString(crypto.FromECDSAPub(&privateKey.PublicKey))
	decPubkeyStr := hex.EncodeToString(crypto.CompressPubkey(&privateKey.PublicKey))
	if err != nil {
		return "0x00", "0x00", "0x00", err
	}
	return priKeyStr, pubKeyStr, decPubkeyStr, nil
}

func SignMessage(privKey string, txMsg string) (string, error) {
	hash := common.HexToHash(txMsg)
	fmt.Println(hash.Hex())
	privByte, err := hex.DecodeString(privKey)
	if err != nil {
		log.Error("decode private key fail", "err", err)
		return "0x00", err
	}
	privKeyEcdsa, err := crypto.ToECDSA(privByte)
	if err != nil {
		log.Error("Byte private key to ecdsa key fail", "err", err)
		return "0x00", err
	}
	signatureByte, err := crypto.Sign(hash[:], privKeyEcdsa)
	if err != nil {
		log.Error("sign transaction fail", "err", err)
		return "0x00", err
	}
	return hex.EncodeToString(signatureByte), nil
}

func VerifySign(pubKey, msgHash, sig string) bool {
	publicKey, _ := hex.DecodeString(pubKey)
	messageHash, _ := hex.DecodeString(msgHash)
	signature, _ := hex.DecodeString(sig)
	fmt.Println(len(publicKey), len(messageHash), len(signature))
	return crypto.VerifySignature(publicKey, messageHash, signature)
}
