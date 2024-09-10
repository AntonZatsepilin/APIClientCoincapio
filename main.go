package main

import (
	"fmt"
	"golangClient/coincap"
	"log"
	"time"
)

func main() {
	coincapClient, err := coincap.NewClient(time.Second * 10)
	if err != nil {
		log.Fatal(err)
	}
	assets, err := coincapClient.GetAssets()
	if err != nil {
		log.Fatal(err)
	}

	for _, asset := range assets {
		fmt.Println(asset.Info())
	}
}
