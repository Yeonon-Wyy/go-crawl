package main

import (
	"go-crawl/distributed/config"
	"go-crawl/distributed/persist"
	"go-crawl/distributed/rpcsupport"
	"log"
	"fmt"
	"github.com/olivere/elastic"
)

func main() {
	log.Fatal(ServrRpc(fmt.Sprintf(":%d", config.ItemSaverPort), config.ItemSaverESIndex))
}

func ServrRpc(host string, index string) error {

	esClient, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(config.ElasticSearchUrl))
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: esClient,
		Index:  index,
	})

}
