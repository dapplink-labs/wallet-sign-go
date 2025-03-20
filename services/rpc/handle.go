package rpc

import (
	"context"

	"github.com/ethereum/go-ethereum/log"
	"github.com/pkg/errors"

	"github.com/dapplink-labs/wallet-sign-go/leveldb"
	"github.com/dapplink-labs/wallet-sign-go/protobuf"
	"github.com/dapplink-labs/wallet-sign-go/protobuf/wallet"
	"github.com/dapplink-labs/wallet-sign-go/ssm"
)

func (s *RpcServer) GetSupportSignWay(ctx context.Context, in *wallet.SupportSignWayRequest) (*wallet.SupportSignWayResponse, error) {
	var signWay []*wallet.SignWay
	signWay = append(signWay, &wallet.SignWay{Schema: "ecdsa"})
	signWay = append(signWay, &wallet.SignWay{Schema: "eddsa"})
	return &wallet.SupportSignWayResponse{
		Code:    wallet.ReturnCode_SUCCESS,
		Msg:     "get sign way success",
		SignWay: signWay,
	}, nil
}

func (s *RpcServer) ExportPublicKeyList(ctx context.Context, in *wallet.ExportPublicKeyRequest) (*wallet.ExportPublicKeyResponse, error) {
	resp := &wallet.ExportPublicKeyResponse{
		Code: wallet.ReturnCode_ERROR,
	}
	cryptoType, err := protobuf.ParseTransactionType(in.Type)
	if err != nil {
		resp.Msg = "input type error"
		return resp, nil
	}
	if in.Number > 10000 {
		resp.Msg = "Number must be less than 100000"
		return resp, nil
	}

	var keyList []leveldb.Key
	var retKeyList []*wallet.PublicKey

	for counter := 0; counter < int(in.Number); counter++ {
		var priKeyStr, pubKeyStr, compressPubkeyStr string
		var err error

		switch cryptoType {
		case protobuf.ECDSA:
			priKeyStr, pubKeyStr, compressPubkeyStr, err = ssm.CreateECDSAKeyPair()
		case protobuf.EDDSA:
			priKeyStr, pubKeyStr, err = ssm.CreateEdDSAKeyPair()
			compressPubkeyStr = pubKeyStr
		default:
			return nil, errors.New("unsupported key type")
		}
		if err != nil {
			log.Error("create key pair fail", "err", err)
			return nil, err
		}

		keyItem := leveldb.Key{
			PrivateKey: priKeyStr,
			Pubkey:     pubKeyStr,
		}
		pukItem := &wallet.PublicKey{
			CompressPubkey: compressPubkeyStr,
			Pubkey:         pubKeyStr,
		}
		retKeyList = append(retKeyList, pukItem)
		keyList = append(keyList, keyItem)
	}
	isOk := s.db.StoreKeys(keyList)
	if !isOk {
		log.Error("store keys fail", "isOk", isOk)
		return nil, errors.New("store keys fail")
	}
	resp.Code = wallet.ReturnCode_SUCCESS
	resp.Msg = "create keys success"
	resp.PublicKey = retKeyList
	return resp, nil
}

func (s *RpcServer) SignTxMessage(ctx context.Context, in *wallet.SignTxMessageRequest) (*wallet.SignTxMessageResponse, error) {
	resp := &wallet.SignTxMessageResponse{
		Code: wallet.ReturnCode_ERROR,
	}
	cryptoType, err := protobuf.ParseTransactionType(in.Type)
	if err != nil {
		resp.Msg = "input type error"
		return resp, nil
	}

	privKey, isOk := s.db.GetPrivKey(in.PublicKey)
	if !isOk {
		return nil, errors.New("get private key by public key fail")
	}

	var signature string
	var err2 error

	switch cryptoType {
	case protobuf.ECDSA:
		signature, err2 = ssm.SignECDSAMessage(privKey, in.MessageHash)
	case protobuf.EDDSA:
		signature, err2 = ssm.SignEdDSAMessage(privKey, in.MessageHash)
	default:
		return nil, errors.New("unsupported key type")
	}
	if err2 != nil {
		return nil, err2
	}
	resp.Msg = "sign tx message success"
	resp.Signature = signature
	resp.Code = wallet.ReturnCode_SUCCESS
	return resp, nil
}
