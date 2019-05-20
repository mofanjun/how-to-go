package persist

import (
	"log"
	"gopkg.in/olivere/elastic.v6"
	"context"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <- out
			log.Printf("Item Saver : got item #%d:%v",itemCount,item)
			itemCount++

			_, err := save(item)

			if err != nil {
				log.Printf("save item got error #Item:%v, #error:%v",item,err)
			}
		}
	}()
	return out
}

func save(item interface{}) (id string, err error)  {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	if err != nil {
		return "", err
	}

	resp, err := client.Index().
		Index("dating_profile").//数据库名字
		Type("zhenai").//表名
		BodyJson(item).
		Do(context.Background())

	if err != nil {
		return "", err
	}

	return resp.Id, nil
}
