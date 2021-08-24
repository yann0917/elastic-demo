package client

import (
	"log"

	"github.com/olivere/elastic/v7"
)

var (
	err     error
	Client  *elastic.Client
	servers = []string{"http://127.0.0.1:9200"}
)

func init() {
	Client, err = elastic.NewClient(
		elastic.SetSniff(false), // 连接docker中的es需要设置为 false，https://github.com/olivere/elastic/issues/810
		elastic.SetURL(servers...),
	)
	if err != nil {
		log.Println(err)
	}
	// exists, err := client.IndexExists(indexName).Do(ctx)
	// if err != nil {
	// 	panic(err)
	// }
	// if !exists {
	// 	_, err := client.CreateIndex(indexName).BodyString(mapping).Do(ctx)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
}
