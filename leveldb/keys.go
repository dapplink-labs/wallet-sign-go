package leveldb

import "github.com/ethereum/go-ethereum/log"

type Keys struct {
	db *LevelStore
}

func NewKeyStore(path string) (*Keys, error) {
	db, err := NewLevelStore(path)
	if err != nil {
		log.Error("Could not create leveldb database.")
		return nil, err
	}
	return &Keys{
		db: db,
	}, nil
}

func (k *Keys) GetPrivKey(publicKey string) (string, bool) {
	key := []byte(publicKey)
	data, err := k.db.Get(key)
	if err != nil {
		return "0x00", false
	}
	bstr := toString(data)
	return bstr, true
}

func (k *Keys) StoreKeys(keyList []Key) bool {
	for _, item := range keyList {
		key := []byte(item.Pubkey)
		value := toBytes(item.PrivateKey)
		err := k.db.Put(key, value)
		if err != nil {
			log.Error("store key value fail", "err", err, "key", key, "value", value)
			return false
		}
	}
	return true
}
