package keychain

import (
	"github.com/99designs/keyring"
)

const serviceName = "totp-bio"

var ring keyring.Keyring

func init() {
	ring, _ = keyring.Open(keyring.Config{
		KeychainTrustApplication: true,
		ServiceName:              serviceName,
	})
}

func AddItem(name, secret string) error {
	return ring.Set(keyring.Item{
		Key:   name,
		Label: name,
		Data:  []byte(secret),
	})
}

func GetItem(name string) (string, error) {
	item, err := ring.Get(name)
	if err != nil {
		return "", err
	}
	return string(item.Data), nil
}

func DeleteItem(name string) error {
	return ring.Remove(name)
}

func GetItems() ([]string, error) {
	items, err := ring.Keys()
	if err != nil {
		return nil, err
	}
	return items, nil
}
