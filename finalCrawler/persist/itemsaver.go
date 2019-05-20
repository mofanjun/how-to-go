package persist

import (
	"log"
	"gopkg.in/olivere/elastic.v6"
	"context"
	"crawler/finalCrawler/engine"
	"errors"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <- out
			log.Printf("Item Saver : got item #%d:%v",itemCount,item)
			itemCount++

			err := save(client, index, item)

			if err != nil {
				log.Printf("save item got error #Item:%v, #error:%v",item,err)
			}
		}
	}()
	return out, nil
}

func save(client *elastic.Client,
	index string,item engine.Item) error {
	//verify Type
	if item.Type == "" {
		return errors.New("must supply Type")
	}

	indexService := client.Index().
		Index(index).//数据库名字-可配置
		Type(item.Type).
		BodyJson(item)

	//verify Id
	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_,err := indexService.
		Do(context.Background())

	if err != nil {
		return err
	}

	return nil
}
