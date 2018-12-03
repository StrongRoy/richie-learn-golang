package client

import (
	"log"
	"richie.com/richie/learn_golang/crawler/engine"
	"richie.com/richie/learn_golang/crawler_distributed/config"
	"richie.com/richie/learn_golang/crawler_distributed/rpcsupport"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item saver: got item #%d: %v",
				itemCount, item)
			itemCount++

			// Call RPC to save item
			result := ""
			err = client.Call(config.ItemSaverRpc, item, &result)
			if err != nil || result != "ok" {
			}
			if err != nil {
				log.Printf("Item Saver: error"+
					"saving item %v: %v", item, err)
			}
		}

	}()
	return out, nil
}
