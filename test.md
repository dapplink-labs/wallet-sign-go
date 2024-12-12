
# api test

## 1.支持算法

- Request
```
grpcurl -plaintext -d '{
  "type": "ecdsa"
}' 127.0.0.1:8983 dapplink.wallet.WalletService.getSupportSignWay
```

- Response
```
{
  "code": "1",
  "msg": "Support this sign way",
  "support": true
}
```


## 2.导出公钥列表

- Request
```
grpcurl -plaintext -d '{
  "type": "ecdsa",
  "number": "5"
}' 127.0.0.1:8983 dapplink.wallet.WalletService.exportPublicKeyList
```

- Response
```
{
  "code": "1",
  "msg": "create keys success",
  "public_key": [
    {
      "compress_pubkey": "0474b6a198569a07efca7817f259124251970d4f6ee80508e6ced40be8895425689629f97abc5e262005960c0117e4a59067a2949d757cbd2d40d6bd9bd1acbef0",
      "decompress_pubkey": "0274b6a198569a07efca7817f259124251970d4f6ee80508e6ced40be889542568"
    },
    {
      "compress_pubkey": "04c2fb70e8bb957e5be63b05dd35be4e62ba0e84330fc82794292ff58447462a53b52d60fe0a78f1218117279830b558245cf7c1f5e82735d2680b394addc7117c",
      "decompress_pubkey": "02c2fb70e8bb957e5be63b05dd35be4e62ba0e84330fc82794292ff58447462a53"
    },
    {
      "compress_pubkey": "048d689651ab8853122ea62fc87e2eaf7397d33b2e24c7f468e2cab19311f677f78ab4ed0e2f7c0c9844c00413dca97fe3db77795a20802d8bc2f4a1d12af4097f",
      "decompress_pubkey": "038d689651ab8853122ea62fc87e2eaf7397d33b2e24c7f468e2cab19311f677f7"
    },
    {
      "compress_pubkey": "04570a453caae2639495b6c26b5a7d7142e7d02ca73410a07bdd984467c6034224744724244e2b71123f3afcd6c3a1b0ee3a67d87aaaf26651fdfc720bcc6b0a8f",
      "decompress_pubkey": "03570a453caae2639495b6c26b5a7d7142e7d02ca73410a07bdd984467c6034224"
    },
    {
      "compress_pubkey": "042aa1b0dcb48cb74a1027285bb5594ff6b2e65a2b6cc2e5ec91b71cb5e60d071e7e3eabd1efbc84a59c8c68f00c78690c55d4086f875511905d47a07780f966a5",
      "decompress_pubkey": "032aa1b0dcb48cb74a1027285bb5594ff6b2e65a2b6cc2e5ec91b71cb5e60d071e"
    }
  ]
}
```

## 3.消息签名

- Request

```
grpcurl -plaintext -d '{
  "messageHash": "0x9ca77bd43a45da2399da96159b554bebdd89839eec73a8ff0626abfb2fb4b538",
  "publicKey": "04c2fb70e8bb957e5be63b05dd35be4e62ba0e84330fc82794292ff58447462a53b52d60fe0a78f1218117279830b558245cf7c1f5e82735d2680b394addc7117c",
  "type": "ecdsa"
}' 127.0.0.1:8983 dapplink.wallet.WalletService.signTxMessage
```

- Response

```
{
  "code": "1",
  "msg": "sign tx message success",
  "signature": "b3cf4df645d385e57c1498ae64cb32f2caf64546bfcb894b70e7dfc454ede9aa32849ffbac8b6605b2febacf8b605f2b688f899a09eb355e8f540aed1125897f00"
}
```
