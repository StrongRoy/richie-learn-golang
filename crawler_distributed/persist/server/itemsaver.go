package main

import (
	"flag"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"richie.com/richie/learn_golang/crawler_distributed/config"
	"richie.com/richie/learn_golang/crawler_distributed/persist"
	"richie.com/richie/learn_golang/crawler_distributed/rpcsupport"
)

func serverRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return err
	}

	log.Printf("Listening on %s",host)

	return rpcsupport.ServeRPC(host,
		&persist.ItemSaverService{
			Client: client,
			Index:  index,
		})

}

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatalln(serverRpc(fmt.Sprintf(":%d", *port),
		config.ElasticIndex))
}
