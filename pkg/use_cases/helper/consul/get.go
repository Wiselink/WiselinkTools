package consul

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

func GetKv(key, Address, dirconsul string) []byte {
	config := api.DefaultConfig()
	config.Token = " "
	config.Address = Address

	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
	kv := client.KV()

	pathconsul := dirconsul + key

	k, _, err2 := kv.Get(pathconsul, nil)

	if err2 != nil {
		fmt.Println("error", err2.Error())
	}

	return k.Value
}
