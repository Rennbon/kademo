package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

func main() {

	conf := api.DefaultConfig()
	fmt.Println(conf)
	// Get a new client
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	// Get a handle to the KV API
	kv := client.KV()

	// PUT a new KV pair
	p := &api.KVPair{Key: "REDIS_MAXCLIENTS", Value: []byte("1000")}
	_, err = kv.Put(p, nil)
	if err != nil {
		panic(err)

	}
	// Lookup the pair
	pair, _, err := kv.Get("REDIS_MAXCLIENTS", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("KV: %v %s\n", pair.Key, pair.Value)

	agent := client.Agent()
	fmt.Println(agent.Checks())
	fmt.Println(agent.Services())

	fmt.Println("---------------------------------------------------------")
	infomap, _ := agent.Self()
	for k, v := range infomap {
		fmt.Println(k, v)
	}

}
